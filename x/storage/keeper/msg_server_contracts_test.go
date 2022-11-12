package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (suite *KeeperTestSuite) TestPostContracts() {
	suite.SetupSuite()
	msgSrvr, sKeeper, goCtx := setupMsgServer(suite)
	creator, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	buyer, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	cases := []struct {
		name      string
		preRun    func() *types.MsgPostContract
		postRun   func()
		expResp   types.MsgPostContractResponse
		expErr    bool
		expErrMsg string
	}{
		{
			name: "provider_doesn't_exist",
			preRun: func() *types.MsgPostContract {
				return &types.MsgPostContract{
					Creator:    creator.String(),
					Priceamt:   "1",
					Pricedenom: "1",
					Merkle:     "1",
					Signee:     "1",
					Duration:   "1",
					Filesize:   "1",
					Fid:        "1",
				}
			},
			expErr:    true,
			expErrMsg: "can't find provider",
		},

		{
			name: "invalid_provider_total_space_format",
			preRun: func() *types.MsgPostContract {
				// Set provider with invalid totalspace string
				p := types.Providers{
					Address:         creator.String(),
					Ip:              "123.0.0.0",
					Totalspace:      "bad_content",
					BurnedContracts: "",
					Creator:         creator.String(),
				}
				sKeeper.SetProviders(suite.ctx, p)
				_, found := sKeeper.GetProviders(suite.ctx, creator.String())
				suite.Require().True(found)
				return &types.MsgPostContract{
					Creator:    creator.String(),
					Priceamt:   "1",
					Pricedenom: "1",
					Merkle:     "1",
					Signee:     "1",
					Duration:   "1",
					Filesize:   "1",
					Fid:        "1",
				}
			},
			postRun: func() {
				// fix the bad format for next tc
				p, found := sKeeper.GetProviders(suite.ctx, creator.String())
				suite.Require().True(found)
				p.Totalspace = "100000"
				sKeeper.SetProviders(suite.ctx, p)
			},
			expErr:    true,
			expErrMsg: "error parsing total space",
		},

		{
			name: "bad_filesize_format",
			preRun: func() *types.MsgPostContract {
				return &types.MsgPostContract{
					Creator:    creator.String(),
					Priceamt:   "1",
					Pricedenom: "1",
					Merkle:     "1",
					Signee:     "1",
					Duration:   "1",
					Filesize:   "bad_filesize",
					Fid:        "1",
				}
			},
			expErr:    true,
			expErrMsg: "error parsing file size",
		},

		{
			name: "not_enough_provider_storage",
			preRun: func() *types.MsgPostContract {
				return &types.MsgPostContract{
					Creator:    creator.String(),
					Priceamt:   "1",
					Pricedenom: "1",
					Merkle:     "1",
					Signee:     "1",
					Duration:   "1",
					Filesize:   "1000001",
					Fid:        "1",
				}
			},
			expErr:    true,
			expErrMsg: "not enough space on provider",
		},

		{
			name: "not_enough_user_storage",
			preRun: func() *types.MsgPostContract {
				// Setup provider storag
				p := types.Providers{
					Address:         creator.String(),
					Ip:              "123.0.0.0",
					Totalspace:      "1000000000000000",
					BurnedContracts: "",
					Creator:         creator.String(),
				}
				sKeeper.SetProviders(suite.ctx, p)
				// start free two gig trial
				suite.ctx = suite.ctx.WithBlockHeight(0)
				err := sKeeper.CreatePayBlock(suite.ctx, buyer.String(), 1, 0)
				suite.Require().NoError(err)
				sKeeper.SetClientUsage(suite.ctx, types.ClientUsage{
					Usage:   "1900000000",
					Address: buyer.String(),
				})
				return &types.MsgPostContract{
					Creator:    creator.String(),
					Priceamt:   "1",
					Pricedenom: "1",
					Merkle:     "1",
					Signee:     "1",
					Duration:   "1",
					Filesize:   "20000000000",
					Fid:        "1",
				}
			},
			expErr:    true,
			expErrMsg: "not enough storage on the users account",
		},

		{
			name: "user_didn't_pay_for_storage",
			preRun: func() *types.MsgPostContract {
				// Start free trial
				suite.ctx = suite.ctx.WithBlockHeight(0)
				err := sKeeper.CreatePayBlock(suite.ctx, buyer.String(), 100, 100000000)
				suite.Require().NoError(err)
				// end free trial and create "not paid" condition
				suite.ctx = suite.ctx.WithBlockHeight(100)
				err = sKeeper.CreatePayBlock(suite.ctx, buyer.String(), 100000, 0)
				suite.Require().NoError(err)
				suite.ctx = suite.ctx.WithBlockHeight(500)
				goCtx = sdk.WrapSDKContext(suite.ctx)
				return &types.MsgPostContract{
					Creator:    creator.String(),
					Priceamt:   "1",
					Pricedenom: "1",
					Merkle:     "1",
					Signee:     buyer.String(),
					Duration:   "1",
					Filesize:   "10000",
					Fid:        "1",
				}
			},
			expErr:    true,
			expErrMsg: "user has not paid for any storage",
		},

		{
			name: "successful_post_contract",
			preRun: func() *types.MsgPostContract {
				err := sKeeper.CreatePayBlock(suite.ctx, buyer.String(), 100000, 10000000000)
				suite.Require().NoError(err)
				return &types.MsgPostContract{
					Creator:    creator.String(),
					Priceamt:   "1",
					Pricedenom: "ujkl",
					Merkle:     "1",
					Signee:     buyer.String(),
					Duration:   "10000",
					Filesize:   "10000",
					Fid:        "123",
				}
			},
			expErr: false,
		},

		{
			name: "cannot_duplicate_contract",
			preRun: func() *types.MsgPostContract {
				return &types.MsgPostContract{
					Creator:    creator.String(),
					Priceamt:   "1",
					Pricedenom: "ujkl",
					Merkle:     "1",
					Signee:     buyer.String(),
					Duration:   "10000",
					Filesize:   "10000",
					Fid:        "123",
				}
			},
			expErr:    true,
			expErrMsg: "cannot post the same contract twice",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			// preRun must be defined to get MsgPostContract
			suite.Require().NotNil(tc.preRun)
			c := tc.preRun()
			_, err := msgSrvr.PostContract(goCtx, c)
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
