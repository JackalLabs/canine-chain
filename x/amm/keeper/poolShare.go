package keeper

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/amm/types"
)

// Calculate amount of pool coins to deposit to get desired amount of PToken.
// This assumes that pool coins are normalized.
// Example: pool has coin x and y.
//
//	This function returns amount of x and y coins to deposit in order to get
//	desired amount of PToken.
func CoinsToDepositForPToken(pool types.Pool, desiredAmount sdk.Int) (sdk.Coins, error) {
	totalPtoken, _ := sdk.NewIntFromString(pool.PTokenBalance)

	if totalPtoken.IsZero() {
		return sdk.NewCoins(), errors.New(
			"pool.PTokenBalance is zero, will not proceed to prevent" +
				" division by zero")
	}
	// Convert [] coin to sdk.Coins
	poolCoins := sdk.NewCoins(pool.Coins...)

	result := sdk.NewCoins()

	// Let a set of coins in pool be poolCoins.
	// For each coins in poolCoins,
	// amount to deposit is: amtToDep = cAmtInPool * desiredAmount / totalPToken
	for _, c := range poolCoins {
		amtToDep := poolCoins.AmountOf(c.GetDenom()).Mul(desiredAmount).Quo(totalPtoken)

		result = result.Add(sdk.NewCoin(c.GetDenom(), amtToDep))
	}

	return result, nil
}

// Calculate amount of other pool coins to deposit given coin x to deposit same coin values.
// Example: Pool has coin x and y.
//
//	User wants to deposit x but can't figure out how much y to deposit to make valid a
//	liquidity pair.
//	This function returns amount of y to make the valid liquidity pair.
func MakeValidPair(pool types.Pool, deposit sdk.Coin) (sdk.Coins, error) {

	poolCoins := sdk.NewCoins(pool.Coins...)
	totalPToken, _ := sdk.NewIntFromString(pool.PTokenBalance)
	if totalPToken.IsZero() {
		return sdk.NewCoins(), errors.New(
			"pool.PTokenBalance is zero, will not proceed to prevent" +
				" division by zero")
	}

	// Let deposit denom be x.
	xDenom := deposit.GetDenom()
	// Then, share = totalPToken * xAmtInDeposit / xAmtInPool
	xAmtInPool := poolCoins.AmountOf(xDenom)
	if xAmtInPool.IsZero() {
		return sdk.NewCoins(), errors.New(
			fmt.Sprintf("coin %s in pool is zero, will not proceed to prevent"+
				" division by zero",
				xDenom))
	}
	share := totalPToken.Mul(deposit.Amount).Quo(xAmtInPool)

	// So, now we know amount of shares.
	// Use the share to calculate coin amount for the rest of pool denoms.
	// Let a set of pool coins be y where coin x is not in y (the set does not contain coin x).
	// e.g. poolDenoms = {a, b, c}
	//		  x = {a}
	// 	  y = {b, c}
	// Then, for every element (call it e) in y, calculate equivalent amount (remember coin{denom:str, amount:sdk.Int}).
	// eAmtToDeposit = share / totalPToken * eAmtInPool

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
		eAmtToDeposit := share.Quo(totalPToken).Mul(poolCoins.AmountOf(e.GetDenom()))
		coin := sdk.NewCoin(e.GetDenom(), eAmtToDeposit)
		// Record the result
		result = result.Add(coin)
	}

	return result, nil
}

// Calculate amount of PToken (sdk.Int) to be given out based on deposits.
// If provided deposits are not same value, it'll return error.
func CalculatePoolShare(pool types.Pool, depositCoins sdk.Coins) (sdk.Int, error) {

	// Check if pool is being initiated
	if pool.PTokenBalance == "" {

		// Using sdk.Dec to use sqrt()
		x := sdk.OneDec()

		// Initial pool token is sqrt(coin0 * ... * coinN)
		for _, c := range depositCoins {
			x = x.MulInt(depositCoins.AmountOf(c.GetDenom()))
		}

		amount, err := x.ApproxSqrt()

		if err != nil {
			return sdk.ZeroInt(), err
		}

		return amount.TruncateInt(), nil

	} else {
		poolCoins := sdk.NewCoins(pool.Coins...)

		// Get total PTokens and convert it to sdk.Dec
		totalPToken, ok := sdk.NewIntFromString(pool.PTokenBalance)

		if !ok {
			return sdk.ZeroInt(), errors.New("Failed to convert" +
				" pool.PTokenBalance to sdk.Int")
		}

		if !depositCoins.DenomsSubsetOf(poolCoins) {
			return sdk.ZeroInt(), sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins,
				"Input Coins are not subset of pool coins."+
					" Pool: %s, Input: %s",
				poolCoins.String(), depositCoins.String(),
			)
		}

		// Using Coin type to return precise error
		var coinShares []sdk.Coin
		coinShares = make([]sdk.Coin, 0)

		for _, x := range depositCoins {
			// Calculate shares and append it.
			// Get amount of coin x.
			totalXInPool := poolCoins.AmountOf(x.GetDenom())

			if totalXInPool.IsZero() {
				return sdk.ZeroInt(), errors.New(fmt.Sprintf("Zero amount of coin %s,"+
					" will not proceed to prevent division by zero",
					x.GetDenom(),
				))
			}

			// share = totalPToken * xAmtInDeposit / totalXInPool
			share := totalPToken.Mul(depositCoins.AmountOf(x.GetDenom())).Quo(totalXInPool)
			coinShare := sdk.NewCoin(x.GetDenom(), share)
			coinShares = append(coinShares, coinShare)
		}

		// Check if all input coins are same value
		// If that is true, all share amount of coins should be same
		// shareX0 == shareX1 == shareX2 ... shareXn
		for _, x := range coinShares {
			if !x.Amount.Equal(coinShares[0].Amount) {
				return sdk.ZeroInt(), errors.New(
					fmt.Sprintf("Same value of coin not provided. denom: %s,"+
						" value: %s",
						x.Denom,
						x.Amount.String(),
					))
			}
		}

		return coinShares[0].Amount, nil
	}
}

func CalculatePoolShareBurnReturn(pool types.Pool, burnAmt sdk.Int) (sdk.Coins, error) {

	poolCoins := sdk.NewCoins(pool.Coins...)

	totalPToken, ok := sdk.NewIntFromString(pool.PTokenBalance)

	if !ok {
		return nil, errors.New(fmt.Sprintf("Failed to convert PTokenBalance to"+
			" sdk.Int: %s", pool.PTokenBalance))
	}

	if totalPToken.IsZero() {
		return nil, errors.New(fmt.Sprintf("Total Ptoken is zero." +
			" Will not proceed to prevent divide by zero"))
	}

	if burnAmt.GT(totalPToken) {
		return nil, errors.New(fmt.Sprint("Burn amount is greater than total" +
			" Ptoken that exists"))
	}

	returns := sdk.NewCoins()

	// Calculate pool coin values in respect to amount of shares burned.
	// return = burnAmt * coinInPool / totalPTokens
	for _, coin := range poolCoins {
		cAmtInPool := poolCoins.AmountOf(coin.GetDenom())
		result := burnAmt.Mul(cAmtInPool).Quo(totalPToken)
		resultCoin := sdk.NewCoin(coin.GetDenom(), result)
		returns = returns.Add(resultCoin)
	}

	return returns, nil
}
