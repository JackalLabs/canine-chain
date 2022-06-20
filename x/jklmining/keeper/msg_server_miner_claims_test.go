package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/x/jklmining/keeper"
	"github.com/jackal-dao/canine/x/jklmining/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestMinerClaimsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.JklminingKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateMinerClaims{Creator: creator,
			Hash: strconv.Itoa(i),
		}
		_, err := srv.CreateMinerClaims(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetMinerClaims(ctx,
			expected.Hash,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestMinerClaimsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateMinerClaims
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateMinerClaims{Creator: creator,
				Hash: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateMinerClaims{Creator: "B",
				Hash: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateMinerClaims{Creator: creator,
				Hash: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.JklminingKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateMinerClaims{Creator: creator,
				Hash: strconv.Itoa(0),
			}
			_, err := srv.CreateMinerClaims(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateMinerClaims(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetMinerClaims(ctx,
					expected.Hash,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestMinerClaimsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteMinerClaims
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteMinerClaims{Creator: creator,
				Hash: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteMinerClaims{Creator: "B",
				Hash: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteMinerClaims{Creator: creator,
				Hash: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.JklminingKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateMinerClaims(wctx, &types.MsgCreateMinerClaims{Creator: creator,
				Hash: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteMinerClaims(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetMinerClaims(ctx,
					tc.request.Hash,
				)
				require.False(t, found)
			}
		})
	}
}
