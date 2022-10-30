package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/x/filetree/types"

	rnsTypes "github.com/jackal-dao/canine/x/rns/types"
)

func (k msgServer) InitAll(goCtx context.Context, msg *types.MsgInitAll) (*types.MsgInitAllResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pubKey := types.Pubkey{

		Key:     msg.Pubkey,
	}

	k.SetPubkey(ctx, pubKey)

	_, found := k.rnsKeeper.GetInit(ctx, msg.Creator)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot initialize more than once")
	}

	i := rnsTypes.Init{
		Address:  msg.Creator,
		Complete: true,
	}

	k.rnsKeeper.SetInit(ctx, i)

	bh := ctx.BlockHeight()
	name := rnsTypes.MakeName(int(bh), bh)

	if strings.Contains(name, ".") {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name cannot contain '.'")
	}

	if len(name) < 6 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name cannot be less than 6 characters")
	}

	whois, isFound := k.rnsKeeper.GetNames(ctx, name, "jkl")
	block_height := ctx.BlockHeight()

	if isFound {
		if block_height < whois.Expires {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name already registered")
		}
	}

	time := 6311520 + block_height

	emptySubdomains := []*rnsTypes.Names{}

	// Create an updated whois record
	newWhois := rnsTypes.Names{
		Name:       name,
		Expires:    time,
		Value:      msg.Creator,
		Data:       "{}",
		Subdomains: emptySubdomains,
		Tld:        "jkl",
		Locked:     time,
	}
	// Write whois information to the store
	k.rnsKeeper.SetNames(ctx, newWhois)

	return &types.MsgInitAllResponse{Name: fmt.Sprintf("%s.jkl", name)}, nil
}
