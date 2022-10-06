package keeper

import (
	"context"

	"github.com/jackal-dao/canine/x/lp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	return nil
}

func (k msgServer) JoinPool(goCtx context.Context, msg *types.MsgJoinPool) (*types.MsgJoinPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.validateJoinPoolMsg(ctx, msg)

	if err != nil {
		return nil, err
	}

	// Calculate amount of pool share

	pool, _ := k.GetLPool(ctx, msg.PoolName)

	coins := sdk.NormalizeCoins(msg.Coins)

	shares, err := CalculatePoolShare(pool, coins)

	if err != nil {
		return nil, err
	}


	// Mint and send pool token to msg creator
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

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

	for _, c := range coins {
		// This works by comparing denoms in poolCoins and doing addition on the first
		// denom match.
		poolCoins = poolCoins.Add(c)
	}

	pool.Coins = poolCoins

	poolTotalToken, _ := sdk.NewIntFromString(pool.LPTokenBalance)
	poolTotalToken = poolTotalToken.Add(shares)

	pool.LPTokenBalance = poolTotalToken.String()

	k.SetLPool(ctx, pool)
	
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
		return nil, err
	}

	EmitPoolJoinedEvent(ctx, creator, pool, coins, msg.LockDuration)

	return &types.MsgJoinPoolResponse{}, nil
}
