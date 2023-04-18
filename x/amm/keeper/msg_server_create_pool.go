package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/amm/types"
)

func (k Keeper) ValidateCreatePoolMsg(ctx sdk.Context, msg *types.MsgCreatePool) error {

	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	params := k.GetParams(ctx)

	minInitLiquidity := params.MinInitLiquidity

	coins := sdk.NewCoins(msg.Coins...)

	if uint32(coins.Len()) > types.MaxPoolDenomCount {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, 
			"too much denoms to create pool")
	}

	for _, c := range coins {
		if c.Amount.LT(sdk.NewInt(int64(minInitLiquidity))) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"provided liquidity is below minimum")
		}
	}

	// Validate AMM id
	_, err := types.GetAMM(msg.AmmId)

	if err != nil {
		return sdkerrors.Wrapf(err, "AMM id (%d) not found",
			msg.AmmId)
	}

	return nil
}

// Creates new liquidity pool with unique coins pairs.
// DecCoins are normalized (converted to smallest unit) and stored as sdk.Coins.
// A provider record is created with contribution and unlock time.
// If pool already exists with coin denoms it returns error.
func (k msgServer) CreatePool(
	goCtx context.Context,
	msg *types.MsgCreatePool,
) (
	*types.MsgCreatePoolResponse,
	error,
) {
	// Pool creation process:
	// 1. Create pool
	// 2. Mint and send pool token
	// 3. Create provider record and engage lock
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.ValidateCreatePoolMsg(ctx, msg); err != nil {
		return nil, err
	}
	poolCoins := sdk.NewCoins(msg.Coins...)

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	pool := NewPool(
		k.GetPoolCountAndIncrement(ctx),
		poolCoins,
		msg.AmmId,
		msg.SwapFeeMulti, 
		msg.PenaltyMulti,
		msg.MinLockDuration,
	)
	k.SetPool(ctx, pool)
	
	k.registerPoolToken(ctx, pool.PoolToken.Denom)
	k.MintAndSendPoolToken(ctx, pool, creator, pool.PoolToken.Amount)

	// Create ProviderRecord
	lockDuration := GetDuration(msg.MinLockDuration)
	err := k.CreateProviderRecord(ctx, creator, pool.Id, lockDuration)

	if err != nil {
		return nil, sdkerrors.Wrapf(
			err,
			"Failed to create liquidity pool. Failed to initialize"+
				" ProviderRecord",
		)
	}

	recordKey := types.ProviderRecordKey(pool.Id, creator.String())

	// Engage lock
	if err := k.EngageLock(ctx, recordKey); err != nil {
		return nil, sdkerrors.Wrapf(
			err,
			"Failed to create liquidity pool. Failed to engage lock",
		)
	}

	// Transfer coins from the creator to module account and give liquidity pool
	// token.
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator,
		types.ModuleName, poolCoins)

	if sdkError != nil {
		return nil, sdkerrors.Wrapf(
			sdkError,
			"failed to create liquidity pool. Failed to retrieve deposit coins "+
				"from sender")
	}

	sdkError = k.MintAndSendPoolToken(ctx, pool, creator, pool.PoolToken.Amount)

	if sdkError != nil {
		return &types.MsgCreatePoolResponse{}, sdkerrors.Wrapf(
			sdkError,
			"Failed to create liquidity pool. Failed to mint and send token",
		)
	}

	EmitPoolCreatedEvent(ctx, creator, pool)

	return &types.MsgCreatePoolResponse{Id: pool.Id}, nil
}
