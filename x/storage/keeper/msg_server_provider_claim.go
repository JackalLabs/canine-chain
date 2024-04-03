package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k msgServer) AddProviderClaimer(goCtx context.Context, msg *types.MsgAddClaimer) (*types.MsgAddClaimerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, found := k.GetProviders(ctx, msg.Creator)

	if !found {
		return nil, types.ErrProviderNotFound
	}

	if provider.AuthClaimers == nil {
		provider.AuthClaimers = []string{}
	}

	for _, claimer := range provider.AuthClaimers {
		if claimer == msg.ClaimAddress {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrConflict, "cannot add the same claimer twice")
		}
	}

	provider.AuthClaimers = append(provider.AuthClaimers, msg.ClaimAddress)

	k.SetProviders(ctx, provider)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgAddClaimerResponse{}, nil
}

func (k msgServer) RemoveProviderClaimer(goCtx context.Context, msg *types.MsgRemoveClaimer) (*types.MsgRemoveClaimerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, found := k.GetProviders(ctx, msg.Creator)
	authClaimers := provider.AuthClaimers

	if !found {
		return nil, types.ErrProviderNotFound
	}
	p := len(authClaimers)
	if p == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrConflict, "Provider has no claimer addresses")
	}

	newClaim := []string{}

	for i := 0; i < len(authClaimers); i++ {
		if authClaimers[i] != msg.ClaimAddress {
			newClaim = append(newClaim, authClaimers[i])
		}
	}
	provider.AuthClaimers = newClaim

	if p == len(provider.AuthClaimers) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "this address is not a claimer")
	}

	k.SetProviders(ctx, provider)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgRemoveClaimerResponse{}, nil
}
