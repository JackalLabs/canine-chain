package keeper

import (
	"context"
	"crypto/sha256"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) PostFile(goCtx context.Context, msg *types.MsgPostFile) (*types.MsgPostFileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//old implementation: hex ( hash ( concatenate (msg.Creator, msg.Hashpath)))
	// h := sha256.New()
	// h.Write([]byte(fmt.Sprintf("%s%s", msg.Creator, msg.Hashpath)))
	// hash := h.Sum(nil)

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("o%s%s", msg.HashParent, msg.Account)))
	hash := h.Sum(nil)
	ownerString := fmt.Sprintf("%x", hash)

	parentFile, found := k.GetFiles(ctx, msg.HashParent, ownerString)
	if !found {
		return nil, types.ErrParentFileNotFound
	}

	hasEdit := HasEditAccess(parentFile, msg.Creator)
	if !hasEdit {
		return nil, types.ErrCannotWrite
	}

	//Make the full path
	merklePath := types.AddToMerkle(msg.HashParent, msg.HashChild)

	//desperately need a 'makeownerString function'
	ha := sha256.New()
	ha.Write([]byte(fmt.Sprintf("o%s%s", merklePath, msg.Account))) //msg.Creator will change to msg.accountAddress soon
	Hash := h.Sum(nil)
	Owner := fmt.Sprintf("%x", Hash)

	file := types.Files{
		Contents:      msg.Contents,
		Owner:         Owner,
		ViewingAccess: msg.Viewers,
		EditAccess:    msg.Editors,
		Address:       merklePath,
	}

	k.SetFiles(ctx, file)

	return &types.MsgPostFileResponse{Path: merklePath}, nil
}
