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

func TestProofsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.StorageKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateProofs{Creator: creator,
			Cid: strconv.Itoa(i),
		}
		_, err := srv.CreateProofs(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetProofs(ctx,
			expected.Cid,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestProofsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateProofs
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateProofs{Creator: creator,
				Cid: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateProofs{Creator: "B",
				Cid: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateProofs{Creator: creator,
				Cid: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.StorageKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateProofs{Creator: creator,
				Cid: strconv.Itoa(0),
			}
			_, err := srv.CreateProofs(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateProofs(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetProofs(ctx,
					expected.Cid,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestProofsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteProofs
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteProofs{Creator: creator,
				Cid: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteProofs{Creator: "B",
				Cid: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteProofs{Creator: creator,
				Cid: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.StorageKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateProofs(wctx, &types.MsgCreateProofs{Creator: creator,
				Cid: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteProofs(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetProofs(ctx,
					tc.request.Cid,
				)
				require.False(t, found)
			}
		})
	}
}
