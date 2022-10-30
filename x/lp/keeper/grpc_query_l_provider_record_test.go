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
	"github.com/jackalLabs/canine-chain/x/lp/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestLProviderRecordQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.LpKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLProviderRecord(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetLProviderRecordRequest
		response *types.QueryGetLProviderRecordResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetLProviderRecordRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetLProviderRecordResponse{LProviderRecord: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetLProviderRecordRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetLProviderRecordResponse{LProviderRecord: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetLProviderRecordRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.LProviderRecord(wctx, tc.request)
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

func TestLProviderRecordQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.LpKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLProviderRecord(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllLProviderRecordRequest {
		return &types.QueryAllLProviderRecordRequest{
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
			resp, err := keeper.LProviderRecordAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LProviderRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LProviderRecord),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LProviderRecordAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LProviderRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LProviderRecord),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.LProviderRecordAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.LProviderRecord),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.LProviderRecordAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}