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

func (k Keeper) FilesAll(c context.Context, req *types.QueryAllFilesRequest) (*types.QueryAllFilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var activeDealss []types.UnifiedFile
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	activeDealsStore := prefix.NewStore(store, types.KeyPrefix(types.FilePrimaryKeyPrefix))

	pageRes, err := query.Paginate(activeDealsStore, req.Pagination, func(key []byte, value []byte) error {
		var activeDeals types.UnifiedFile
		if err := k.cdc.Unmarshal(value, &activeDeals); err != nil {
			return err
		}

		activeDealss = append(activeDealss, activeDeals)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFilesResponse{Files: activeDealss, Pagination: pageRes}, nil
}

// OpenFiles returns a paginated list of files with space that providers have yet to fill
//
// TODO: Create unit-test cases for this
func (k Keeper) OpenFiles(c context.Context, req *types.QueryOpenFilesRequest) (*types.QueryAllFilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var files []types.UnifiedFile
	ctx := sdk.UnwrapSDKContext(c)

	var i uint64
	k.IterateFilesByMerkle(ctx, req.Pagination.Reverse, func(key []byte, val []byte) bool {
		if i >= req.Pagination.Limit {
			return true
		}

		var file types.UnifiedFile
		if err := k.cdc.Unmarshal(val, &file); err != nil {
			return false
		}

		if len(file.Proofs) < int(file.MaxProofs) {
			files = append(files, file)
		}

		i++

		return false
	})

	qpr := query.PageResponse{
		NextKey: nil,
		Total:   i,
	}

	return &types.QueryAllFilesResponse{Files: files, Pagination: &qpr}, nil
}

func (k Keeper) File(c context.Context, req *types.QueryFileRequest) (*types.QueryFileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetFile(
		ctx,
		req.Merkle,
		req.Owner,
		req.Start,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryFileResponse{File: val}, nil
}

func (k Keeper) Proof(c context.Context, req *types.QueryProofRequest) (*types.QueryProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProof(
		ctx,
		req.ProviderAddress,
		req.Merkle,
		req.Owner,
		req.Start,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryProofResponse{Proof: val}, nil
}

func (k Keeper) ProofsAll(c context.Context, req *types.QueryProofsByAddressRequest) (*types.QueryProofsByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var proofs []types.FileProof
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	proofStore := prefix.NewStore(store, types.ProofPrefix(req.ProviderAddress))

	pageRes, err := query.Paginate(proofStore, req.Pagination, func(key []byte, value []byte) error {
		var proof types.FileProof
		if err := k.cdc.Unmarshal(value, &proof); err != nil {
			return err
		}

		proofs = append(proofs, proof)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryProofsByAddressResponse{Proof: proofs, Pagination: pageRes}, nil
}
