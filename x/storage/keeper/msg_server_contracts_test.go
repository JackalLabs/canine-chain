package keeper_test

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	testutil "github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

const tst = "testownercid"

func (suite *KeeperTestSuite) TestPostContracts() {
	suite.SetupSuite()
	msgSrvr, sKeeper, goCtx := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	creator := testAddresses[0]
	buyer := testAddresses[1]

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
					Creator:  creator,
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
					Address:         creator,
					Ip:              "123.0.0.0",
					Totalspace:      "bad_content",
					BurnedContracts: "",
					Creator:         creator,
				}
				sKeeper.SetProviders(suite.ctx, p)
				_, found := sKeeper.GetProviders(suite.ctx, creator)
				suite.Require().True(found)
				return &types.MsgPostContract{
					Creator:  creator,
					Merkle:   "1",
					Signee:   "1",
					Filesize: "1",
					Fid:      "1",
				}
			},
			postRun: func() {
				// fix the bad format for next tc
				p, found := sKeeper.GetProviders(suite.ctx, creator)
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
					Creator:  creator,
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
					Creator:  creator,
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
			name: "successful_post_contract",
			preRun: func() *types.MsgPostContract {
				suite.Require().NoError(err)
				return &types.MsgPostContract{
					Creator:  creator,
					Merkle:   "1",
					Signee:   buyer,
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
					Creator:  creator,
					Merkle:   "1",
					Signee:   buyer,
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

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	provider := testAddresses[0]
	user := testAddresses[1]
	coins, err := sdk.ParseCoinsNormalized("500000000ujkl")
	suite.Require().NoError(err)

	address, err := sdk.AccAddressFromBech32(user)
	suite.Require().NoError(err)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, banktypes.ModuleName, address, coins)
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
					Creator: provider,
				}
			},
			expErr:    true,
			expErrMsg: "contract not found",
		},

		{
			name: "invalid_permission_to_sign_contract",
			preRun: func() *types.MsgSignContract {
				// creating a test contract to sign
				c := types.Contracts{
					Cid:        "123",
					Creator:    provider,
					Priceamt:   "1",
					Pricedenom: "ujkl",
					Merkle:     "1",
					Signee:     user,
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
			name: "not enough storage",
			preRun: func() *types.MsgSignContract {
				// create a test StoragePaymentInfo
				spi := types.StoragePaymentInfo{
					SpaceAvailable: 200_000_000,
					SpaceUsed:      200_000_000,
					Address:        user,
				}
				sKeeper.SetStoragePaymentInfo(suite.ctx, spi)
				_, found := sKeeper.GetStoragePaymentInfo(suite.ctx, user)
				suite.Require().True(found)
				return &types.MsgSignContract{
					Cid:     "123",
					Creator: user,
				}
			},
			expErr:    true,
			expErrMsg: "not enough storage space",
			postRun: func() {
				sKeeper.RemoveStoragePaymentInfo(suite.ctx, user)
			},
		},
		{
			name: "expired storage subscription",
			preRun: func() *types.MsgSignContract {
				// create a test StoragePaymentInfo
				spi := types.StoragePaymentInfo{
					SpaceAvailable: 200_000_000,
					SpaceUsed:      0,
					// set expiration date to yesterday
					End:     time.Now().AddDate(0, -1, 0),
					Address: user,
				}
				sKeeper.SetStoragePaymentInfo(suite.ctx, spi)
				_, found := sKeeper.GetStoragePaymentInfo(suite.ctx, user)
				suite.Require().True(found)
				return &types.MsgSignContract{
					Cid:     "123",
					Creator: user,
				}
			},
			expErr:    true,
			expErrMsg: "storage subscription has expired",
			postRun: func() {
				sKeeper.RemoveStoragePaymentInfo(suite.ctx, user)
			},
		},
		{
			name: "successful_contract_signed",
			preRun: func() *types.MsgSignContract {
				spi := types.StoragePaymentInfo{
					SpaceAvailable: 200_000_000,
					SpaceUsed:      0,
					End:            time.Now().AddDate(0, 10, 0),
					Address:        user,
				}
				sKeeper.SetStoragePaymentInfo(suite.ctx, spi)
				_, found := sKeeper.GetStoragePaymentInfo(suite.ctx, user)
				suite.Require().True(found)
				return &types.MsgSignContract{
					Cid:     "123",
					Creator: user,
				}
			},
			expErr: false,
		},
		{
			name: "pay_once",
			preRun: func() *types.MsgSignContract {
				// creating a test contract to sign
				c := types.Contracts{
					Cid:        "456",
					Creator:    provider,
					Priceamt:   "1",
					Pricedenom: "ujkl",
					Merkle:     "1",
					Signee:     user,
					Duration:   "10000",
					Filesize:   "10000",
					Fid:        "123",
				}
				sKeeper.SetContracts(suite.ctx, c)
				_, found := sKeeper.GetContracts(suite.ctx, c.Cid)
				suite.Require().True(found)
				return &types.MsgSignContract{
					Cid:     "456",
					Creator: user,
					PayOnce: true,
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

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)
	user := testAddresses[0]

	suite.storageKeeper.SetStoragePaymentInfo(suite.ctx, types.StoragePaymentInfo{
		Start:          time.Now(),
		End:            time.Now().AddDate(1, 0, 0),
		SpaceAvailable: 1000000000,
		SpaceUsed:      0,
		Address:        user,
	})

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
					Creator: user,
					Cid:     "foo",
				}
			},
			expErr:    true,
			expErrMsg: "cid does not exist",
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
					Cid:      dcid,
					Signee:   user,
					Creator:  user,
					Filesize: "10",
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
						Cid:      scid,
						Signee:   user,
						Creator:  user,
						Filesize: "10",
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
			expErrMsg: fmt.Sprintf("cannot cancel a contract that isn't yours. foo is not %s: unauthorized", user),
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

				d := types.ActiveDeals{
					Cid:      dcid,
					Creator:  user,
					Signee:   user,
					Fid:      "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
					Filesize: "10",
				}
				sKeeper.SetActiveDeals(suite.ctx, d)

				b, err := json.Marshal(cids)
				suite.Require().NoError(err)

				ftc := types.FidCid{
					Fid:  "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
					Cids: string(b),
				}
				sKeeper.SetFidCid(suite.ctx, ftc)

				suite.Require().NoError(err)

				return &types.MsgCancelContract{
					Creator: user,
					Cid:     d.Cid,
				}
			},
			expErr: false,
		},

		{
			name: "successfully_cancelled_contract_with_strays",
			preRun: func() *types.MsgCancelContract {
				dcid := tst
				h := sha256.New()
				_, err := io.WriteString(h, dcid)
				suite.Require().NoError(err)
				hashName := h.Sum(nil)

				dcid, err = keeper.MakeCid(hashName)
				suite.Require().NoError(err)

				cids := []string{dcid}

				d := types.Strays{
					Cid:      dcid,
					Fid:      "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
					Signee:   user,
					Filesize: "10",
				}
				sKeeper.SetStrays(suite.ctx, d)

				b, err := json.Marshal(cids)
				suite.Require().NoError(err)

				ftc := types.FidCid{
					Fid:  "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
					Cids: string(b),
				}
				sKeeper.SetFidCid(suite.ctx, ftc)

				suite.Require().NoError(err)

				return &types.MsgCancelContract{
					Creator: user,
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
				fidCid, found := suite.storageKeeper.GetFidCid(suite.ctx, "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x")
				suite.Require().True(found)
				var cids []string
				err := json.Unmarshal([]byte(fidCid.Cids), &cids) // getting all cids from the existing fid_cid
				if err != nil {
					suite.Require().NoError(err)
				}
				suite.Require().Equal(0, len(cids))
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

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 10)
	suite.Require().NoError(err)

	provider := testAddresses[0]
	provider2 := testAddresses[1]

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
					Creator:    provider,
					Cid:        "foo",
					ForAddress: provider,
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
					Cid:        s.Cid,
					Creator:    provider,
					ForAddress: provider,
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
					Address: provider,
					Creator: provider,
				}
				sKeeper.SetProviders(suite.ctx, p)
				ad := types.ActiveDeals{
					Fid:      s.Fid,
					Provider: p.Address,
				}
				sKeeper.SetActiveDeals(suite.ctx, ad)
				return &types.MsgClaimStray{
					Cid:        s.Cid,
					Creator:    provider,
					ForAddress: provider,
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
					Address: provider2,
					Creator: provider2,
				}
				s := types.Strays{
					Cid: "foo",
				}
				sKeeper.SetStrays(suite.ctx, s)
				sKeeper.SetProviders(suite.ctx, p)
				return &types.MsgClaimStray{
					Cid:        "foo",
					Creator:    provider2,
					ForAddress: provider2,
				}
			},
			expErr: false,
		},
		{
			name: "successfully_claimed_stray_with_auth_claim",

			preRun: func() *types.MsgClaimStray {
				p := types.Providers{
					Ip:           "192.168.0.40",
					Address:      testAddresses[5],
					Creator:      testAddresses[5],
					AuthClaimers: []string{testAddresses[6]},
				}
				s := types.Strays{
					Cid: "quoz",
				}
				sKeeper.SetStrays(suite.ctx, s)
				sKeeper.SetProviders(suite.ctx, p)
				return &types.MsgClaimStray{
					Cid:        "quoz",
					Creator:    testAddresses[6],
					ForAddress: testAddresses[5],
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
