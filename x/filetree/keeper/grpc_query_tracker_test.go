package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func TestTrackerQuery(t *testing.T) {
	keeper, ctx := keepertest.FiletreeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestTracker(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTrackerRequest
		response *types.QueryGetTrackerResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTrackerRequest{},
			response: &types.QueryGetTrackerResponse{Tracker: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Tracker(wctx, tc.request)
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
