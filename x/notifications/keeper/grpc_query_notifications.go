package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v5/x/notifications/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllNotificationsByAddress(c context.Context, req *types.QueryAllNotificationsByAddress) (*types.QueryAllNotificationsByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	page, limit, err := query.ParsePagination(req.Pagination)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "cannot parse pagination")
	}
	offset := (page - 1) * limit

	notifications := k.GetAllNotificationsByAddress(ctx, req.To)

	if offset > len(notifications) {
		pres := query.PageResponse{
			NextKey: nil,
			Total:   0,
		}
		return &types.QueryAllNotificationsByAddressResponse{Notifications: make([]types.Notification, 0), Pagination: &pres}, nil
	}

	notifications = notifications[offset:]

	if len(notifications) > limit {
		notifications = notifications[:limit]
	}
	pres := query.PageResponse{
		NextKey: nil,
		Total:   uint64(len(notifications)),
	}
	return &types.QueryAllNotificationsByAddressResponse{Notifications: notifications, Pagination: &pres}, nil
}

func (k Keeper) AllNotifications(c context.Context, req *types.QueryAllNotifications) (*types.QueryAllNotificationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var notifications []types.Notification
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	keyPrefix := types.NotificationsKeyPrefix

	notificationsStore := prefix.NewStore(store, types.KeyPrefix(keyPrefix))

	pageRes, err := query.Paginate(notificationsStore, req.Pagination, func(_ []byte, value []byte) error {
		var notification types.Notification
		if err := k.cdc.Unmarshal(value, &notification); err != nil {
			return err
		}

		notifications = append(notifications, notification)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNotificationsResponse{Notifications: notifications, Pagination: pageRes}, nil
}

func (k Keeper) Notification(c context.Context, req *types.QueryNotification) (*types.QueryNotificationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNotification(
		ctx,
		req.To,
		req.From,
		req.Time,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryNotificationResponse{Notification: val}, nil
}
