package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v3/x/rns/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllNames(c context.Context, req *types.QueryAllNames) (*types.QueryAllNamesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var namess []types.Names
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	namesStore := prefix.NewStore(store, types.KeyPrefix(types.NamesKeyPrefix))

	pageRes, err := query.Paginate(namesStore, req.Pagination, func(key []byte, value []byte) error {
		var names types.Names
		if err := k.cdc.Unmarshal(value, &names); err != nil {
			return err
		}

		namess = append(namess, names)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNamesResponse{Name: namess, Pagination: pageRes}, nil
}

func (k Keeper) Name(c context.Context, req *types.QueryName) (*types.QueryNameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	n, tld, err := GetNameAndTLD(req.Index)
	if err != nil {
		return nil, err
	}

	sub, name, hasSub := GetSubdomain(n)
	if hasSub {
		n = name
	}

	val, found := k.GetNames(ctx, n, tld)

	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	if hasSub {
		for _, domain := range val.Subdomains {
			if domain.Name == sub {
				return &types.QueryNameResponse{Name: *domain}, nil
			}
		}
	}

	return &types.QueryNameResponse{Name: val}, nil
}
