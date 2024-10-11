package keeper

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllFiles(c context.Context, req *types.QueryAllFiles) (*types.QueryAllFilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var files []types.UnifiedFile
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilePrimaryKeyPrefix))

	pageRes, err := query.Paginate(store, req.Pagination, func(_ []byte, value []byte) error {
		var file types.UnifiedFile
		if err := k.cdc.Unmarshal(value, &file); err != nil {
			return err
		}

		files = append(files, file)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFilesResponse{Files: files, Pagination: pageRes}, nil
}

func (k Keeper) FilesFromNote(c context.Context, req *types.QueryFilesFromNote) (*types.QueryFilesFromNoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var files []types.UnifiedFile
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FilePrimaryKeyPrefix))

	pageRes, err := query.Paginate(store, req.Pagination, func(_ []byte, value []byte) error {
		var file types.UnifiedFile
		if err := k.cdc.Unmarshal(value, &file); err != nil {
			return err
		}

		var kv map[string]string
		err := json.Unmarshal([]byte(file.Note), &kv)
		if err != nil {
			return nil
		}

		r, exists := kv[req.Key]
		if !exists {
			return nil
		}

		if r != req.Value {
			return nil
		}

		files = append(files, file)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryFilesFromNoteResponse{Files: files, Pagination: pageRes}, nil
}

func (k Keeper) AllFilesByMerkle(c context.Context, req *types.QueryAllFilesByMerkle) (*types.QueryAllFilesByMerkleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var files []types.UnifiedFile
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.FilesMerklePrefix(req.Merkle))

	pageRes, err := query.Paginate(store, req.Pagination, func(_ []byte, value []byte) error {
		var file types.UnifiedFile
		if err := k.cdc.Unmarshal(value, &file); err != nil {
			return err
		}

		files = append(files, file)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFilesByMerkleResponse{Files: files, Pagination: pageRes}, nil
}

func (k Keeper) AllFilesByOwner(c context.Context, req *types.QueryAllFilesByOwner) (*types.QueryAllFilesByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var files []types.UnifiedFile
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.FilesOwnerPrefix(req.Owner))

	pageRes, err := query.Paginate(store, req.Pagination, func(_ []byte, value []byte) error {
		var file types.UnifiedFile
		if err := k.cdc.Unmarshal(value, &file); err != nil {
			return err
		}

		files = append(files, file)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFilesByOwnerResponse{Files: files, Pagination: pageRes}, nil
}

// OpenFiles returns a paginated list of files with space that providers have yet to fill
//
// TODO: Create unit-test cases for this
func (k Keeper) OpenFiles(c context.Context, req *types.QueryOpenFiles) (*types.QueryAllFilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var files []types.UnifiedFile
	ctx := sdk.UnwrapSDKContext(c)

	reverse := false
	var limit uint64 = 100
	if req.Pagination != nil { // HERE IS THE FIX
		reverse = req.Pagination.Reverse
		limit = req.Pagination.Limit
	}

	var i uint64
	var total uint64
	k.IterateFilesByMerkle(ctx, reverse, func(_ []byte, val []byte) bool {

		var file types.UnifiedFile
		if err := k.cdc.Unmarshal(val, &file); err != nil {
			return false
		}

		if file.ContainsProver(req.ProviderAddress) {
			return false
		}

		if len(file.Proofs) < int(file.MaxProofs) {
			total++
			if i >= limit {
				return false
			}
			files = append(files, file)
		} else {
			return false
		}

		i++

		return false
	})

	qpr := query.PageResponse{
		NextKey: nil,
		Total:   total,
	}

	return &types.QueryAllFilesResponse{Files: files, Pagination: &qpr}, nil
}

func (k Keeper) File(c context.Context, req *types.QueryFile) (*types.QueryFileResponse, error) {
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

func (k Keeper) Proof(c context.Context, req *types.QueryProof) (*types.QueryProofResponse, error) {
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

func (k Keeper) AllProofs(c context.Context, req *types.QueryAllProofs) (*types.QueryAllProofsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var proofs []types.FileProof
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	proofStore := prefix.NewStore(store, types.KeyPrefix(types.ProofKeyPrefix))

	pageRes, err := query.Paginate(proofStore, req.Pagination, func(_ []byte, value []byte) error {
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

	return &types.QueryAllProofsResponse{Proofs: proofs, Pagination: pageRes}, nil
}

func (k Keeper) ProofsByAddress(c context.Context, req *types.QueryProofsByAddress) (*types.QueryProofsByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var proofs []types.FileProof
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	proofStore := prefix.NewStore(store, types.ProofPrefix(req.ProviderAddress))

	pageRes, err := query.Paginate(proofStore, req.Pagination, func(_ []byte, value []byte) error {
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

	return &types.QueryProofsByAddressResponse{Proofs: proofs, Pagination: pageRes}, nil
}
