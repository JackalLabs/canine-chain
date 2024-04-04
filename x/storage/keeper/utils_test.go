package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
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
			paidAmt: 0,
			free:    true,
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			// preRun must be defined to get MsgPostContract
			suite.Require().NotNil(tc.preRun)
			addr, _ := tc.preRun()
			rPaidAmt := sKeeper.GetPaidAmount(suite.ctx, addr)

			suite.Require().Equal(tc.paidAmt, rPaidAmt)
		})
	}
}

func (suite *KeeperTestSuite) TestGetProviderUsing() {
	suite.SetupSuite()
	setupMsgServer(suite)

	cases := []struct {
		name      string
		preRun    func()
		expReturn int64
	}{
		{
			name: "No_provider_found",
			preRun: func() {
				ad := types.UnifiedFile{
					Merkle:    []byte("merkle"),
					Owner:     "owner",
					Start:     0,
					FileSize:  100000,
					MaxProofs: 3,
				}
				suite.storageKeeper.SetFile(suite.ctx, ad)
			},
			expReturn: 0,
		},

		{
			name: "valid_active_deal_file_size",
			preRun: func() {
				ad := types.UnifiedFile{
					Merkle:    []byte("merkle"),
					Owner:     "owner",
					Start:     0,
					FileSize:  100000,
					MaxProofs: 3,
				}
				suite.storageKeeper.SetFile(suite.ctx, ad)
				ad.AddProver(suite.ctx, suite.storageKeeper, "prover1")
			},
			expReturn: 100000,
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			suite.Require().NotNil(tc.preRun)
			tc.preRun()
			result := suite.storageKeeper.GetProviderUsing(suite.ctx, "prover1")

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
			gbs:      10,
			months:   5,
			expected: sdk.NewInt(555555),
		},
		{
			name:     "5GB for 24months",
			gbs:      5,
			months:   24,
			expected: sdk.NewInt(1111111),
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			cost := sKeeper.GetStorageCost(suite.ctx, tc.gbs, tc.months*720)
			suite.Require().Equal(tc.expected, cost)
		})
	}
}
