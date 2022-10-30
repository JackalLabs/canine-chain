package keeper_test

import (
	"dsig/testutil/nullify"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "dsig/testutil/keeper"

	"github.com/jackalLabs/canine-chain/x/dsig/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestUserUploadsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DsigKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNUserUploads(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetUserUploadsRequest
		response *types.QueryGetUserUploadsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetUserUploadsRequest{
				Fid: msgs[0].Fid,
			},
			response: &types.QueryGetUserUploadsResponse{UserUploads: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetUserUploadsRequest{
				Fid: msgs[1].Fid,
			},
			response: &types.QueryGetUserUploadsResponse{UserUploads: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetUserUploadsRequest{
				Fid: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.UserUploads(wctx, tc.request)
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

func TestUserUploadsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DsigKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNUserUploads(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllUserUploadsRequest {
		return &types.QueryAllUserUploadsRequest{
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
			resp, err := keeper.UserUploadsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UserUploads), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UserUploads),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.UserUploadsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UserUploads), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UserUploads),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.UserUploadsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.UserUploads),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.UserUploadsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
