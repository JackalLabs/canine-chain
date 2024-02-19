package keeper_test

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgRemoveViewers() {
	suite.SetupSuite()
	msgSrvr, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 3)
	suite.Require().NoError(err)

	alice := testAddresses[0]
	bob := testAddresses[1]
	charlie := testAddresses[2]

	// Let it be that bob has posted a public key after signing
	// with his keyring backend using the CLI.

	bobPrivateKey, err := types.MakePrivateKey("bob")
	suite.Require().NoError(err)

	bobPubKey := bobPrivateKey.PublicKey.Bytes(false) // to hex
	pubKeyStruct := types.Pubkey{
		Address: bob,
		Key:     fmt.Sprintf("%x", bobPubKey),
	}
	suite.filetreeKeeper.SetPubkey(suite.ctx, pubKeyStruct)

	// set root folder for alice
	aliceRootFolder, err := types.CreateRootFolder(alice)
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceRootFolder)

	// set home folder for alice
	aliceHomeFolder, err := types.CreateFolderOrFile(alice, strings.Split(alice, ","), strings.Split(alice, ","), "s/home/")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *aliceHomeFolder)

	// add bob as a viewer for pepe

	viewerIDs := strings.Split(alice, ",")
	viewerIDs = append(viewerIDs, bob)

	// put pepe in home
	pepejpg, err := types.CreateFolderOrFile(alice, strings.Split(alice, ","), viewerIDs, "s/home/pepe.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *pepejpg)

	pepeMerklePath := types.MerklePath("s/home/pepe.jpg")
	aliceAccountHash := types.HashThenHex(alice)
	aliceOwnerAddress := types.MakeOwnerAddress(pepeMerklePath, aliceAccountHash)

	// Let's query the file after it was set to confirm that alice and bob are viewers

	fileReq := types.QueryFile{
		Address:      pepeMerklePath,
		OwnerAddress: aliceOwnerAddress,
	}

	res, err := suite.queryClient.File(suite.ctx.Context(), &fileReq)
	suite.Require().NoError(err)

	bobIsViewer, err := keeper.HasViewingAccess(res.File, alice)
	suite.Require().NoError(err)
	suite.Require().Equal(bobIsViewer, true)

	aliceIsViewer, err := keeper.HasViewingAccess(res.File, alice)
	suite.Require().NoError(err)
	suite.Require().Equal(aliceIsViewer, true)

	bobViewerAddress := keeper.MakeViewerAddress(res.File.TrackingNumber, bob)

	cases := []struct {
		preRun    func() *types.MsgRemoveViewers
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // charlie fails to remove bob from alice's viewing access
			preRun: func() *types.MsgRemoveViewers {
				return types.NewMsgRemoveViewers(
					charlie,
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
					alice,
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
					alice,
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

				fileReq := types.QueryFile{
					Address:      pepeMerklePath,
					OwnerAddress: aliceOwnerAddress,
				}
				res, err := suite.queryClient.File(suite.ctx.Context(), &fileReq)
				suite.Require().NoError(err)

				bobIsViewer, err := keeper.HasViewingAccess(res.File, bob)
				suite.Require().NoError(err)
				suite.Require().EqualValues(bobIsViewer, false)

				aliceIsViewer, err := keeper.HasViewingAccess(res.File, alice)
				suite.Require().NoError(err)
				suite.Require().EqualValues(aliceIsViewer, true)

				pvacc := res.File.ViewingAccess
				jvacc := make(map[string]string)

				err = json.Unmarshal([]byte(pvacc), &jvacc)
				suite.Require().NoError(err)
				suite.Require().EqualValues(len(jvacc), 1)

			}
		})
	}
}
