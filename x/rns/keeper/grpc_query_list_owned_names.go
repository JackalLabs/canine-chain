package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListOwnedNames(goCtx context.Context, req *types.QueryListOwnedNames) (*types.QueryListOwnedNamesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var namess []types.Names
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	namesStore := prefix.NewStore(store, types.KeyPrefix(types.NamesKeyPrefix))

	pageRes, err := query.Paginate(namesStore, req.Pagination, func(_ []byte, value []byte) error {
		var names types.Names
		if err := k.cdc.Unmarshal(value, &names); err != nil {
			return err
		}

		if names.Value == req.Address {
			namess = append(namess, names)
		}

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListOwnedNamesResponse{Names: namess, Pagination: pageRes}, nil
}
