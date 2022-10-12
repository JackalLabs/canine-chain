package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/filetree/types"
	notiTypes "github.com/jackal-dao/canine/x/notifications/types"
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

	ownerAddress := MakeOwnerAddress(msg.RootHashpath, msg.Account)

	file := types.Files{
		Contents:       "Root/", //might hex this later but leaving it here for now to see it in swagger
		Owner:          ownerAddress,
		ViewingAccess:  fmt.Sprintf("%x", "NONE"), //dummy var, no viewing access
		EditAccess:     msg.Editors,
		Address:        msg.RootHashpath,
		TrackingNumber: msg.TrackingNumber, //place holder
	}

	k.SetFiles(ctx, file)

	//Set notiCounter
	// Check if the counter already exists
	_, isFound := k.notiKeeper.GetNotiCounter(
		ctx,
		msg.Creator,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "notiCounter already set")
	}

	//Add yourself as a permitted Sender in the beginning so you can notify yourself

	placeholderMap := make([]string, 0, 2000)
	placeholderMap = append(placeholderMap, msg.Creator)
	marshalledSenders, err := json.Marshal(placeholderMap)
	if err != nil {
		return nil, types.ErrCantUnmarshall
	}

	updatedSenders := string(marshalledSenders)

	var counter = notiTypes.NotiCounter{
		Address:          msg.Creator,
		Counter:          0,
		PermittedSenders: updatedSenders,
	}

	k.notiKeeper.SetNotiCounter(
		ctx,
		counter,
	)

	return &types.MsgInitAccountResponse{ /*Don't really need tracking number or anything here*/ }, nil
}
