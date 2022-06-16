package keeper

import (
	"context"

	"crypto/sha256"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/jklmining/types"
)

func (k msgServer) ClaimSave(goCtx context.Context, msg *types.MsgClaimSave) (*types.MsgClaimSaveResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	savefile, _ := k.GetSaveRequests(
		ctx,
		msg.Saveindex,
	)

	if savefile.Approved == "true" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Already claimed.")
	}

	sum := sha256.Sum256([]byte(msg.Key))
	s := strings.ToUpper(fmt.Sprintf("%x", sum))
	i := strings.ToUpper(savefile.Index)
	if s != i {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s is not %s", s, i))
	}

	var saveRequests = types.SaveRequests{
		Creator:  savefile.Creator,
		Index:    savefile.Index,
		Size_:    savefile.Size_,
		Approved: "true",
	}

	k.SetSaveRequests(
		ctx,
		saveRequests,
	)

	return &types.MsgClaimSaveResponse{}, nil
}
