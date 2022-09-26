package keeper

import (
	"errors"
	"fmt"

	"github.com/jackal-dao/canine/x/lp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Calculate amount of pool coins to deposit to get desired amount of LPToken.
// This assumes that pool coins are normalized.
// Example: pool has coin x and y.
// 	This function returns amount of x and y coins to deposit in order to get
// 	desired amount of LPToken.
func CoinsToDepositForLPToken(pool types.LPool, desiredAmount sdk.Int) (sdk.Coins, error) {
	totalLPtoken, _ := sdk.NewIntFromString(pool.LPTokenBalance)
	// Convert [] coin to sdk.Coins
	poolCoins := sdk.NewCoins(pool.Coins...)

	result := sdk.NewCoins()

	// Let a set of coins in pool be poolCoins.
	// For each coins in poolCoins,
	// amount to deposit is: amtToDep = cAmtInPool * desiredAmount / totalLPToken
	for _, c := range poolCoins {
		amtToDep := poolCoins.AmountOf(c.GetDenom()).Mul(desiredAmount).Quo(totalLPtoken)

		result = result.Add(sdk.NewCoin(c.GetDenom(), amtToDep))
	}

	return result, nil
}

// Calculate amount of other pool coins to deposit given coin x to deposit same coin values.
// Example: Pool has coin x and y.
//	User wants to deposit x but can't figure out how much y to deposit to make valid a
//	liquidity pair.
//	This function returns amount of y to make the valid liquidity pair.
func MakeValidPair(pool types.LPool, deposit sdk.Coin) (sdk.Coins, error) {

	poolCoins := sdk.NewCoins(pool.Coins...)
	totalLPToken, _ := sdk.NewIntFromString(pool.LPTokenBalance)

	// Let deposit denom be x.
	xDenom := deposit.GetDenom()
	// Then, share = totalLPToken * xAmtInDeposit / xAmtInPool
	xAmtInPool := poolCoins.AmountOf(xDenom)
	share := totalLPToken.Mul(deposit.Amount).Quo(xAmtInPool)

	// So, now we know amount of shares.
	// Use the share to calculate coin amount for the rest of pool denoms.
	// Let a set of pool coins be y where coin x is not in y (the set does not contain coin x).
	// e.g. poolDenoms = {a, b, c}
	//		  x = {a}
	// 	  y = {b, c}
	// Then, for every element (call it e) in y, calculate equivalent amount (remember coin{denom:str, amount:sdk.Int}).
	// eAmtToDeposit = share / totalLPToken * eAmtInPool

	setU := sdk.NewCoins(pool.Coins...)
	// Removing x from the set to get y.
	// e.g. setU = {10x, 20b, 30c}
	// We want a set without x. So, we subtract the amount of x in setU 
	// from setU.
	// setY = {10x, 20b, 30c} - {10x} = {20b, 30c}
	xCoins := sdk.NewCoins(sdk.NewCoin(xDenom, xAmtInPool))
	setY := setU.Sub(xCoins)

	result := sdk.NewCoins()

	for _, e := range setY {
		eAmtToDeposit := share.Quo(totalLPToken).Mul(poolCoins.AmountOf(e.GetDenom()))
		coin := sdk.NewCoin(e.GetDenom(), eAmtToDeposit)
		// Record the result
		result = result.Add(coin)
	}

	return result, nil
}

// Calculate amount of shares (sdk.Int) to be given out based on deposits.
// If provided deposits are not same value, it'll return error.
func CalculatePoolShare(pool types.LPool, depositCoins sdk.Coins) (sdk.Int, error) {

	shareAmount := sdk.ZeroDec()

	// Check if pool is being initiated
	if pool.LPTokenBalance == "" {

		x := sdk.OneDec()

		// Initial pool token is sqrt(coin0 * ... * coinN)
		for _, c := range depositCoins {
			x = x.MulInt(depositCoins.AmountOf(c.GetDenom()))
		}

		amount, err := x.ApproxSqrt()

		if err != nil {
			return sdk.ZeroInt(), sdkerrors.Wrapf(err,
				"Error occured while calculating pool share",
			)
		}

		shareAmount = amount
	} else {
		poolCoins := sdk.NewCoins(pool.Coins...)
		// WARNING: This is not be safe.
		totalLPToken, _ := sdk.NewDecFromStr(pool.LPTokenBalance)

		if len(poolCoins) != len(depositCoins) {
			return sdk.ZeroInt(), sdkerrors.Wrapf(sdkerrors.ErrConflict,
				"Number of pool coins does not match number of deposit coins",
			)
		}

		var shares []sdk.Dec
		shares = make([]sdk.Dec, 0)

		for _, x := range depositCoins {
			// Calculate shares and append it.
			// Get amount of coin x.
			totalXInPool := sdk.NewDecFromInt(poolCoins.AmountOf(x.GetDenom()))
			// NOTE: Make this code more readable.
			// share = totalLPToken * xAmtInDeposit / totalXInPool
			share := totalLPToken.MulInt(depositCoins.AmountOf(x.GetDenom())).Quo(totalXInPool)
			shares = append(shares, share)
		}

		mul := shares[0].MulInt64(int64(len(shares))).TruncateInt()

		add := sdk.NewDec(0)

		for _, v := range shares {
			add = add.Add(v)
		}

		if !mul.Equal(add.TruncateInt()) {
			return sdk.ZeroInt(), sdkerrors.Wrapf(
				sdkerrors.ErrInvalidCoins,
				"Deposits are not same value. Added: %d, multiplied: %d, share[0]: %d, share[1]: %d",
				mul,
				add.TruncateInt64(),
				shares[0].TruncateInt64(),
				shares[1].TruncateInt64(),
			)
		}

		shareAmount = shares[0]
	}

	return shareAmount.TruncateInt(), nil
}

func CalculatePoolShareBurnReturn(pool types.LPool, burnAmt sdk.Int) (sdk.Coins, error) {
	
	poolCoins := sdk.NewCoins(pool.Coins...)

	totalLPToken, ok := sdk.NewIntFromString(pool.LPTokenBalance)

	if !ok {
		return nil, errors.New(fmt.Sprintf("Failed to convert LPTokenBalance to" +
			" sdk.Int: %s", pool.LPTokenBalance))
	}

	if burnAmt.GT(totalLPToken) {
		return nil, errors.New(fmt.Sprint("Burn amount is greater than total" + 
			" liquidity pool token exists"))
	}

	returns := sdk.NewCoins()

	// Calculate pool coin values in respect to amount of shares burned.
	// return = burnAmt * coinInPool / totalLPTokens
	for _, coin := range poolCoins {
		cAmtInPool := poolCoins.AmountOf(coin.GetDenom())	
		result := burnAmt.Mul(cAmtInPool).Quo(totalLPToken)
		resultCoin := sdk.NewCoin(coin.GetDenom(), result)
		returns = returns.Add(resultCoin)
	}

	return returns, nil
}
