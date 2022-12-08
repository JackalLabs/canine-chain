package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createStoragePaymentInfo(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoragePaymentInfo {
	items := make([]types.StoragePaymentInfo, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		keeper.SetStoragePaymentInfo(ctx, items[i])
	}
	return items
}

func (suite *KeeperTestSuite) TestGetStoragePaymentInfo() {
	k := suite.storageKeeper
	ctx := suite.ctx

	items := createStoragePaymentInfo(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetStoragePaymentInfo(ctx, item.Address)
		suite.Require().True(found)
		suite.Equal(item, rst)
	}
}

func (suite *KeeperTestSuite) TestRemoveStoragePaymentInfo() {
	k := suite.storageKeeper
	ctx := suite.ctx

	items := createStoragePaymentInfo(k, ctx, 10)
	for _, item := range items {
		k.RemoveStoragePaymentInfo(ctx, item.Address)

		rst, found := k.GetStoragePaymentInfo(ctx, item.Address)
		suite.Require().Empty(rst)
		suite.Require().False(found)
	}
}

// fix this last test boi!
func (suite *KeeperTestSuite) TestGetAllStoragePaymentInfo() {
	k := suite.storageKeeper
	ctx := suite.ctx

	items := createStoragePaymentInfo(k, ctx, 10)
	suite.Require().Equal(items, k.GetAllStoragePaymentInfo(ctx))
}
