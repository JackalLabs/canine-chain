package keeper

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackal-dao/canine/x/lp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) validateSwapMsg(ctx sdk.Context, msg *types.MsgSwap) error {
	if err := msg.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	pool, found := k.GetLPool(ctx, msg.PoolName)

	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			types.ErrLiquidityPoolNotFound.Error())
	}

	// Convert DecCoin to Normalized Coin
	coin, _ := sdk.NormalizeDecCoin(msg.Coin).TruncateDecimal()

	if !coin.IsValid() || coin.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins,
			"coin is invalid or has zero amount")
	}

	poolCoins := sdk.NewCoins(pool.Coins...)

	// Check if msg denoms match pool coin denoms.
	if !sdk.NewCoins(coin).DenomsSubsetOf(poolCoins) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"Provided coin denoms does not match pool denoms",
		)
	}

	return nil
}

func GetSwapFeeCost(swapFee string, coin sdk.Coin) (sdk.Coin, error) {
	sfm, err := sdk.NewDecFromStr(swapFee)
	if err != nil {
		return coin, err
	}
		
	feeAmt := sfm.MulInt(coin.Amount)
	return sdk.NewCoin(coin.GetDenom(), feeAmt.RoundInt()), nil
}

// Creator deposits a coin and receives coin.
// Swap fee is subtracted from the deposit before swap is calculated.
func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	emptyMsgResponse := types.MsgSwapResponse{}

	if err := k.validateSwapMsg(ctx, msg); err != nil {
		return &emptyMsgResponse, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			err.Error())
	}

	pool, _ := k.GetLPool(ctx, msg.PoolName)
	poolCoins := sdk.NewCoins(pool.Coins...)
	depositCoin, _ := sdk.NormalizeDecCoin(msg.Coin).TruncateDecimal()

	for _, pCoin := range poolCoins {
		if pCoin.Amount.Equal(sdk.OneInt()) {
			return &emptyMsgResponse, sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"Pool is empty",
			)
		}
	}

	swapFeeCoin, err := GetSwapFeeCost(pool.SwapFeeMulti, depositCoin) 
	// Something went wrong when LPool was initialized
	// SwapFeeMulti saved in string format that could not be parsed
	// by sdk.Dec NewDecFromStr()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Internal error! Location: Swap()" + 
			" Failed to parse SwapFeeMulti: %s, err: %s", pool.SwapFeeMulti, err))
	}

	subtractedDeposit := depositCoin.Sub(swapFeeCoin)

	AMM, err := types.GetAMM(pool.AMM_Id)

	if err != nil {
		return &emptyMsgResponse, err
	}

	swapReturnCoins, err := AMM.EstimateReturn(poolCoins, sdk.NewCoins(subtractedDeposit))

	if err != nil {
		return &emptyMsgResponse, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			err.Error(),
		)
	}

	// Transfer money
	creatorAcc, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Send deposit to pool
	sdkErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAcc, types.ModuleName, sdk.NewCoins(depositCoin))

	if sdkErr != nil {
		return &emptyMsgResponse, sdkerrors.Wrap(sdkErr, "swap failed")
	}

	// Send other coin to creator
	sdkErr = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAcc, swapReturnCoins)

	if sdkErr != nil {
		return &emptyMsgResponse, sdkerrors.Wrap(sdkErr, "swap failed")
	}

	// Update pool balance
	poolCoins = poolCoins.Sub(swapReturnCoins)
	poolCoins = poolCoins.Add(depositCoin)

	pool.Coins = poolCoins

	k.SetLPool(ctx, pool)

	EmitCoinSwappedEvent(
		ctx, 
		creatorAcc, 
		pool, 
		sdk.NewCoins(depositCoin), 
		swapReturnCoins,
		sdk.NewCoins(swapFeeCoin))

	return &emptyMsgResponse, nil
}
