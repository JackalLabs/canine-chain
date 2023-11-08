package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PubKeyAll(c context.Context, req *types.QueryAllPubKeysRequest) (*types.QueryAllPubKeysResponse, error) {
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

	return &types.QueryAllPubKeysResponse{PubKey: pubkeys, Pagination: pageRes}, nil
}

func (k Keeper) PubKey(c context.Context, req *types.QueryPubKeyRequest) (*types.QueryPubKeyResponse, error) {
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

	return &types.QueryPubKeyResponse{PubKey: val}, nil
}
