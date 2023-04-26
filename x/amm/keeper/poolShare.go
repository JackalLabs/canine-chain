package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/amm/types"
)

// Calculate amount of pool token to be given out based on provided liquidity.
// If provided liquidity are not same value, it'll return error.
func CalcShareJoin(poolToken sdk.Coin, poolCoins, liquidity sdk.Coins) (shareAmt sdk.Int, excess sdk.Coins, err error) {

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

	return minRatio.MulInt(poolToken.Amount).TruncateInt(), excess, nil
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
