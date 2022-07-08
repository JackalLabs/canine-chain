package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/themarstonconnell/telescope/testutil/keeper"
	"github.com/themarstonconnell/telescope/testutil/nullify"
	"github.com/themarstonconnell/telescope/x/telescope/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestWhoisQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.TelescopeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNWhois(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetWhoisRequest
		response *types.QueryGetWhoisResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetWhoisRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetWhoisResponse{Whois: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetWhoisRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetWhoisResponse{Whois: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetWhoisRequest{
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
			response, err := keeper.Whois(wctx, tc.request)
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

func TestWhoisQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.TelescopeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNWhois(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllWhoisRequest {
		return &types.QueryAllWhoisRequest{
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
			resp, err := keeper.WhoisAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Whois), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Whois),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.WhoisAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Whois), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Whois),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.WhoisAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Whois),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.WhoisAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
