package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetPayData(goCtx context.Context, req *types.QueryPayDataRequest) (*types.QueryPayDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	paid := k.GetPaidAmount(ctx, req.Address)

	payInfo, found := k.GetStoragePaymentInfo(ctx, req.Address)
	if !found {
		return &types.QueryPayDataResponse{TimeRemaining: -1, Bytes: 0}, nil
	}

	left := payInfo.End.Unix() - ctx.BlockTime().Unix()

	return &types.QueryPayDataResponse{TimeRemaining: left, Bytes: paid}, nil
}

// func (k Keeper) GetStoragePayementInfoList(goCtx context.Context, req *types.QueryStoragePaymentInfoRequest) (*types.QueryStoragePaymentInfoResponse, error) {
// 	if req == nil {
// 		return nil, status.Error(codes.InvalidArgument, "invalid request")
// 	}
//
// 	ctx := sdk.UnwrapSDKContext(goCtx)
//
// 	info := k.GetAllStoragePaymentInfo(ctx)
//
// 	return &types.QueryStoragePaymentInfoResponse{PaymentInfo: info}, nil
// }
