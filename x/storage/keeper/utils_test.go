package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	module "github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (suite *KeeperTestSuite) TestGetPaidAmount() {
	suite.SetupSuite()
	_, sKeeper, _ := setupMsgServer(suite)

	cases := []struct {
		name    string
		preRun  func() (string, int64)
		paidAmt int64
		free    bool
	}{
		{
			name: "no_payblock",
			preRun: func() (string, int64) {
				suite.ctx = suite.ctx.WithBlockHeight(100)
				return "cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl", 1
			},
			paidAmt: module.TwoGigs,
			free:    true,
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			// preRun must be defined to get MsgPostContract
			suite.Require().NotNil(tc.preRun)
			addr, _ := tc.preRun()
			rPaidAmt, rFree := sKeeper.GetPaidAmount(suite.ctx, addr)

			suite.Require().Equal(tc.paidAmt, rPaidAmt)
			suite.Require().Equal(tc.free, rFree)
		})
	}
}

func (suite *KeeperTestSuite) TestGetProviderUsing() {
	suite.SetupSuite()
	_, sKeeper, _ := setupMsgServer(suite)

	cases := []struct {
		name      string
		preRun    func() string
		expReturn int64
	}{
		{
			name: "No_provider_found",
			preRun: func() string {
				return "a"
			},
			expReturn: 0,
		},

		{
			name: "invalid_active_deal_file_size",
			preRun: func() string {
				ad := types.ActiveDeals{
					Provider: "a",
					Filesize: "aaaaa",
					Cid:      "abc",
				}
				sKeeper.SetActiveDeals(suite.ctx, ad)
				return "a"
			},
			expReturn: 0,
		},

		{
			name: "valid_active_deal_file_size",
			preRun: func() string {
				ad := types.ActiveDeals{
					Provider: "a",
					Filesize: "100000",
					Cid:      "bbb",
				}
				sKeeper.SetActiveDeals(suite.ctx, ad)
				return "a"
			},
			expReturn: 100000,
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			suite.Require().NotNil(tc.preRun)
			provider := tc.preRun()
			result := sKeeper.GetProviderUsing(suite.ctx, provider)

			suite.Require().Equal(tc.expReturn, result)
		})
	}
}

func (suite *KeeperTestSuite) TestGetJklPrice() {
	suite.SetupSuite()
	_, sKeeper, _ := setupMsgServer(suite)

	price := sKeeper.GetJklPrice(suite.ctx)
	expected, err := sdk.NewDecFromStr("0.24")
	suite.Require().NoError(err)
	suite.Require().Equal(expected, price)
}

func (suite *KeeperTestSuite) TestGetStorageCost() {
	suite.SetupSuite()
	_, sKeeper, _ := setupMsgServer(suite)

	cases := []struct {
		name     string
		gbs      int64
		months   int64
		expected sdk.Int
	}{
		{
			name:     "10GB for 5months",
			gbs:      10_000_000,
			months:   5,
			expected: sdk.NewInt(1666666666666),
		},
		{
			name:     "5GB for 24months",
			gbs:      5_000_000,
			months:   24,
			expected: sdk.NewInt(4000000000000),
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			cost := sKeeper.GetStorageCost(suite.ctx, tc.gbs, tc.months)
			suite.Require().Equal(tc.expected, cost)
		})
	}
}
