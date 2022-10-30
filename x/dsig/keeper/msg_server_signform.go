package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackalLabs/canine-chain/x/dsig/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Signform(goCtx context.Context, msg *types.MsgSignform) (*types.MsgSignformResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// finding the form
	iform, err := k.GetForm(ctx, msg.Ffid)
	if !err {
		return nil, sdkerrors.Wrapf(types.NoForm, "The selected form could not be found.")
	}

	// unmarshalling the state of the votes
	voteState := make(map[string]uint64, 8)
	jsonbool := json.Unmarshal([]byte(iform.Signees), &voteState)
	// dont know if this is handled well
	if jsonbool != nil {
		return nil, sdkerrors.Wrapf(types.NoForm, "The selected form's json was malformed")
	}

	// validating user and updating their voting status
	if _, ok := voteState[msg.Creator]; ok {
		if msg.Vote > 4 || msg.Vote < 0 {
			return nil, sdkerrors.Wrapf(types.BadVote, "Unknown vote operation")
		}
		// updating the vote
		voteState[msg.Creator] = uint64(msg.Vote)
	} else {
		return nil, sdkerrors.Wrapf(types.BadUser, "You are not a signee")
	}

	// executing function if consensus is reached
	consensus := true
	for _, val := range voteState {
		if val != 1 {
			consensus = false
			break
		}
	}
	if consensus {
		// more creative functions to be implemented here
		fmt.Println("consensus reached!")
	}

	// marshalling the updated vote status
	updatedState, _ := json.Marshal(voteState)
	updatedStatestr := string(updatedState)

	// uploading back into the form
	updatedForm := types.Form{
		Ffid:    iform.Ffid,
		Cid:     iform.Cid,
		Fid:     iform.Fid,
		Signees: updatedStatestr,
	}
	k.SetForm(ctx, updatedForm)

	return &types.MsgSignformResponse{}, nil
}
