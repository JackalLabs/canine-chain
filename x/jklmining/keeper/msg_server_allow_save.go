package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/jklmining/types"
)

func (k msgServer) AllowSave(goCtx context.Context, msg *types.MsgAllowSave) (*types.MsgAllowSaveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAllowSaveResponse{}, nil
}
