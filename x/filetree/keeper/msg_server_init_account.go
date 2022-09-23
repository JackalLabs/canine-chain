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

	//msg.Account was already hex(hashed) before it go to here.
	//make the full OwnerAddress
	// H := sha256.New()
	// H.Write([]byte(fmt.Sprintf("o%s%s", msg.RootHashpath, msg.Account)))
	// Hash := H.Sum(nil)
	// ownerAddress := fmt.Sprintf("%x", Hash)

	ownerAddress := MakeOwnerAddress(msg.RootHashpath, msg.Account)

	file := types.Files{
		Contents:       "home/", //might hex this later but leaving it here for now to see it in swagger
		Owner:          ownerAddress,
		ViewingAccess:  fmt.Sprintf("%x", "NONE"), //dummy var, no viewing access
		EditAccess:     msg.Editors,
		Address:        msg.RootHashpath,
		TrackingNumber: msg.TrackingNumber, //place holder
	}

	updatedTrackingNumber := msg.TrackingNumber + 1

	//need to double check this number
	if msg.TrackingNumber == 18446744073709551615 {
		updatedTrackingNumber = 0
		k.SetTracker(ctx, types.Tracker{
			TrackingNumber: uint64(updatedTrackingNumber),
		})
	} else {
		k.SetTracker(ctx, types.Tracker{
			TrackingNumber: uint64(updatedTrackingNumber),
		})
	}

	k.SetFiles(ctx, file)

	return &types.MsgInitAccountResponse{TrackingNumber: updatedTrackingNumber}, nil
}
