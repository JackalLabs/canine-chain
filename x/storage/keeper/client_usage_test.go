package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createClientUsage(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ClientUsage {
	items := make([]types.ClientUsage, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		items[i].Usage = strconv.Itoa(i)

		keeper.SetClientUsage(ctx, items[i])
	}
	return items
}

func (suite *KeeperTestSuite) TestGetClientUsage() {
	k := suite.storageKeeper
	ctx := suite.ctx

	items := createClientUsage(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetClientUsage(ctx, item.Address)
		suite.Require().True(found)
		suite.Equal(item, rst)
	}
}

func (suite *KeeperTestSuite) TestRemoveClientUsage() {
	k := suite.storageKeeper
	ctx := suite.ctx

	items := createClientUsage(k, ctx, 10)
	for _, item := range items {
		k.RemoveClientUsage(ctx, item.Address)

		rst, found := k.GetClientUsage(ctx, item.Address)
		suite.Require().Empty(rst)
		suite.Require().False(found)
	}
}

func (suite *KeeperTestSuite) TestGetAllClientUsage() {
	k := suite.storageKeeper
	ctx := suite.ctx

	items := createClientUsage(k, ctx, 10)
	suite.Require().Equal(items, k.GetAllClientUsage(ctx))
}
