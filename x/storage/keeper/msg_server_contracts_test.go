package keeper_test

import (
	"encoding/json"

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

func (suite *KeeperTestSuite) TestSignContract() {
	suite.SetupSuite()
	msgSrvr, sKeeper, goCtx := setupMsgServer(suite)
	provider, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	cases := []struct {
		name      string
		preRun    func() *types.MsgSignContract
		postRun   func()
		expResp   types.MsgSignContractResponse
		expErr    bool
		expErrMsg string
	}{
		{
			name: "contract_not_found",
			preRun: func() *types.MsgSignContract {
				return &types.MsgSignContract{
					Cid:     "contract_that_doesn't_exist",
					Creator: provider.String(),
				}
			},
			expErr:    true,
			expErrMsg: "contract not found",
		},

		{
			name: "invalid_permission_to_sign_contract",
			preRun: func() *types.MsgSignContract {
				c := types.Contracts{
					Cid:        "123",
					Creator:    provider.String(),
					Priceamt:   "1",
					Pricedenom: "ujkl",
					Merkle:     "1",
					Signee:     user.String(),
					Duration:   "10000",
					Filesize:   "10000",
					Fid:        "123",
				}
				sKeeper.SetContracts(suite.ctx, c)
				_, found := sKeeper.GetContracts(suite.ctx, c.Cid)
				suite.Require().True(found)
				return &types.MsgSignContract{
					Cid:     c.Cid,
					Creator: "invalid_creator",
				}
			},
			expErr:    true,
			expErrMsg: "you do not have permission to approve this contract",
		},

		{
			name: "successful_contract_signed",
			preRun: func() *types.MsgSignContract {
				return &types.MsgSignContract{
					Cid:     "123",
					Creator: user.String(),
				}
			},
			expErr: false,
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			// preRun must be defined to get MsgPostContract
			suite.Require().NotNil(tc.preRun)
			c := tc.preRun()
			_, err := msgSrvr.SignContract(goCtx, c)
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

func (suite *KeeperTestSuite) TestCancelContract() {
	suite.SetupSuite()
	msgSrvr, sKeeper, goCtx := setupMsgServer(suite)

	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	cases := []struct {
		name      string
		preRun    func() *types.MsgCancelContract
		postRun   func()
		expResp   types.MsgCancelContractResponse
		expErr    bool
		expErrMsg string
	}{
		{
			name: "active_deal_not_found",
			preRun: func() *types.MsgCancelContract {
				return &types.MsgCancelContract{
					Creator: user.String(),
					Cid:     "foo",
				}
			},
			expErr:    true,
			expErrMsg: "can't find contract",
		},

		{
			name: "invalid_deal_owner",
			preRun: func() *types.MsgCancelContract {
				d := types.ActiveDeals{
					Cid:     "100",
					Creator: user.String(),
				}
				sKeeper.SetActiveDeals(suite.ctx, d)
				_, found := sKeeper.GetActiveDeals(suite.ctx, "100")
				suite.Require().True(found)
				return &types.MsgCancelContract{
					Creator: "foo",
					Cid:     d.Cid,
				}
			},
			expErr:    true,
			expErrMsg: "you don't own this deal",
		},

		{
			name: "fid_not_found",
			preRun: func() *types.MsgCancelContract {
				d, found := sKeeper.GetActiveDeals(suite.ctx, "100")
				suite.Require().True(found)
				d.Fid = "100"
				sKeeper.SetActiveDeals(suite.ctx, d)
				return &types.MsgCancelContract{
					Creator: user.String(),
					Cid:     d.Cid,
				}
			},
			expErr:    true,
			expErrMsg: "no fid found",
		},

		{
			name: "invalid_cid_json",
			preRun: func() *types.MsgCancelContract {
				ftc := types.FidCid{
					Fid:  "100",
					Cids: "100",
				}
				sKeeper.SetFidCid(suite.ctx, ftc)
				return &types.MsgCancelContract{
					Creator: user.String(),
					Cid:     ftc.Cids,
				}
			},
			expErr:    true,
			expErrMsg: "cannot unmarshal number into Go value of type []string",
		},

		{
			name: "successfully_cancelled_contract",
			preRun: func() *types.MsgCancelContract {
				ncids := []string{"abc", "def", "foo", "123_bar"}
				b, err := json.Marshal(ncids)
				suite.Require().NoError(err)
				ftc := types.FidCid{
					Fid:  "100",
					Cids: string(b),
				}
				sKeeper.SetFidCid(suite.ctx, ftc)
				return &types.MsgCancelContract{
					Creator: user.String(),
					Cid:     "100",
				}
			},
			expErr: false,
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			// preRun must be defined to get MsgPostContract
			suite.Require().NotNil(tc.preRun)
			c := tc.preRun()
			_, err := msgSrvr.CancelContract(goCtx, c)
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

func (suite *KeeperTestSuite) TestClaimStray() {
	suite.SetupSuite()
	msgSrvr, sKeeper, goCtx := setupMsgServer(suite)
	provider, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	provider2, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	cases := []struct {
		name      string
		preRun    func() *types.MsgClaimStray
		postRun   func()
		expResp   types.MsgClaimStrayResponse
		expErr    bool
		expErrMsg string
	}{
		{
			name: "stray_not_found",
			preRun: func() *types.MsgClaimStray {
				return &types.MsgClaimStray{
					Creator: provider.String(),
					Cid:     "foo",
				}
			},
			expErr:    true,
			expErrMsg: "stray contract either no longer is stray, or has been removed by the user",
		},

		{
			name: "not_a_provider",
			preRun: func() *types.MsgClaimStray {
				s := types.Strays{
					Cid: "foo",
				}
				sKeeper.SetStrays(suite.ctx, s)
				return &types.MsgClaimStray{
					Cid:     s.Cid,
					Creator: provider.String(),
				}
			},
			expErr:    true,
			expErrMsg: "not a provider",
		},

		{
			name: "cannot_claim_your_own_stray",
			preRun: func() *types.MsgClaimStray {
				s, found := sKeeper.GetStrays(suite.ctx, "foo")
				suite.Require().True(found)
				s.Fid = "some_fid"
				sKeeper.SetStrays(suite.ctx, s)
				p := types.Providers{
					Ip:      "0.0.0.0",
					Address: provider.String(),
					Creator: provider.String(),
				}
				sKeeper.SetProviders(suite.ctx, p)
				ad := types.ActiveDeals{
					Fid:      s.Fid,
					Provider: p.Address,
				}
				sKeeper.SetActiveDeals(suite.ctx, ad)
				return &types.MsgClaimStray{
					Cid:     s.Cid,
					Creator: provider.String(),
				}
			},
			expErr:    true,
			expErrMsg: "cannot claim a stray you own.",
		},

		{
			name: "successfully_claimed_stray",
			preRun: func() *types.MsgClaimStray {
				p := types.Providers{
					Ip:      "123.0.0.0",
					Address: provider2.String(),
					Creator: provider.String(),
				}
				sKeeper.SetProviders(suite.ctx, p)
				return &types.MsgClaimStray{
					Cid:     "foo",
					Creator: provider2.String(),
				}
			},
			expErr: false,
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			// preRun must be defined to get MsgPostContract
			suite.Require().NotNil(tc.preRun)
			c := tc.preRun()
			_, err := msgSrvr.ClaimStray(goCtx, c)
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
