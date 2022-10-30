package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/jackalLabs/canine-chain/testutil/keeper"
	"github.com/jackalLabs/canine-chain/x/notifications/keeper"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestNotificationsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.NotificationsKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateNotifications{
			Creator: creator,
			Count:   uint64(i),
		}
		_, err := srv.CreateNotifications(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetNotifications(ctx,
			expected.Count,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestNotificationsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateNotifications
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateNotifications{
				Creator: creator,
				Count:   0,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateNotifications{
				Creator: "B",
				Count:   0,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateNotifications{
				Creator: creator,
				Count:   100000,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NotificationsKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateNotifications{
				Creator: creator,
				Count:   0,
			}
			_, err := srv.CreateNotifications(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateNotifications(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetNotifications(ctx,
					expected.Count,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestNotificationsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteNotifications
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteNotifications{
				Creator: creator,
				Count:   0,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteNotifications{
				Creator: "B",
				Count:   0,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteNotifications{
				Creator: creator,
				Count:   100000,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NotificationsKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateNotifications(wctx, &types.MsgCreateNotifications{
				Creator: creator,
				Count:   0,
			})
			require.NoError(t, err)
			_, err = srv.DeleteNotifications(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetNotifications(ctx,
					tc.request.Count,
				)
				require.False(t, found)
			}
		})
	}
}
