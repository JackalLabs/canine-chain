package keeper

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
)

func getPage(pg *query.PageRequest) (uint64, uint64) {
	// Determine pagination parameters
	var limit uint64 = 100
	var offset uint64 = 0
	if pg != nil {
		limit = pg.Limit
		if pg.Offset > 0 {
			offset = pg.Offset
		} else if pg.Key != nil {
			parts := strings.Split(string(pg.Key), ":")
			if len(parts) == 2 {
				parsedOffset, err := strconv.ParseUint(parts[0], 10, 64)
				if err == nil {
					offset = parsedOffset
				}
			}
		}
	}

	return limit, offset
}

func (k Keeper) AllFiles(ctx context.Context, req *types.QueryAllFiles) (*types.QueryAllFilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	limit, offset := getPage(req.Pagination)

	files, total := k.GetAllFileByMerklePg(offset, limit)

	// Pagination response
	nextOffset := offset + uint64(len(files))
	nextKey := []byte(fmt.Sprintf("%d:%d", nextOffset, limit))
	if nextOffset >= uint64(total) {
		nextKey = nil
	}

	pageRes := &query.PageResponse{
		NextKey: nextKey,
		Total:   uint64(total),
	}

	return &types.QueryAllFilesResponse{Files: files, Pagination: pageRes}, nil
}

func (k Keeper) FilesFromNote(c context.Context, req *types.QueryFilesFromNote) (*types.QueryFilesFromNoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	limit, offset := getPage(req.Pagination)

	files, total := k.GetAllFileByMerklePgWithJSONFilter(offset, limit, req.Key, req.Value)

	// Pagination response
	nextOffset := offset + uint64(len(files))
	nextKey := []byte(fmt.Sprintf("%d:%d", nextOffset, limit))
	if nextOffset >= uint64(total) {
		nextKey = nil
	}

	pageRes := &query.PageResponse{
		NextKey: nextKey,
		Total:   uint64(total),
	}

	return &types.QueryFilesFromNoteResponse{Files: files, Pagination: pageRes}, nil
}

func (k Keeper) AllFilesByMerkle(c context.Context, req *types.QueryAllFilesByMerkle) (*types.QueryAllFilesByMerkleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	limit, offset := getPage(req.Pagination)

	files, total := k.GetAllFilesWithMerklePg(req.Merkle, limit, offset)

	// Pagination response
	nextOffset := offset + uint64(len(files))
	nextKey := []byte(fmt.Sprintf("%d:%d", nextOffset, limit))
	if nextOffset >= uint64(total) {
		nextKey = nil
	}

	pageRes := &query.PageResponse{
		NextKey: nextKey,
		Total:   uint64(total),
	}

	return &types.QueryAllFilesByMerkleResponse{Files: files, Pagination: pageRes}, nil
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
