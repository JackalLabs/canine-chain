package keeper

import (
	"context"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/amm/types"
)

func (k Keeper) validateExitPool(ctx sdk.Context, msg *types.MsgExitPool) error {
	pool, found := k.GetPool(ctx, msg.PoolId)

	if !found {
		return types.ErrLiquidityPoolNotFound
	}

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	recordKey := types.ProviderRecordKey(pool.Id, creator.String())
	_, found = k.GetProviderRecord(ctx, recordKey)

	if !found {
		return fmt.Errorf("provider record not found")
	}

	if msg.Amount< 0 {
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

	pool, _ := k.GetPool(ctx, msg.PoolId)
	poolCoins := sdk.NewCoins(pool.Coins...)
	totalShares := pool.PoolToken.Amount

	// Calculate tokens to return
	// If PToken is still locked, apply panelty.
	recordKey := types.ProviderRecordKey(pool.Id, creatorAcc.String())
	record, _ := k.GetProviderRecord(ctx, recordKey)

	unlockTime, _ := StringToTime(record.UnlockTime)

	// This is used to calculate amount of coins to return
	burningAmt := sdk.NewInt(msg.Amount)

	penaltyAmt := sdk.ZeroInt()



	if ctx.BlockTime().Before(unlockTime) {
		pm, err := sdk.NewDecFromStr(pool.PenaltyMulti)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Failed to convert penalty"+
				" multiplier; saved in invalid format: %s err: %s",
				pool.PenaltyMulti, err))
		}
		penaltyAmt = pm.MulInt(burningAmt).RoundInt()
	}

	burningAmt = burningAmt.Sub(penaltyAmt)

	coinsOut, err := CalcShareExit(pool, burningAmt)

	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"Failed to calculate pool share burn return",
		)
	}

	burningCoin := sdk.NewCoin(pool.PoolToken.Denom, sdk.NewInt(msg.Amount))
	burningCoins := sdk.NewCoins(burningCoin)

	// Transfer PToken to module
	sdkErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAcc, types.ModuleName, burningCoins)

	if sdkErr != nil {
		return nil, sdkErr
	}

	// Send return coins
	sdkErr = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAcc, coinsOut)

	if sdkErr != nil {
		return nil, sdkErr
	}

	// Burn the PToken
	sdkErr = k.bankKeeper.BurnCoins(ctx, types.ModuleName, burningCoins)

	if sdkErr != nil {
		return nil, sdkErr
	}

	// Update amount on pool
	poolCoins = poolCoins.Sub(coinsOut)

	pool.Coins = poolCoins
	pool.PoolToken.Amount = totalShares.Sub(sdk.NewInt(msg.Amount))

	k.SetPool(ctx, pool)

	EmitPoolExitedEvent(
		ctx,
		creatorAcc,
		pool,
		sdk.NewCoin(pool.PoolToken.Denom, sdk.NewInt(msg.Amount)),
		coinsOut,
		sdk.NewCoin(pool.PoolToken.Denom, penaltyAmt))

	return nil, nil
}
