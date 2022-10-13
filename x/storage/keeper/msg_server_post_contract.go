package keeper

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
)

func (k msgServer) PostContract(goCtx context.Context, msg *types.MsgPostContract) (*types.MsgPostContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, ok := k.GetProviders(ctx, msg.Creator)
	if !ok {
		return nil, fmt.Errorf("can't find provider")
	}

	ts, ok := sdk.NewIntFromString(provider.Totalspace)

	if !ok {
		return nil, fmt.Errorf("error parsing total space")
	}

	fs, ok := sdk.NewIntFromString(msg.Filesize)

	if !ok {
		return nil, fmt.Errorf("error parsing file size")
	}

	if k.GetProviderUsing(ctx, msg.Creator)+fs.Int64() > ts.Int64() {
		return nil, fmt.Errorf("not enough space on provider")
	}

	paidAMT := k.GetPaidAmount(ctx, msg.Signee, ctx.BlockHeight())

	if paidAMT <= 0 {
		return nil, fmt.Errorf("user has not paid for any storage")
	}

	usage, found := k.GetClientUsage(ctx, msg.Signee)
	if !found {
		usage = types.ClientUsage{
			Usage:   "0",
			Address: msg.Signee,
		}
	}

	bytesUsed, ok := sdk.NewIntFromString(usage.Usage)
	if !ok {
		return nil, fmt.Errorf("failed to parse usage")
	}

	filesize, ok := sdk.NewIntFromString(msg.Filesize)
	if !ok {
		return nil, fmt.Errorf("cannot parse filesize")
	}

	if bytesUsed.Int64()+filesize.Int64() > paidAMT {
		return nil, fmt.Errorf("not enough storage on the users account")
	}

	h := sha256.New()
	io.WriteString(h, msg.Creator+msg.Fid)
	hashName := h.Sum(nil)

	newContract := types.Contracts{
		Cid:        fmt.Sprintf("%x", hashName),
		Priceamt:   msg.Priceamt,
		Pricedenom: msg.Pricedenom,
		Signee:     msg.Signee,
		Duration:   msg.Duration,
		Fid:        msg.Fid,
		Filesize:   msg.Filesize,
		Creator:    msg.Creator,
		Merkle:     msg.Merkle,
	}

	k.SetContracts(ctx, newContract)

	return &types.MsgPostContractResponse{}, nil
}
