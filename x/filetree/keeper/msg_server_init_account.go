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
		Address: msg.Creator, //create public key for message caller
		Key:     msg.Key,
	}

	k.SetPubkey(ctx, pubKey)

	//free RNS name

	//msg.Account was already hex(hashed) before it go to here.
	//make the full OwnerAddress

	ownerAddress := MakeOwnerAddress(msg.RootHashpath, msg.Account)

	//These addresses in the viewer access and editor access below are not one
	//And the same as the wallet address
	file := types.Files{
		Contents:       "Root/", //might hex this later but leaving it here for now to see it in swagger
		Owner:          ownerAddress,
		ViewingAccess:  fmt.Sprintf("%x", "NONE"), //dummy var, no viewing access
		EditAccess:     msg.Editors,
		Address:        msg.RootHashpath,
		TrackingNumber: msg.TrackingNumber, //place holder
	}

	k.SetFiles(ctx, file)

	return &types.MsgInitAccountResponse{}, nil
}
