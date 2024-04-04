package keeper_test

import (
	"encoding/json"
	"strings"

	"github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgResetEditors() {
	suite.SetupSuite()
	msgSrvr, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 3)
	suite.Require().NoError(err)

	alice := testAddresses[0]
	bob := testAddresses[1]
	charlie := testAddresses[2]

	// set root folder for alice
	aliceRootFolder, err := types.CreateRootFolder(alice)
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceRootFolder)

	// set home folder for alice
	aliceHomeFolder, err := types.CreateFolderOrFile(alice, strings.Split(alice, ","), strings.Split(alice, ","), "s/home/")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceHomeFolder)

	// add bob and charlie as editors for pepe

	EditorIDs := strings.Split(alice, ",")
	EditorIDs = append(EditorIDs, bob, charlie)

	// put pepe in home
	pepejpg, err := types.CreateFolderOrFile(alice, EditorIDs, strings.Split(alice, ","), "s/home/pepe.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *pepejpg)

	pepeMerklePath := types.MerklePath("s/home/pepe.jpg")
	aliceAccountHash := types.HashThenHex(alice)
	aliceOwnerAddress := types.MakeOwnerAddress(pepeMerklePath, aliceAccountHash)

	// Let's query the file after it was set to confirm that alice and bob are editors

	fileReq := types.QueryFile{
		Address:      pepeMerklePath,
		OwnerAddress: aliceOwnerAddress,
	}

	res, err := suite.queryClient.File(suite.ctx.Context(), &fileReq)
	suite.Require().NoError(err)

	bobIsEditor, err := keeper.HasEditAccess(res.File, bob)
	suite.Require().NoError(err)
	suite.Require().Equal(bobIsEditor, true)

	aliceIsEditor, err := keeper.HasEditAccess(res.File, alice)
	suite.Require().NoError(err)
	suite.Require().Equal(aliceIsEditor, true)

	charlieIsEditor, err := keeper.HasEditAccess(res.File, charlie)
	suite.Require().NoError(err)
	suite.Require().Equal(charlieIsEditor, true)

	peacc := res.File.EditAccess
	jeacc := make(map[string]string)

	err = json.Unmarshal([]byte(peacc), &jeacc)
	suite.Require().NoError(err)
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
					charlie,
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
					alice,
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
					alice,
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

				fileReq := types.QueryFile{
					Address:      pepeMerklePath,
					OwnerAddress: aliceOwnerAddress,
				}
				res, err := suite.queryClient.File(suite.ctx.Context(), &fileReq)
				suite.Require().NoError(err)

				bobIsEditor, err := keeper.HasEditAccess(res.File, bob)
				suite.Require().NoError(err)
				suite.Require().EqualValues(bobIsEditor, false)

				charlieIsEditor, err := keeper.HasEditAccess(res.File, charlie)
				suite.Require().NoError(err)
				suite.Require().EqualValues(charlieIsEditor, false)

				aliceIsEditor, err := keeper.HasEditAccess(res.File, alice)
				suite.Require().NoError(err)
				suite.Require().EqualValues(aliceIsEditor, true)

				peacc := res.File.EditAccess
				jeacc := make(map[string]string)

				err = json.Unmarshal([]byte(peacc), &jeacc)
				suite.Require().NoError(err)
				suite.Require().EqualValues(len(jeacc), 1)

			}
		})
	}
}
