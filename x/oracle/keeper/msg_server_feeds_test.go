package keeper_test

import (
	"github.com/jackalLabs/canine-chain/x/oracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

func (suite *KeeperTestSuite) TestCreateFeed() {
	genesisFeed := types.Feed{
		Owner: suite.testAccs[0].String(),
		Data:  "",
		Name:  "conflict",
	}

	cases := map[string]struct {
		creator   sdk.AccAddress
		name      string
		expErr    bool
		expErrMsg string
	}{
		"create feed": {
			creator: suite.testAccs[0],
			name:    "foo",
			expErr:  false,
		},

		"cannot overwrite feed": {
			creator:   suite.testAccs[0],
			name:      "conflict",
			expErr:    true,
			expErrMsg: "overwrite",
		},

		"not enough balance": {
			creator:   suite.testAccs[1],
			name:      "bar",
			expErr:    true,
			expErrMsg: "not enough balance",
		},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
			suite.SetupSuite()
			msgSrvr, _, _ := setupMsgServer(suite)
			suite.oracleKeeper.SetFeed(suite.ctx, genesisFeed)

			// TODO: setup simulation and use that instead
			// Fund account
			coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(1000000000000)))
			err := suite.bankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
			suite.NoError(err)
			err = suite.bankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, coins)
			suite.NoError(err)

			coin := sdk.NewInt64Coin("ujkl", 100000000)
			err = suite.bankKeeper.SendCoinsFromModuleToAccount(
				suite.ctx,
				types.ModuleName,
				suite.testAccs[0],
				sdk.NewCoins(coin))
			suite.Require().NoError(err)

			result, err := msgSrvr.CreateFeed(sdk.WrapSDKContext(suite.ctx), &types.MsgCreateFeed{
				Creator: tc.creator.String(),
				Name:    tc.name,
			})

			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
				suite.Require().Nil(result)
			}
		})
	}
}
