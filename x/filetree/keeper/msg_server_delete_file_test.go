package keeper_test

import (
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
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

	editorIds := strings.Split(alice.String(), ",")
	editorIds = append(editorIds, bob.String())

	// set home folder for alice and add bob as an editor
	aliceHomeFolder, err := types.CreateFolderOrFile(alice.String(), editorIds, "s/home/")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceHomeFolder)

	// put pepe in home
	pepejpg, err := types.CreateFolderOrFile(alice.String(), strings.Split(alice.String(), ","), "s/home/pepe.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *pepejpg)

	// put hasbullah in home
	hasbullahjpg, err := types.CreateFolderOrFile(alice.String(), strings.Split(alice.String(), ","), "s/home/hasbullah.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *hasbullahjpg)

	aliceAccountHash := types.HashThenHex(alice.String())
	pepeMerklePath := types.MerklePath("s/home/pepe.jpg")
	hasbullahMerklePath := types.MerklePath("s/home/hasbullah.jpg")

	// Let's confirm that bob has edit access before moving on

	fileReq := types.QueryGetFilesRequest{
		Address:      types.MerklePath("s/home/"),
		OwnerAddress: types.MakeOwnerAddress(types.MerklePath("s/home/"), aliceAccountHash),
	}

	res, err := suite.queryClient.Files(suite.ctx.Context(), &fileReq)
	suite.Require().NoError(err)
	suite.Require().Equal(res.Files, *aliceHomeFolder)

	validEditor, err := keeper.HasEditAccess(res.Files, bob.String())
	suite.Require().NoError(err)
	suite.Require().Equal(validEditor, true)

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
			name:   "alice successfully deletes pepe",
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
		{ // bob tries to delete alice's hasbullah, but fails to do so even though he is an editor in alice's home folder.
			// This confirms that only the owner of a file can delete it.
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
		{ // alice deletes s/home/
			preRun: func() *types.MsgDeleteFile {
				return types.NewMsgDeleteFile(
					alice.String(),
					types.MerklePath("s/home/"),
					aliceAccountHash,
				)
			},
			expErr: false,
			name:   "alice successfully deletes home folder",
		},
		{ // Confirm alice's s/home/ has been deleted
			preRun: func() *types.MsgDeleteFile {
				return types.NewMsgDeleteFile(
					alice.String(),
					types.MerklePath("s/home/"),
					aliceAccountHash,
				)
			},
			expErr:    true,
			expErrMsg: "file not found",
			name:      "alice home already deleted",
		},
		{ // Even though s/home/ has been deleted, hasbullah is still reachable and can be deleted. This design choice is intentional
			preRun: func() *types.MsgDeleteFile {
				return types.NewMsgDeleteFile(
					alice.String(),
					hasbullahMerklePath,
					aliceAccountHash,
				)
			},
			expErr: false,
			name:   "alice deletes hasbullah",
		},
		{ // Confirm hasbullah has been deleted
			preRun: func() *types.MsgDeleteFile {
				return types.NewMsgDeleteFile(
					alice.String(),
					hasbullahMerklePath,
					aliceAccountHash,
				)
			},
			expErr:    true,
			expErrMsg: "file not found",
			name:      "alice can not delete hasbullah twice",
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
