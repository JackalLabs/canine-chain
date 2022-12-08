package keeper_test

import (
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

func (suite *KeeperTestSuite) TestCreatePayBlock() {
	suite.SetupSuite()
	//_, sKeeper, _ := setupMsgServer(suite)

	cases := []struct {
		name      string
		preRun    func() (string, int64, int64)
		check     func()
		expErr    bool
		expErrMsg string
	}{

		{
			name: "buying_within_existing_storage_window_payblock",
			preRun: func() (string, int64, int64) {
				return "a", 10000, 10000
			},
			check:     func() {},
			expErr:    true,
			expErrMsg: "can't buy storage within another storage window",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			suite.Require().NotNil(tc.preRun)
			//addr, length, bytes := tc.preRun()
			//err := sKeeper.CreatePayBlock(suite.ctx, addr, length, bytes)
			tc.check()

			if tc.expErr {
				//suite.Require().Error(err)
				//suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				//suite.Require().NoError(err)
			}
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
