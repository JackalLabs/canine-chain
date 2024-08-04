package keeper_test

import (
	"math/rand"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
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
	suite.SetupSuite()
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
	suite.SetupSuite()
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
	suite.SetupSuite()
	k := suite.storageKeeper
	ctx := suite.ctx

	items := createStoragePaymentInfo(k, ctx, 10)
	suite.Require().Equal(items, k.GetAllStoragePaymentInfo(ctx))
}

func (suite *KeeperTestSuite) TestIterateGauges() {
	suite.SetupSuite()
	k := suite.storageKeeper
	ctx := suite.ctx

	for i := 0; i < 50; i++ {
		ls := make([][]byte, i)
		for m := 0; m < i; m++ {
			ls[m] = k.NewGauge(ctx, sdk.NewCoins(sdk.NewInt64Coin("ujkl", rand.Int63())), time.Now().Add(time.Hour*20)).Id
		}

		is := 0
		k.IterateGauges(ctx, func(_ types.PaymentGauge) {
			is++
		})

		suite.Require().Equal(i, is)

		for _, l := range ls {
			k.RemoveGauge(ctx, l)
		}
	}
}
