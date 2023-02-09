package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/testutil"
	types "github.com/jackalLabs/canine-chain/x/rns/types"
)

func (suite *KeeperTestSuite) TestDelistMsg() {
	suite.SetupSuite()
	msgSrvr, _, ctx := setupMsgServer(suite)
	// ctx = suite.ctx.WithBlockHeight(100)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	nameOwner, err := sdk.AccAddressFromBech32(testAddresses[0])
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
		testName  string
		preRun    func() *types.MsgDelist
		postRun   func()
		expErr    bool
		expErrMsg string
	}{
		{
			testName: "Name_listed",
			preRun: func() *types.MsgDelist {
				// Check if name is actually saved
				name, found := keeper.GetNames(suite.ctx, "Nuggie", "jkl")
				suite.Require().True(found)
				// Set the name for sale in KVStore
				newsale := types.Forsale{
					Name:  fmt.Sprintf("%s,%s", name.Name, name.Tld),
					Price: "100000000ujkl",
					Owner: nameOwner.String(),
				}
				keeper.SetForsale(suite.ctx, newsale)
				return &types.MsgDelist{
					Creator: nameOwner.String(),
					Name:    fmt.Sprintf("%s,%s", name.Name, name.Tld),
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
			expErr:    false,
			expErrMsg: "Name is listed and can be delisted.",
		},

		{
			testName: "name_not_found",
			preRun: func() *types.MsgDelist {
				return &types.MsgDelist{
					Creator: nameOwner.String(),
					Name:    "nonexistent.jkl",
				}
			},
			expErr:    true,
			expErrMsg: "Name isn't listed.: unauthorized",
		},

		{
			testName: "wrong_onwer",
			preRun: func() *types.MsgDelist {
				return &types.MsgDelist{
					Creator: "wrong_account",
					Name:    "Nuggie.jkl",
				}
			},
			expErr:    true,
			expErrMsg: "Name isn't listed.: unauthorized",
		},

		{
			testName: "cannot_transfer_free_name",
			preRun: func() *types.MsgDelist {
				blockHeight := suite.ctx.BlockHeight()
				names, found := keeper.GetNames(suite.ctx, "Nuggie", "jkl")
				suite.Require().True(found)
				names.Locked = blockHeight + 1
				keeper.SetNames(suite.ctx, names)
				return &types.MsgDelist{
					Creator: nameOwner.String(),
					Name:    "Nuggie.jkl",
				}
			},
			postRun: func() {
				// Turn back to original
				names, found := keeper.GetNames(suite.ctx, "Nuggie", "jkl")
				suite.Require().True(found)
				names.Locked--
				keeper.SetNames(suite.ctx, names)
			},
			expErr:    true,
			expErrMsg: "Name isn't listed.: unauthorized",
		},

		{
			testName: "expired_name",
			preRun: func() *types.MsgDelist {
				blockHeight := suite.ctx.BlockHeight()
				names, found := keeper.GetNames(suite.ctx, "Nuggie", "jkl")
				suite.Require().True(found)
				names.Expires = blockHeight - 1
				keeper.SetNames(suite.ctx, names)
				return &types.MsgDelist{
					Creator: nameOwner.String(),
					Name:    "Nuggie.jkl",
				}
			},
			postRun: func() {
				// Turn back to original
				names, found := keeper.GetNames(suite.ctx, "Nuggie", "jkl")
				suite.Require().True(found)
				names.Expires++
				keeper.SetNames(suite.ctx, names)
			},
			expErr:    true,
			expErrMsg: "Name isn't listed.: unauthorized",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.testName, func() {
			msg := tc.preRun()

			_, err := msgSrvr.Delist(ctx, msg)
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
