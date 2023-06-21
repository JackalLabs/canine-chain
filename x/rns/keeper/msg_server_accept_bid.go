package keeper

import (
	"context"
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/rns/types"
)

func (k Keeper) AcceptOneBid(ctx sdk.Context, sender string, name string, bidder string) error {
	name = strings.ToLower(name)
	owner, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return err
	}

	n, tld, err := GetNameAndTLD(name)
	if err != nil {
		return err
	}

	whois, isFound := k.GetNames(ctx, n, tld)

	blockHeight := ctx.BlockHeight()

	if !isFound {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
	}

	if blockHeight > whois.Expires {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
	}

	if whois.Value != owner.String() {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You are not the owner of that name.")
	}

	if whois.Locked > blockHeight {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "cannot transfer free name")
	}

	bid, bidFound := k.GetBids(ctx, fmt.Sprintf("%s%s", bidder, name))

	if !bidFound {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "Bid does not exist or has expired.")
	}

	price, err := sdk.ParseCoinsNormalized(bid.Price)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, price)
	if err != nil {
		return err
	}

	k.RemoveBids(ctx, fmt.Sprintf("%s%s", bidder, name))

	whois.Value = bid.Bidder
	whois.Data = "{}"

	// Write whois information to the store
	k.SetNames(ctx, whois)

	return nil
}

func (k msgServer) AcceptBid(goCtx context.Context, msg *types.MsgAcceptBid) (*types.MsgAcceptBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.AcceptOneBid(ctx, msg.Creator, msg.Name, msg.From)

	return &types.MsgAcceptBidResponse{}, err
}
