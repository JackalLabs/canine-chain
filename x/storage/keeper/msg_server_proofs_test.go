package keeper_test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/testutil"
	k "github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"

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

func CreateMerkleForProof(file TestFile) (string, string, error) {
	f := []byte(file.Data)
	index := 0
	var data [][]byte
	item := f

	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%d%x", index, f))
	if err != nil {
		return "", "", err
	}
	hashName := h.Sum(nil)

	data = append(data, hashName)

	tree, err := merkletree.NewUsing(data, sha3.New512(), false)
	if err != nil {
		return "", "", err
	}

	h = sha256.New()
	_, err = io.WriteString(h, fmt.Sprintf("%d%x", index, item))
	if err != nil {
		return "", "", err
	}
	ditem := h.Sum(nil)

	proof, err := tree.GenerateProof(ditem, 0)
	if err != nil {
		return "", "", err
	}

	jproof, err := json.Marshal(*proof)
	if err != nil {
		return "", "", err
	}

	e := hex.EncodeToString(tree.Root())

	k, _ := hex.DecodeString(e)

	verified, err := merkletree.VerifyProofUsing(ditem, false, proof, [][]byte{k}, sha3.New512())
	if err != nil {
		return "", "", err
	}

	if !verified {
		return "", "", types.ErrCannotVerifyProof
	}

	return fmt.Sprintf("%x", item), string(jproof), nil
}

func makeContract(file TestFile) (string, string, error) {
	f := []byte(file.Data)
	var list [][]byte

	size := len(f)

	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%d%x", 0, f))
	if err != nil {
		return "", "", err
	}
	hashName := h.Sum(nil)

	list = append(list, hashName)

	t, err := merkletree.NewUsing(list, sha3.New512(), false)
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(t.Root()), fmt.Sprintf("%d", size), nil
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

	suite.storageKeeper.SetParams(suite.ctx, types.Params{
		DepositAccount:         depoAccount,
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     8,
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

	suite.Require().Equal("11", filesize)

	_, found := keeper.GetStoragePaymentInfo(suite.ctx, user.String())
	suite.Require().Equal(true, found)
	// Post Contract
	_, err = msgSrvr.PostContract(context, &types.MsgPostContract{
		Creator:  testProvider.String(),
		Signee:   user.String(),
		Filesize: filesize,
		Fid:      "fid",
		Merkle:   merkleroot,
	})
	suite.Require().NoError(err)
	h := sha256.New()
	_, err = io.WriteString(h, fmt.Sprintf("%s%s%s", user.String(), testProvider.String(), "fid"))
	suite.Require().NoError(err)
	hashName := h.Sum(nil)
	cid1, err := k.MakeCid(hashName)
	suite.Require().NoError(err)

	// Post Contract #2
	_, err = msgSrvr.PostContract(context, &types.MsgPostContract{
		Creator:  testProvider.String(),
		Signee:   user.String(),
		Filesize: "1000",
		Fid:      "fid2",
		Merkle:   "invalid_merkleroot",
	})
	suite.Require().NoError(err)
	h2 := sha256.New()
	_, err = io.WriteString(h2, fmt.Sprintf("%s%s%s", user.String(), testProvider.String(), "fid"))
	suite.Require().NoError(err)
	hashName2 := h.Sum(nil)
	cid2, err := k.MakeCid(hashName2)
	suite.Require().NoError(err)
	// Sign Contract for active deal
	_, err = msgSrvr.SignContract(context, &types.MsgSignContract{
		Creator: user.String(),
		Cid:     cid1,
	})
	suite.Require().NoError(err)

	// Sign Contract #2 for active deal
	_, err = msgSrvr.SignContract(context, &types.MsgSignContract{
		Creator: user.String(),
		Cid:     cid2,
	})
	suite.Require().Error(err)

	// Storage Provider get file and create merkle for proof
	// for tc 1 and 2
	item, hashlist, err := CreateMerkleForProof(fileFromSP)
	suite.Require().NoError(err)

	// for tc 3: post proof from a different file
	item2, hashlist2, err2 := CreateMerkleForProof(randomFile)
	suite.Require().NoError(err2)

	cases := []struct {
		testName  string
		msg       types.MsgPostproof
		expErr    bool
		expErrMsg string
		postRun   func()
	}{
		{
			testName: "proof successfully verified",
			msg: types.MsgPostproof{
				Creator:  testProvider.String(),
				Cid:      cid1,
				Item:     item,
				Hashlist: hashlist,
			},
			expErr:    false,
			expErrMsg: "",
		},
		{
			testName: "postproof for the same file again",
			msg: types.MsgPostproof{
				Creator:  testProvider.String(),
				Cid:      cid1,
				Item:     item,
				Hashlist: hashlist,
			},
			expErr:    true,
			expErrMsg: "proof already verified",
			postRun: func() {
				// Set Proofverified back to false
				contract, _ := keeper.GetActiveDeals(suite.ctx, CID)
				contract.ProofVerified = false
				keeper.SetActiveDeals(suite.ctx, contract)
			},
		},
		{
			testName: "proof fail to verify",
			msg: types.MsgPostproof{
				Creator:  testProvider.String(),
				Cid:      cid1,
				Item:     item2,
				Hashlist: hashlist2,
			},
			expErr:    true,
			expErrMsg: "file chunk was not verified",
		},
		{
			testName: "nonexisting contract",
			msg: types.MsgPostproof{
				Creator:  testProvider.String(),
				Cid:      "fakecontractid",
				Item:     item,
				Hashlist: hashlist,
			},
			expErr:    true,
			expErrMsg: "contract not found",
		},
		{
			testName: "contract with invalid merkleroot",
			msg: types.MsgPostproof{
				Creator:  testProvider.String(),
				Cid:      cid2,
				Item:     item2,
				Hashlist: hashlist2,
			},
			expErr:    true,
			expErrMsg: "could not build merkle tree",
		},
	}

	for _, tc := range cases {
		suite.Run(
			tc.testName, func() {
				res, err := msgSrvr.Postproof(context, &tc.msg)
				if tc.expErr {
					suite.Require().Equal(false, res.Success)
				} else {
					contract, _ := keeper.GetActiveDeals(suite.ctx, cid1)
					suite.Require().Equal(true, contract.ProofVerified)
					suite.Require().NoError(err)
				}
				if tc.postRun != nil {
					tc.postRun()
				}
			},
		)
	}
}
