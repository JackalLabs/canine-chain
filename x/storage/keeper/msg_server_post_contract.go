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

	miner, ok := k.GetMiners(ctx, msg.Creator)
	if !ok {
		return nil, fmt.Errorf("can't find miner")
	}

	ts, ok := sdk.NewIntFromString(miner.Totalspace)

	if !ok {
		return nil, fmt.Errorf("error parsing total space")
	}

	fs, ok := sdk.NewIntFromString(msg.Filesize)

	if !ok {
		return nil, fmt.Errorf("error parsing file size")
	}

	if k.GetMinerUsing(ctx, msg.Creator)+fs.Int64() > ts.Int64() {
		return nil, fmt.Errorf("not enough space on miner")
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
