package keeper_test

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	testutil "github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (suite *KeeperTestSuite) TestDecimals() {
	suite.SetupSuite()

	denom, err := sdk.NewDecFromStr("260040000")
	suite.Require().NoError(err)

	nom, err := sdk.NewDecFromStr("10300000")
	suite.Require().NoError(err)

	res := nom.Quo(denom)

	coins := sdk.NewCoin("ujkl", sdk.NewInt(6000000))

	res = res.Mul(coins.Amount.ToDec())

	suite.Require().Equal(int64(237655), res.TruncateInt64())
}

func (suite *KeeperTestSuite) TestReward() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	signer := testAddresses[0]
	suite.storageKeeper.SetStoragePaymentInfo(suite.ctx, types.StoragePaymentInfo{
		Start:          time.Now(),
		End:            time.Now().AddDate(1, 0, 0),
		SpaceAvailable: 1000000000,
		SpaceUsed:      0,
		Address:        signer,
	})
	providerOne := testAddresses[1]

	dealOne := types.UnifiedFile{
		Merkle:        []byte("merkle"),
		Owner:         signer,
		Start:         0,
		Expires:       0,
		FileSize:      1000,
		ProofInterval: 100,
		ProofType:     0,
		Proofs:        make([]string, 0),
		MaxProofs:     3,
		Note:          "test",
	}
	suite.storageKeeper.SetFile(suite.ctx, dealOne)
	dealOne.AddProver(suite.ctx, suite.storageKeeper, providerOne)

	acc := suite.accountKeeper.GetModuleAddress(types.ModuleName)

	bal := suite.bankKeeper.GetBalance(suite.ctx, acc, "ujkl")
	suite.Require().Equal(int64(0), bal.Amount.Int64())

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(6000000)))

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, acc, coins)
	suite.NoError(err)

	bal = suite.bankKeeper.GetBalance(suite.ctx, acc, "ujkl")
	suite.Require().Equal(int64(6000000), bal.Amount.Int64())

	var blocks int64 = 10 * 5

	ctx := suite.ctx.WithBlockHeight(blocks).WithHeaderHash([]byte{10, 15, 16, 20})

	suite.Require().Equal(blocks, ctx.BlockHeight())
	suite.Require().Equal(ctx.BlockHeight()%blocks, int64(0))

	suite.storageKeeper.ManageRewards(ctx)

	pOneAcc, err := sdk.AccAddressFromBech32(providerOne)
	suite.NoError(err)
	bal = suite.bankKeeper.GetBalance(suite.ctx, pOneAcc, "ujkl")
	suite.Require().Equal(int64(6000000), bal.Amount.Int64())

	bal = suite.bankKeeper.GetBalance(ctx, acc, "ujkl")
	suite.Require().Equal(int64(0), bal.Amount.Int64())
}

func (suite *KeeperTestSuite) TestMultiReward() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	signer := testAddresses[0]
	suite.storageKeeper.SetStoragePaymentInfo(suite.ctx, types.StoragePaymentInfo{
		Start:          time.Now(),
		End:            time.Now().AddDate(1, 0, 0),
		SpaceAvailable: 1000000000,
		SpaceUsed:      0,
		Address:        signer,
	})

	const l = 50

	providers := make([]sdk.AccAddress, l)

	for i := 0; i < l; i++ {
		acc, err := sdk.AccAddressFromHex(fmt.Sprintf("%08x", i))
		suite.Require().NoError(err)
		providers[i] = acc
	}

	deals := make([]types.UnifiedFile, l*2)

	total := 0

	for i := 0; i < l*2; i++ {
		p := providers[i%l]
		deal := types.UnifiedFile{
			Merkle:        []byte("merkle"),
			Owner:         signer,
			Start:         0,
			Expires:       0,
			FileSize:      1000,
			ProofInterval: 100,
			ProofType:     0,
			Proofs:        make([]string, 0),
			MaxProofs:     3,
			Note:          "test",
		}
		suite.storageKeeper.SetFile(suite.ctx, deal)
		deal.AddProver(suite.ctx, suite.storageKeeper, p.String())

		total += i

		deals[i] = deal
	}

	acc := suite.accountKeeper.GetModuleAddress(types.ModuleName)

	bal := suite.bankKeeper.GetBalance(suite.ctx, acc, "ujkl")
	suite.Require().Equal(int64(0), bal.Amount.Int64())

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(6000000)))

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, acc, coins)
	suite.NoError(err)

	bal = suite.bankKeeper.GetBalance(suite.ctx, acc, "ujkl")
	suite.Require().Equal(int64(6000000), bal.Amount.Int64())

	var blocks int64 = 10 * 5

	ctx := suite.ctx.WithBlockHeight(blocks).WithHeaderHash([]byte{10, 15, 16, 20})

	suite.Require().Equal(blocks, ctx.BlockHeight())
	suite.Require().Equal(ctx.BlockHeight()%blocks, int64(0))

	suite.storageKeeper.ManageRewards(ctx)

	nom := sdk.NewDec(20)
	den := sdk.NewDec(int64(total))

	r := nom.Quo(den).Mul(sdk.NewDec(6000000))
	m := r.TruncateInt64()

	nom = sdk.NewDec(70)
	den = sdk.NewDec(int64(total))

	r = nom.Quo(den).Mul(sdk.NewDec(6000000))

	m += r.TruncateInt64()

	bal = suite.bankKeeper.GetBalance(suite.ctx, providers[20], "ujkl")
	suite.Require().Equal(m, bal.Amount.Int64())

	bal = suite.bankKeeper.GetBalance(ctx, acc, "ujkl")
	suite.Require().Equal(int64(0), bal.Amount.Int64())
}
