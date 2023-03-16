package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func createPubkey(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Pubkey {
	items := make([]types.Pubkey, n)

	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetPubkey(ctx, items[i])
	}
	return items
}

func (suite *KeeperTestSuite) TestGetPubkey() {
	k := suite.filetreeKeeper
	ctx := suite.ctx

	items := createPubkey(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetPubkey(ctx, item.Address)
		suite.Require().True(found)
		suite.Equal(item, rst)
	}
}

func (suite *KeeperTestSuite) TestRemovePubkey() {
	k := suite.filetreeKeeper
	ctx := suite.ctx

	items := createPubkey(k, ctx, 10)
	for _, item := range items {
		k.RemovePubkey(ctx, item.Address)

		rst, found := k.GetPubkey(ctx, item.Address)
		suite.Require().Empty(rst)
		suite.Require().False(found)
	}
}

func (suite *KeeperTestSuite) TestGetAllPubkey() {
	k := suite.filetreeKeeper
	ctx := suite.ctx

	items := createPubkey(k, ctx, 10)
	suite.Require().Equal(items, k.GetAllPubkey(ctx))
}
