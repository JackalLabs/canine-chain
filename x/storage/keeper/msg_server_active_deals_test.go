package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/x/storage/keeper"
	"github.com/jackal-dao/canine/x/storage/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestActiveDealsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.StorageKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateActiveDeals{Creator: creator,
			Cid: strconv.Itoa(i),
		}
		_, err := srv.CreateActiveDeals(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetActiveDeals(ctx,
			expected.Cid,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestActiveDealsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateActiveDeals
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateActiveDeals{Creator: creator,
				Cid: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateActiveDeals{Creator: "B",
				Cid: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateActiveDeals{Creator: creator,
				Cid: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.StorageKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateActiveDeals{Creator: creator,
				Cid: strconv.Itoa(0),
			}
			_, err := srv.CreateActiveDeals(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateActiveDeals(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetActiveDeals(ctx,
					expected.Cid,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestActiveDealsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteActiveDeals
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteActiveDeals{Creator: creator,
				Cid: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteActiveDeals{Creator: "B",
				Cid: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteActiveDeals{Creator: creator,
				Cid: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.StorageKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateActiveDeals(wctx, &types.MsgCreateActiveDeals{Creator: creator,
				Cid: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteActiveDeals(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetActiveDeals(ctx,
					tc.request.Cid,
				)
				require.False(t, found)
			}
		})
	}
}
