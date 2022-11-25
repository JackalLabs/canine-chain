package keeper_test

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgDeleteFile() {
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

	// set home folder for alice
	aliceHomeFolder, err := types.CreateFolderOrFile(alice.String(), "s/home/")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceHomeFolder)

	// put pepe in home
	pepejpg, err := types.CreateFolderOrFile(alice.String(), "s/home/pepe.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *pepejpg)

	// put hasbullah in home
	hasbullahjpg, err := types.CreateFolderOrFile(alice.String(), "s/home/hasbullah.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *hasbullahjpg)

	aliceAccountHash := types.HashThenHex(alice.String())
	pepeMerklePath := types.MerklePath("s/home/pepe.jpg")
	hasbullahMerklePath := types.MerklePath("s/home/hasbullah.jpg")

	cases := []struct {
		preRun    func() *types.MsgDeleteFile
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // alice deletes pepe
			preRun: func() *types.MsgDeleteFile {
				return types.NewMsgDeleteFile(
					alice.String(),
					pepeMerklePath,
					aliceAccountHash,
				)
			},
			expErr: false,
			name:   "alice successfully deletes",
		},
		{ // alice tries to delete a file that doesn't exist
			preRun: func() *types.MsgDeleteFile {
				return types.NewMsgDeleteFile(
					alice.String(),
					types.MerklePath("s/home/ghost.png"),
					aliceAccountHash,
				)
			},
			expErr:    true,
			expErrMsg: "file not found",
			name:      "can't delete ghosts",
		},
		{ // bob tries to delete alice's hasbullah
			preRun: func() *types.MsgDeleteFile {
				return types.NewMsgDeleteFile(
					bob.String(),
					hasbullahMerklePath,
					aliceAccountHash,
				)
			},
			expErr:    true,
			expErrMsg: "You are not authorized to delete this file",
			name:      "bob cannot delete alice's hasbullah",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.DeleteFile(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgDeleteFileResponse{}, *res)

			}
		})
	}
}
