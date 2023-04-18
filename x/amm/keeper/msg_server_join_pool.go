package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/amm/types"
)

func (k Keeper) validateJoinPoolMsg(ctx sdk.Context, msg *types.MsgJoinPool) error {
	pool, found := k.GetPool(ctx, msg.PoolId)

	if !found {
		return types.ErrLiquidityPoolNotFound
	}

	coins := sdk.NewCoins(msg.Coins...)
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
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Get amount of PToken to send
	pool, _ := k.GetPool(ctx, msg.PoolId)

	coins := sdk.NewCoins(msg.Coins...)

	shares, excess, err := CalcShareJoin(pool, coins)

	if err != nil {
		return nil, err
	}

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Initialize ProviderRecord
	lockDuration := GetDuration(pool.MinLockDuration)

	recordKey := types.ProviderRecordKey(pool.Id, creator.String())
	record, found := k.GetProviderRecord(ctx, recordKey)

	if !found {
		err = k.CreateProviderRecord(ctx, creator, pool.Id, lockDuration)

		if err != nil {
			return nil, err
		}
	} else {
		record.LockDuration = lockDuration.String()
		k.SetProviderRecord(ctx, record)
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

	if err := k.MintAndSendPoolToken(ctx, pool, creator, shares); err != nil {
		return nil, err
	}

	// Update pool liquidity
	poolCoins := sdk.NewCoins(pool.Coins...)
	// Add liquidity to pool
	for _, c := range coins {
		poolCoins = poolCoins.Add(c)
	}

	pool.Coins = poolCoins

	pool.PoolToken.Amount = pool.PoolToken.Amount.Add(shares)

	k.SetPool(ctx, pool)

	EmitPoolJoinedEvent(ctx, creator, pool, coins, pool.MinLockDuration)

	return &types.MsgJoinPoolResponse{}, nil
}
