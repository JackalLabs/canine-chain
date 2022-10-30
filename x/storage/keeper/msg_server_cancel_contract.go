package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k msgServer) CancelContract(goCtx context.Context, msg *types.MsgCancelContract) (*types.MsgCancelContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	deal, found := k.GetActiveDeals(ctx, msg.Cid)
	if !found {
		return nil, fmt.Errorf("can't find contract")
	}

	if deal.Creator != msg.Creator {
		return nil, fmt.Errorf("you don't own this deal")
	}

	ftc, found := k.GetFidCid(ctx, deal.Fid)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "no fid found")
	}
	var ncids []string
	err := json.Unmarshal([]byte(ftc.Cids), &ncids)
	if err != nil {
		return nil, err
	}
	cids := make([]string, 0)
	for _, v := range ncids {
		if v != msg.Cid {
			cids = append(cids, v)
		}
	}
	b, err := json.Marshal(cids)
	if err != nil {
		return nil, err
	}
	ftc.Cids = string(b)

	k.SetFidCid(ctx, ftc)

	k.RemoveActiveDeals(ctx, deal.Cid)

	return &types.MsgCancelContractResponse{}, nil
}
