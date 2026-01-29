package keeper

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// paginationParams holds extracted pagination parameters
type paginationParams struct {
	reverse bool
	limit   uint64
	offset  uint64
}

// extractPaginationParams extracts pagination parameters with defaults
func extractPaginationParams(pagination *query.PageRequest) paginationParams {
	p := paginationParams{
		reverse: false,
		limit:   100,
		offset:  0,
	}
	if pagination != nil {
		p.reverse = pagination.Reverse
		if pagination.Limit > 0 {
			p.limit = pagination.Limit
		}
		p.offset = pagination.Offset
	}
	return p
}

// filterFilesWithPagination iterates over files and returns those matching the filter with pagination
func (k Keeper) filterFilesWithPagination(ctx sdk.Context, params paginationParams, filter func(*types.UnifiedFile) bool) ([]types.UnifiedFile, *query.PageResponse) {
	var files []types.UnifiedFile
	var skipped, collected, total uint64

	k.IterateFilesByMerkle(ctx, params.reverse, func(_ []byte, val []byte) bool {
		var file types.UnifiedFile
		if err := k.cdc.Unmarshal(val, &file); err != nil {
			return false
		}

		if filter(&file) {
			total++
			if skipped < params.offset {
				skipped++
				return false
			}
			if collected < params.limit {
				files = append(files, file)
				collected++
			}
		}
		return false
	})

	return files, &query.PageResponse{Total: total}
}

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

		var kv map[string]any
		err := json.Unmarshal([]byte(file.Note), &kv)
		if err != nil {
			return nil
		}

		r, exists := kv[req.Key]
		if !exists {
			return nil
		}

		s, ok := r.(string)
		if !ok {
			return nil
		}

		if s != req.Value {
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

// AllFilesByOwner returns a paginated list of files owned by a specific address
func (k Keeper) AllFilesByOwner(c context.Context, req *types.QueryAllFilesByOwner) (*types.QueryAllFilesByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.Owner == "" {
		return nil, status.Error(codes.InvalidArgument, "owner address is required")
	}

	if _, err := sdk.AccAddressFromBech32(req.Owner); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid owner address format")
	}

	ctx := sdk.UnwrapSDKContext(c)
	params := extractPaginationParams(req.Pagination)

	files, pageRes := k.filterFilesWithPagination(ctx, params, func(file *types.UnifiedFile) bool {
		return file.Owner == req.Owner
	})

	return &types.QueryAllFilesByOwnerResponse{Files: files, Pagination: pageRes}, nil
}

// OpenFiles returns a paginated list of files with space that providers have yet to fill
func (k Keeper) OpenFiles(c context.Context, req *types.QueryOpenFiles) (*types.QueryAllFilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	params := extractPaginationParams(req.Pagination)

	files, pageRes := k.filterFilesWithPagination(ctx, params, func(file *types.UnifiedFile) bool {
		return !file.ContainsProver(req.ProviderAddress) && len(file.Proofs) < int(file.MaxProofs)
	})

	return &types.QueryAllFilesResponse{Files: files, Pagination: pageRes}, nil
}

// EndangeredFiles returns a paginated list of files with only 1x redundancy
func (k Keeper) EndangeredFiles(c context.Context, req *types.QueryOpenFiles) (*types.QueryAllFilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	params := extractPaginationParams(req.Pagination)

	files, pageRes := k.filterFilesWithPagination(ctx, params, func(file *types.UnifiedFile) bool {
		return !file.ContainsProver(req.ProviderAddress) && len(file.Proofs) == 1
	})

	return &types.QueryAllFilesResponse{Files: files, Pagination: pageRes}, nil
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
