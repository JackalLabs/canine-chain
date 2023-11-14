package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k Keeper) DoReport(ctx sdk.Context, prover string, merkle []byte, owner string, start int64, creator string) error {
	form, found := k.GetReportForm(ctx, prover, merkle, owner, start)
	if !found {
		return sdkerrors.Wrapf(types.ErrAttestInvalid, "cannot find this report")
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
		k.SetReportForm(ctx, form)
		return nil
	}

	deal, found := k.GetFile(ctx, merkle, owner, start)

	if !found {
		return sdkerrors.Wrapf(types.ErrDealNotFound, "cannot find active deal from form")
	}

	k.RemoveReport(ctx, prover, merkle, owner, start)

	deal.RemoveProver(ctx, k, prover)
	return nil
}

func (k msgServer) Report(goCtx context.Context, msg *types.MsgReport) (*types.MsgReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.DoReport(ctx, msg.Prover, msg.Merkle, msg.Owner, msg.Start, msg.Creator)
	if err != nil {
		return nil, err
	}

	return &types.MsgReportResponse{}, nil
}

func (k Keeper) RequestReport(ctx sdk.Context, prover string, merkle []byte, owner string, start int64) ([]string, error) {
	deal, found := k.GetFile(ctx, merkle, owner, start)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrDealNotFound, "cannot find active deal for report form")
	}

	_, found = k.GetReportForm(ctx, prover, merkle, owner, start)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrAttestAlreadyExists, "report form already exists")
	}

	_, err := deal.GetProver(ctx, k, prover)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "not a provider of this file")
	}

	provider, found := k.GetProviders(ctx, prover)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrProviderNotFound, "cannot find provider matching deal")
	}

	providers := k.GetActiveProviders(ctx, provider.Ip) // get a random list of active providers
	params := k.GetParams(ctx)

	if len(providers) < int(params.AttestFormSize) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidLengthQuery, "not enough providers online")
	}

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

	form := types.ReportForm{
		Attestations: attestations,
		Prover:       prover,
		Merkle:       merkle,
		Owner:        owner,
		Start:        start,
	}

	k.SetReportForm(ctx, form)

	return providerAddresses, nil
}

func (k msgServer) RequestReportForm(goCtx context.Context, msg *types.MsgRequestReportForm) (*types.MsgRequestReportFormResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	providerAddresses, err := k.RequestReport(ctx, msg.Prover, msg.Merkle, msg.Owner, msg.Start)

	success := true

	errorString := ""

	if err != nil {
		success = false
		errorString = err.Error()
	}

	return &types.MsgRequestReportFormResponse{
		Providers: providerAddresses,
		Success:   success,
		Error:     errorString,
	}, nil
}
