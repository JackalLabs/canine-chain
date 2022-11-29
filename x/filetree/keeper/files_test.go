package keeper_test

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func createFiles(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Files {
	items := make([]types.Files, n)
	for i := range items {
		priv := secp256k1.GenPrivKey()
		address := sdk.AccAddress(priv.PubKey().Address()).String()

		items[i].Address = strconv.Itoa(i)
		items[i].Owner = address

		keeper.SetFiles(ctx, items[i])
	}
	return items
}

func (suite *KeeperTestSuite) TestGetFiles() {
	k := suite.filetreeKeeper
	ctx := suite.ctx

	items := createFiles(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetFiles(ctx, item.Address, item.Owner)
		suite.Require().True(found)
		suite.Equal(item, rst)
	}
}

func (suite *KeeperTestSuite) TestRemoveClientUsage() {
	k := suite.filetreeKeeper
	ctx := suite.ctx

	items := createFiles(k, ctx, 10)
	for _, item := range items {
		k.RemoveFiles(ctx, item.Address, item.Owner)

		rst, found := k.GetFiles(ctx, item.Address, item.Owner)
		suite.Require().Empty(rst)
		suite.Require().False(found)
	}
}

func (suite *KeeperTestSuite) TestGetAllFiles() {
	k := suite.filetreeKeeper
	ctx := suite.ctx

	items := createFiles(k, ctx, 10)
	suite.Require().Equal(items, k.GetAllFiles(ctx))
}
