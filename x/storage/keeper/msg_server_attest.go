package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func (k Keeper) Attest(ctx sdk.Context, prover string, merkle []byte, owner string, start int64, creator string) error {
	form, found := k.GetAttestationForm(ctx, prover, merkle, owner, start)
	if !found {
		return sdkerrors.Wrapf(types.ErrAttestInvalid, "cannot find this attestation")
	}

	done := false

	var count int64

	attestations := form.Attestations
	for _, attestation := range attestations {
		if attestation.Provider == creator {
			attestation.Complete = true
			done = true
		}

		if attestation.Complete {
			count++
		}
	}

	if !done {
		return sdkerrors.Wrapf(types.ErrAttestInvalid, "you cannot attest to this deal")
	}

	if count < k.GetParams(ctx).AttestMinToPass {
		form.Attestations = attestations
		k.SetAttestationForm(ctx, form)
		return nil
	}

	deal, found := k.GetFile(ctx, form.Merkle, form.Owner, form.Start)

	if !found {
		return sdkerrors.Wrapf(types.ErrDealNotFound, "cannot find active deal from form")
	}

	proof, err := deal.GetProver(ctx, k, form.Prover)
	if err != nil {
		return sdkerrors.Wrapf(err, "can't find proof for attestation")
	}

	proof.LastProven = ctx.BlockHeight()

	k.SetProof(ctx, *proof)

	k.RemoveAttestation(ctx, form.Prover, form.Merkle, form.Owner, form.Start)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return nil
}

func (k msgServer) Attest(goCtx context.Context, msg *types.MsgAttest) (*types.MsgAttestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.Attest(ctx, msg.Prover, msg.Merkle, msg.Owner, msg.Start, msg.Creator)
	if err != nil {
		ctx.Logger().Debug(err.Error())
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJackalMessage,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	return &types.MsgAttestResponse{}, nil
}

func (k Keeper) RequestAttestation(ctx sdk.Context, merkle []byte, owner string, start int64, creator string) ([]string, error) {
	deal, found := k.GetFile(ctx, merkle, owner, start)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrDealNotFound, "cannot find active deal for attestation form")
	}

	_, err := deal.GetProver(ctx, k, creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "you are not a prover for this file")
	}

	_, found = k.GetAttestationForm(ctx, creator, merkle, owner, start)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrAttestAlreadyExists, "attestation form already exists")
	}

	provider, found := k.GetProviders(ctx, creator)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrProviderNotFound, "cannot find provider matching deal")
	}

	providers := k.GetActiveProviders(ctx, provider.Ip) // get a random list of active providers
	params := k.GetParams(ctx)

	if len(providers) < int(params.AttestFormSize) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidLengthQuery, "not enough providers online")
	}

	rand.Seed(ctx.BlockHeight())

	attestations := make([]*types.Attestation, params.AttestFormSize)

	providerAddresses := make([]string, params.AttestFormSize)

	for i := 0; i < int(params.AttestFormSize); i++ {
		p := providers[i]

		providerAddresses[i] = p.Address

		attestations[i] = &types.Attestation{
			Provider: p.Address,
			Complete: false,
		}
	}

	form := types.AttestationForm{
		Attestations: attestations,
		Prover:       creator,
		Merkle:       merkle,
		Owner:        owner,
		Start:        start,
	}

	k.SetAttestationForm(ctx, form)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return providerAddresses, nil
}

func (k msgServer) RequestAttestationForm(goCtx context.Context, msg *types.MsgRequestAttestationForm) (*types.MsgRequestAttestationFormResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	creator := msg.Creator

	merkle := msg.Merkle
	owner := msg.Owner
	start := msg.Start

	providerAddresses, err := k.RequestAttestation(ctx, merkle, owner, start, creator)

	success := true

	errorString := ""

	if err != nil {
		success = false
		errorString = err.Error()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJackalMessage,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
		),
	)

	return &types.MsgRequestAttestationFormResponse{
		Providers: providerAddresses,
		Success:   success,
		Error:     errorString,
	}, nil
}
