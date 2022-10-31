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
	suite.Require().Equal(sdk.NewInt(4000000), feeBalanceAfter.Balance.Amount)
	supplyAfter, err := app.BankKeeper.TotalSupply(sdk.WrapSDKContext(ctx), &types.QueryTotalSupplyRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(1, len(supplyAfter.Supply))
	suite.Require().Equal(sdk.NewInt(4000000), supplyAfter.Supply.AmountOf(denom))
	// After BlockMint we now have exactly 4JKL in the fee collector account
}
