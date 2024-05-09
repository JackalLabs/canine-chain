package keeper

import (
	"context"
	"encoding/hex"
	"strconv"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) InternalActiveDealsAll(c context.Context, req *types.QueryAllInternalActiveDealsRequest) (*types.QueryAllInternalActiveDealsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var activeDealss []types.ActiveDeals
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	activeDealsStore := prefix.NewStore(store, types.KeyPrefix(types.ActiveDealsKeyPrefix))

	pageRes, err := query.Paginate(activeDealsStore, req.Pagination, func(key []byte, value []byte) error {
		var activeDeal types.ActiveDeals
		if err := k.cdc.Unmarshal(value, &activeDeal); err != nil {
			return err
		}

		activeDealss = append(activeDealss, activeDeal)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInternalActiveDealsResponse{ActiveDeals: activeDealss, Pagination: pageRes}, nil
}

func (k Keeper) ActiveDealsAll(c context.Context, req *types.QueryAllActiveDealsRequest) (*types.QueryAllActiveDealsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var activeDealss []types.LegacyActiveDeals
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	activeDealsStore := prefix.NewStore(store, types.KeyPrefix(types.ActiveDealsKeyPrefix))

	p := k.GetParams(ctx)

	pageRes, err := query.Paginate(activeDealsStore, req.Pagination, func(key []byte, value []byte) error {
		var activeDeal types.ActiveDeals
		if err := k.cdc.Unmarshal(value, &activeDeal); err != nil {
			return err
		}

		ver := "false"
		v := activeDeal.IsVerified(ctx.BlockHeight(), p.ProofWindow)
		if v {
			ver = "true"
		}

		lad := types.LegacyActiveDeals{
			Cid:           activeDeal.Cid,
			Signee:        activeDeal.Signee,
			Provider:      activeDeal.Provider,
			Startblock:    activeDeal.Startblock,
			Endblock:      activeDeal.Endblock,
			Filesize:      activeDeal.Filesize,
			Proofverified: ver,
			Proofsmissed:  activeDeal.Proofsmissed,
			Blocktoprove:  activeDeal.Blocktoprove,
			Creator:       activeDeal.Creator,
			Merkle:        activeDeal.Merkle,
			Fid:           activeDeal.Fid,
		}

		activeDealss = append(activeDealss, lad)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActiveDealsResponse{ActiveDeals: activeDealss, Pagination: pageRes}, nil
}

func (k Keeper) ActiveDeals(c context.Context, req *types.QueryActiveDealRequest) (*types.QueryActiveDealResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	activeDeal, found := k.GetActiveDeals(
		ctx,
		req.Cid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	p := k.GetParams(ctx)

	ver := "false"
	v := activeDeal.IsVerified(ctx.BlockHeight(), p.ProofWindow)
	if v {
		ver = "true"
	}

	lad := types.LegacyActiveDeals{
		Cid:           activeDeal.Cid,
		Signee:        activeDeal.Signee,
		Provider:      activeDeal.Provider,
		Startblock:    activeDeal.Startblock,
		Endblock:      activeDeal.Endblock,
		Filesize:      activeDeal.Filesize,
		Proofverified: ver,
		Proofsmissed:  activeDeal.Proofsmissed,
		Blocktoprove:  activeDeal.Blocktoprove,
		Creator:       activeDeal.Creator,
		Merkle:        activeDeal.Merkle,
		Fid:           activeDeal.Fid,
	}

	return &types.QueryActiveDealResponse{ActiveDeals: lad}, nil
}

func (k Keeper) File(c context.Context, req *types.QueryFile) (*types.QueryFileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	activeDeal, err := k.FindDealFromUF(ctx, req.Merkle, req.Owner, req.Start)
	if err != nil {
		return nil, err
	}

	p := k.GetParams(ctx)

	expire, err := strconv.ParseInt(activeDeal.Endblock, 10, 64)
	if err != nil {
		return nil, err
	}

	size, err := strconv.ParseInt(activeDeal.Filesize, 10, 64)
	if err != nil {
		return nil, err
	}

	merkle, err := hex.DecodeString(activeDeal.Merkle)
	if err != nil {
		return nil, err
	}

	proof := types.ProofKey(activeDeal.Provider, merkle, activeDeal.Signee, req.Start)

	proofs := make([]string, 1)
	proofs[0] = string(proof)

	lad := types.UnifiedFile{
		Owner:         activeDeal.Signee,
		Start:         req.Start,
		Expires:       expire,
		FileSize:      size,
		ProofInterval: p.ProofWindow,
		ProofType:     0,
		Proofs:        proofs,
		MaxProofs:     3,
		Note:          "",
	}

	return &types.QueryFileResponse{File: lad}, nil
}

func (k Keeper) Proof(c context.Context, req *types.QueryProof) (*types.QueryProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	activeDeal, err := k.FindDealFromUF(ctx, req.Merkle, req.Owner, req.Start)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot find deal")
	}

	blockToProve, err := strconv.ParseInt(activeDeal.Blocktoprove, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot parse block to prove")
	}

	merkle, err := hex.DecodeString(activeDeal.Merkle)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot parse merkle")
	}

	proof := types.FileProof{
		Prover:       activeDeal.Provider,
		Merkle:       merkle,
		Owner:        activeDeal.Signee,
		Start:        req.Start,
		LastProven:   activeDeal.LastProof,
		ChunkToProve: blockToProve,
	}

	return &types.QueryProofResponse{Proof: proof}, nil
}

func (k Keeper) OpenFiles(c context.Context, req *types.QueryOpenFiles) (*types.QueryAllFilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var files []types.UnifiedFile
	ctx := sdk.UnwrapSDKContext(c)

	strays := k.GetAllStrays(ctx)
	for _, stray := range strays {

		m, err := hex.DecodeString(stray.Merkle)
		if err != nil {
			continue
		}

		fileSize, err := strconv.ParseInt(stray.Filesize, 10, 64)
		if err != nil {
			continue
		}

		uf := types.UnifiedFile{
			Merkle:        m,
			Owner:         stray.Signee,
			Start:         ctx.BlockHeight(),
			Expires:       stray.End,
			FileSize:      fileSize,
			ProofInterval: 3600,
			ProofType:     0,
			Proofs:        make([]string, 0),
			MaxProofs:     3,
			Note:          "",
		}
		files = append(files, uf)
	}

	qpr := query.PageResponse{
		NextKey: nil,
		Total:   uint64(len(files)),
	}

	return &types.QueryAllFilesResponse{Files: files, Pagination: &qpr}, nil
}
