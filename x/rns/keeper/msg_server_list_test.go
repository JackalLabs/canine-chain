package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/jackalLabs/canine-chain/x/rns/types"
)

func (suite *KeeperTestSuite) TestListMsg () {
	suite.SetupSuite()
	msgSrvr, _, ctx := setupMsgServer(suite)
	//ctx = suite.ctx.WithBlockHeight(100)

	// Create name owner account
	nameOwner, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)
	coin := sdk.NewCoin("ujkl", sdk.NewInt(100000000)) // Send some coins to their account
	coins := sdk.NewCoins(coin)
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, nameOwner, coins)
	suite.Require().NoError(err)

	// Init rns account and register rns
	rnsName := "Nuggie.jkl"
	suite.rnsKeeper.SetInit(suite.ctx, types.Init{Address: nameOwner.String(), Complete: true})
	err = suite.rnsKeeper.RegisterName(suite.ctx, nameOwner.String(), rnsName, "{}", "2")
	suite.Require().NoError(err)

	keeper := suite.rnsKeeper

	// Each test cases are independent hence test env must be returned to original
	// Use postRun to return it to original state
	cases := []struct {
		testName string
		preRun func() *types.MsgList
		postRun func()
		expErr bool
		expErrMsg string
	}{
		{
			testName: "Name_already_listed",
			preRun: func() *types.MsgList {
				// Check if name is actually saved
				name, found := keeper.GetNames(suite.ctx, "Nuggie", "jkl")
				suite.Require().True(found)
				// Set the name for sale in KVStore
				newsale := types.Forsale{
					Name: name.Name,
					Price: "100000000ujkl",
					Owner: nameOwner.String(),
				}
				keeper.SetForsale(suite.ctx, newsale)
				return &types.MsgList{
					Creator: nameOwner.String(),
					Name: name.Name,
					Price: "100000000ujkl",
				}
			},
			postRun: func() {
				// Clean up Forsale
				name, found := keeper.GetNames(suite.ctx, "Nuggie", "jkl")
				suite.Require().True(found)
				keeper.RemoveForsale(suite.ctx, name.Name)
				_, found = keeper.GetForsale(suite.ctx, name.Name)
				suite.Require().False(found)
			},
			expErr: true,
			expErrMsg: "Name already listed.",
		},

		{
			testName: "name_not_found",
			preRun: func() *types.MsgList {
				return &types.MsgList{
					Creator: nameOwner.String(),
					Name: "nonexistent.jkl",
					Price: "100ujkl",
				}
			},
			expErr: true,
			expErrMsg: "Name does not exist or has expired.",
		},

 		{
			testName: "wrong_onwer",
			preRun: func() *types.MsgList {
				return &types.MsgList{
					Creator: "wrong_account",
					Name: "Nuggie.jkl", 
					Price: "100ujkl",
 				}
			},
			expErr: true,
			expErrMsg: "You do not own this name.",
		},

		{
			testName: "cannot_transfer_free_name",
			preRun: func() *types.MsgList {
				blockHeight := suite.ctx.BlockHeight()
				names, found := keeper.GetNames(suite.ctx, "Nuggie", "jkl")
				suite.Require().True(found)
				names.Locked = blockHeight + 1
				keeper.SetNames(suite.ctx, names)
				return &types.MsgList{
					Creator: nameOwner.String(),
					Name: "Nuggie.jkl",
					Price: "100ujkl",
				}
			},
			postRun: func() {
				// Turn back to original
				names, found := keeper.GetNames(suite.ctx, "Nuggie", "jkl")
				suite.Require().True(found)
				names.Locked = suite.ctx.BlockHeight() - 1
				keeper.SetNames(suite.ctx, names)
			},
			expErr: true,
			expErrMsg: "cannot transfer free name",
		},

		{
			testName: "expired_name",
			preRun: func() *types.MsgList{
				blockHeight := suite.ctx.BlockHeight()
				names, found := keeper.GetNames(suite.ctx, "Nuggie", "jkl")
				suite.Require().True(found)
				names.Expires = blockHeight - 1
				keeper.SetNames(suite.ctx, names)
				return &types.MsgList{
					Creator: nameOwner.String(),
					Name: "Nuggie.jkl",
					Price: "100ujkl",
				}
			},
			postRun: func() {
				// Turn back to original
				names, found := keeper.GetNames(suite.ctx, "Nuggie", "jkl")
				suite.Require().True(found)
				names.Expires = names.Expires + 1
				keeper.SetNames(suite.ctx, names)
			},
			expErr: true,
			expErrMsg: "Name does not exist or has expired.",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.testName, func(){
			msg := tc.preRun()

			_, err := msgSrvr.List(ctx, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}

			if tc.postRun != nil {
				tc.postRun()
			}
		})
	}
}
