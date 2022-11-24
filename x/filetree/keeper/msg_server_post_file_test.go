package keeper_test

import (
	"crypto/sha256"
	"fmt"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgPostFile() {
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

	alice, err := sdkTypes.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	bob, err := sdkTypes.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	// set root folder for alice
	aliceRootFolder, err := types.CreateRootFolder(alice.String())
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceRootFolder)

	// set root folder for bob
	bobRootFolder, err := types.CreateRootFolder(bob.String())
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *bobRootFolder)

	// arguments for home folder
	parentHash, childHash := types.MerkleHelper("s/home/")
	aliceHomeTrackingNumber := uuid.NewString()
	bobHomeTrackingNumber := uuid.NewString()

	// hash alice account address
	H := sha256.New()
	H.Write([]byte(alice.String()))
	hash := H.Sum(nil)
	aliceAccountHash := fmt.Sprintf("%x", hash)

	// hash bob account address
	H1 := sha256.New()
	H1.Write([]byte(bob.String()))
	hash1 := H1.Sum(nil)
	bobAccountHash := fmt.Sprintf("%x", hash1)

	// arguments for non existent root
	ghostParentHash, ghostChildHash := types.MerkleHelper("g/home/")

	cases := []struct {
		preRun    func() *types.MsgPostFile
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // alice makes her home folder
			preRun: func() *types.MsgPostFile {
				return types.NewMsgPostFile(
					alice.String(),
					aliceAccountHash,
					parentHash,
					childHash,
					"contents: FID goes here",
					"viewers",
					"editors",
					aliceHomeTrackingNumber,
				)
			},
			expErr: false,
			name:   "alice successfully makes her home folder",
		},
		{ // alice fails to make a home folder in a non existent root
			preRun: func() *types.MsgPostFile {
				return types.NewMsgPostFile(
					alice.String(),
					aliceAccountHash,
					ghostParentHash,
					ghostChildHash,
					"contents: FID goes here",
					"viewers",
					"editors",
					"none",
				)
			},
			expErr:    true,
			name:      "post file fail",
			expErrMsg: "Parent folder does not exist",
		},
		{ // bob fails to make a home folder inside of alice's root folder, i.e., inside of alice's account
			preRun: func() *types.MsgPostFile {
				return types.NewMsgPostFile(
					bob.String(),
					aliceAccountHash,
					parentHash,
					childHash,
					"contents: FID goes here",
					"viewers",
					"editors",
					"none",
				)
			},
			expErr:    true,
			name:      "fail to write to root in other account",
			expErrMsg: "You are not permitted to write to this folder",
		},
		{ // bob can make a home folder inside of his root in his account
			preRun: func() *types.MsgPostFile {
				return types.NewMsgPostFile(
					bob.String(),
					bobAccountHash,
					parentHash,
					childHash,
					"contents: FID goes here",
					"viewers",
					"editors",
					bobHomeTrackingNumber,
				)
			},
			expErr: false,
			name:   "bob makes his root folder",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.PostFile(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				// If creator has permissions in parent folder, the full merkle path of the incoming child will be created
				// and sent back as a response
				fullMerklePath := types.AddToMerkle(msg.HashParent, msg.HashChild)
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgPostFileResponse{Path: fullMerklePath}, *res)

			}
		})
	}
}
