package keeper_test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	merkletree "github.com/wealdtech/go-merkletree"
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
	CID  = "6ef1cf960c0b1e257049645ff13ed890c2d4ef69c62165bf4e090ec480770d67"
	CID2 = "a257645c433288e20c23d4fc8ce823c66558e5398c98b140d7ccfb39e8a3c70e"
)

func CreateMerkleForProof(ctx sdk.Context, file TestFile) (string, string, error) {
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

	tree, err := merkletree.New(data)
	if err != nil {
		return "", "", err
	}

	h = sha256.New()
	_, err = io.WriteString(h, fmt.Sprintf("%d%x", index, item))
	if err != nil {
		return "", "", err
	}
	ditem := h.Sum(nil)

	proof, err := tree.GenerateProof(ditem)
	if err != nil {
		return "", "", err
	}

	jproof, err := json.Marshal(*proof)
	if err != nil {
		return "", "", err
	}

	e := hex.EncodeToString(tree.Root())

	k, _ := hex.DecodeString(e)

	verified, err := merkletree.VerifyProof(ditem, proof, k)
	if err != nil {
		ctx.Logger().Error("%v\n", err)
	}

	if !verified {
		ctx.Logger().Error("%s\n", "Cannot verify")
	}

	return fmt.Sprintf("%x", item), string(jproof), nil
}

func makeContract(ctx sdk.Context, file TestFile) (string, string, error) {
	f := []byte(file.Data)
	size := 0
	var list [][]byte

	size += len(f)

	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%d%x", 0, f))
	if err != nil {
		return "", "", err
	}
	hashName := h.Sum(nil)

	list = append(list, hashName)

	t, err := merkletree.New(list)
	if err != nil {
		ctx.Logger().Debug("%v\n", err)
	}

	return hex.EncodeToString(t.Root()), fmt.Sprintf("%d", size), nil
}

func (suite *KeeperTestSuite) TestPostProof() {
	suite.SetupSuite()

	msgSrvr, keeper, context := setupMsgServer(suite)

	// Create user account
	user, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	// Create provider account
	testProvider, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	// Init Provider
	_, err = msgSrvr.InitProvider(context, &types.MsgInitProvider{
		Creator:    testProvider.String(),
		Ip:         "198.0.0.1",
		Totalspace: "1_000_000",
	})
	suite.Require().NoError(err)

	// Storage Provider receives file and make merkleroot for contract
	merkleroot, filesize, err := makeContract(suite.ctx, originalFile)
	suite.Require().NoError(err)

	// Post Contract
	_, err = msgSrvr.PostContract(context, &types.MsgPostContract{
		Creator:  testProvider.String(),
		Signee:   user.String(),
		Duration: "1",
		Filesize: filesize,
		Fid:      "fid",
		Merkle:   merkleroot,
	})
	suite.Require().NoError(err)

	// Post Contract #2
	_, err = msgSrvr.PostContract(context, &types.MsgPostContract{
		Creator:  testProvider.String(),
		Signee:   user.String(),
		Duration: "10",
		Filesize: "1000",
		Fid:      "fid2",
		Merkle:   "invalid_merkleroot",
	})
	suite.Require().NoError(err)

	// Sign Contract for active deal
	_, err = msgSrvr.SignContract(context, &types.MsgSignContract{
		Creator: user.String(),
		Cid:     CID,
	})
	suite.Require().NoError(err)

	// Sign Contract #2 for active deal
	_, err = msgSrvr.SignContract(context, &types.MsgSignContract{
		Creator: user.String(),
		Cid:     CID2,
	})
	suite.Require().NoError(err)

	// Storage Provider get file and create merkle for proof
	// for tc 1 and 2
	item, hashlist, err := CreateMerkleForProof(suite.ctx, fileFromSP)
	suite.Require().NoError(err)

	// for tc 3: post proof from a different file
	item2, hashlist2, err2 := CreateMerkleForProof(suite.ctx, randomFile)
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
				Cid:      CID,
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
				Cid:      CID,
				Item:     item,
				Hashlist: hashlist,
			},
			expErr:    true,
			expErrMsg: "proof already verified",
			postRun: func() {
				// Set Proofverified back to false
				contract, _ := keeper.GetActiveDeals(suite.ctx, CID)
				contract.Proofverified = "false"
				keeper.SetActiveDeals(suite.ctx, contract)
			},
		},
		{
			testName: "proof fail to verify",
			msg: types.MsgPostproof{
				Creator:  testProvider.String(),
				Cid:      CID,
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
				Cid:      CID2,
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
				_, err := msgSrvr.Postproof(context, &tc.msg)
				if tc.expErr {
					suite.Require().EqualError(err, tc.expErrMsg)
				} else {
					contract, _ := keeper.GetActiveDeals(suite.ctx, CID)
					suite.Require().Equal("true", contract.Proofverified)
					suite.Require().NoError(err)
				}
				if tc.postRun != nil {
					tc.postRun()
				}
			},
		)
	}
}
