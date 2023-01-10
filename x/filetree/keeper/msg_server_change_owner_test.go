package keeper_test

import (
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgChangeOwners() {
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

	// put pepe in home of alice
	pepejpg, err := types.CreateFolderOrFile(alice.String(), strings.Split(alice.String(), ","), strings.Split(alice.String(), ","), "s/home/pepe.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *pepejpg)

	pepeMerklePath := types.MerklePath("s/home/pepe.jpg")
	aliceAccountHash := types.HashThenHex(alice.String())
	aliceOwnerAddress := types.MakeOwnerAddress(pepeMerklePath, aliceAccountHash)

	// Let's query the file after it was set to confirm that alice is the owner

	fileReq := types.QueryFileRequest{
		Address:      pepeMerklePath,
		OwnerAddress: aliceOwnerAddress,
	}

	res, err := suite.queryClient.Files(suite.ctx.Context(), &fileReq)
	suite.Require().NoError(err)

	aliceIsOwner := keeper.IsOwner(res.Files, alice.String())
	suite.Require().Equal(aliceIsOwner, true)

	// we make a pepe.jpg for charlie as well to show that alice cannot give charlie her 'pepe.jpg' if he already has one--i.e., duplicates are not permitted
	// set root folder for charlie
	charlieRootFolder, err := types.CreateRootFolder(charlie.String())
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *charlieRootFolder)

	// set home folder for charlie
	charlieHomeFolder, err := types.CreateFolderOrFile(charlie.String(), strings.Split(charlie.String(), ","), strings.Split(charlie.String(), ","), "s/home/")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *charlieHomeFolder)

	// put pepe in home of charlie
	charliePepejpg, err := types.CreateFolderOrFile(charlie.String(), strings.Split(charlie.String(), ","), strings.Split(charlie.String(), ","), "s/home/pepe.jpg")
	suite.Require().NoError(err)
	suite.filetreeKeeper.SetFiles(suite.ctx, *charliePepejpg)

	cases := []struct {
		preRun    func() *types.MsgChangeOwner
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // charlie can't give away a file he doesn't own
			preRun: func() *types.MsgChangeOwner {
				return types.NewMsgChangeOwner(
					charlie.String(),
					pepeMerklePath,
					aliceAccountHash,
					types.HashThenHex(bob.String()),
				)
			},
			expErr:    true,
			expErrMsg: "You do not own this file and cannot give it away",
			name:      "charlie can't give away a file he doesn't own",
		},
		{ // alice can't give pepe.jpg to charlie because he already owns a pepe.jpg
			preRun: func() *types.MsgChangeOwner {
				return types.NewMsgChangeOwner(
					alice.String(),
					pepeMerklePath,
					aliceAccountHash,
					types.HashThenHex(charlie.String()),
				)
			},
			expErr:    true,
			expErrMsg: "Proposed new owner already has a file set with this path name. No duplicates allowed.",
			name:      "alice can't give pepe.jpg to charlie",
		},
		{ // alice can't give away a file that doesn't exist
			preRun: func() *types.MsgChangeOwner {
				return types.NewMsgChangeOwner(
					alice.String(),
					types.MerklePath("s/home/ghost.jpg"),
					aliceAccountHash,
					types.HashThenHex(bob.String()),
				)
			},
			expErr:    true,
			expErrMsg: "file not found",
			name:      "alice can't give away non existent file",
		},
		{ // alice can give pepe.jpg to bob
			preRun: func() *types.MsgChangeOwner {
				return types.NewMsgChangeOwner(
					alice.String(),
					pepeMerklePath,
					aliceAccountHash,
					types.HashThenHex(bob.String()),
				)
			},
			expErr: false,
			name:   "alice gives pepe.jpg to bob",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.ChangeOwner(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {

				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgChangeOwnerResponse{}, *res)

				// Because filetree entries are indexed (keyed) by address and ownerAddress, querying for a pepe.jpg that belongs to alice as an owner
				// should fail here because alice gave away pepe.jpg to bob

				fileReq1 := types.QueryFileRequest{
					Address:      pepeMerklePath,
					OwnerAddress: aliceOwnerAddress,
				}
				_, err := suite.queryClient.Files(suite.ctx.Context(), &fileReq1)
				suite.Require().Error(err)

				// we will find a pepe.jpg that belongs to bob

				bobAccountHash := types.HashThenHex(bob.String())
				bobOwnerAddress := types.MakeOwnerAddress(pepeMerklePath, bobAccountHash)

				fileReq2 := types.QueryFileRequest{
					Address:      pepeMerklePath,
					OwnerAddress: bobOwnerAddress,
				}
				res, error := suite.queryClient.Files(suite.ctx.Context(), &fileReq2)
				suite.Require().NoError(error)

				bobIsOwner := keeper.IsOwner(res.Files, bob.String())
				suite.Require().Equal(bobIsOwner, true)

				aliceIsOwner := keeper.IsOwner(res.Files, alice.String())
				suite.Require().Equal(aliceIsOwner, false)

			}
		})
	}
}
