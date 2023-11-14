package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil "github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

// testing msg server files for: init_provider, set_provider_ip, set_provider_totalspace

func (suite *KeeperTestSuite) TestMsgInitProvider() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)
	user := testAddresses[0]

	userAdr, err := sdk.AccAddressFromBech32(user)
	suite.Require().NoError(err)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, userAdr, sdk.NewCoins(sdk.NewInt64Coin("ujkl", 10000000000)))
	suite.Require().NoError(err)

	cases := []struct {
		preRun    func() *types.MsgInitProvider
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgInitProvider {
				return types.NewMsgInitProvider(
					user,
					"127.0.0.1",
					1000000000,
					"test-key",
				)
			},
			expErr: false,
			name:   "init provider success",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.InitProvider(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgInitProviderResponse{}, *res)

			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgSetProviderIP() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	user := testAddresses[0]

	provider := types.Providers{
		Address:         user,
		Ip:              "",
		Totalspace:      "1000000",
		Creator:         user,
		BurnedContracts: "0",
	}

	suite.storageKeeper.SetProviders(suite.ctx, provider)

	cases := []struct {
		preRun    func() *types.MsgSetProviderIP
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgSetProviderIP {
				return types.NewMsgSetProviderIP(
					user,
					"127.0.0.1",
				)
			},
			expErr: false,
			name:   "set provider ip success",
		},
		{
			preRun: func() *types.MsgSetProviderIP {
				return types.NewMsgSetProviderIP(
					"wrong address",
					"127.0.0.1",
				)
			},
			expErr:    true,
			expErrMsg: "provider not found please init your provider",
			name:      "set provider ip fail",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.SetProviderIP(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgSetProviderIPResponse{}, *res)

			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgSetProviderTotalSpace() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	user := testAddresses[0]

	provider := types.Providers{
		Address:         user,
		Ip:              "127.0.0.1",
		Totalspace:      "",
		Creator:         user,
		BurnedContracts: "0",
	}

	suite.storageKeeper.SetProviders(suite.ctx, provider)

	cases := []struct {
		preRun    func() *types.MsgSetProviderTotalSpace
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgSetProviderTotalSpace {
				return types.NewMsgSetProviderTotalSpace(
					user,
					1000000,
				)
			},
			expErr: false,
			name:   "set provider total space success",
		},
		{
			preRun: func() *types.MsgSetProviderTotalSpace {
				return types.NewMsgSetProviderTotalSpace(
					"wrong address",
					1000000,
				)
			},
			expErr:    true,
			expErrMsg: "provider not found please init your provider",
			name:      "set provider total space fail",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.SetProviderTotalSpace(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgSetProviderIPResponse{}, *res)

			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgSetProviderKeybase() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	user := testAddresses[0]

	provider := types.Providers{
		Address:         user,
		Ip:              "127.0.0.1",
		Totalspace:      "",
		Creator:         user,
		BurnedContracts: "0",
	}

	suite.storageKeeper.SetProviders(suite.ctx, provider)

	cases := []struct {
		preRun    func() *types.MsgSetProviderKeybase
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgSetProviderKeybase {
				return types.NewMsgSetProviderKeybase(
					user,
					"test_key_1234",
				)
			},
			expErr: false,
			name:   "set provider keybase success",
		},
		{
			preRun: func() *types.MsgSetProviderKeybase {
				return types.NewMsgSetProviderKeybase(
					"wrong address",
					"test_key_1234",
				)
			},
			expErr:    true,
			expErrMsg: "provider not found please init your provider",
			name:      "set provider keybase fail",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.SetProviderKeybase(context, msg)
			if tc.expErr {
				suite.Require().EqualError(err, tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgSetProviderKeybaseResponse{}, *res)

			}
		})
	}
}
