package keeper_test

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

const tst = "testownercid"

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
					Creator:  creator.String(),
					Merkle:   "1",
					Signee:   "1",
					Filesize: "1",
					Fid:      "1",
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
					Creator:  creator.String(),
					Merkle:   "1",
					Signee:   "1",
					Filesize: "1",
					Fid:      "1",
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
					Creator:  creator.String(),
					Merkle:   "1",
					Signee:   "1",
					Filesize: "bad_filesize",
					Fid:      "1",
				}
			},
			expErr:    true,
			expErrMsg: "error parsing file size",
		},

		{
			name: "not_enough_provider_storage",
			preRun: func() *types.MsgPostContract {
				return &types.MsgPostContract{
					Creator:  creator.String(),
					Merkle:   "1",
					Signee:   "1",
					Filesize: "1000001",
					Fid:      "1",
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
					Creator:  creator.String(),
					Merkle:   "1",
					Signee:   "1",
					Filesize: "20000000000",
					Fid:      "1",
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
					Creator:  creator.String(),
					Merkle:   "1",
					Signee:   buyer.String(),
					Filesize: "10000",
					Fid:      "1",
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
					Creator:  creator.String(),
					Merkle:   "1",
					Signee:   buyer.String(),
					Filesize: "10000",
					Fid:      "123",
				}
			},
			expErr: false,
		},

		{
			name: "cannot_duplicate_contract",
			preRun: func() *types.MsgPostContract {
				return &types.MsgPostContract{
					Creator:  creator.String(),
					Merkle:   "1",
					Signee:   buyer.String(),
					Filesize: "10000",
					Fid:      "123",
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
				dcid := tst
				h := sha256.New()
				_, err := io.WriteString(h, dcid)
				suite.Require().NoError(err)
				hashName := h.Sum(nil)

				dcid, err = keeper.MakeCid(hashName)
				suite.Require().NoError(err)

				d := types.ActiveDeals{
					Cid:     dcid,
					Creator: user.String(),
				}
				sKeeper.SetActiveDeals(suite.ctx, d)

				for i := 0; i < 2; i++ {
					h := sha256.New()
					_, err := io.WriteString(h, fmt.Sprintf("%s%d", d.Cid, i))
					suite.Require().NoError(err)
					hashName := h.Sum(nil)

					scid, err := keeper.MakeCid(hashName)
					suite.Require().NoError(err)

					k := types.ActiveDeals{
						Cid:     scid,
						Creator: user.String(),
					}
					sKeeper.SetActiveDeals(suite.ctx, k)
				}

				_, found := sKeeper.GetActiveDeals(suite.ctx, d.Cid)
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
				dcid := tst
				h := sha256.New()
				_, err := io.WriteString(h, dcid)
				suite.Require().NoError(err)
				hashName := h.Sum(nil)

				dcid, err = keeper.MakeCid(hashName)
				suite.Require().NoError(err)

				d := types.ActiveDeals{
					Cid:     dcid,
					Creator: user.String(),
				}
				sKeeper.SetActiveDeals(suite.ctx, d)

				for i := 0; i < 2; i++ {
					h := sha256.New()
					_, err := io.WriteString(h, fmt.Sprintf("%s%d", d.Cid, i))
					suite.Require().NoError(err)
					hashName := h.Sum(nil)

					scid, err := keeper.MakeCid(hashName)
					suite.Require().NoError(err)

					k := types.ActiveDeals{
						Cid:     scid,
						Creator: user.String(),
					}
					sKeeper.SetActiveDeals(suite.ctx, k)
				}

				d, found := sKeeper.GetActiveDeals(suite.ctx, d.Cid)
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
				dcid := tst
				h := sha256.New()
				_, err := io.WriteString(h, dcid)
				suite.Require().NoError(err)
				hashName := h.Sum(nil)

				dcid, err = keeper.MakeCid(hashName)
				suite.Require().NoError(err)

				cids := []string{dcid}
				fmt.Println(dcid)

				d := types.ActiveDeals{
					Cid:     dcid,
					Creator: user.String(),
					Fid:     "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
				}
				sKeeper.SetActiveDeals(suite.ctx, d)

				h = sha256.New()
				_, err = io.WriteString(h, fmt.Sprintf("%s%d", dcid, 0))
				suite.Require().NoError(err)
				hashName = h.Sum(nil)

				left, err := keeper.MakeCid(hashName)
				suite.Require().NoError(err)

				cids = append(cids, left)
				k := types.ActiveDeals{
					Cid:     left,
					Creator: user.String(),
					Fid:     "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
				}
				sKeeper.SetActiveDeals(suite.ctx, k)

				h = sha256.New()
				_, err = io.WriteString(h, fmt.Sprintf("%s%d", dcid, 1))
				suite.Require().NoError(err)
				hashName = h.Sum(nil)

				right, err := keeper.MakeCid(hashName)
				suite.Require().NoError(err)

				cids = append(cids, right)
				k = types.ActiveDeals{
					Cid:     right,
					Creator: user.String(),
					Fid:     "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
				}
				sKeeper.SetActiveDeals(suite.ctx, k)

				b, err := json.Marshal(cids)
				suite.Require().NoError(err)

				ftc := types.FidCid{
					Fid:  "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
					Cids: string(b),
				}
				sKeeper.SetFidCid(suite.ctx, ftc)

				suite.Require().NoError(err)

				deals := sKeeper.GetAllActiveDeals(suite.ctx)
				for _, v := range deals {
					fmt.Println(v)
				}

				return &types.MsgCancelContract{
					Creator: user.String(),
					Cid:     right,
				}
			},
			expErr:    true,
			expErrMsg: "can't find contract",
		},

		{
			name: "successfully_cancelled_contract",
			preRun: func() *types.MsgCancelContract {
				dcid := tst
				h := sha256.New()
				_, err := io.WriteString(h, dcid)
				suite.Require().NoError(err)
				hashName := h.Sum(nil)

				dcid, err = keeper.MakeCid(hashName)
				suite.Require().NoError(err)

				cids := []string{dcid}
				fmt.Println(dcid)

				d := types.ActiveDeals{
					Cid:     dcid,
					Creator: user.String(),
					Fid:     "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
				}
				sKeeper.SetActiveDeals(suite.ctx, d)

				h = sha256.New()
				_, err = io.WriteString(h, fmt.Sprintf("%s%d", dcid, 0))
				suite.Require().NoError(err)
				hashName = h.Sum(nil)

				left, err := keeper.MakeCid(hashName)
				suite.Require().NoError(err)

				cids = append(cids, left)
				k := types.ActiveDeals{
					Cid:     left,
					Creator: user.String(),
					Fid:     "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
				}
				sKeeper.SetActiveDeals(suite.ctx, k)

				h = sha256.New()
				_, err = io.WriteString(h, fmt.Sprintf("%s%d", dcid, 1))
				suite.Require().NoError(err)
				hashName = h.Sum(nil)

				right, err := keeper.MakeCid(hashName)
				suite.Require().NoError(err)

				cids = append(cids, right)
				k = types.ActiveDeals{
					Cid:     right,
					Creator: user.String(),
					Fid:     "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
				}
				sKeeper.SetActiveDeals(suite.ctx, k)

				b, err := json.Marshal(cids)
				suite.Require().NoError(err)

				ftc := types.FidCid{
					Fid:  "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
					Cids: string(b),
				}
				sKeeper.SetFidCid(suite.ctx, ftc)

				suite.Require().NoError(err)

				deals := sKeeper.GetAllActiveDeals(suite.ctx)
				for _, v := range deals {
					fmt.Println(v)
				}

				return &types.MsgCancelContract{
					Creator: user.String(),
					Cid:     d.Cid,
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
