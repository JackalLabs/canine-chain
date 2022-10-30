package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/rns/types"
)

func (k msgServer) AcceptBid(goCtx context.Context, msg *types.MsgAcceptBid) (*types.MsgAcceptBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, _ := sdk.AccAddressFromBech32(msg.Creator)

	n, tld, err := getNameAndTLD(msg.Name)
	if err != nil {
		return nil, err
	}

	whois, isFound := k.GetNames(ctx, n, tld)

	blockHeight := ctx.BlockHeight()

	if isFound {

		if blockHeight > whois.Expires {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
		}

		if whois.Value != owner.String() {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You are not the owner of that name.")
		}

		if whois.Locked > blockHeight {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "cannot transfer free name")
		}

		bid, bidFound := k.GetBids(ctx, msg.From+msg.Name)

		if bidFound {
			cost, _ := sdk.NewIntFromString(bid.Price)
			price := sdk.Coins{sdk.NewInt64Coin("ujkl", cost.Int64())}
			err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, price)
			if err != nil {
				return nil, err
			}

			k.RemoveBids(ctx, msg.From+msg.Name)

			whois.Value = bid.Bidder
			whois.Data = "{}"

			// Write whois information to the store
			k.SetNames(ctx, whois)
		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Bid does not exist or has expired.")
		}

	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
	}

	return &types.MsgAcceptBidResponse{}, nil
}
