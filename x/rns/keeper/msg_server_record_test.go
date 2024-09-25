package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

// testing msg server files for: addRecord, deleteRecord
func (suite *KeeperTestSuite) TestMsgAddRecord() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	owner, err := sdk.AccAddressFromBech32(testAddresses[0])
	suite.Require().NoError(err)

	coin := sdk.NewCoin("ujkl", sdk.NewInt(10000000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, owner, coins)
	suite.Require().NoError(err)

	err = suite.rnsKeeper.RegisterRNSName(suite.ctx, owner.String(), "BiPhan.jkl", "{}", 2, true)
	suite.Require().NoError(err)

	cases := []struct {
		preRun    func() *types.MsgAddRecord
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgAddRecord {
				return types.NewMsgAddRecord(
					owner.String(),
					"BiPhan.jkl",
					"app.BiPhan.jkl",
					owner.String(),
					"{}",
				)
			},
			expErr: false,
			name:   "successfully added record",
		},
		{
			preRun: func() *types.MsgAddRecord {
				return types.NewMsgAddRecord(
					owner.String(),
					"Nuggie.jkl",
					"BiPhanApp.Nuggie.jkl",
					owner.String(),
					"{}",
				)
			},
			expErr:    true,
			expErrMsg: "name does not exist or has expired: not found",
			name:      "cannot add record to name that doesn't exist",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.AddRecord(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgAddRecordResponse{}, *res)

			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgDelRecord() {
	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	owner, err := sdk.AccAddressFromBech32(testAddresses[0])
	suite.Require().NoError(err)

	coin := sdk.NewCoin("ujkl", sdk.NewInt(10000000000))
	coins := sdk.NewCoins(coin)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, owner, coins)
	suite.Require().NoError(err)

	err = suite.rnsKeeper.RegisterRNSName(suite.ctx, owner.String(), "BiPhan.jkl", "{}", 2, true)
	suite.Require().NoError(err)

	_, _ = msgSrvr.AddRecord(context, types.NewMsgAddRecord(owner.String(), "BiPhan.jkl", "app", owner.String(), "{}"))

	cases := []struct {
		preRun    func() *types.MsgDelRecord
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgDelRecord {
				return types.NewMsgDelRecord(
					owner.String(),
					"app.BiPhan.jkl",
				)
			},
			expErr: false,
			name:   "successfully deleted record",
		},
		{
			preRun: func() *types.MsgDelRecord {
				return types.NewMsgDelRecord(
					owner.String(),
					"juno.BiPhan.jkl",
				)
			},
			expErr:    true,
			expErrMsg: "record does not exist for this name: unauthorized",
			name:      "can't delete a record that doesn't exist",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.DelRecord(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgDelRecordResponse{}, *res)

			}
		})
	}
}
