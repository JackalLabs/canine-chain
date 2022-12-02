package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/jklmint/types"
	storeTypes "github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k Keeper) BlockMint(ctx sdk.Context) {
	var validatorRatio int64 = 4
	var providerRatio int64 = 6

	denom := k.GetParams(ctx).MintDenom

	totalCoin := sdk.NewCoin(denom, sdk.NewInt((validatorRatio)*1000000))

	// mint coins, update supply
	mintedCoin := sdk.NewCoin(denom, sdk.NewInt(validatorRatio*1000000))
	mintedCoins := sdk.NewCoins(mintedCoin)

	err := k.MintCoins(ctx, sdk.NewCoins(totalCoin))
	if err != nil {
		panic(err)
	}

	// send the minted coins to the fee collector account
	err = k.AddCollectedFees(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// mint coins, update supply
	providerCoin := sdk.NewCoin(denom, sdk.NewInt(providerRatio*1000000))
	providerCoins := sdk.NewCoins(providerCoin)

	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, storeTypes.ModuleName, providerCoins)
	if err != nil {
		panic(err)
	}

	if mintedCoin.Amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(totalCoin.Amount.Int64()), "minted_tokens")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(sdk.AttributeKeyAmount, fmt.Sprintf("%d", totalCoin.Amount.Int64())),
		),
	)
}
