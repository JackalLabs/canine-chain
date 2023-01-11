package keeper_test

import (
	"fmt"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgAddEditors() {
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

	aliceViewerID := strings.Split(alice.String(), ",")
	aliceEditorID := aliceViewerID

	// set home folder for alice
	aliceHomeFolder, err := types.CreateFolderOrFile(alice.String(), aliceEditorID, aliceViewerID, "s/home/")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceHomeFolder)

	aliceAccountHash := types.HashThenHex(alice.String())
	aliceHomeMerklePath := types.MerklePath("s/home/")

	ownerAddress := types.MakeOwnerAddress(aliceHomeMerklePath, aliceAccountHash)
	bobEditorAddress := keeper.MakeEditorAddress(aliceHomeFolder.TrackingNumber, bob.String())

	cases := []struct {
		preRun    func() *types.MsgAddEditors
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // alice adds an editor
			preRun: func() *types.MsgAddEditors {
				return types.NewMsgAddEditors(
					alice.String(),
					bobEditorAddress,
					fmt.Sprintf("%x", "encryptedHomeFolderAESKeyAndIV"),
					aliceHomeMerklePath,
					ownerAddress,
				)
			},
			expErr: false,
			name:   "alice adds an editor",
		},
		{ // alice cannot add an editor to a non existent file
			preRun: func() *types.MsgAddEditors {
				return types.NewMsgAddEditors(
					alice.String(),
					bobEditorAddress,
					fmt.Sprintf("%x", "encryptedAESKeyAndIV"),
					types.MerklePath("s/DNE/"),
					ownerAddress,
				)
			},
			expErr:    true,
			name:      "alice fails to add editor",
			expErrMsg: "file not found",
		},
		{ // bob can't add himself as an editor to alice's home folder
			preRun: func() *types.MsgAddEditors {
				return types.NewMsgAddEditors(
					bob.String(),
					bobEditorAddress,
					fmt.Sprintf("%x", "encryptedHomeFolderAESKeyAndIV"),
					aliceHomeMerklePath,
					ownerAddress,
				)
			},
			expErr:    true,
			name:      "bob cannot add himself as editor",
			expErrMsg: "Unathorized. Only the owner can add an editor.",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.AddEditors(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgAddEditorsResponse{}, *res)

				fileReq := types.QueryFileRequest{
					Address:      aliceHomeMerklePath,
					OwnerAddress: ownerAddress,
				}

				res, err := suite.queryClient.Files(suite.ctx.Context(), &fileReq)
				suite.Require().NoError(err)

				validEditor, err := keeper.HasEditAccess(res.Files, bob.String())
				suite.Require().NoError(err)
				suite.Require().Equal(validEditor, true)

			}
		})
	}
}
