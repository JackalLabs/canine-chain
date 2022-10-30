package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/lp/types"
)

func (k Keeper) validateExitPool(ctx sdk.Context, msg *types.MsgExitPool) error {
	pool, found := k.GetLPool(ctx, msg.PoolName)

	if !found {
		return types.ErrLiquidityPoolNotFound
	}

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	recordKey := types.LProviderRecordKey(pool.Name, creator.String())
	_, found = k.GetLProviderRecord(ctx, recordKey)

	if !found {
		return types.ErrLProviderRecordNotFound
	}

	if msg.Shares < 0 {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"Burn amount cannot be negative",
		)
	}

	return nil
}

func (k msgServer) ExitPool(goCtx context.Context, msg *types.MsgExitPool) (*types.MsgExitPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	creatorAcc, _ := sdk.AccAddressFromBech32(msg.Creator)

	if err := k.validateExitPool(ctx, msg); err != nil {
		return nil, err
	}

	pool, _ := k.GetLPool(ctx, msg.PoolName)
	poolCoins := sdk.NewCoins(pool.Coins...)
	totalShares, _ := sdk.NewIntFromString(pool.LPTokenBalance)

	// Calculate tokens to return
	// If LPToken is still locked, apply panelty.
	recordKey := types.LProviderRecordKey(pool.Name, creatorAcc.String())
	record, _ := k.GetLProviderRecord(ctx, recordKey)

	unlockTime, _ := StringToTime(record.UnlockTime)

	// This is used to calculate amount of coins to return
	burningAmt := sdk.NewInt(msg.Shares)

	penaltyAmt := sdk.ZeroInt()

	pm, err := sdk.NewDecFromStr(pool.PenaltyMulti)
	if err != nil {
		return nil, fmt.Errorf("failed to convert penalty"+
			" multiplier; saved in invalid format: %s err: %s",
			pool.PenaltyMulti, err)
	}

	if ctx.BlockTime().Before(unlockTime) {
		penaltyAmt = pm.MulInt(burningAmt).RoundInt()
	}

	burningAmt = burningAmt.Sub(penaltyAmt)

	coinsOut, err := CalculatePoolShareBurnReturn(pool, burningAmt)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"Failed to calculate pool share burn return",
		)
	}

	burningCoin := sdk.NewCoin(pool.LptokenDenom, sdk.NewInt(msg.Shares))
	burningCoins := sdk.NewCoins(burningCoin)

	// Transfer LPToken to module
	sdkErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAcc, types.ModuleName, burningCoins)

	if sdkErr != nil {
		return nil, sdkErr
	}

	// Send return coins
	sdkErr = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAcc, coinsOut)

	if sdkErr != nil {
		return nil, sdkErr
	}

	// Burn the LPToken
	sdkErr = k.bankKeeper.BurnCoins(ctx, types.ModuleName, burningCoins)

	if sdkErr != nil {
		return nil, sdkErr
	}

	// Update amount on pool
	poolCoins = poolCoins.Sub(coinsOut)

	totalShares = totalShares.Sub(sdk.NewInt(msg.Shares))

	pool.Coins = poolCoins
	pool.LPTokenBalance = totalShares.String()

	k.SetLPool(ctx, pool)

	EmitPoolExitedEvent(
		ctx,
		creatorAcc,
		pool,
		sdk.NewCoin(pool.LptokenDenom, sdk.NewInt(msg.Shares)),
		coinsOut,
		sdk.NewCoin(pool.LptokenDenom, penaltyAmt))

	return nil, nil
}
