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

const CID = "6ef1cf960c0b1e257049645ff13ed890c2d4ef69c62165bf4e090ec480770d67"

// filename := []byte(fmt.Sprintf("%s%s", "FILE-", CID))

func CreateMerkleForProof() (string, string, error) {
	// example file
	f := []byte("jackal maxi")
	index := 0
	var data [][]byte
	var item []byte = f

	h := sha256.New()
	io.WriteString(h, fmt.Sprintf("%d%x", index, f))
	hashName := h.Sum(nil)

	data = append(data, hashName)

	tree, err := merkletree.New(data)
	if err != nil {
		fmt.Println(err)
	}

	h = sha256.New()
	io.WriteString(h, fmt.Sprintf("%d%x", index, item))
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
		fmt.Printf("%v\n", err)
	}

	if !verified {
		fmt.Printf("%s\n", "Cannot verify")
	}

	return fmt.Sprintf("%x", item), string(jproof), nil
}

func makeContract() (string, string) {
	// example file
	f := []byte("jackal maxi")
	size := 0
	var list [][]byte

	size = size + len(f)

	h := sha256.New()
	io.WriteString(h, fmt.Sprintf("%d%x", 0, f))
	hashName := h.Sum(nil)

	list = append(list, hashName)

	t, err := merkletree.New(list)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	return hex.EncodeToString(t.Root()), fmt.Sprintf("%d", size)
}

func (suite *KeeperTestSuite) TestPostProof() {
	suite.SetupSuite()

	// Create user account
	user, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	// Create provider account
	testProvider, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	msgSrvr, keeper, context := setupMsgServer(suite)

	// Init Provider
	_, err = msgSrvr.InitProvider(context, &types.MsgInitProvider{
		Creator:    testProvider.String(),
		Ip:         "198.0.0.1",
		Totalspace: "1_000_000",
	})
	if err != nil {
		fmt.Println(err)
	}

	// run makeContract
	merkleroot, filesize := makeContract()
	// fmt.Println(merkleroot)
	// fmt.Println(filesize)

	// Post Contract
	_, err = msgSrvr.PostContract(context, &types.MsgPostContract{
		Creator:  testProvider.String(),
		Signee:   user.String(),
		Duration: "1",
		Filesize: filesize,
		Fid:      "fid",
		Merkle:   merkleroot,
	})
	if err != nil {
		fmt.Println(err)
	}

	// Sign Contract for active deal
	_, err = msgSrvr.SignContract(context, &types.MsgSignContract{
		Creator: user.String(),
		Cid:     CID,
	})
	if err != nil {
		fmt.Println(err)
	}

	item, hashlist, err := CreateMerkleForProof()
	if err != nil {
		fmt.Println(err)
	}

	cases := []struct {
		testName  string
		msg       types.MsgPostproof
		expErr    bool
		expErrMsg string
	}{
		{
			testName: "create 1 proof",
			msg: types.MsgPostproof{
				Creator:  testProvider.String(),
				Cid:      CID,
				Item:     item,
				Hashlist: hashlist,
			},
			expErr:    false,
			expErrMsg: "",
		},
	}

	for _, tc := range cases {
		suite.Run(
			tc.testName, func() {
				_, err := msgSrvr.Postproof(context, &tc.msg)
				fmt.Println(err)
			},
		)
		fmt.Println(keeper.GetAllActiveDeals(suite.ctx)[0].Proofverified)
	}
}
