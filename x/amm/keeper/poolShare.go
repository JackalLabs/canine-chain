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

// Calculate amount of pool token to be given out based on provided liquidity.
// If provided liquidity are not same value, it'll return error.
func CalcShareForJoin(pool types.Pool, liquidity sdk.Coins) (shareAmt sdk.Int, excess sdk.Coins, err error) {

	if pool.PoolToken.Amount.Equal(sdk.ZeroInt()) {
		err = errors.New("pool has zero outstanding pool tokens")
		return
	}

	poolCoins := sdk.NewCoins(pool.Coins...)

	if !liquidity.DenomsSubsetOf(poolCoins) {
		err = errors.New("provided liquidity is not pool coins")
		return
	}

	minRatio := sdk.ZeroDec()
	maxRatio := sdk.ZeroDec()

	coinRatio := make([]sdk.Dec, len(poolCoins))

	for i, coin := range liquidity {
		r := coin.Amount.ToDec().QuoInt(poolCoins.AmountOf(coin.Denom))
		coinRatio[i] = r
		if minRatio.GT(r){
			minRatio = r
		}
		if maxRatio.LT(r){
			maxRatio = r
		}
	}

	// use min ratio as a base to calculate other coins used
	if !minRatio.Equal(maxRatio) {
		for i, coin := range liquidity {
			if coinRatio[i].Equal(minRatio){
				continue
			}

			useAmt := minRatio.MulInt(poolCoins.AmountOf(coin.Denom)).Ceil().TruncateInt()
			excessAmt := coin.Amount.Sub(useAmt)
			if !excessAmt.IsZero(){
				excess = excess.Add(sdk.NewCoin(coin.Denom, excessAmt))
			}
		}
	}

	return minRatio.MulInt(pool.PoolToken.Amount).TruncateInt(), excess, nil
}

func CalcShareExit(pool types.Pool, exitAmt sdk.Int) (sdk.Coins, error) {

	ratio := sdk.NewDecFromInt(exitAmt).QuoInt(pool.PoolToken.Amount)
	retCoins := sdk.Coins{}

	for _, coin := range pool.Coins {
		amt := ratio.MulInt(coin.Amount).TruncateInt()
		if amt.IsZero(){
			continue
		}
		if amt.GTE(coin.Amount){
			return sdk.Coins{}, errors.New("exit amount is too high")
		}
		retCoins = retCoins.Add(sdk.NewCoin(coin.Denom, amt))
	}

	return retCoins, nil
}
