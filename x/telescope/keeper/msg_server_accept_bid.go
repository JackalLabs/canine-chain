package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/telescope/types"
)

func (k msgServer) AcceptBid(goCtx context.Context, msg *types.MsgAcceptBid) (*types.MsgAcceptBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, _ := sdk.AccAddressFromBech32(msg.Creator)

	whois, isFound := k.GetNames(ctx, msg.Name)

	block_height := ctx.BlockHeight()

	if isFound {
		expires, _ := sdk.NewIntFromString(whois.Expires)

		if block_height > expires.Int64() {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
		}

		if whois.Value != owner.String() {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You are not the owner of that name.")
		}

		bid, bidFound := k.GetBids(ctx, msg.From+msg.Name)

		if bidFound {
			cost, _ := sdk.NewIntFromString(bid.Price)
			price := sdk.Coins{sdk.NewInt64Coin("ujkl", cost.Int64())}
			k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, price)

			k.RemoveBids(ctx, msg.From+msg.Name)

			// Create an updated whois record
			newWhois := types.Names{
				Index:   msg.Name,
				Name:    msg.Name,
				Expires: whois.Expires,
				Value:   bid.Bidder,
				Data:    whois.Data,
			}
			// Write whois information to the store
			k.SetNames(ctx, newWhois)
		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Bid does not exist or has expired.")
		}

	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
	}

	return &types.MsgAcceptBidResponse{}, nil
}
