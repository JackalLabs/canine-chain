package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
)

func (k msgServer) SetMinerTotalspace(goCtx context.Context, msg *types.MsgSetMinerTotalspace) (*types.MsgSetMinerTotalspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	miner, found := k.GetMiners(ctx, msg.Creator)

	if !found {
		miner = types.Miners{
			Address:    msg.Creator,
			Ip:         "",
			Totalspace: "0",
			Creator:    msg.Creator,
		}
	}

	miner.Totalspace = msg.Space

	k.SetMiners(ctx, miner)

	return &types.MsgSetMinerTotalspaceResponse{}, nil
}
