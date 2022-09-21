package keeper

import (
	"context"
	"crypto/sha256"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) InitAccount(goCtx context.Context, msg *types.MsgInitAccount) (*types.MsgInitAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pubKey := types.Pubkey{
		Address: msg.Creator,
		Key:     msg.Key,
	}
	k.SetPubkey(ctx, pubKey)

	pathString := msg.RootHashpath

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("o%s%s", pathString, msg.Creator))) //msg.Creator will change to msg.accountAddress soon
	hash := h.Sum(nil)

	ownerString := fmt.Sprintf("%x", hash)

	file := types.Files{
		Contents:      "home/",
		Owner:         ownerString,
		ViewingAccess: "NONE",
		EditAccess:    msg.Editors,
		Address:       pathString,
	}

	k.SetFiles(ctx, file)

	return &types.MsgInitAccountResponse{}, nil
}
