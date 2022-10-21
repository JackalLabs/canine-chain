package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) ResetViewers(goCtx context.Context, msg *types.MsgResetViewers) (*types.MsgResetViewersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	logger, logFile := createLogger()

	file, found := k.GetFiles(ctx, msg.Address, msg.Fileowner)
	if !found {
		return nil, types.ErrFileNotFound
	}

	isOwner := IsOwner(file, msg.Creator)
	if !isOwner {
		return nil, types.ErrNotOwner
	}

	logger.Println("The current viewer Access is", file.ViewingAccess)

	ownerViewerAddress := MakeViewerAddress(file.TrackingNumber, msg.Creator)

	pvacc := file.ViewingAccess
	//Unmarshall current edit access to this blank map
	jvacc := make(map[string]string)
	json.Unmarshal([]byte(pvacc), &jvacc)

	ownerKey := jvacc[ownerViewerAddress]

	resetViewers := make(map[string]string)
	resetViewers[ownerViewerAddress] = ownerKey

	vaccbytes, err := json.Marshal(resetViewers)
	if err != nil {
		return nil, types.ErrCantMarshall
	}
	newViewers := string(vaccbytes)

	file.ViewingAccess = newViewers

	k.SetFiles(ctx, file)

	logger.Println("The reset viewer Access is", file.ViewingAccess)
	logger.Println()
	logFile.Close()

	return &types.MsgResetViewersResponse{}, nil
}
