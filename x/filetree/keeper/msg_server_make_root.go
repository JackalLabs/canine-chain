package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/filetree/types"
	notiTypes "github.com/jackal-dao/canine/x/notifications/types"
)

func (k msgServer) MakeRoot(goCtx context.Context, msg *types.MsgMakeRoot) (*types.MsgMakeRootResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)


	//make the full OwnerAddress

	ownerAddress := MakeOwnerAddress(msg.RootHashPath, msg.Account)

	file := types.Files{

		Owner:          ownerAddress,

		EditAccess:     msg.Editors,
		Address:        msg.RootHashPath,
		TrackingNumber: msg.TrackingNumber,
	}

	k.SetFiles(ctx, file)


	// Check if the counter already exists
	_, isFound := k.notiKeeper.GetNotiCounter(
		ctx,
		msg.Creator,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "notiCounter already set")
	}

	placeholderMap := make([]string, 0, 2000)
	placeholderMap = append(placeholderMap, msg.Creator)
	marshalledSenders, err := json.Marshal(placeholderMap)
	if err != nil {
		return nil, types.ErrCantUnmarshall
	}

	updatedSenders := string(marshalledSenders)

	counter := notiTypes.NotiCounter{
		Address:          msg.Creator,
		Counter:          0,
		PermittedSenders: updatedSenders,
	}

	k.notiKeeper.SetNotiCounter(
		ctx,
		counter,
	)

	return &types.MsgMakeRootResponse{}, nil
}
