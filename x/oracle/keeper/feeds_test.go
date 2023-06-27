package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/oracle/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/oracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createFeed(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Feed {
	items := make([]types.Feed, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)
		keeper.SetFeed(ctx, items[i])
	}
	return items
}

func (suite *KeeperTestSuite) TestGetFeed() {
	k := suite.oracleKeeper
	ctx := suite.ctx

	items := createFeed(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetFeed(ctx, item.Name)
		suite.Require().True(found)
		suite.Equal(item, rst)
	}
}

func (suite *KeeperTestSuite) TestRemoveFeed() {
	k := suite.oracleKeeper
	ctx := suite.ctx

	items := createFeed(k, ctx, 10)
	for _, item := range items {
		k.RemoveFeed(ctx, item.Name)

		rst, found := k.GetFeed(ctx, item.Name)
		suite.Require().Empty(rst)
		suite.Require().False(found)
	}
}

func (suite *KeeperTestSuite) TestGetAllFeeds() {
	k := suite.oracleKeeper
	ctx := suite.ctx

	items := createFeed(k, ctx, 10)
	suite.Require().Equal(items, k.GetAllFeeds(ctx))
}
