package keeper

import (
	"context"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/rns/types"
)

func (k msgServer) Init(goCtx context.Context, msg *types.MsgInit) (*types.MsgInitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetInit(ctx, msg.Creator)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot initialize more than once")
	}

	i := types.Init{
		Address:  msg.Creator,
		Complete: true,
	}

	k.SetInit(ctx, i)

	bh := ctx.BlockHeight()

	name := types.MakeName(int(bh), bh)

	if strings.Contains(name, ".") {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name cannot contain '.'")
	}

	if len(name) < 6 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name cannot be less than 6 characters")
	}

	whois, isFound := k.GetNames(ctx, name, "jkl")
	var block_height = ctx.BlockHeight()

	if isFound {

		if block_height < whois.Expires {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name already registered")
		}

	}

	var time = 6311520 + block_height

	emptySubdomains := []*types.Names{}

	// Create an updated whois record
	newWhois := types.Names{
		Name:       name,
		Expires:    time,
		Value:      msg.Creator,
		Data:       "{}",
		Subdomains: emptySubdomains,
		Tld:        "jkl",
		Locked:     time,
	}
	// Write whois information to the store
	k.SetNames(ctx, newWhois)

	return &types.MsgInitResponse{}, nil
}
