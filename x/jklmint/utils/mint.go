package utils

import sdk "github.com/cosmos/cosmos-sdk/types"

func int64ToDec(i int64) sdk.Dec {
	return sdk.NewDecFromInt(sdk.NewInt(i))
}

func GetMintForBlock(mintedLastBlock int64, blocksPerYear int64, mintDecrease int64) int64 {
	lastBlockTokens := int64ToDec(mintedLastBlock)
	blockPerYearDec := int64ToDec(blocksPerYear)
	decrease := int64ToDec(mintDecrease)
	return lastBlockTokens.Sub(decrease.Quo(blockPerYearDec)).TruncateInt64()
}

func GetTokensOwed(totalTokens int64, ratio int64) int64 {
	tokensDec := int64ToDec(totalTokens)
	ratioDec := int64ToDec(ratio)
	r := ratioDec.Quo(int64ToDec(100))
	return tokensDec.Mul(r).TruncateInt64()
}
