package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil "github.com/jackalLabs/canine-chain/v3/testutil"
	jklminttypes "github.com/jackalLabs/canine-chain/v3/x/jklmint/types"
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

	s := suite.ctx.BlockTime()

	t := s.AddDate(1, 0, 0)
	suite.storageKeeper.SetStoragePaymentInfo(suite.ctx, types.StoragePaymentInfo{
		Start:          s,
		End:            t,
		SpaceAvailable: 1000000000,
		SpaceUsed:      0,
		Address:        signer,
	})

	acc, err := types.GetTokenHolderAccount()
	suite.Require().NoError(err)

	bal := suite.bankKeeper.GetBalance(suite.ctx, acc, "ujkl")
	suite.Require().Equal(int64(0), bal.Amount.Int64())

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(6000000)))

	suite.storageKeeper.NewGauge(suite.ctx, coins, t)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, jklminttypes.ModuleName, acc, coins)
	suite.NoError(err)

	bal = suite.bankKeeper.GetBalance(suite.ctx, acc, "ujkl")
	suite.Require().Equal(int64(6000000), bal.Amount.Int64())

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

	blocks := dealOne.ProofInterval * 3

	suite.storageKeeper.SetFile(suite.ctx, dealOne)
	dealOne.AddProver(suite.ctx, suite.storageKeeper, providerOne)
	suite.storageKeeper.SetProof(suite.ctx, types.FileProof{
		Prover:       providerOne,
		Merkle:       dealOne.Merkle,
		Owner:        dealOne.Owner,
		Start:        dealOne.Start,
		LastProven:   blocks - 1,
		ChunkToProve: 0,
	})

	_, found := suite.storageKeeper.GetProof(suite.ctx, providerOne, dealOne.Merkle, dealOne.Owner, dealOne.Start)
	suite.Require().True(found)

	ctx := suite.ctx.WithBlockHeight(blocks).WithHeaderHash([]byte{10, 15, 16, 20}).WithBlockTime(s.AddDate(0, 3, 0))

	suite.Require().Equal(blocks, ctx.BlockHeight())
	suite.Require().Equal(ctx.BlockHeight()%blocks, int64(0))

	suite.storageKeeper.ManageRewards(ctx)

	pOneAcc, err := sdk.AccAddressFromBech32(providerOne)
	suite.NoError(err)
	bal = suite.bankKeeper.GetBalance(suite.ctx, pOneAcc, "ujkl")
	suite.Require().Equal(int64(1512328), bal.Amount.Int64())
}
