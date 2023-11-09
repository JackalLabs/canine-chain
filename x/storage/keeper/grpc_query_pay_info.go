package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PaymentInfo(goCtx context.Context, req *types.QueryStoragePaymentInfo) (*types.QueryStoragePaymentInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	payInfo, found := k.GetStoragePaymentInfo(ctx, req.Address)
	if !found {
		t := types.StoragePaymentInfo{
			Start:          time.UnixMicro(0),
			End:            time.UnixMicro(0),
			Address:        req.Address,
			SpaceUsed:      0,
			SpaceAvailable: 0,
		}
		return &types.QueryStoragePaymentInfoResponse{
			StoragePaymentInfo: t,
		}, nil
	}

	return &types.QueryStoragePaymentInfoResponse{
		StoragePaymentInfo: payInfo,
	}, nil
}
