package keeper_test

import (
	"encoding/json"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgResetEditors() {
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

	alice, err := sdkTypes.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	bob, err := sdkTypes.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	charlie, err := sdkTypes.AccAddressFromBech32("cosmos1xetrp5dwjplsn4lev5r2cu8en5qsq824vza9nu")
	suite.Require().NoError(err)

	// set root folder for alice
	aliceRootFolder, err := types.CreateRootFolder(alice.String())
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceRootFolder)

	// set home folder for alice
	aliceHomeFolder, err := types.CreateFolderOrFile(alice.String(), strings.Split(alice.String(), ","), strings.Split(alice.String(), ","), "s/home/")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceHomeFolder)

	// add bob and charlie as editors for pepe

	EditorIds := strings.Split(alice.String(), ",")
	EditorIds = append(EditorIds, bob.String(), charlie.String())

	// put pepe in home
	pepejpg, err := types.CreateFolderOrFile(alice.String(), EditorIds, strings.Split(alice.String(), ","), "s/home/pepe.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *pepejpg)

	pepeMerklePath := types.MerklePath("s/home/pepe.jpg")
	aliceAccountHash := types.HashThenHex(alice.String())
	aliceOwnerAddress := types.MakeOwnerAddress(pepeMerklePath, aliceAccountHash)

	// Let's query the file after it was set to confirm that alice and bob are editors

	fileReq := types.QueryFileRequest{
		Address:      pepeMerklePath,
		OwnerAddress: aliceOwnerAddress,
	}

	res, err := suite.queryClient.Files(suite.ctx.Context(), &fileReq)
	suite.Require().NoError(err)

	bobIsEditor, err := keeper.HasEditAccess(res.Files, bob.String())
	suite.Require().NoError(err)
	suite.Require().Equal(bobIsEditor, true)

	aliceIsEditor, err := keeper.HasEditAccess(res.Files, alice.String())
	suite.Require().NoError(err)
	suite.Require().Equal(aliceIsEditor, true)

	charlieIsEditor, err := keeper.HasEditAccess(res.Files, charlie.String())
	suite.Require().NoError(err)
	suite.Require().Equal(charlieIsEditor, true)

	peacc := res.Files.EditAccess
	jeacc := make(map[string]string)

	error := json.Unmarshal([]byte(peacc), &jeacc)
	suite.Require().NoError(error)
	suite.Require().EqualValues(len(jeacc), 3)

	cases := []struct {
		preRun    func() *types.MsgResetEditors
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // charlie fails to reset alice's editing access
			preRun: func() *types.MsgResetEditors {
				return types.NewMsgResetEditors(
					charlie.String(),
					pepeMerklePath,
					aliceOwnerAddress,
				)
			},
			expErr:    true,
			expErrMsg: "Not permitted to remove or reset edit/view access. You are not the owner of this file",
			name:      "charlie fails to reset alice's editing permissions",
		},
		{ // alice fails to reset edit permissions for a non existent file
			preRun: func() *types.MsgResetEditors {
				return types.NewMsgResetEditors(
					alice.String(),
					types.MerklePath("s/home/ghost.jpg"),
					aliceOwnerAddress,
				)
			},
			expErr:    true,
			expErrMsg: "file not found",
			name:      "alice can't reset editing access for a file that doesn't exist",
		},
		{ // alice resets editing access
			preRun: func() *types.MsgResetEditors {
				return types.NewMsgResetEditors(
					alice.String(),
					pepeMerklePath,
					aliceOwnerAddress,
				)
			},
			expErr: false,
			name:   "alice resets editing permissions",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.ResetEditors(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {

				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgRemoveEditorsResponse{}, *res)
				// Let's confirm that bob is no longer an editor

				fileReq := types.QueryFileRequest{
					Address:      pepeMerklePath,
					OwnerAddress: aliceOwnerAddress,
				}
				res, err := suite.queryClient.Files(suite.ctx.Context(), &fileReq)
				suite.Require().NoError(err)

				bobIsEditor, err := keeper.HasEditAccess(res.Files, bob.String())
				suite.Require().NoError(err)
				suite.Require().EqualValues(bobIsEditor, false)

				charlieIsEditor, err := keeper.HasEditAccess(res.Files, charlie.String())
				suite.Require().NoError(err)
				suite.Require().EqualValues(charlieIsEditor, false)

				aliceIsEditor, err := keeper.HasEditAccess(res.Files, alice.String())
				suite.Require().NoError(err)
				suite.Require().EqualValues(aliceIsEditor, true)

				peacc := res.Files.EditAccess
				jeacc := make(map[string]string)

				error := json.Unmarshal([]byte(peacc), &jeacc)
				suite.Require().NoError(error)
				suite.Require().EqualValues(len(jeacc), 1)

			}
		})
	}
}
