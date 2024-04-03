package keeper

import (
	"context"
	"encoding/json"
	"strconv"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k Keeper) CanContract(ctx sdk.Context, root string, creator string) error {
	d, dealFound := k.GetActiveDeals(ctx, root)
	s, strayFound := k.GetStrays(ctx, root)

	var fileSize int64
	var fid string
	var signee string
	var err error

	// nolint
	if dealFound {
		signee = d.Signee
		fid = d.Fid
		fileSize, err = strconv.ParseInt(d.Filesize, 10, 64)
		if err != nil {
			return sdkerrors.Wrapf(err, "cannot parse file size from deal")
		}
		endBlock, err := strconv.ParseInt(d.Endblock, 10, 64)
		if err != nil {
			return sdkerrors.Wrapf(err, "cannot parse end block from deal")
		}
		if endBlock > ctx.BlockHeight() {
			return sdkerrors.Wrapf(types.ErrCancelContract, "this is a persistent file that can't be cancelled until it expires")
		}
	} else if strayFound {
		signee = s.Signee
		fid = s.Fid
		fileSize, err = strconv.ParseInt(s.Filesize, 10, 64)
		if err != nil {
			return err
		}
		if s.End > ctx.BlockHeight() {
			return sdkerrors.Wrapf(types.ErrCancelContract, "this is a persistent file that can't be cancelled until it expires")
		}
	} else {
		return types.ErrNoCid
	}

	if creator != signee {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "cannot cancel a contract that isn't yours. %s is not %s", creator, signee)
	}

	k.RemoveStrays(ctx, root)
	k.RemoveActiveDeals(ctx, root)

	newFidCid := types.FidCid{
		Fid:  fid,
		Cids: "",
	}

	ftc, found := k.GetFidCid(ctx, fid) // get existing FIDCID Mapping
	cids := make([]string, 0)           // create new home for CID list
	if found {                          // if found we remove the existing cid from the list
		var ncids []string
		err := json.Unmarshal([]byte(ftc.Cids), &ncids) // getting all cids from the existing fid_cid
		if err != nil {
			return err
		}

		for _, v := range ncids { // all all cids to the list again if they aren't the root
			if v != root {
				cids = append(cids, v)
			}
		}
	}

	b, err := json.Marshal(cids) // put em all back
	if err != nil {
		return err
	}
	newFidCid.Cids = string(b)

	k.SetFidCid(ctx, newFidCid)

	info, found := k.GetStoragePaymentInfo(ctx, creator)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "cannot find storage payment")
	}
	info.SpaceUsed -= fileSize
	k.SetStoragePaymentInfo(ctx, info)

	return nil
}

func (k msgServer) CancelContract(goCtx context.Context, msg *types.MsgCancelContract) (*types.MsgCancelContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	root := msg.Cid

	err := k.Keeper.CanContract(ctx, root, msg.Creator)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCancelContract,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyContract, msg.Cid),
		),
	)

	return &types.MsgCancelContractResponse{}, err
}
