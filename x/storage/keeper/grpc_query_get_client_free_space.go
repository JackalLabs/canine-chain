package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetClientFreeSpace(goCtx context.Context, req *types.QueryClientFreeSpaceRequest) (*types.QueryClientFreeSpaceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	payInfo, found := k.GetStoragePaymentInfo(ctx, req.Address)
	if !found {
		return &types.QueryClientFreeSpaceResponse{Bytesfree: 0}, nil
	}

	return &types.QueryClientFreeSpaceResponse{Bytesfree: payInfo.SpaceAvailable - payInfo.SpaceUsed}, nil
}
