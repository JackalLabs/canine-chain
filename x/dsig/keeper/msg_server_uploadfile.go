package keeper

import (
	"context"

	"github.com/jackalLabs/canine-chain/x/dsig/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Uploadfile(goCtx context.Context, msg *types.MsgUploadfile) (*types.MsgUploadfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// creating the file upload type
	uF := types.UserUploads{
		Fid:       msg.Fid,
		Cid:       msg.Creator,
		CreatedAt: ctx.BlockHeight(),
	}
	// checking if file exists for selected fid
	if _, ok := k.GetUserUploads(ctx, uF.Fid); ok {
		return nil, sdkerrors.Wrapf(types.DuplicateFid, "File already exists for selected FID")
	}
	// Uploading to the local store
	k.SetUserUploads(ctx, uF)

	return &types.MsgUploadfileResponse{}, nil
}
