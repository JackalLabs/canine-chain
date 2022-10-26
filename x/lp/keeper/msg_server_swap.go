package keeper

import (
	"context"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/lp/types"
)

func (k Keeper) validateSwapMsg(ctx sdk.Context, msg *types.MsgSwap) error {
	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	pool, found := k.GetLPool(ctx, msg.PoolName)

	if !found {
		return types.ErrLiquidityPoolNotFound
	}

	// Convert DecCoin to Normalized Coin
	coin, _ := sdk.NormalizeDecCoin(msg.CoinInput).TruncateDecimal()

	// Convert DecCoin to Normalized MinCoinOutput
	minCoinOut, _ := sdk.NormalizeDecCoin(msg.MinCoinOutput).TruncateDecimal()

	if !minCoinOut.IsValid() || coin.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins,
			"coin is invalid or has zero amount")
	}

	poolCoins := sdk.NewCoins(pool.Coins...)

	// Check if msg denoms match pool coin denoms.
	if !sdk.NewCoins(coin, minCoinOut).DenomsSubsetOf(poolCoins) {
		return errors.New("Provided coin denoms does not match pool denoms")
	}

	return nil
}

// Returns swap fee
// Panics if swapFeeRate couldn't be converted to sdk.Dec
func GetSwapFee(swapFeeRate string, coin sdk.Coin) sdk.Coin {
	sfm, err := sdk.NewDecFromStr(swapFeeRate)

	// Something went wrong when LPool was initialized
	// SwapFeeMulti saved in string format that could not be parsed
	// by sdk.Dec NewDecFromStr()
	if err != nil {
		panic(fmt.Errorf("Internal error! Location: Swap()"+
				" Failed to parse SwapFeeMulti: %s, err: %s", swapFeeRate, err))

	}

	feeAmt := sfm.MulInt(coin.Amount)
	return sdk.NewCoin(coin.GetDenom(), feeAmt.RoundInt())
}

// Returns protocol fee
func GetProtocolFee(coin sdk.Coin) sdk.Coin {
	rate := types.ProtocolFeeRate()
	fee := rate.MulInt(coin.Amount) 

	return sdk.NewCoin(coin.GetDenom(), fee.RoundInt())
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
	swapIn, _ := sdk.NormalizeDecCoin(msg.CoinInput).TruncateDecimal()

	for _, pCoin := range poolCoins {
		if pCoin.Amount.Equal(sdk.OneInt()) {
			return &emptyMsgResponse, sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"Pool is empty",
			)
		}
	}

	swapFee := GetSwapFee(pool.SwapFeeMulti, swapIn)
	protocolFee := GetProtocolFee(swapIn)

	deductedCoinIn := swapIn.Sub(swapFee).Sub(protocolFee)

	AMM, err := types.GetAMM(pool.AMM_Id)

	// Something went wrong when LPool was initialized
	// Registered invalid AMM
	if err != nil {
		panic(err)
	}

	swapOut, err := AMM.EstimateReturn(poolCoins, sdk.NewCoins(deductedCoinIn))

	if err != nil {
		return &emptyMsgResponse, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			err.Error(),
		)
	}

	creatorAcc, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Check slippage
	minCoinOut, _ := sdk.NormalizeDecCoin(msg.MinCoinOutput).TruncateDecimal()
	if swapOut.IsAllLT(sdk.NewCoins(minCoinOut)) {
		EmitCoinSwapFailedEvent(
			ctx,
			creatorAcc,
			pool,
			sdk.NewCoins(swapIn),
			swapOut,
			sdk.NewCoins(minCoinOut),
		)
		return &emptyMsgResponse, nil
	}

	// Transfer money

	// Send coin input to pool
	sdkErr := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, 
		creatorAcc,
		types.ModuleName, 
		sdk.NewCoins(swapIn))

	if sdkErr != nil {
		return &emptyMsgResponse, sdkerrors.Wrap(sdkErr, "swap failed")
	}
	
	// Send protocol fee to protocol fee to account
	sdkErr = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, 
		types.ModuleName, 
		types.ProtocolFeeToAcc(), 
		sdk.NewCoins(protocolFee))

	if sdkErr != nil {
		return &emptyMsgResponse, sdkerrors.Wrap(sdkErr, "swap failed")
	}

	// Send swap out to the creator
	sdkErr = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAcc, swapOut)

	if sdkErr != nil {
		return &emptyMsgResponse, sdkerrors.Wrap(sdkErr, "swap failed")
	}

	// Update pool balance
	poolCoins = poolCoins.Sub(swapOut)
	poolCoins = poolCoins.Add(swapIn)

	pool.Coins = poolCoins

	k.SetLPool(ctx, pool)

	EmitCoinSwappedEvent(
		ctx,
		creatorAcc,
		pool,
		sdk.NewCoins(swapIn),
		swapOut,
		sdk.NewCoins(swapFeeCoin))

	return &emptyMsgResponse, nil
}
