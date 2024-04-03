package keeper

import (
	"context"
	"encoding/hex"
	"encoding/json"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k msgServer) PostFile(goCtx context.Context, msg *types.MsgPostFile) (*types.MsgPostFileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !json.Valid([]byte(msg.Note)) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "note is not valid json `%s`", msg.Note)
	}

	window := k.GetParams(ctx).ProofWindow
	// if msg.ProofInterval != window {
	//	return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "cannot create a file with a window different than %d", window)
	//}

	providers := k.GetActiveProviders(ctx, "")
	if len(providers) == 0 {
		allProviders := k.GetRandomizedProviders(ctx)

		l := make([]types.ActiveProviders, len(allProviders))
		for i, provider := range allProviders {
			l[i] = types.ActiveProviders{Address: provider.Address}
		}

		providers = l
	}

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

	for i, provider := range providers { // adding all provers
		if i >= int(msg.MaxProofs) {
			break
		}
		file.AddProver(ctx, k, provider.Address)

		prv, found := k.GetProviders(ctx, provider.Address)
		if !found {
			continue
		}

		ips = append(ips, prv.Ip)
	}

	res := &types.MsgPostFileResponse{ProviderIps: ips, StartBlock: ctx.BlockHeight()}

	totalSize := msg.FileSize * msg.MaxProofs
	if msg.Expires > 0 { // if the file is posted as a one-time payment
		kbs := totalSize / 1000
		var kbMin int64 = 1024
		if kbs < kbMin { // minimum amount of kbs to post
			kbs = kbMin
		}

		var minBlocks int64 = 600
		blockDuration := msg.Expires - ctx.BlockHeight()
		if blockDuration < minBlocks {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "cannot post a file for less than %d blocks", minBlocks)
		}

		seconds := blockDuration * 6
		minutes := seconds / 60
		hours := minutes / 60

		cost := k.GetStorageCostKbs(ctx, kbs, hours)

		// TODO: charge user for cost using USDC
		_ = cost

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

	return res, nil
}
