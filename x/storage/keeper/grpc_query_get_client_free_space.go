package keeper

import (
	"context"
	"fmt"

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

	clientUsage, found := k.GetClientUsage(ctx, req.Address)
	if !found {
		clientUsage = types.ClientUsage{
			Address: req.Address,
			Usage:   "0",
		}
	}

	usage, ok := sdk.NewIntFromString(clientUsage.Usage)
	if !ok {
		return nil, fmt.Errorf("cannot parse client usage")
	}

	paid, _, _ := k.GetPaidAmount(ctx, req.Address, ctx.BlockHeight())

	if paid < usage.Int64() {
		return nil, fmt.Errorf("paid amount cannot be smaller than usage")
	}

	bfree := paid - usage.Int64()

	return &types.QueryClientFreeSpaceResponse{Bytesfree: fmt.Sprintf("%d", bfree)}, nil
}
