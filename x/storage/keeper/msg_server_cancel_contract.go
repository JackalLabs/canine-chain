package keeper

import (
	"context"
	"encoding/json"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

type contract struct {
	Signee   string
	Cid      string
	Fid      string
	Filesize string
}

func CanContract(ctx sdk.Context, root string, creator string, k Keeper) error {
	var fid string
	var c contract

	d, dealFound := k.GetActiveDeals(ctx, root)

	s, found := k.GetStrays(ctx, root)
	if !found {
		if !dealFound {
			return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "no deal found")
		}
		c.Cid = d.Cid
		c.Signee = d.Signee
		c.Fid = d.Fid
		c.Filesize = d.Filesize
	} else {
		c.Cid = s.Cid
		c.Signee = s.Signee
		c.Fid = s.Fid
		c.Filesize = s.Filesize
	}

	if creator != c.Signee {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "cannot cancel a contract that isn't yours. %s is not %s", creator, c.Signee)
	}

	k.RemoveStrays(ctx, c.Cid)
	k.RemoveActiveDeals(ctx, c.Cid)

	ftc, found := k.GetFidCid(ctx, fid)
	if found {
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
	}

	info, found := k.GetStoragePaymentInfo(ctx, creator)
	if found {
		size, ok := sdk.NewIntFromString(c.Filesize)
		if !ok {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "cannot parse file size")
		}
		info.SpaceUsed -= size.Int64()
		k.SetStoragePaymentInfo(ctx, info)
	}

	return nil
}

func (k msgServer) CancelContract(goCtx context.Context, msg *types.MsgCancelContract) (*types.MsgCancelContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	root := msg.Cid

	err := CanContract(ctx, root, msg.Creator, k.Keeper)

	return &types.MsgCancelContractResponse{}, err
}
