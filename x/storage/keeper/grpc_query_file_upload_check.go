package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FileUploadCheck(c context.Context, req *types.QueryFileUploadCheck) (*types.QueryFileUploadCheckResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	spi, found := k.GetStoragePaymentInfo(ctx, req.Address)

	if !found {
		return nil, status.Error(codes.NotFound, "address not found")
	}

	if req.Bytes < 0 {
		return nil, status.Error(codes.InvalidArgument, "bytes cannot be negative")
	}

	if (spi.SpaceAvailable - spi.SpaceUsed) < req.Bytes {
		return &types.QueryFileUploadCheckResponse{Valid: false}, nil
	}

	return &types.QueryFileUploadCheckResponse{Valid: true}, nil
}
