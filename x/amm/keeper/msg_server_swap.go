package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/amm/types"
)

func (k Keeper) validateSwapMsg(ctx sdk.Context, msg *types.MsgSwap) error {
	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	_, found := k.GetPool(ctx, msg.PoolId)

	if !found {
		return types.ErrLiquidityPoolNotFound
	}

	return nil
}

// Returns swap fee
// Panics if swapFeeRate couldn't be converted to sdk.Dec
func GetSwapFee(swapFeeRate string, coin sdk.Coin) sdk.Coin {
	sfm, err := sdk.NewDecFromStr(swapFeeRate)

	// Something went wrong when Pool was initialized
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
func (k Keeper) GetProtocolFee(ctx sdk.Context, coin sdk.Coin) sdk.Coin {
	rate := k.ProtocolFeeRate(ctx)
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

	pool, _ := k.GetPool(ctx, msg.PoolId)
	poolCoins := sdk.NewCoins(pool.Coins...)

	for _, pCoin := range poolCoins {
		if pCoin.Amount.Equal(sdk.OneInt()) {
			return &emptyMsgResponse, sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"Pool is empty",
			)
		}
	}

	swapFee := GetSwapFee(pool.SwapFeeMulti, msg.CoinInput)
	protocolFee := k.GetProtocolFee(ctx, msg.CoinInput)

	deductedCoinIn := msg.CoinInput.Sub(swapFee).Sub(protocolFee)

	AMM, err := types.GetAMM(pool.AMM_Id)

	// Something went wrong when Pool was initialized
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
	if swapOut.IsAllLT(sdk.NewCoins(msg.MinCoinOutput)) {
		EmitCoinSwapFailedEvent(
			ctx,
			creatorAcc,
			pool,
			sdk.NewCoins(msg.CoinInput),
			swapOut,
			sdk.NewCoins(msg.MinCoinOutput),
		)
		return &emptyMsgResponse, nil
	}

	// Transfer money

	// Send coin input to pool
	sdkErr := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, 
		creatorAcc,
		types.ModuleName, 
		sdk.NewCoins(msg.CoinInput))

	if sdkErr != nil {
		return &emptyMsgResponse, sdkerrors.Wrap(sdkErr, "swap failed")
	}
	
	// Send protocol fee to protocol fee to account
	sdkErr = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, 
		types.ModuleName, 
		k.ProtocolFeeToAcc(ctx), 
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
	poolCoins = poolCoins.Add(msg.CoinInput)

	pool.Coins = poolCoins

	k.SetPool(ctx, pool)

	EmitCoinSwappedEvent(
		ctx,
		creatorAcc,
		pool,
		sdk.NewCoins(msg.CoinInput),
		swapOut,
		sdk.NewCoins(swapFee),
		sdk.NewCoins(protocolFee))

	return &emptyMsgResponse, nil
}
