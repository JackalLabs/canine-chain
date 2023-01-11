package keeper

import (
	"context"
	"encoding/json"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func CanContract(ctx sdk.Context, root string, creator string, k Keeper) error {
	var fid string

	d, found := k.GetActiveDeals(ctx, root)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "no deal found")
	}

	if creator != d.Signee {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "cannot cancel a contract that isn't yours")
	}

	k.RemoveStrays(ctx, d.Cid)
	k.RemoveActiveDeals(ctx, d.Cid)

	ftc, found := k.GetFidCid(ctx, fid)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "no fid found")
	}

	var ncids []string
	err := json.Unmarshal([]byte(ftc.Cids), &ncids)
	if err != nil {
		return err
	}

	cids := make([]string, 0)
	for _, v := range ncids {
		if v != root {
			cids = append(cids, v)
		}
	}
	b, err := json.Marshal(cids)
	if err != nil {
		return err
	}
	ftc.Cids = string(b)

	k.SetFidCid(ctx, ftc)
	return nil
}

func (k msgServer) CancelContract(goCtx context.Context, msg *types.MsgCancelContract) (*types.MsgCancelContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	root := msg.Cid

	err := CanContract(ctx, root, msg.Creator, k.Keeper)

	return &types.MsgCancelContractResponse{}, err
}
