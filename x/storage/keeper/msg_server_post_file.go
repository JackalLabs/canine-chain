package keeper

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (k msgServer) PostFile(goCtx context.Context, msg *types.MsgPostFile) (*types.MsgPostFileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	if !json.Valid([]byte(msg.Note)) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "note is not valid json `%s`", msg.Note)
	}

	window := k.GetParams(ctx).ProofWindow

	file := types.UnifiedFile{
		Merkle:        msg.Merkle,
		Owner:         msg.Creator,
		Start:         ctx.BlockHeight(),
		Expires:       msg.Expires,
		FileSize:      msg.FileSize,
		ProofInterval: window,
		ProofType:     msg.ProofType,
		Proofs:        make([]string, 0),
		MaxProofs:     msg.MaxProofs,
		Note:          msg.Note,
	}
	k.SetFile(ctx, file)

	ips := make([]string, 0)

	res := &types.MsgPostFileResponse{ProviderIps: ips, StartBlock: ctx.BlockHeight()}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	b := False // pay once event
	if msg.Expires > 0 {
		b = True
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSignContract,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyContract, hex.EncodeToString(msg.Merkle)),
			sdk.NewAttribute(types.AttributeKeyPayOnce, b),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJackalMessage,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	totalSize := msg.FileSize * msg.MaxProofs
	if msg.Expires > 0 { // if the file is posted as a one-time payment
		kbs := totalSize / 1000
		var kbMin int64 = 1024
		if kbs < kbMin { // minimum amount of kbs to post
			kbs = kbMin
		}

		blockDuration := msg.Expires - ctx.BlockHeight()
		seconds := blockDuration * 6
		minutes := seconds / 60
		hours := minutes / 60
		days := hours / 24
		if days <= 0 {
			return nil, fmt.Errorf("cannot pay for less than a day")
		}

		cost := k.GetStorageCostKbs(ctx, kbs, hours)

		toPay := sdk.NewCoin("ujkl", cost)

		refDec := sdk.NewDec(params.ReferralCommission).QuoInt64(100)
		pol := sdk.NewDec(params.PolRatio).QuoInt64(100)

		spr := sdk.NewDec(1).Sub(refDec).Sub(pol) // whatever is left from pol and referrals

		storageProviderCut := toPay.Amount.ToDec().Mul(spr)
		spcToken := sdk.NewCoin(toPay.Denom, storageProviderCut.TruncateInt())
		spcTokens := sdk.NewCoins(spcToken)

		end := ctx.BlockTime().AddDate(0, 0, int(days))
		gauge := k.NewGauge(ctx, spcTokens, end) // creating new payment gauge
		addr, err := sdk.AccAddressFromBech32(msg.Creator)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "cannot get address from message creator")
		}

		acc, err := types.GetGaugeAccount(gauge)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "cannot get gauge account")
		}

		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(toPay)) // taking money from user
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "cannot send tokens from %s", msg.Creator)
		}

		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, acc, spcTokens)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "cannot send tokens to token holder account")
		}

		return res, nil
	}

	// traditional storage plan payment info

	paymentInfo, found := k.GetStoragePaymentInfo(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "storage account does not exist")
	}
	if paymentInfo.End.Before(ctx.BlockTime()) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "storage account is expired")
	}

	paymentInfo.SpaceUsed += totalSize
	if paymentInfo.SpaceUsed > paymentInfo.SpaceAvailable {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "storage account does not have enough space available %d > %d", paymentInfo.SpaceUsed, paymentInfo.SpaceAvailable)
	}

	k.SetStoragePaymentInfo(ctx, paymentInfo)

	return res, nil
}
