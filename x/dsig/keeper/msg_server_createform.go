package keeper

import (
	"context"
	"crypto/md5" //nolint:gosec // NOTE: THIS IS FLAGGED AS A WEAK CRYPTOGRAPHIC PRIMITIVE.
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/jackal-dao/canine/x/dsig/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Createform(goCtx context.Context, msg *types.MsgCreateform) (*types.MsgCreateformResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// created by Maximilian Pochapski, July 28, 2022

	// checking if the fid exists
	filedata, exists := k.GetUserUploads(ctx, msg.Fid)
	if !exists {
		return nil, sdkerrors.Wrapf(types.NoFile, "The selected file does not exist")
	}
	// checking if the file was made by the creator
	if filedata.Cid != msg.Creator {
		return nil, sdkerrors.Wrapf(types.NotOwner, "Permission denied: not owner")
	}

	// converting the createdAt signature
	bd := make([]byte, 8)
	binary.BigEndian.PutUint64(bd, uint64(filedata.CreatedAt))
	// creating the ffid of the form type
	// seed is the concatenation of the fid, blockheight, and cid
	seed := append(bd, []byte(filedata.Fid)...)
	seed = append(seed, []byte(filedata.Cid)...)
	ffid := md5.Sum(seed)

	// formatting the signees
	// signees delimited as "uid1 uid2 uid3 uid4"
	signees := strings.Split(msg.Signees, " ")
	// verifying the users exist
	for _, signee := range signees {
		// formatting the search term
		searchuser, _ := sdk.AccAddressFromBech32(signee)
		// using the accountkeeper to verify the signee exists on-chain
		verif := k.accountKeeper.GetAccount(ctx, searchuser)
		if verif == nil {
			return nil, sdkerrors.Wrapf(types.InvalidSignee, "Signee accounts do not exist")
		}
	}
	// marshalling the signees into json format
	formMap := make(map[string]uint64)
	for _, signee := range signees {
		// encoding the signature state of the signees
		// 0 = disapproved, 1 = approved, 2 = abstained, 3 = no response
		formMap[signee] = 3
	}
	// maybe do error correcting instead of _?
	marshalledsignee, _ := json.Marshal(formMap)
	signeejson := string(marshalledsignee)

	// creating the form type
	newForm := types.Form{
		Cid:     msg.Creator,
		Fid:     msg.Fid,
		Signees: signeejson,
		Ffid:    hex.EncodeToString(ffid[:]),
	}

	// checking if form exists
	if _, ok := k.GetForm(ctx, newForm.Ffid); ok {
		// in the future, make the program determine a new FFID for the file
		return nil, sdkerrors.Wrapf(types.DuplicateForm, "Form already exists")
	}
	k.SetForm(ctx, newForm)

	return &types.MsgCreateformResponse{Ffid: newForm.Ffid}, nil
}
