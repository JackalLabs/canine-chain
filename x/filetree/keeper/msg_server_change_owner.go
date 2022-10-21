package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) ChangeOwner(goCtx context.Context, msg *types.MsgChangeOwner) (*types.MsgChangeOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	log, logfile := createLogger()

	file, found := k.GetFiles(ctx, msg.Address, msg.FileOwner)
	if !found {
		return nil, types.ErrFileNotFound
	}
	log.Printf("current owner is %s\n", file.Owner)
	//Only the owner of a file can give it away
	isOwner := IsOwner(file, msg.Creator)
	if !isOwner {
		return nil, types.ErrCantGiveAway
	}

	file.Owner = msg.NewOwner

	k.SetFiles(ctx, file)
	//Delete old file
	k.RemoveFiles(ctx, msg.Address, msg.FileOwner)

	log.Printf("new owner is %s\n", file.Owner)

	logfile.Close()

	return &types.MsgChangeOwnerResponse{}, nil
}
