package keeper

import (
	"context"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/lp/types"
)

func (k Keeper) validateJoinPoolMsg(ctx sdk.Context, msg *types.MsgJoinPool) error {
	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	pool, found := k.GetLPool(ctx, msg.PoolName)

	if !found {
		return types.ErrLiquidityPoolNotFound
	}

	coins := sdk.NormalizeCoins(msg.Coins)
	poolCoins := sdk.NewCoins(pool.Coins...)

	if !coins.DenomsSubsetOf(poolCoins) {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInvalidCoins,
			"Deposit coins are not pool coins",
		)
	}

	// Check if existing record unlock time is later than this message
	recordKey := types.LProviderRecordKey(pool.Name, msg.Creator)
	record, found := k.GetLProviderRecord(ctx, recordKey)
	if found {
		oldUnlockTime, _ := StringToTime(record.UnlockTime)

		newDuration := GetDuration(msg.LockDuration)

		newUnlockTime := ctx.BlockTime().Add(newDuration)

		if newUnlockTime.Before(oldUnlockTime) {
			return errors.New(
				fmt.Sprintf("new unlock time must be after old." +
				" new: %s, old %s", TimeToString(newUnlockTime), 
				TimeToString(oldUnlockTime))) 
		}
	}

	return nil
}

func (k msgServer) JoinPool(goCtx context.Context, msg *types.MsgJoinPool) (*types.MsgJoinPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.validateJoinPoolMsg(ctx, msg)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}


	// Get amount of LPToken to send
	pool, _ := k.GetLPool(ctx, msg.PoolName)

	coins := sdk.NormalizeCoins(msg.Coins)

	shares, err := CalculatePoolShare(pool, coins)

	if err != nil {
		return nil, err
	}

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Initialize LProviderRecord
	lockDuration := GetDuration(msg.LockDuration)

	recordKey := types.LProviderRecordKey(pool.Name, creator.String())
	record, found := k.GetLProviderRecord(ctx, recordKey)

	if !found {
		err = k.InitLProviderRecord(ctx, creator, pool.Name, lockDuration)

		if err != nil {
			return nil, err
		}
	} else {
		record.LockDuration = lockDuration.String()
		k.SetLProviderRecord(ctx, record)
	}

	err = k.EngageLock(ctx, recordKey)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Transfer liquidity from the creator account to module account
	sdkErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, coins)

	if sdkErr != nil {
		return nil, sdkErr
	}

	if err := k.MintAndSendLPToken(ctx, pool, creator, shares); err != nil {
		return nil, err
	}

	// Update pool liquidity
	poolCoins := sdk.NewCoins(pool.Coins...)
	// Add liquidity to pool
	for _, c := range coins {
		poolCoins = poolCoins.Add(c)
	}

	pool.Coins = poolCoins

	// Update LPTokens
	netLPToken, _ := sdk.NewIntFromString(pool.LPTokenBalance)
	netLPToken = netLPToken.Add(shares)

	pool.LPTokenBalance = netLPToken.String()

	k.SetLPool(ctx, pool)
	
	EmitPoolJoinedEvent(ctx, creator, pool, coins, msg.LockDuration)

	return &types.MsgJoinPoolResponse{}, nil
}
