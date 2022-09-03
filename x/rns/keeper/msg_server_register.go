package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackal-dao/canine/x/rns/types"
)

func (k msgServer) Register(goCtx context.Context, msg *types.MsgRegister) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Try getting a name from the store

	name, tld, err := getNameAndTLD(msg.Name)
	if err != nil {
		return nil, err
	}

	if types.IS_RESERVED[tld] {
		return nil, types.ErrReserved
	}

	whois, isFound := k.GetNames(ctx, name, tld)
	// Set the price at which the name has to be bought if it didn't have an owner before

	chars := strings.Count(name, "")

	var baseCost int64 = getCost(tld)

	var cost int64

	switch chars {
	case 0:
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Must be 1 or more characters.")
	case 1:
		cost = baseCost * 32
	case 2:
		cost = baseCost * 16
	case 3:
		cost = baseCost * 8
	case 4:
		cost = baseCost * 4
	case 5:
		cost = baseCost * 2
	default:
		cost = baseCost
	}

	price := sdk.Coins{sdk.NewInt64Coin("ujkl", cost)}

	num_years, _ := sdk.NewIntFromString(msg.Years)

	var block_height = ctx.BlockHeight()

	var time = num_years.Int64() * 6311520

	owner, _ := sdk.AccAddressFromBech32(msg.Creator)
	// If a name is found in store
	if isFound {

		if whois.Value == owner.String() {
			time = whois.Expires + time
		} else {
			if block_height < whois.Expires {
				return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name already registered")
			}
		}

	} else {
		time = time + block_height
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, price)
	if err != nil {
		return nil, err
	}

	emptySubdomains := []*types.Names{}

	// Create an updated whois record
	newWhois := types.Names{
		Name:       name,
		Expires:    time,
		Value:      owner.String(),
		Data:       msg.Data,
		Subdomains: emptySubdomains,
		Tld:        tld,
		Locked:     0,
	}
	// Write whois information to the store
	k.SetNames(ctx, newWhois)
	return &types.MsgRegisterResponse{}, nil

}
