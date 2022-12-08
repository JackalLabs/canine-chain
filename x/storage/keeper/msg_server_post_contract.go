package keeper

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k msgServer) PostContract(goCtx context.Context, msg *types.MsgPostContract) (*types.MsgPostContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, ok := k.GetProviders(ctx, msg.Creator)
	if !ok {
		return nil, fmt.Errorf("can't find provider")
	}

	ts, ok := sdk.NewIntFromString(provider.Totalspace)

	if !ok {
		return nil, fmt.Errorf("error parsing total space")
	}

	fs, ok := sdk.NewIntFromString(msg.Filesize)

	if !ok {
		return nil, fmt.Errorf("error parsing file size")
	}

	if k.GetProviderUsing(ctx, msg.Creator)+fs.Int64() > ts.Int64() {
		return nil, fmt.Errorf("not enough space on provider")
	}

	filesize, ok := sdk.NewIntFromString(msg.Filesize)
	if !ok {
		return nil, fmt.Errorf("cannot parse filesize")
	}

	payInfo, found := k.GetStoragePaymentInfo(ctx, msg.Signee)
	if !found {
		if filesize.Int64() > TwoGigs {
			return nil, fmt.Errorf("cannot save files greater than 2gb on a free plan")
		}
	} else {
		if payInfo.SpaceUsed+filesize.Int64() > payInfo.SpaceAvailable {
			return nil, fmt.Errorf("not enough storage on the users account")
		}
	}

	h := sha256.New()
	_, err := io.WriteString(h, fmt.Sprintf("%s%s%s", msg.Signee, msg.Creator, msg.Fid))
	if err != nil {
		return nil, err
	}
	hashName := h.Sum(nil)

	cid, err := MakeCid(hashName)
	if err != nil {
		return nil, err
	}

	_, cidtaken := k.GetContracts(ctx, cid)
	if cidtaken {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "cannot post the same contract twice")
	}

	newContract := types.Contracts{
		Cid:      cid,
		Signee:   msg.Signee,
		Fid:      msg.Fid,
		Filesize: msg.Filesize,
		Creator:  msg.Creator,
		Merkle:   msg.Merkle,
	}

	k.SetContracts(ctx, newContract)

	return &types.MsgPostContractResponse{}, nil
}
