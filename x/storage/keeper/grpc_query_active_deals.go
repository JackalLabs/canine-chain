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
