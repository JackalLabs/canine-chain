package keeper

import (
	"strconv"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/amm/types"
)

// New event that has current state of pool balance
func newPoolBalanceEvent(pool types.Pool) sdk.Event {
	return sdk.NewEvent(
		types.TypedEventPoolInfo,
		sdk.NewAttribute(types.AttrKeyPoolId, strconv.FormatUint(pool.Id, 10)),
		sdk.NewAttribute(
			types.AttrKeyPoolBalance,
			sdk.NewCoins(pool.Coins...).String()),
	)
}

func EmitPoolCreatedEvent(ctx sdk.Context, sender sdk.AccAddress, pool types.Pool) {
	if ctx.EventManager() == nil {
		return
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{newPoolCreatedEvent(sender, pool), newPoolBalanceEvent(pool)},
	)
}

func newPoolCreatedEvent(sender sdk.AccAddress, pool types.Pool) sdk.Event {
	return sdk.NewEvent(
		types.TypedEventPoolCreated,
		sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttrValueModule),
		sdk.NewAttribute(types.AttrKeyPoolId, strconv.FormatUint(pool.Id, 10)),
		sdk.NewAttribute(types.AttrKeySwapFeeMulti, pool.SwapFeeMulti),
		sdk.NewAttribute(types.AttrKeyPenaltyMulti, pool.PenaltyMulti),
		sdk.NewAttribute(types.AttrKeyPoolTokenDenom, pool.PoolToken.Denom),
		sdk.NewAttribute(
			types.AttrKeyLockDuration,
			ToSecondsStr(pool.MinLockDuration)),
	)
}

func EmitPoolJoinedEvent(
	ctx sdk.Context,
	sender sdk.AccAddress,
	pool types.Pool,
	coins sdk.Coins,
	lockDuration int64,
) {
	if ctx.EventManager() == nil {
		return
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			newPoolJoinedEvent(sender, pool, coins, lockDuration),
			newPoolBalanceEvent(pool)},
	)
}

func newPoolJoinedEvent(
	sender sdk.AccAddress,
	pool types.Pool,
	coins sdk.Coins,
	lockDuration int64,
) sdk.Event {
	return sdk.NewEvent(
		types.TypedEventPoolJoined,
		sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttrValueModule),
		sdk.NewAttribute(types.AttrKeyPoolId, strconv.FormatUint(pool.Id, 10)),
		sdk.NewAttribute(types.AttrKeyCoinsIn, coins.String()),
		sdk.NewAttribute(types.AttrKeyLockDuration, ToSecondsStr(lockDuration)),
	)
}

func EmitPoolExitedEvent(
	ctx sdk.Context,
	sender sdk.AccAddress,
	pool types.Pool,
	amount sdk.Coin,
	coinsOut sdk.Coins,
	penaltyFee sdk.Coin,
) {
	if ctx.EventManager() == nil {
		return
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			newPoolExitedEvent(sender, pool, amount, coinsOut, penaltyFee),
			newPoolBalanceEvent(pool)},
	)
}

func newPoolExitedEvent(
	sender sdk.AccAddress,
	pool types.Pool,
	amount sdk.Coin,
	coinsOut sdk.Coins,
	penaltyFee sdk.Coin,
) sdk.Event {
	return sdk.NewEvent(
		types.TypedEventPoolExited,
		sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttrValueModule),
		sdk.NewAttribute(types.AttrKeyPoolId, strconv.FormatUint(pool.Id, 10)),
		sdk.NewAttribute(types.AttrKeyCoinsIn, amount.String()),
		sdk.NewAttribute(types.AttrKeyCoinsOut, coinsOut.String()),
		sdk.NewAttribute(types.AttrKeyPenaltyFee, penaltyFee.String()),
	)
}

func EmitCoinSwappedEvent(
	ctx sdk.Context,
	sender sdk.AccAddress,
	pool types.Pool,
	coinsIn sdk.Coins,
	coinsOut sdk.Coins,
	swapFee sdk.Coins,
	protocolFee sdk.Coins,
) {
	if ctx.EventManager() == nil {
		return
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			newCoinSwappedEvent(
				sender, 
				pool, 
				coinsIn, 
				coinsOut, 
				swapFee, 
				protocolFee),
			newPoolBalanceEvent(pool)},
	)
}

func newCoinSwappedEvent(
	sender sdk.AccAddress,
	pool types.Pool,
	coinsIn sdk.Coins,
	coinsOut sdk.Coins,
	swapFee sdk.Coins,
	protocolFee sdk.Coins,
) sdk.Event {
	return sdk.NewEvent(
		types.TypedEventCoinSwapped,
		sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttrValueModule),
		sdk.NewAttribute(types.AttrKeyCoinsIn, coinsIn.String()),
		sdk.NewAttribute(types.AttrKeyCoinsOut, coinsOut.String()),
		sdk.NewAttribute(types.AttrKeySwapFee, swapFee.String()),
		sdk.NewAttribute(types.AttrKeyProtocolFee, protocolFee.String()),
	)
}

func EmitCoinSwapFailedEvent(
	ctx sdk.Context,
	sender sdk.AccAddress,
	pool types.Pool,
	coinsIn sdk.Coins,
	coinsOut sdk.Coins,
	minCoinsOut sdk.Coins,
) {
	if ctx.EventManager() == nil {
		return
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			newCoinSwapFailedEvent(sender, pool, coinsIn, coinsOut, minCoinsOut),
			newPoolBalanceEvent(pool)},
	)
}

func newCoinSwapFailedEvent(
	sender sdk.AccAddress,
	pool types.Pool,
	coinsIn sdk.Coins,
	coinsOut sdk.Coins,
	minCoinsOut sdk.Coins,
) sdk.Event {
	return sdk.NewEvent(
		types.TypedEventCoinSwapFailed,
		sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttrValueModule),
		sdk.NewAttribute(types.AttrKeyCoinsIn, coinsIn.String()),
		sdk.NewAttribute(types.AttrKeyCoinsOut, coinsOut.String()),
		sdk.NewAttribute(types.AttrKeyMinCoinsOut, minCoinsOut.String()),
	)
}
