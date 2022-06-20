package jklmint

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/jklmint/keeper"
	"github.com/jackal-dao/canine/x/jklmint/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	var validator_ratio int64 = 4
	var miner_ratio int64 = 6

	// mint coins, update supply
	mintedCoin := sdk.NewCoin("ujkl", sdk.NewInt(validator_ratio*1000000))
	mintedCoins := sdk.NewCoins(mintedCoin)

	err := k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// send the minted coins to the fee collector account
	err = k.AddCollectedFees(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// mint coins, update supply
	minerCoin := sdk.NewCoin("ujkl", sdk.NewInt(miner_ratio*1000000))
	minerCoins := sdk.NewCoins(minerCoin)

	err = k.MintCoins(ctx, minerCoins)
	if err != nil {
		panic(err)
	}

	err = k.SendToMiners(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	totalCoin := minerCoin.Amount.Int64() + mintedCoin.Amount.Int64()

	if mintedCoin.Amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(totalCoin), "minted_tokens")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(sdk.AttributeKeyAmount, fmt.Sprintf("%d", totalCoin)),
		),
	)
}
