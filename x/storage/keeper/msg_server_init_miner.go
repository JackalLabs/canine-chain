package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/storage/types"
)

func (k msgServer) InitMiner(goCtx context.Context, msg *types.MsgInitMiner) (*types.MsgInitMinerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	miner := types.Miners{
		Address:    msg.Creator,
		Ip:         msg.Ip,
		Totalspace: msg.Totalspace,
		Creator:    msg.Creator,
	}

	k.SetMiners(ctx, miner)

	return &types.MsgInitMinerResponse{}, nil
}
