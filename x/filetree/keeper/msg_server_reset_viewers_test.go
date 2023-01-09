package keeper_test

import (
	"encoding/json"
	"fmt"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgResetViewers() {
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

	// add bob and charlie as viewers for pepe

	viewerIds := strings.Split(alice.String(), ",")
	viewerIds = append(viewerIds, bob.String(), charlie.String())

	// put pepe in home
	pepejpg, err := types.CreateFolderOrFile(alice.String(), strings.Split(alice.String(), ","), viewerIds, "s/home/pepe.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *pepejpg)

	pepeMerklePath := types.MerklePath("s/home/pepe.jpg")
	aliceAccountHash := types.HashThenHex(alice.String())
	aliceOwnerAddress := types.MakeOwnerAddress(pepeMerklePath, aliceAccountHash)

	// Let's query the file after it was set to confirm that alice, bob, and charlie are viewers

	fileReq := types.QueryFileRequest{
		Address:      pepeMerklePath,
		OwnerAddress: aliceOwnerAddress,
	}

	res, err := suite.queryClient.Files(suite.ctx.Context(), &fileReq)
	suite.Require().NoError(err)

	bobIsViewer, err := keeper.HasViewingAccess(res.Files, bob.String())
	suite.Require().NoError(err)
	suite.Require().Equal(bobIsViewer, true)

	aliceIsViewer, err := keeper.HasViewingAccess(res.Files, alice.String())
	suite.Require().NoError(err)
	suite.Require().Equal(aliceIsViewer, true)

	charlieIsViewer, err := keeper.HasViewingAccess(res.Files, charlie.String())
	suite.Require().NoError(err)
	suite.Require().Equal(charlieIsViewer, true)

	pvacc := res.Files.ViewingAccess
	jvacc := make(map[string]string)

	error := json.Unmarshal([]byte(pvacc), &jvacc)
	suite.Require().NoError(error)
	suite.Require().EqualValues(len(jvacc), 3)

	cases := []struct {
		preRun    func() *types.MsgResetViewers
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // charlie fails to reset alice's viewers
			preRun: func() *types.MsgResetViewers {
				return types.NewMsgResetViewers(
					charlie.String(),
					pepeMerklePath,
					aliceOwnerAddress,
				)
			},
			expErr:    true,
			expErrMsg: "Not permitted to remove or reset edit/view access. You are not the owner of this file",
			name:      "charlie fails to reset alice's viewing permissions",
		},
		{ // alice fails to reset her viewing access
			preRun: func() *types.MsgResetViewers {
				return types.NewMsgResetViewers(
					alice.String(),
					types.MerklePath("s/home/ghost.jpg"),
					aliceOwnerAddress,
				)
			},
			expErr:    true,
			expErrMsg: "file not found",
			name:      "alice can't reset viewing access for a file that doesn't exist",
		},
		{ // alice resets her viewing access
			preRun: func() *types.MsgResetViewers {
				return types.NewMsgResetViewers(
					alice.String(),
					pepeMerklePath,
					aliceOwnerAddress,
				)
			},
			expErr: false,
			name:   "alice resets her viewing permissions",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.ResetViewers(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {

				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgRemoveViewersResponse{}, *res)
				// Let's confirm that bob and charlie are no longer viewers

				fileReq := types.QueryFileRequest{
					Address:      pepeMerklePath,
					OwnerAddress: aliceOwnerAddress,
				}
				res, err := suite.queryClient.Files(suite.ctx.Context(), &fileReq)
				suite.Require().NoError(err)

				bobIsViewer, err := keeper.HasViewingAccess(res.Files, bob.String())
				suite.Require().NoError(err)
				suite.Require().EqualValues(bobIsViewer, false)

				charlieIsViewer, err := keeper.HasViewingAccess(res.Files, charlie.String())
				suite.Require().NoError(err)
				suite.Require().EqualValues(charlieIsViewer, false)

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
