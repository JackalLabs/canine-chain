package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PriceCheck(c context.Context, req *types.QueryPriceCheck) (*types.QueryPriceCheckResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	duration := time.Duration(req.Duration) * time.Hour * 24

	timeMonth := time.Hour * 24 * 30
	if duration.Truncate(timeMonth) <= 0 {
		return nil, fmt.Errorf("duration can't be less than 1 month")
	}

	bytes := req.Bytes

	size := sdk.NewInt(bytes)
	s := size.Quo(sdk.NewInt(1_000_000)).Int64() // round to mbs
	if s <= 0 {
		s = 1
	}

	hours := sdk.NewDec(duration.Milliseconds()).Quo(sdk.NewDec(60 * 60 * 1000))

	cost := k.GetStorageCostKbs(ctx, s*1000, hours.TruncateInt64()) // pay for 200 years in mbs

	return &types.QueryPriceCheckResponse{Price: cost.Int64()}, nil
}
