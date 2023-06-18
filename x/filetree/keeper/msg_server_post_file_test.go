package keeper_test

import (
	"strings"

	"github.com/google/uuid"
	"github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgPostFile() {
	suite.SetupSuite()
	msgSrvr, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	alice := testAddresses[0]
	bob := testAddresses[1]

	// set root folder for alice
	aliceRootFolder, err := types.CreateRootFolder(alice)
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceRootFolder)

	// set root folder for bob
	bobRootFolder, err := types.CreateRootFolder(bob)
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *bobRootFolder)

	// arguments for home folder
	// alice and bob can both have merkelPath("s/home/") as the address of their home folder. File structs are indexed by both path address
	// and owner address so collisions are prevented
	parentHash, childHash := types.MerkleHelper("s/home/")

	aliceHomeTrackingNumber := uuid.NewString()
	aliceEditorAccess, err := types.MakeEditorAccessMap(aliceHomeTrackingNumber, strings.Split(alice, ","), "place holder key")
	suite.Require().NoError(err)

	bobHomeTrackingNumber := uuid.NewString()
	bobEditorAccess, err := types.MakeEditorAccessMap(bobHomeTrackingNumber, strings.Split(bob, ","), "place holder key")
	suite.Require().NoError(err)

	// hash alice account address
	aliceAccountHash := types.HashThenHex(alice)

	// hash bob account address
	bobAccountHash := types.HashThenHex(bob)

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
					alice,
					aliceAccountHash,
					parentHash,
					childHash,
					"contents: FID goes here",
					"viewers",
					string(aliceEditorAccess),
					aliceHomeTrackingNumber,
				)
			},
			expErr: false,
			name:   "alice successfully makes her home folder",
		},
		{ // alice fails to make a home folder in a non existent root
			preRun: func() *types.MsgPostFile {
				return types.NewMsgPostFile(
					alice,
					aliceAccountHash,
					ghostParentHash,
					ghostChildHash,
					"contents: FID goes here",
					"viewers",
					string(aliceEditorAccess),
					"none",
				)
			},
			expErr:    true,
			name:      "post file fail",
			expErrMsg: "Parent folder does not exist",
		},
		{ // alice makes pepe.jpg inside of her home folder
			preRun: func() *types.MsgPostFile {
				pepeTrackingNumber := uuid.NewString()
				pepeEditorAccess, err := types.MakeEditorAccessMap(pepeTrackingNumber, strings.Split(alice, ","), "place holder key")
				suite.Require().NoError(err)

				msg, err := types.CreateMsgPostFile(alice, "s/home/pepe.jpg", pepeEditorAccess, pepeTrackingNumber)
				suite.Require().NoError(err)
				return msg
			},
			expErr: false,
			name:   "alice successfully puts pepe in home",
		},
		{ // alice can't put pepe.jpg inside of s/videos/ because this folder doesn't exist
			preRun: func() *types.MsgPostFile {
				pepeTrackingNumber := uuid.NewString()
				pepeEditorAccess, err := types.MakeEditorAccessMap(pepeTrackingNumber, strings.Split(alice, ","), "place holder key")
				suite.Require().NoError(err)

				msg, err := types.CreateMsgPostFile(alice, "s/videos/pepe.jpg", pepeEditorAccess, pepeTrackingNumber)
				suite.Require().NoError(err)
				return msg
			},
			expErr:    true,
			expErrMsg: "Parent folder does not exist",
			name:      "alice fails to put pepe in videos",
		},
		{ // bob fails to make a home folder inside of alice's root folder, i.e., inside of alice's account
			preRun: func() *types.MsgPostFile {
				return types.NewMsgPostFile(
					bob,
					aliceAccountHash,
					parentHash,
					childHash,
					"contents: FID goes here",
					"viewers",
					string(bobEditorAccess),
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
					bob,
					bobAccountHash,
					parentHash,
					childHash,
					"contents: FID goes here",
					"viewers",
					string(bobEditorAccess),
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
