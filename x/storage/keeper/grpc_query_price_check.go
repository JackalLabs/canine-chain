package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
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

	const gb int64 = 1000000000

	gbs := bytes / gb
	if gbs <= 0 {
		return nil, fmt.Errorf("cannot buy less than a gb")
	}

	hours := sdk.NewDec(duration.Milliseconds()).Quo(sdk.NewDec(60 * 60 * 1000))
	p := k.GetStorageCost(ctx, gbs, hours.TruncateInt().Int64())

	return &types.QueryPriceCheckResponse{Price: p.Int64()}, nil
}
