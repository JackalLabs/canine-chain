package keeper_test

import (
	"encoding/json"
	"strings"

	"github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgRemoveEditors() {
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

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

	// add bob as a editor for pepe

	EditorIds := strings.Split(alice, ",")
	EditorIds = append(EditorIds, bob)

	// put pepe in home
	pepejpg, err := types.CreateFolderOrFile(alice, EditorIds, strings.Split(alice, ","), "s/home/pepe.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *pepejpg)

	pepeMerklePath := types.MerklePath("s/home/pepe.jpg")
	aliceAccountHash := types.HashThenHex(alice)
	aliceOwnerAddress := types.MakeOwnerAddress(pepeMerklePath, aliceAccountHash)

	// Let's query the file after it was set to confirm that alice and bob are editors

	fileReq := types.QueryFileRequest{
		Address:      pepeMerklePath,
		OwnerAddress: aliceOwnerAddress,
	}

	res, err := suite.queryClient.Files(suite.ctx.Context(), &fileReq)
	suite.Require().NoError(err)

	bobIsEditor, err := keeper.HasEditAccess(res.Files, bob)
	suite.Require().NoError(err)
	suite.Require().Equal(bobIsEditor, true)

	aliceIsEditor, err := keeper.HasEditAccess(res.Files, alice)
	suite.Require().NoError(err)
	suite.Require().Equal(aliceIsEditor, true)

	bobEditorAddress := keeper.MakeEditorAddress(res.Files.TrackingNumber, bob)

	cases := []struct {
		preRun    func() *types.MsgRemoveEditors
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // charlie fails to remove bob from alice's editing access
			preRun: func() *types.MsgRemoveEditors {
				return types.NewMsgRemoveEditors(
					charlie,
					bobEditorAddress,
					pepeMerklePath,
					aliceOwnerAddress,
				)
			},
			expErr:    true,
			expErrMsg: "Not permitted to remove or reset edit/view access. You are not the owner of this file",
			name:      "charlie fails to remove bob from alice's viewing permissions",
		},
		{ // alice removes bob from editing access for a non existent file
			preRun: func() *types.MsgRemoveEditors {
				return types.NewMsgRemoveEditors(
					alice,
					bobEditorAddress,
					types.MerklePath("s/home/ghost.jpg"),
					aliceOwnerAddress,
				)
			},
			expErr:    true,
			expErrMsg: "file not found",
			name:      "alice can't edit editing access for a file that doesn't exist",
		},
		{ // alice removes bob from viewing access
			preRun: func() *types.MsgRemoveEditors {
				return types.NewMsgRemoveEditors(
					alice,
					bobEditorAddress,
					pepeMerklePath,
					aliceOwnerAddress,
				)
			},
			expErr: false,
			name:   "alice removes bob from editing permissions",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.RemoveEditors(context, msg)
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

				bobIsEditor, err := keeper.HasEditAccess(res.Files, bob)
				suite.Require().NoError(err)
				suite.Require().EqualValues(bobIsEditor, false)

				aliceIsEditor, err := keeper.HasEditAccess(res.Files, alice)
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
