package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) InitAccount(goCtx context.Context, msg *types.MsgInitAccount) (*types.MsgInitAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pubKey := types.Pubkey{
		Address: msg.Creator, //this is hex(hashed)
		Key:     msg.Key,
	}
	k.SetPubkey(ctx, pubKey)

	merklePath := msg.RootHashpath
	ownerString := MakeOwnerAddress(merklePath, msg.Creator)

	file := types.Files{
		Contents:      fmt.Sprintf("%x", "homeContents"), //dummy contents
		Owner:         ownerString,
		ViewingAccess: fmt.Sprintf("%x", "NONE"), //dummy var, no viewing access
		EditAccess:    msg.Editors,
		Address:       merklePath,
	}

	k.SetFiles(ctx, file)

	return &types.MsgInitAccountResponse{}, nil
}
