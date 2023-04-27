package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (k msgServer) MakeRoot(goCtx context.Context, msg *types.MsgMakeRoot) (*types.MsgMakeRootResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	merklePath := types.MerklePath("s")

	// msg.Account was already hex(hashed) before it go to here.
	// make the full OwnerAddress

	ownerAddress := MakeOwnerAddress(msg.RootHashPath, msg.Account)

	file := types.Files{
		Contents:       msg.Contents,
		Owner:          ownerAddress,
		ViewingAccess:  msg.Viewers,
		EditAccess:     msg.Editors,
		Address:        merklePath,
		TrackingNumber: msg.TrackingNumber,
	}

	k.SetFiles(ctx, file)

	return &types.MsgMakeRootResponse{}, nil
}
