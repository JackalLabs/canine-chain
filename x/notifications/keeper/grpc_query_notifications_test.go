package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/notifications/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestNotificationsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NotificationsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNotifications(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNotificationsRequest
		response *types.QueryGetNotificationsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetNotificationsRequest{
				Count: msgs[0].Count,
			},
			response: &types.QueryGetNotificationsResponse{Notifications: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetNotificationsRequest{
				Count: msgs[1].Count,
			},
			response: &types.QueryGetNotificationsResponse{Notifications: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetNotificationsRequest{
				Count: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Notifications(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestNotificationsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NotificationsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNotifications(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllNotificationsRequest {
		return &types.QueryAllNotificationsRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.NotificationsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Notifications), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Notifications),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.NotificationsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Notifications), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Notifications),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.NotificationsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Notifications),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.NotificationsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
