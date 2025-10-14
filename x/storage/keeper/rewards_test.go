package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil "github.com/jackalLabs/canine-chain/v4/testutil"
	jklminttypes "github.com/jackalLabs/canine-chain/v4/x/jklmint/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"github.com/tendermint/tendermint/libs/rand"
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

func (suite *KeeperTestSuite) TestEmptyReward() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	signer := testAddresses[0]

	s := suite.ctx.BlockTime()

	t := s.AddDate(1, 0, 0) // simulate buying a whole year
	suite.storageKeeper.SetStoragePaymentInfo(suite.ctx, types.StoragePaymentInfo{
		Start:          s,
		End:            t,
		SpaceAvailable: 1000000000,
		SpaceUsed:      0,
		Address:        signer,
	})

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(6000000)))

	gauge := suite.storageKeeper.NewGauge(suite.ctx, coins, t)
	gaugeAccount, err := types.GetGaugeAccount(gauge)
	suite.NoError(err)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, jklminttypes.ModuleName, gaugeAccount, coins)
	suite.NoError(err)

	bal := suite.bankKeeper.GetBalance(suite.ctx, gaugeAccount, "ujkl")
	suite.Require().Equal(int64(6000000), bal.Amount.Int64())

	providerOne := testAddresses[1]
	pOneAcc, err := sdk.AccAddressFromBech32(providerOne)
	suite.NoError(err)

	bal = suite.bankKeeper.GetBalance(suite.ctx, pOneAcc, "ujkl")
	suite.Require().Equal(int64(0), bal.Amount.Int64())

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

	_, found := suite.storageKeeper.GetProof(suite.ctx, providerOne, dealOne.Merkle, dealOne.Owner, dealOne.Start)
	suite.Require().True(found)

	newTime := s.AddDate(0, 3, 0)
	ctx := suite.ctx.WithBlockHeight(blocks).WithHeaderHash([]byte{10, 15, 16, 20}).WithBlockTime(newTime)

	suite.Require().Equal(blocks, ctx.BlockHeight())
	suite.Require().Equal(ctx.BlockHeight()%blocks, int64(0))

	suite.storageKeeper.ManageRewards(ctx)

	bal = suite.bankKeeper.GetBalance(suite.ctx, pOneAcc, "ujkl")
	suite.Require().Equal(int64(0), bal.Amount.Int64())
}

func (suite *KeeperTestSuite) TestReward() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	signer := testAddresses[0]

	s := suite.ctx.BlockTime()

	t := s.AddDate(1, 0, 0) // simulate buying a whole year
	suite.storageKeeper.SetStoragePaymentInfo(suite.ctx, types.StoragePaymentInfo{
		Start:          s,
		End:            t,
		SpaceAvailable: 1000000000,
		SpaceUsed:      0,
		Address:        signer,
	})

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(6000000)))

	gauge := suite.storageKeeper.NewGauge(suite.ctx, coins, t)
	gaugeAccount, err := types.GetGaugeAccount(gauge)
	suite.NoError(err)

	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, jklminttypes.ModuleName, gaugeAccount, coins)
	suite.NoError(err)

	bal := suite.bankKeeper.GetBalance(suite.ctx, gaugeAccount, "ujkl")
	suite.Require().Equal(int64(6000000), bal.Amount.Int64())

	providerOne := testAddresses[1]
	pOneAcc, err := sdk.AccAddressFromBech32(providerOne)
	suite.NoError(err)

	bal = suite.bankKeeper.GetBalance(suite.ctx, pOneAcc, "ujkl")
	suite.Require().Equal(int64(0), bal.Amount.Int64())

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
	proof := types.FileProof{
		Prover:       providerOne,
		Merkle:       dealOne.Merkle,
		Owner:        dealOne.Owner,
		Start:        dealOne.Start,
		LastProven:   blocks - 1,
		ChunkToProve: 0,
	}
	suite.storageKeeper.UpdateProof(suite.ctx, &proof, &dealOne)

	_, found := suite.storageKeeper.GetProof(suite.ctx, providerOne, dealOne.Merkle, dealOne.Owner, dealOne.Start)
	suite.Require().True(found)

	newTime := s.AddDate(0, 3, 0)
	ctx := suite.ctx.WithBlockHeight(blocks).WithHeaderHash([]byte{10, 15, 16, 20}).WithBlockTime(newTime)

	suite.Require().Equal(blocks, ctx.BlockHeight())
	suite.Require().Equal(ctx.BlockHeight()%blocks, int64(0))

	testDiff := gauge.End.Sub(s)
	realDiff := newTime.Sub(s)

	ratio := float64(realDiff.Microseconds()) / float64(testDiff.Microseconds())
	r := ratio * float64(6000000)

	suite.storageKeeper.ManageRewards(ctx)

	bal = suite.bankKeeper.GetBalance(suite.ctx, pOneAcc, "ujkl")
	suite.Require().Equal(int64(r), bal.Amount.Int64())
}

func (suite *KeeperTestSuite) TestLongTermReward() {
	for i := 1; i < 10; i++ {
		suite.SetupSuite()

		testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
		suite.Require().NoError(err)

		signer := testAddresses[0]

		s := suite.ctx.BlockTime()

		totalBlocks := int64(1600)

		timePerBlock := int64(i)

		totalBlockTime := totalBlocks * timePerBlock
		t := s.Add(time.Second * time.Duration(totalBlockTime) / 2) // simulate buying only half the simulation time
		suite.storageKeeper.SetStoragePaymentInfo(suite.ctx, types.StoragePaymentInfo{
			Start:          s,
			End:            t,
			SpaceAvailable: 1000000000,
			SpaceUsed:      0,
			Address:        signer,
		})

		coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(6000000)))

		gauge := suite.storageKeeper.NewGauge(suite.ctx, coins, t)
		gaugeAccount, err := types.GetGaugeAccount(gauge)
		suite.NoError(err)

		err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, jklminttypes.ModuleName, gaugeAccount, coins)
		suite.NoError(err)

		totalSpend := int64(6000000)

		bal := suite.bankKeeper.GetBalance(suite.ctx, gaugeAccount, "ujkl")
		suite.Require().Equal(totalSpend, bal.Amount.Int64())

		providerOne := testAddresses[1]
		pOneAcc, err := sdk.AccAddressFromBech32(providerOne)
		suite.NoError(err)

		bal = suite.bankKeeper.GetBalance(suite.ctx, pOneAcc, "ujkl")
		suite.Require().Equal(int64(0), bal.Amount.Int64())

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

		var b int64
		blockTime := suite.ctx.BlockTime()
		for i := int64(1); i < totalBlocks; i++ {
			suite.T().Logf("Block: %d", i)
			p, found := suite.storageKeeper.GetProof(suite.ctx, providerOne, dealOne.Merkle, dealOne.Owner, dealOne.Start)
			suite.Require().True(found)
			p.LastProven = i
			suite.storageKeeper.UpdateProof(suite.ctx, &p, &dealOne)

			blockTime = blockTime.Add(time.Second * time.Duration(timePerBlock)) // step forward 6 seconds
			suite.ctx = suite.ctx.WithBlockHeight(i).WithHeaderHash(rand.Bytes(20)).WithBlockTime(blockTime)

			testDiff := gauge.End.Sub(s)
			realDiff := time.Second * time.Duration(timePerBlock) // step forward 6 seconds

			ratio := float64(realDiff.Microseconds()) / float64(testDiff.Microseconds())
			r := ratio * float64(totalSpend)
			suite.T().Logf("Ratio: %f -> %f", ratio, r)
			suite.storageKeeper.ManageRewards(suite.ctx)

			bal = suite.bankKeeper.GetBalance(suite.ctx, pOneAcc, "ujkl")
			ibal := bal.Amount.Int64()
			suite.T().Logf("Account has %d tokens", ibal)
			dif := ibal - b
			if i > totalBlocks/2 {
				suite.Require().Equal(int64(0), dif)
			} else {
				suite.Require().Equal(int64(r), dif)
			}
			b = ibal

		}
	}
}

func (suite *KeeperTestSuite) TestLongTermRewardWithWindows() {
	for j := int64(1); j < 10; j++ {
		suite.SetupSuite()

		testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
		suite.Require().NoError(err)

		signer := testAddresses[0]

		s := suite.ctx.BlockTime()

		totalBlocks := int64(1600)

		timePerBlock := int64(6)

		totalBlockTime := totalBlocks * timePerBlock
		t := s.Add(time.Second * time.Duration(totalBlockTime) / 2) // simulate buying only half the simulation time
		suite.storageKeeper.SetStoragePaymentInfo(suite.ctx, types.StoragePaymentInfo{
			Start:          s,
			End:            t,
			SpaceAvailable: 1000000000,
			SpaceUsed:      0,
			Address:        signer,
		})

		coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(6000000)))

		gauge := suite.storageKeeper.NewGauge(suite.ctx, coins, t)
		gaugeAccount, err := types.GetGaugeAccount(gauge)
		suite.NoError(err)

		err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, jklminttypes.ModuleName, gaugeAccount, coins)
		suite.NoError(err)

		totalSpend := int64(6000000)

		bal := suite.bankKeeper.GetBalance(suite.ctx, gaugeAccount, "ujkl")
		suite.Require().Equal(totalSpend, bal.Amount.Int64())

		providerOne := testAddresses[1]
		pOneAcc, err := sdk.AccAddressFromBech32(providerOne)
		suite.NoError(err)

		bal = suite.bankKeeper.GetBalance(suite.ctx, pOneAcc, "ujkl")
		suite.Require().Equal(int64(0), bal.Amount.Int64())

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

		var b int64
		blockTime := suite.ctx.BlockTime()
		for i := int64(1); i < totalBlocks; i++ {
			suite.T().Logf("Block: %d", i)
			p, found := suite.storageKeeper.GetProof(suite.ctx, providerOne, dealOne.Merkle, dealOne.Owner, dealOne.Start)
			suite.Require().True(found)
			p.LastProven = i
			suite.storageKeeper.UpdateProof(suite.ctx, &p, &dealOne)

			blockTime = blockTime.Add(time.Second * time.Duration(timePerBlock)) // step forward 6 seconds
			suite.ctx = suite.ctx.WithBlockHeight(i).WithHeaderHash(rand.Bytes(20)).WithBlockTime(blockTime)

			if suite.ctx.BlockHeight()%j > 0 {
				continue
			}

			testDiff := gauge.End.Sub(s)
			realDiff := time.Second * time.Duration(timePerBlock*j) // step forward 6 seconds

			ratio := float64(realDiff.Microseconds()) / float64(testDiff.Microseconds())
			r := ratio * float64(totalSpend)
			suite.T().Logf("Ratio: %f -> %f", ratio, r)
			suite.storageKeeper.ManageRewards(suite.ctx)

			bal = suite.bankKeeper.GetBalance(suite.ctx, pOneAcc, "ujkl")
			ibal := bal.Amount.Int64()
			suite.T().Logf("Account has %d tokens", ibal)
			dif := ibal - b
			if i > totalBlocks/2 {
				suite.Require().Equal(int64(0), dif)
			} else {
				suite.Require().Equal(int64(r), dif)
			}
			b = ibal

		}
	}
}
