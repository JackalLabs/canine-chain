package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k msgServer) CreateContracts(goCtx context.Context, msg *types.MsgCreateContracts) (*types.MsgCreateContractsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetContracts(
		ctx,
		msg.Cid,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	contracts := types.Contracts{
		Creator:    msg.Creator,
		Cid:        msg.Cid,
		Priceamt:   msg.Priceamt,
		Pricedenom: msg.Pricedenom,
		Merkle:     msg.Merkle,
		Signee:     msg.Signee,
		Duration:   msg.Duration,
		Filesize:   msg.Filesize,
		Fid:        msg.Fid,
	}

	k.SetContracts(
		ctx,
		contracts,
	)
	return &types.MsgCreateContractsResponse{}, nil
}
