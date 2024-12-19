package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v4/x/jklmint/types"
)

func (k Keeper) send(ctx sdk.Context, denom string, amount int64, receiver string) error {
	adr, err := sdk.AccAddressFromBech32(receiver)
	if err != nil {
		return sdkerrors.Wrapf(err, "cannot parse receiver address")
	}

	c := sdk.NewInt64Coin(denom, amount)
	cs := sdk.NewCoins(c)

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, adr, cs)
	if err != nil {
		return sdkerrors.Wrapf(err, "cannot send coins from the mint module to %s", receiver)
	}

	return nil
}

func (k Keeper) mintStaker(ctx sdk.Context, mintTokens int64, denom string, params types.Params) error {
	stakerRatio := sdk.NewDec(params.StakerRatio).QuoInt64(100)

	stakerCoinValue := stakerRatio.MulInt64(mintTokens).TruncateInt64()
	stakerCoins := sdk.NewCoins(sdk.NewInt64Coin(denom, stakerCoinValue))

	// send the minted validator coins to the fee collector account
	err := k.AddCollectedFees(ctx, stakerCoins)
	if err != nil {
		return sdkerrors.Wrapf(err, "cannot send tokens to stakers & community pool")
	}

	return nil
}

func GetDevGrantsAccount() (sdk.AccAddress, error) {
	return GetAccount(types.DevGrantsPool)
}

func GetAccount(name string) (sdk.AccAddress, error) {
	s := sha256.New()
	s.Write([]byte(name))
	m := s.Sum(nil)
	mh := hex.EncodeToString(m)
	adr, err := sdk.AccAddressFromHex(mh)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot get account account")
	}
	return adr, nil
}

func (k Keeper) mintDevGrants(ctx sdk.Context, mintTokens int64, denom string, params types.Params) error {
	devGrantRatio := sdk.NewDec(params.DevGrantsRatio).QuoInt64(100)

	devGrantTokenAmount := devGrantRatio.MulInt64(mintTokens).TruncateInt64()

	address, err := GetDevGrantsAccount()
	if err != nil {
		return sdkerrors.Wrapf(err, "cannot create dev grants address")
	}

	err = k.send(ctx, denom, devGrantTokenAmount, address.String())
	if err != nil {
		return sdkerrors.Wrapf(err, "cannot send tokens to dev grants")
	}

	return nil
}

func (k Keeper) mintStorageProviderStipend(ctx sdk.Context, mintTokens int64, denom string, params types.Params) error {
	provRatio := sdk.NewDec(params.StorageProviderRatio).QuoInt64(100)

	provTokens := provRatio.MulInt64(mintTokens).TruncateInt64()

	err := k.send(ctx, denom, provTokens, params.StorageStipendAddress)
	if err != nil {
		return sdkerrors.Wrapf(err, "cannot send tokens to storage provider stipends")
	}

	return nil
}

func (k Keeper) BlockMint(ctx sdk.Context) {
	params := k.GetParams(ctx)

	mintedNum := params.TokensPerBlock
	// minted, found := k.GetMintedBlock(ctx, ctx.BlockHeight()-1)
	//if found {
	//	mintedNum = minted.Minted
	//}
	//var bpy int64 = (365 * 24 * 60 * 60) / 6
	//
	//newMintForBlock := utils.GetMintForBlock(mintedNum, bpy, params.MintDecrease)
	//
	//mintTokens := newMintForBlock
	denom := k.GetParams(ctx).MintDenom
	if denom == "" { // error handling mostly built for tests
		denom = "ujkl"
	}

	totalCoin := sdk.NewInt64Coin(denom, mintedNum)
	coins := sdk.NewCoins(totalCoin)

	err := k.MintCoins(ctx, coins)
	if err != nil {
		ctx.Logger().Error(sdkerrors.Wrapf(err, "could not mint tokens at block %d", ctx.BlockHeight()).Error())
		return
	}

	err = k.mintStaker(ctx, mintedNum, denom, params)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	err = k.mintDevGrants(ctx, mintedNum, denom, params)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	err = k.mintStorageProviderStipend(ctx, mintedNum, denom, params)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	k.SetMintedBlock(ctx, types.MintedBlock{
		Height: ctx.BlockHeight(),
		Minted: mintedNum,
		Denom:  "ujkl",
	})

	// alerting network on mint amount
	defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedNum), "minted_tokens")

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(sdk.AttributeKeyAmount, fmt.Sprintf("%d", mintedNum)),
		),
	)
}
