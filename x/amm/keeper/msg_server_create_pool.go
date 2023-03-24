package keeper

import (
	"context"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/amm/types"
)

/*
	 Checks if MsgCreatePool is a valid request
		It validates:
			1. Coin format
			2. Max denom count allowed in a pool
			3. Coin deposit is more than minimum pool creation amount
			4. Able to create pool name
			5.	The new pool is not a duplicate
			6. AMM model id is valid and exists
			7. Swap fee, lock duration and penalty percentage are non-negative
*/
func (k Keeper) ValidateCreatePoolMsg(ctx sdk.Context, msg *types.MsgCreatePool) error {

	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	params := k.GetParams(ctx)

	minInitPoolDeposit := params.MinInitPoolDeposit
	maxPoolDenomCount := params.MaxPoolDenomCount

	coins := sdk.NormalizeCoins(msg.Coins)

	if uint32(coins.Len()) > maxPoolDenomCount {
		return sdkerrors.Wrap(errors.New(fmt.Sprintf(
			"pool can only balance %d coins", maxPoolDenomCount)),
			sdkerrors.ErrInvalidRequest.Error())
	}

	for _, c := range coins {
		if c.Amount.LT(sdk.NewInt(int64(minInitPoolDeposit))) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"You need to deposit at least %v amount to create a liquidity pool",
				minInitPoolDeposit)
		}
	}

	poolName := generatePoolName(coins)

	_, found := k.GetPool(ctx, poolName)

	if found {
		return sdkerrors.Wrap(types.ErrLiquidityPoolExists,
			sdkerrors.ErrInvalidRequest.Error())
	}

	// Validate AMM id
	_, err := types.GetAMM(msg.Amm_Id)

	if err != nil {
		return sdkerrors.Wrapf(err, "AMM with id %v does not exist",
			msg.Amm_Id)
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

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.ValidateCreatePoolMsg(ctx, msg); err != nil {
		return nil, sdkerrors.Wrapf(
			err,
			"failed to create a liquidity pool")
	}

	pool := k.NewPool(ctx, msg)

	creatorAccAddr, _ := sdk.AccAddressFromBech32(msg.Creator)
	normCoins := sdk.NormalizeCoins(msg.Coins)

	shareAmount, err := CalculatePoolShare(pool, normCoins)

	if err != nil {
		return nil, sdkerrors.Wrapf(
			err,
			"Failed to create liquidity pool, error during CalculatePoolShare()")
	}

	pool.Coins = normCoins
	pool.PTokenBalance = shareAmount.String()

	k.SetPool(ctx, pool)

	// Create ProviderRecord
	lockDuration := GetDuration(msg.MinLockDuration)
	err = k.InitProviderRecord(ctx, creatorAccAddr, pool.Name, lockDuration)

	if err != nil {
		k.RemovePool(ctx, pool.Index)

		return nil, sdkerrors.Wrapf(
			err,
			"Failed to create liquidity pool. Failed to initialize"+
				" ProviderRecord",
		)
	}

	recordKey := types.ProviderRecordKey(pool.Name, creatorAccAddr.String())

	// Engage lock
	if err := k.EngageLock(ctx, recordKey); err != nil {
		k.RemovePool(ctx, pool.Index)
		k.EraseProviderRecord(ctx, creatorAccAddr, pool.Name)

		return nil, sdkerrors.Wrapf(
			err,
			"Failed to create liquidity pool. Failed to engage lock",
		)
	}

	// Transfer coins from the creator to module account and give liquidity pool
	// token.
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAccAddr,
		types.ModuleName, normCoins)

	if sdkError != nil {
		k.RemovePool(ctx, pool.Index)
		k.EraseProviderRecord(ctx, creatorAccAddr, pool.Name)

		return nil, sdkerrors.Wrapf(
			sdkError,
			"failed to create liquidity pool. Failed to retrieve deposit coins "+
				"from sender")
	}

	sdkError = k.MintAndSendPToken(ctx, pool, creatorAccAddr, shareAmount)

	if sdkError != nil {
		k.RemovePool(ctx, pool.Index)
		k.EraseProviderRecord(ctx, creatorAccAddr, pool.Name)

		return &types.MsgCreatePoolResponse{}, sdkerrors.Wrapf(
			sdkError,
			"Failed to create liquidity pool. Failed to mint and send token",
		)
	}

	EmitPoolCreatedEvent(ctx, creatorAccAddr, pool)

	return &types.MsgCreatePoolResponse{Id: pool.Index}, nil
}
