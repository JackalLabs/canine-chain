package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PubkeyAll(c context.Context, req *types.QueryAllPubkeysRequest) (*types.QueryAllPubkeysResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pubkeys []types.Pubkey
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pubkeyStore := prefix.NewStore(store, types.KeyPrefix(types.PubkeyKeyPrefix))

	pageRes, err := query.Paginate(pubkeyStore, req.Pagination, func(key []byte, value []byte) error {
		var pubkey types.Pubkey
		if err := k.cdc.Unmarshal(value, &pubkey); err != nil {
			return err
		}

		pubkeys = append(pubkeys, pubkey)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPubkeysResponse{Pubkey: pubkeys, Pagination: pageRes}, nil
}

func (k Keeper) Pubkey(c context.Context, req *types.QueryPubkeyRequest) (*types.QueryPubkeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPubkey(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryPubkeyResponse{Pubkey: val}, nil
}
