package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllReports(c context.Context, req *types.QueryAllReports) (*types.QueryAllReportsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var reports []types.ReportForm
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	reportStore := prefix.NewStore(store, types.KeyPrefix(types.ReportKeyPrefix))

	pageRes, err := query.Paginate(reportStore, req.Pagination, func(_ []byte, value []byte) error {
		var providers types.ReportForm
		if err := k.cdc.Unmarshal(value, &providers); err != nil {
			return err
		}

		reports = append(reports, providers)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllReportsResponse{Reports: reports, Pagination: pageRes}, nil
}

func (k Keeper) Report(c context.Context, req *types.QueryReport) (*types.QueryReportResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetReportForm(
		ctx,
		req.Prover,
		req.Merkle,
		req.Owner,
		req.Start,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryReportResponse{Report: val}, nil
}
