package keeper_test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"

	"github.com/wealdtech/go-merkletree"
	"github.com/wealdtech/go-merkletree/sha3"
)

type TestFile struct {
	Name string
	Data string
}

var originalFile = TestFile{
	Name: "jackal_file",
	Data: "jackal maxi",
}

var fileFromSP = TestFile{
	Name: "jackal_file",
	Data: "jackal maxi",
}

var randomFile = TestFile{
	Name: "random_file",
	Data: "hello world",
}

const (
	CID  = "jklc1dmcul9svpv0z2uzfv30lz0kcjrpdfmmfccskt06wpy8vfqrhp4nsgvgz32"
	CID2 = "jklc15ftkghzrx2ywyrpr6n7ge6prcej43efe3jvtzsxhenann69rcu8q7jl5uh"
)

func CreateMerkleForProof(file TestFile) (string, []byte, error) {
	f := []byte(file.Data)
	index := 0
	var data [][]byte
	item := f

	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%d%x", index, f))
	if err != nil {
		return "", nil, err
	}
	hashName := h.Sum(nil)

	data = append(data, hashName)

	tree, err := merkletree.NewUsing(data, sha3.New512(), false)
	if err != nil {
		return "", nil, err
	}

	h = sha256.New()
	_, err = io.WriteString(h, fmt.Sprintf("%d%x", index, item))
	if err != nil {
		return "", nil, err
	}
	ditem := h.Sum(nil)

	proof, err := tree.GenerateProof(ditem, 0)
	if err != nil {
		return "", nil, err
	}

	jproof, err := json.Marshal(*proof)
	if err != nil {
		return "", nil, err
	}

	e := hex.EncodeToString(tree.Root())

	k, _ := hex.DecodeString(e)

	verified, err := merkletree.VerifyProofUsing(ditem, false, proof, [][]byte{k}, sha3.New512())
	if err != nil {
		return "", nil, err
	}

	if !verified {
		return "", nil, types.ErrCannotVerifyProof
	}

	return fmt.Sprintf("%x", item), jproof, nil
}

func makeContract(file TestFile) ([]byte, int64, error) {
	f := []byte(file.Data)
	var list [][]byte

	size := len(f)

	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%d%x", 0, f))
	if err != nil {
		return nil, 0, err
	}
	hashName := h.Sum(nil)

	list = append(list, hashName)

	t, err := merkletree.NewUsing(list, sha3.New512(), false)
	if err != nil {
		return nil, 0, err
	}

	return t.Root(), int64(size), nil
}

func (suite *KeeperTestSuite) TestPostProof() {
	suite.SetupSuite()

	msgSrvr, keeper, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 4)
	suite.Require().NoError(err)

	depoAccount := testAddresses[0]

	// harded coded accounts to keep CIDs static for testing
	// Create user account
	user, err := sdk.AccAddressFromBech32(testAddresses[1])
	suite.Require().NoError(err)

	// Create provider account
	testProvider, err := sdk.AccAddressFromBech32(testAddresses[2])
	suite.Require().NoError(err)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, testProvider, sdk.NewCoins(sdk.NewInt64Coin("ujkl", 100000000)))
	suite.Require().NoError(err)

	suite.storageKeeper.SetParams(suite.ctx, types.Params{
		DepositAccount:         depoAccount,
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     8,
		CollateralPrice:        2,
	})

	// Init Provider
	_, err = msgSrvr.InitProvider(context, &types.MsgInitProvider{
		Creator:    testProvider.String(),
		Ip:         "192.168.0.1",
		Totalspace: "1_000_000",
	})
	suite.Require().NoError(err)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, user, sdk.NewCoins(sdk.NewInt64Coin("ujkl", 100000000)))
	suite.Require().NoError(err)

	_, err = msgSrvr.BuyStorage(context, &types.MsgBuyStorage{
		Creator:      user.String(),
		ForAddress:   user.String(),
		Bytes:        "4000000000",
		Duration:     "720h",
		PaymentDenom: "ujkl",
	})
	suite.Require().NoError(err)

	// Storage Provider receives file and make merkleroot for contract
	merkleroot, filesize, err := makeContract(originalFile)
	suite.Require().NoError(err)

	suite.Require().Equal(int64(11), filesize)

	_, found := keeper.GetStoragePaymentInfo(suite.ctx, user.String())
	suite.Require().Equal(true, found)
	// Post Contract
	_, err = msgSrvr.PostFile(context, &types.MsgPostFile{
		Creator:       user.String(),
		Merkle:        merkleroot,
		FileSize:      filesize,
		ProofInterval: 1800,
		ProofType:     0,
		MaxProofs:     3,
		Note:          "",
	})
	suite.Require().NoError(err)
	h := sha256.New()
	_, err = io.WriteString(h, fmt.Sprintf("%s%s%s", user.String(), testProvider.String(), "fid"))
	suite.Require().NoError(err)

	// Post Contract #2
	_, err = msgSrvr.PostFile(context, &types.MsgPostFile{
		Creator:       user.String(),
		Merkle:        []byte{},
		FileSize:      1000,
		ProofInterval: 1800,
		ProofType:     0,
		MaxProofs:     3,
		Note:          "",
	})
	suite.Require().NoError(err)
	h2 := sha256.New()
	_, err = io.WriteString(h2, fmt.Sprintf("%s%s%s", user.String(), testProvider.String(), "fid"))
	suite.Require().NoError(err)

	// Storage Provider get file and create merkle for proof
	// for tc 1 and 2
	item, hashlist, err := CreateMerkleForProof(fileFromSP)
	suite.Require().NoError(err)

	// for tc 3: post proof from a different file
	item2, hashlist2, err2 := CreateMerkleForProof(randomFile)
	suite.Require().NoError(err2)

	cases := []struct {
		testName  string
		msg       types.MsgPostProof
		expErr    bool
		expErrMsg string
		postRun   func()
	}{
		{
			testName: "proof successfully verified",
			msg: types.MsgPostProof{
				Creator:  testProvider.String(),
				Item:     item,
				HashList: hashlist,
				Merkle:   merkleroot,
				Owner:    user.String(),
				Start:    0,
			},
			expErr:    false,
			expErrMsg: "",
		},
		{
			testName: "postproof for the same file again",
			msg: types.MsgPostProof{
				Creator:  testProvider.String(),
				Item:     item,
				HashList: hashlist,
				Merkle:   merkleroot,
				Owner:    user.String(),
				Start:    0,
			},
			expErr:    false,
			expErrMsg: "proof already verified",
		},
		{
			testName: "proof fail to verify",
			msg: types.MsgPostProof{
				Creator:  testProvider.String(),
				Item:     item,
				HashList: hashlist2,
				Merkle:   merkleroot,
				Owner:    user.String(),
				Start:    0,
			},
			expErr:    true,
			expErrMsg: "file chunk was not verified",
		},
		{
			testName: "nonexisting contract",
			msg: types.MsgPostProof{
				Creator:  testProvider.String(),
				Item:     item,
				HashList: hashlist,
				Merkle:   []byte{},
				Owner:    user.String(),
				Start:    0,
			},

			expErr:    true,
			expErrMsg: "contract not found",
		},
		{
			testName: "contract with invalid merkleroot",
			msg: types.MsgPostProof{
				Creator:  testProvider.String(),
				Item:     item2,
				HashList: hashlist2,
				Merkle:   []byte{},
				Owner:    user.String(),
				Start:    0,
			},
			expErr:    true,
			expErrMsg: "could not build merkle tree",
		},
	}

	for _, tcs := range cases {
		tc := tcs
		suite.Run(
			tc.testName, func() {
				res, _ := msgSrvr.PostProof(context, &tc.msg)
				if tc.expErr {
					suite.Require().Equal(tc.expErrMsg, res.ErrorMessage)
					suite.Require().Equal(false, res.Success)
				}
				if tc.postRun != nil {
					tc.postRun()
				}
			},
		)
	}
}
