package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) MakeFolder(goCtx context.Context, msg *types.MsgMakeFolder) (*types.MsgMakeFolderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//msg.Account was already hex(hashed) before it go to here.
	//make the full OwnerAddress

	ownerAddress := MakeOwnerAddress(msg.RootHashPath, msg.Account)

	file := types.Files{
		Contents:       msg.Contents, //might hex this later but leaving it here for now to see it in swagger
		Owner:          ownerAddress,
		ViewingAccess:  fmt.Sprintf("%x", "NONE"), //dummy var, no viewing access
		EditAccess:     msg.Editors,
		Address:        msg.RootHashPath,
		TrackingNumber: msg.TrackingNumber, //place holder
	}

	k.SetFiles(ctx, file)

	return &types.MsgMakeFolderResponse{}, nil
}
