package keeper_test

import (
	"encoding/json"
	"fmt"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgRemoveViewers() {
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

	alice, err := sdkTypes.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	bob, err := sdkTypes.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	charlie, err := sdkTypes.AccAddressFromBech32("cosmos1xetrp5dwjplsn4lev5r2cu8en5qsq824vza9nu")
	suite.Require().NoError(err)

	// Let it be that bob has posted a public key after signing
	// with his keyring backend using the CLI.

	bobPrivateKey, err := types.MakePrivateKey("bob")
	suite.Require().NoError(err)

	bobPubKey := bobPrivateKey.PublicKey.Bytes(false) // to hex
	pubKeyStruct := types.Pubkey{
		Address: bob.String(),
		Key:     fmt.Sprintf("%x", bobPubKey),
	}
	suite.filetreeKeeper.SetPubkey(suite.ctx, pubKeyStruct)

	// set root folder for alice
	aliceRootFolder, err := types.CreateRootFolder(alice.String())
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceRootFolder)

	// set home folder for alice
	aliceHomeFolder, err := types.CreateFolderOrFile(alice.String(), strings.Split(alice.String(), ","), strings.Split(alice.String(), ","), "s/home/")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceHomeFolder)

	// add bob as a viewer for pepe

	viewerIds := strings.Split(alice.String(), ",")
	viewerIds = append(viewerIds, bob.String())

	// put pepe in home
	pepejpg, err := types.CreateFolderOrFile(alice.String(), strings.Split(alice.String(), ","), viewerIds, "s/home/pepe.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *pepejpg)

	pepeMerklePath := types.MerklePath("s/home/pepe.jpg")
	aliceAccountHash := types.HashThenHex(alice.String())
	aliceOwnerAddress := types.MakeOwnerAddress(pepeMerklePath, aliceAccountHash)

	// Let's query the file after it was set to confirm that alice and bob are viewers

	fileReq := types.QueryGetFilesRequest{
		Address:      pepeMerklePath,
		OwnerAddress: aliceOwnerAddress,
	}

	res, err := suite.queryClient.Files(suite.ctx.Context(), &fileReq)
	suite.Require().NoError(err)

	bobIsViewer, err := keeper.HasViewingAccess(res.Files, alice.String())
	suite.Require().NoError(err)
	suite.Require().Equal(bobIsViewer, true)

	aliceIsViewer, err := keeper.HasViewingAccess(res.Files, alice.String())
	suite.Require().NoError(err)
	suite.Require().Equal(aliceIsViewer, true)

	bobViewerAddress := keeper.MakeViewerAddress(res.Files.TrackingNumber, bob.String())

	cases := []struct {
		preRun    func() *types.MsgRemoveViewers
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // charlie fails to remove bob from alice's viewing access
			preRun: func() *types.MsgRemoveViewers {
				return types.NewMsgRemoveViewers(
					charlie.String(),
					bobViewerAddress,
					pepeMerklePath,
					aliceOwnerAddress,
				)
			},
			expErr:    true,
			expErrMsg: "Not permitted to remove or reset edit/view access. You are not the owner of this file",
			name:      "charlie fails to remove bob from alice's viewing permissions",
		},
		{ // alice removes bob from viewing access
			preRun: func() *types.MsgRemoveViewers {
				return types.NewMsgRemoveViewers(
					alice.String(),
					bobViewerAddress,
					types.MerklePath("s/home/ghost.jpg"),
					aliceOwnerAddress,
				)
			},
			expErr:    true,
			expErrMsg: "file not found",
			name:      "alice can't edit viewing access for a file that doesn't exist",
		},
		{ // alice removes bob from viewing access
			preRun: func() *types.MsgRemoveViewers {
				return types.NewMsgRemoveViewers(
					alice.String(),
					bobViewerAddress,
					pepeMerklePath,
					aliceOwnerAddress,
				)
			},
			expErr: false,
			name:   "alice removes bob from viewing permissions",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.RemoveViewers(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {

				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgRemoveViewersResponse{}, *res)
				// Let's confirm that bob is no longer a viewer

				fileReq := types.QueryGetFilesRequest{
					Address:      pepeMerklePath,
					OwnerAddress: aliceOwnerAddress,
				}
				res, err := suite.queryClient.Files(suite.ctx.Context(), &fileReq)
				suite.Require().NoError(err)

				bobIsViewer, err := keeper.HasViewingAccess(res.Files, bob.String())
				suite.Require().NoError(err)
				suite.Require().EqualValues(bobIsViewer, false)

				aliceIsViewer, err := keeper.HasViewingAccess(res.Files, alice.String())
				suite.Require().NoError(err)
				suite.Require().EqualValues(aliceIsViewer, true)

				pvacc := res.Files.ViewingAccess
				jvacc := make(map[string]string)

				error := json.Unmarshal([]byte(pvacc), &jvacc)
				suite.Require().NoError(error)
				suite.Require().EqualValues(len(jvacc), 1)

			}
		})
	}
}
