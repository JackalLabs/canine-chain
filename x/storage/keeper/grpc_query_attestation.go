package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllAttestations(c context.Context, req *types.QueryAllAttestations) (*types.QueryAllAttestationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var attestations []types.AttestationForm
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	attestationStore := prefix.NewStore(store, types.KeyPrefix(types.AttestationKeyPrefix))

	pageRes, err := query.Paginate(attestationStore, req.Pagination, func(key []byte, value []byte) error {
		var providers types.AttestationForm
		if err := k.cdc.Unmarshal(value, &providers); err != nil {
			return err
		}

		attestations = append(attestations, providers)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAttestationsResponse{Attestations: attestations, Pagination: pageRes}, nil
}

func (k Keeper) Attestation(c context.Context, req *types.QueryAttestation) (*types.QueryAttestationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAttestationForm(
		ctx,
		req.Prover,
		req.Merkle,
		req.Owner,
		req.Start,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryAttestationResponse{Attestation: val}, nil
}
