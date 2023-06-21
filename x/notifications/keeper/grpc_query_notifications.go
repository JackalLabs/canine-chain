package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v3/x/notifications/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NotificationsByAddress(c context.Context, req *types.QueryAllNotificationsByAddressRequest) (*types.QueryAllNotificationsByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	notificationss := []types.Notifications{}
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	keyPrefix := fmt.Sprintf("%s%s/", types.NotificationsKeyPrefix, req.Address)

	notificationsStore := prefix.NewStore(store, types.KeyPrefix(keyPrefix))

	pageRes, err := query.Paginate(notificationsStore, req.Pagination, func(key []byte, value []byte) error {
		var notifications types.Notifications
		if err := k.cdc.Unmarshal(value, &notifications); err != nil {
			return err
		}

		notificationss = append(notificationss, notifications)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNotificationsByAddressResponse{Notifications: notificationss, Pagination: pageRes}, nil
}

func (k Keeper) NotificationsAll(c context.Context, req *types.QueryAllNotificationsRequest) (*types.QueryAllNotificationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	notificationss := []types.Notifications{}
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	keyPrefix := types.NotificationsKeyPrefix

	notificationsStore := prefix.NewStore(store, types.KeyPrefix(keyPrefix))

	pageRes, err := query.Paginate(notificationsStore, req.Pagination, func(key []byte, value []byte) error {
		var notifications types.Notifications
		if err := k.cdc.Unmarshal(value, &notifications); err != nil {
			return err
		}

		notificationss = append(notificationss, notifications)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNotificationsResponse{Notifications: notificationss, Pagination: pageRes}, nil
}

// This one is querying a single notification given its index--it was auto generated and is a little bit useless
func (k Keeper) Notifications(c context.Context, req *types.QueryGetNotificationsRequest) (*types.QueryGetNotificationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNotifications(
		ctx,
		req.Count,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNotificationsResponse{Notifications: val}, nil
}
