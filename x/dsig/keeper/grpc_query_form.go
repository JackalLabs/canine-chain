package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackal-dao/canine/x/dsig/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FormAll(c context.Context, req *types.QueryAllFormRequest) (*types.QueryAllFormResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var forms []types.Form
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	formStore := prefix.NewStore(store, types.KeyPrefix(types.FormKeyPrefix))

	pageRes, err := query.Paginate(formStore, req.Pagination, func(key []byte, value []byte) error {
		var form types.Form
		if err := k.cdc.Unmarshal(value, &form); err != nil {
			return err
		}

		forms = append(forms, form)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFormResponse{Form: forms, Pagination: pageRes}, nil
}

func (k Keeper) Form(c context.Context, req *types.QueryGetFormRequest) (*types.QueryGetFormResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetForm(
		ctx,
		req.Ffid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetFormResponse{Form: val}, nil
}
