package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (suite *MintTestSuite) TestBlockMint() {
	suite.SetupTest()
	app, ctx, k := suite.app, suite.ctx, suite.app.MintKeeper
	denom := k.GetParams(ctx).MintDenom
	feeAccount := app.AccountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName)
	feeBalanceBefore, err := app.BankKeeper.Balance(sdk.WrapSDKContext(ctx), &types.QueryBalanceRequest{
		Address: feeAccount.GetAddress().String(),
		Denom:   denom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(sdk.ZeroInt(), feeBalanceBefore.Balance.Amount)
	supplyBefore, err := app.BankKeeper.TotalSupply(sdk.WrapSDKContext(ctx), &types.QueryTotalSupplyRequest{})
	suite.Require().NoError(err)
	suite.Require().True(supplyBefore.Supply.Empty())
	// We have now proved we started with nothing

	k.BlockMint(ctx)

	feeBalanceAfter, err := app.BankKeeper.Balance(sdk.WrapSDKContext(ctx), &types.QueryBalanceRequest{
		Address: feeAccount.GetAddress().String(),
		Denom:   denom,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(sdk.NewInt(335999), feeBalanceAfter.Balance.Amount)
	supplyAfter, err := app.BankKeeper.TotalSupply(sdk.WrapSDKContext(ctx), &types.QueryTotalSupplyRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(1, len(supplyAfter.Supply))
	suite.Require().Equal(sdk.NewInt(419999), supplyAfter.Supply.AmountOf(denom))
	// After BlockMint we now have exactly 3.6JKL in the fee collector account
}

func (suite *MintTestSuite) TestNoProviderBlockMint() {
	suite.SetupTest()
	app, ctx, k := suite.app, suite.ctx, suite.app.MintKeeper

	params := k.GetParams(ctx)
	params.StorageProviderRatio = 0
	k.SetParams(ctx, params)

	denom := k.GetParams(ctx).MintDenom

	pr := k.GetParams(ctx).StorageProviderRatio
	suite.Require().Equal(int64(0), pr)

	feeAccount := app.AccountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName)
	feeBalanceBefore, err := app.BankKeeper.Balance(sdk.WrapSDKContext(ctx), &types.QueryBalanceRequest{
		Address: feeAccount.GetAddress().String(),
		Denom:   denom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(sdk.ZeroInt(), feeBalanceBefore.Balance.Amount)
	supplyBefore, err := app.BankKeeper.TotalSupply(sdk.WrapSDKContext(ctx), &types.QueryTotalSupplyRequest{})
	suite.Require().NoError(err)
	suite.Require().True(supplyBefore.Supply.Empty())
	// We have now proved we started with nothing

	k.BlockMint(ctx)

	feeBalanceAfter, err := app.BankKeeper.Balance(sdk.WrapSDKContext(ctx), &types.QueryBalanceRequest{
		Address: feeAccount.GetAddress().String(),
		Denom:   denom,
	})

	suite.Require().NoError(err)
	suite.Require().Equal(sdk.NewInt(335999), feeBalanceAfter.Balance.Amount)
	supplyAfter, err := app.BankKeeper.TotalSupply(sdk.WrapSDKContext(ctx), &types.QueryTotalSupplyRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(1, len(supplyAfter.Supply))
	suite.Require().Equal(sdk.NewInt(419999), supplyAfter.Supply.AmountOf(denom))
	// After BlockMint we now have exactly 3.6JKL in the fee collector account
}

func (suite *MintTestSuite) TestDecRatios() {
	suite.SetupTest()

	stakerRatio := sdk.NewDec(80).QuoInt64(100)

	s, err := sdk.NewDecFromStr("0.8")
	suite.Require().NoError(err)

	suite.Require().Equal(s, stakerRatio)

	i := stakerRatio.MulInt64(4_200_000)

	suite.Require().Equal(int64(3360000), i.TruncateInt64())
}
