package keeper

import (
	"context"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"

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

	reverse := false
	var limit uint64 = 100
	if req.Pagination != nil { // HERE IS THE FIX
		reverse = req.Pagination.Reverse
		limit = req.Pagination.Limit
	}

	var iterator storetypes.Iterator
	if reverse {
		iterator = sdk.KVStoreReversePrefixIterator(namesStore, []byte{})
	} else {
		iterator = sdk.KVStorePrefixIterator(namesStore, []byte{})
	}

	defer iterator.Close()
	var i uint64

	for ; iterator.Valid(); iterator.Next() {

		if i > limit {
			break
		}

		var names types.Names
		if err := k.cdc.Unmarshal(iterator.Value(), &names); err != nil {
			continue
		}

		if names.Value == req.Address {
			namess = append(namess, names)
			i++
		}

	}

	qpr := query.PageResponse{
		NextKey: nil,
		Total:   uint64(len(namess)),
	}

	return &types.QueryListOwnedNamesResponse{Names: namess, Pagination: &qpr}, nil
}
