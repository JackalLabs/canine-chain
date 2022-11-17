package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (k msgServer) MakeRoot(goCtx context.Context, msg *types.MsgMakeRoot) (*types.MsgMakeRootResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//msg.Account was already hex(hashed) before it go to here.
	//make the full OwnerAddress

	ownerAddress := MakeOwnerAddress(msg.RootHashPath, msg.Account)

	file := types.Files{
		Contents:       msg.Contents, //This won't be used for now, but we're leaving it in as a stub in case it's needed
		Owner:          ownerAddress,
		ViewingAccess:  "NONE", //This won't be used for now, but we're leaving it in as a stub in case it's needed
		EditAccess:     msg.Editors,
		Address:        msg.RootHashPath,
		TrackingNumber: msg.TrackingNumber,
	}

	k.SetFiles(ctx, file)

	// placeholderMap := make([]string, 0, 2000)
	// placeholderMap = append(placeholderMap, msg.Creator)

	return &types.MsgMakeRootResponse{}, nil
}
