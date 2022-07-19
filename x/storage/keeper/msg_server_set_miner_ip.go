package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
)

func (k msgServer) SetMinerIp(goCtx context.Context, msg *types.MsgSetMinerIp) (*types.MsgSetMinerIpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	miner, found := k.GetMiners(ctx, msg.Creator)

	if !found {
		miner = types.Miners{
			Address:         msg.Creator,
			Ip:              "",
			Totalspace:      "0",
			Creator:         msg.Creator,
			BurnedContracts: "0",
		}
	}

	miner.Ip = msg.Ip

	k.SetMiners(ctx, miner)

	return &types.MsgSetMinerIpResponse{}, nil
}
