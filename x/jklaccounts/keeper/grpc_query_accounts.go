package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackal-dao/canine/x/jklaccounts/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AccountsAll(c context.Context, req *types.QueryAllAccountsRequest) (*types.QueryAllAccountsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var accountss []types.Accounts
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	accountsStore := prefix.NewStore(store, types.KeyPrefix(types.AccountsKeyPrefix))

	pageRes, err := query.Paginate(accountsStore, req.Pagination, func(key []byte, value []byte) error {
		var accounts types.Accounts
		if err := k.cdc.Unmarshal(value, &accounts); err != nil {
			return err
		}

		accountss = append(accountss, accounts)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAccountsResponse{Accounts: accountss, Pagination: pageRes}, nil
}

func (k Keeper) Accounts(c context.Context, req *types.QueryGetAccountsRequest) (*types.QueryGetAccountsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAccounts(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAccountsResponse{Accounts: val}, nil
}
