package keeper

import (
	"context"
	"crypto/sha256"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func (k Keeper) MakeRootFolder(ctx sdk.Context, creator string, viewers string, editors string, trackingNumber string) {
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

func (k msgServer) ProvisionFileTree(goCtx context.Context, msg *types.MsgProvisionFileTree) (*types.MsgProvisionFileTreeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.MakeRootFolder(ctx, msg.Creator, msg.Viewers, msg.Editors, msg.TrackingNumber)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMakeRoot,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJackalMessage,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)
	return &types.MsgProvisionFileTreeResponse{}, nil
}
