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

func TestContractsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.StorageKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateContracts{Creator: creator,
			Cid: strconv.Itoa(i),
		}
		_, err := srv.CreateContracts(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetContracts(ctx,
			expected.Cid,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestContractsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateContracts
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateContracts{Creator: creator,
				Cid: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateContracts{Creator: "B",
				Cid: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateContracts{Creator: creator,
				Cid: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.StorageKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateContracts{Creator: creator,
				Cid: strconv.Itoa(0),
			}
			_, err := srv.CreateContracts(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateContracts(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetContracts(ctx,
					expected.Cid,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestContractsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteContracts
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteContracts{Creator: creator,
				Cid: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteContracts{Creator: "B",
				Cid: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteContracts{Creator: creator,
				Cid: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.StorageKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateContracts(wctx, &types.MsgCreateContracts{Creator: creator,
				Cid: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteContracts(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetContracts(ctx,
					tc.request.Cid,
				)
				require.False(t, found)
			}
		})
	}
}
