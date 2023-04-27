package keeper

import (
	"context"
	"crypto/sha256"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func (k msgServer) MakeRootFolder(ctx sdk.Context, creator string, viewers string, editors string, trackingNumber string) {
	merklePath := types.MerklePath("s")

	h1 := sha256.New() // making full address
	h1.Write([]byte(creator))
	hash1 := h1.Sum(nil)

	accountHash := fmt.Sprintf("%x", hash1)

	ownerAddress := MakeOwnerAddress(merklePath, accountHash)

	file := types.Files{
		Contents:       "",
		Owner:          ownerAddress,
		ViewingAccess:  viewers,
		EditAccess:     editors,
		Address:        merklePath,
		TrackingNumber: trackingNumber,
	}

	k.SetFiles(ctx, file)
}

func (k msgServer) MakeRoot(goCtx context.Context, msg *types.MsgMakeRoot) (*types.MsgMakeRootResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.MakeRootFolder(ctx, msg.Creator, msg.Viewers, msg.Editors, msg.TrackingNumber)

	return &types.MsgMakeRootResponse{}, nil
}

func (k msgServer) MakeRootV2(goCtx context.Context, msg *types.MsgMakeRootV2) (*types.MsgMakeRootV2Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.MakeRootFolder(ctx, msg.Creator, msg.Viewers, msg.Editors, msg.TrackingNumber)

	return &types.MsgMakeRootV2Response{}, nil
}
