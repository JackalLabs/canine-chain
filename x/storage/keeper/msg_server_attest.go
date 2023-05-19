package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

const ( // TODO: Figure out optimal values for these and replace them with chain params
	FormSize  = 5
	MinToPass = 3
	True      = "true"
)

func (k msgServer) Attest(goCtx context.Context, msg *types.MsgAttest) (*types.MsgAttestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	form, found := k.GetAttestationForm(ctx, msg.Cid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrFormNotFound, "cannot find attestation form")
	}

	done := false

	count := 0

	attestations := form.Attestations
	for _, attestation := range attestations {
		if attestation.Provider == msg.Creator {
			attestation.Complete = true
			done = true
		}

		if attestation.Complete {
			count++
		}
	}

	if !done {
		return nil, sdkerrors.Wrapf(types.ErrAttestInvalid, "you cannot attest to this deal")
	}

	if count < MinToPass {
		form.Attestations = attestations
		k.SetAttestationForm(ctx, form)
		return &types.MsgAttestResponse{}, nil
	}

	deal, found := k.GetActiveDeals(ctx, msg.Cid)

	if !found {
		return nil, sdkerrors.Wrapf(types.ErrDealNotFound, "cannot find active deal from form")
	}

	deal.Proofverified = True // flip deal to verified if the minimum attestation threshold is met
	k.SetActiveDeals(ctx, deal)
	k.RemoveAttestation(ctx, msg.Cid)

	return &types.MsgAttestResponse{}, nil
}

func (k msgServer) RequestAttestationForm(goCtx context.Context, msg *types.MsgRequestAttestationForm) (*types.MsgRequestAttestationFormResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cid := msg.Cid

	deal, found := k.GetActiveDeals(ctx, cid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrDealNotFound, "cannot find active deal for attestation form")
	}

	if deal.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "you do not own this deal")
	}

	providers := k.GetActiveProviders(ctx) // get a random list of active providers

	rand.Seed(ctx.BlockTime().UnixNano())

	attestations := make([]*types.Attestation, FormSize)

	providerAddresses := make([]string, FormSize)

	for i := 0; i < FormSize; i++ {
		p := providers[i]

		providerAddresses[i] = p.Address

		attestations[i] = &types.Attestation{
			Provider: p.Address,
			Complete: false,
		}
	}

	form := types.AttestationForm{
		Attestations: attestations,
		Cid:          cid,
	}

	k.SetAttestationForm(ctx, form)

	return &types.MsgRequestAttestationFormResponse{Providers: providerAddresses}, nil
}
