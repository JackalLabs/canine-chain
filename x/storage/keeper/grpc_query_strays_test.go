package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/jackalLabs/canine-chain/testutil/keeper"
	"github.com/jackalLabs/canine-chain/testutil/nullify"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestStraysQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStrays(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetStraysRequest
		response *types.QueryGetStraysResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetStraysRequest{
				Cid: msgs[0].Cid,
			},
			response: &types.QueryGetStraysResponse{Strays: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetStraysRequest{
				Cid: msgs[1].Cid,
			},
			response: &types.QueryGetStraysResponse{Strays: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetStraysRequest{
				Cid: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Strays(wctx, tc.request)
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

func TestStraysQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStrays(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllStraysRequest {
		return &types.QueryAllStraysRequest{
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
			resp, err := keeper.StraysAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Strays), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Strays),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.StraysAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Strays), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Strays),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.StraysAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Strays),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.StraysAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
