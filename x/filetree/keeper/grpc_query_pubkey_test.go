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
	"github.com/jackal-dao/canine/x/filetree/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPubkeyQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.FiletreeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPubkey(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetPubkeyRequest
		response *types.QueryGetPubkeyResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPubkeyRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetPubkeyResponse{Pubkey: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPubkeyRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetPubkeyResponse{Pubkey: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPubkeyRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Pubkey(wctx, tc.request)
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

func TestPubkeyQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.FiletreeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPubkey(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPubkeyRequest {
		return &types.QueryAllPubkeyRequest{
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
			resp, err := keeper.PubkeyAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Pubkey), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Pubkey),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PubkeyAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Pubkey), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Pubkey),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PubkeyAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Pubkey),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PubkeyAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
