package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ContractsAll(c context.Context, req *types.QueryAllContractsRequest) (*types.QueryAllContractsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var contractss []types.Contracts
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	contractsStore := prefix.NewStore(store, types.KeyPrefix(types.ContractsKeyPrefix))

	pageRes, err := query.Paginate(contractsStore, req.Pagination, func(key []byte, value []byte) error {
		var contracts types.Contracts
		if err := k.cdc.Unmarshal(value, &contracts); err != nil {
			return err
		}

		contractss = append(contractss, contracts)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllContractsResponse{Contracts: contractss, Pagination: pageRes}, nil
}

func (k Keeper) Contracts(c context.Context, req *types.QueryContractRequest) (*types.QueryContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetContracts(
		ctx,
		req.Cid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryContractResponse{Contracts: val}, nil
}
