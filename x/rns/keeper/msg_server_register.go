package keeper

import (
	"context"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackal-dao/canine/x/rns/types"
)

func (k msgServer) Register(goCtx context.Context, msg *types.MsgRegister) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Try getting a name from the store
	whois, isFound := k.GetNames(ctx, msg.Name)
	// Set the price at which the name has to be bought if it didn't have an owner before

	chars := strings.Count(msg.Name, "")

	cost := 1000000

	switch chars {
	case 0:
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Must be 1 or more characters.")
	case 1:
		cost = 12000000
		break
	case 2:
		cost = 6000000
		break
	case 3:
		cost = 3000000
		break
	case 4:
		cost = 1500000
		break
	case 5:
		cost = 750000
		break
	default:
		cost = 375000
		break
	}

	price := sdk.Coins{sdk.NewInt64Coin("ujkl", int64(cost))}

	num_years, _ := sdk.NewIntFromString(msg.Years)

	var block_height = ctx.BlockHeight()

	var time = num_years.Int64() * 6311520

	owner, _ := sdk.AccAddressFromBech32(msg.Creator)
	// If a name is found in store
	if isFound {

		expires, _ := sdk.NewIntFromString(whois.Expires)

		if whois.Value == owner.String() {
			time = expires.Int64() + time
		} else {
			if block_height < expires.Int64() {
				return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name already registered")
			}
		}

	} else {
		time = time + block_height
	}

	k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, price)

	// Create an updated whois record
	newWhois := types.Names{
		Index:   msg.Name,
		Name:    msg.Name,
		Expires: strconv.FormatInt(time, 10),
		Value:   owner.String(),
		Data:    msg.Data,
	}
	// Write whois information to the store
	k.SetNames(ctx, newWhois)
	return &types.MsgRegisterResponse{}, nil

}
