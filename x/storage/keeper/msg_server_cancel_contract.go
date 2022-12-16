package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func CanContract(ctx sdk.Context, root string, creator string, k Keeper) error {
	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%s%d", root, 0))
	if err != nil {
		return err
	}
	hashName := h.Sum(nil)

	left, err := MakeCid(hashName)
	if err != nil {
		return err
	}
	h = sha256.New()
	_, err = io.WriteString(h, fmt.Sprintf("%s%d", root, 1))
	if err != nil {
		return err
	}
	hashName = h.Sum(nil)

	right, err := MakeCid(hashName)
	if err != nil {
		return err
	}

	var fid string

	d, found := k.GetActiveDeals(ctx, root)
	if !found {
		s, found := k.GetStrays(ctx, root)
		if !found {
			return fmt.Errorf("can't find contract")
		}
		k.RemoveStrays(ctx, s.Cid)
	} else {
		k.RemoveActiveDeals(ctx, d.Cid)
	}

	d, found = k.GetActiveDeals(ctx, left)
	if !found {
		s, found := k.GetStrays(ctx, left)
		if !found {
			return fmt.Errorf("can't find contract")
		}
		k.RemoveStrays(ctx, s.Cid)
	} else {
		k.RemoveActiveDeals(ctx, d.Cid)
	}

	d, found = k.GetActiveDeals(ctx, right)
	if !found {
		s, found := k.GetStrays(ctx, right)
		if !found {
			return fmt.Errorf("can't find contract")
		}
		fid = s.Fid
		k.RemoveStrays(ctx, s.Cid)

	} else {
		fid = d.Fid
		k.RemoveActiveDeals(ctx, d.Cid)
		if !(d.Creator == creator || d.Creator == "admin") {
			return fmt.Errorf("you don't own this deal")
		}
	}

	ftc, found := k.GetFidCid(ctx, fid)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "no fid found")
	}

	var ncids []string
	err = json.Unmarshal([]byte(ftc.Cids), &ncids)
	if err != nil {
		return err
	}

	cids := make([]string, 0)
	for _, v := range ncids {
		if v != root && v != left && v != right {
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
