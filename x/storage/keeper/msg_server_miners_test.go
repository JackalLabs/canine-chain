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

func TestMinersMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.StorageKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateMiners{Creator: creator,
			Address: strconv.Itoa(i),
		}
		_, err := srv.CreateMiners(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetMiners(ctx,
			expected.Address,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestMinersMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateMiners
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateMiners{Creator: creator,
				Address: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateMiners{Creator: "B",
				Address: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateMiners{Creator: creator,
				Address: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.StorageKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateMiners{Creator: creator,
				Address: strconv.Itoa(0),
			}
			_, err := srv.CreateMiners(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateMiners(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetMiners(ctx,
					expected.Address,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestMinersMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteMiners
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteMiners{Creator: creator,
				Address: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteMiners{Creator: "B",
				Address: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteMiners{Creator: creator,
				Address: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.StorageKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateMiners(wctx, &types.MsgCreateMiners{Creator: creator,
				Address: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteMiners(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetMiners(ctx,
					tc.request.Address,
				)
				require.False(t, found)
			}
		})
	}
}
