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

func TestProofsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProofs(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetProofsRequest
		response *types.QueryGetProofsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetProofsRequest{
				Cid: msgs[0].Cid,
			},
			response: &types.QueryGetProofsResponse{Proofs: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetProofsRequest{
				Cid: msgs[1].Cid,
			},
			response: &types.QueryGetProofsResponse{Proofs: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetProofsRequest{
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
			response, err := keeper.Proofs(wctx, tc.request)
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

func TestProofsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProofs(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllProofsRequest {
		return &types.QueryAllProofsRequest{
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
			resp, err := keeper.ProofsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Proofs), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Proofs),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ProofsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Proofs), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Proofs),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ProofsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Proofs),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ProofsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
