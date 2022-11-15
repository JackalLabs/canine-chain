package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createPayBlocks(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PayBlocks {
	items := make([]types.PayBlocks, n)
	for i := range items {
		items[i].Blockid = strconv.Itoa(i)
		items[i].Bytes = strconv.Itoa(i)
		items[i].Blocktype = strconv.Itoa(i)
		items[i].Blocknum = strconv.Itoa(i)

		keeper.SetPayBlocks(ctx, items[i])
	}
	return items
}

func (suite *KeeperTestSuite) TestGetPayBlocks() {
	k := suite.storageKeeper
	ctx := suite.ctx

	items := createPayBlocks(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetPayBlocks(ctx, item.Blockid)
		suite.Require().True(found)
		suite.Equal(item, rst)
	}
}

func (suite *KeeperTestSuite) TestRemovePayBlocks() {
	k := suite.storageKeeper
	ctx := suite.ctx

	items := createPayBlocks(k, ctx, 10)
	for _, item := range items {
		k.RemovePayBlocks(ctx, item.Blockid)

		rst, found := k.GetPayBlocks(ctx, item.Blockid)
		suite.Require().Empty(rst)
		suite.Require().False(found)
	}
}

func (suite *KeeperTestSuite) TestGetAllPayBlocks() {
	k := suite.storageKeeper
	ctx := suite.ctx

	items := createPayBlocks(k, ctx, 10)
	suite.Require().Equal(items, k.GetAllPayBlocks(ctx))
}
