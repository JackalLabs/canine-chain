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
	"github.com/jackal-dao/canine/x/storage/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestActiveDealsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActiveDeals(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetActiveDealsRequest
		response *types.QueryGetActiveDealsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetActiveDealsRequest{
				Cid: msgs[0].Cid,
			},
			response: &types.QueryGetActiveDealsResponse{ActiveDeals: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetActiveDealsRequest{
				Cid: msgs[1].Cid,
			},
			response: &types.QueryGetActiveDealsResponse{ActiveDeals: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetActiveDealsRequest{
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
			response, err := keeper.ActiveDeals(wctx, tc.request)
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

func TestActiveDealsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActiveDeals(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllActiveDealsRequest {
		return &types.QueryAllActiveDealsRequest{
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
			resp, err := keeper.ActiveDealsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActiveDeals), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActiveDeals),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActiveDealsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActiveDeals), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActiveDeals),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ActiveDealsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ActiveDeals),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ActiveDealsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
