package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
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

	k.RemoveActiveDeals(ctx, deal.Cid)

	return &types.MsgCancelContractResponse{}, nil
}
