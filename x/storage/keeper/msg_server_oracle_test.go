package keeper_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/utils"

	"github.com/stretchr/testify/require"
)

func TestMakeAddress(t *testing.T) {
	address, err := keeper.MakeAddress()
	require.NoError(t, err)
	t.Log(address.String())
}

func (suite *KeeperTestSuite) TestRequestAndFulfillOracle() {
	suite.SetupSuite()
	msgSrvr, k, ctx := setupMsgServer(suite)

	replicaFileData := []byte("this is test data for a tree I will be testing against. BLOCKCHAIN AND STORAGE NETWORK DELIVERING ON-CHAIN ACCESS TO DATA STORAGE EVERYWHERE.")
	fileData := []byte("this is test data for a tree I will be testing against. BLOCKCHAIN AND STORAGE NETWORK DELIVERING ON-CHAIN ACCESS TO DATA STORAGE EVERYWHERE.")
	buf := bytes.NewBuffer(fileData)
	var chunkSize int64 = 2
	tree, _, _, err := utils.BuildJustTree(buf, chunkSize)
	suite.Require().NoError(err)

	suite.T().Log(tree.Data)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 3)
	suite.Require().NoError(err)

	testAccount := testAddresses[0]
	testOracleAccount := testAddresses[2]
	depoAccount := testAddresses[1]

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(100000000000))) // Send some coins to their account
	testAcc, _ := sdk.AccAddressFromBech32(testAccount)
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, testAcc, coins)
	suite.Require().NoError(err)

	suite.storageKeeper.SetParams(suite.ctx, types.Params{
		DepositAccount:         depoAccount,
		ProofWindow:            50,
		ChunkSize:              chunkSize,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     8,
		CollateralPrice:        2,
		CheckWindow:            10,
	})

	testMerkle := tree.Root()
	var chunk int64 = 10
	index := chunk * chunkSize
	chunkData := replicaFileData[index : index+chunkSize]

	h := sha256.New()
	_, err = h.Write([]byte(fmt.Sprintf("%d%x", chunk, chunkData)))
	suite.Require().NoError(err)

	c := h.Sum(nil)

	proof, err := tree.GenerateProof(c, 0)
	suite.Require().NoError(err)
	jproof, err := json.Marshal(*proof)
	suite.Require().NoError(err)

	bid, err := sdk.ParseCoinNormalized("500ujkl")
	suite.Require().NoError(err)

	requestChunkMessage := types.MsgRequestChunk{
		Creator: testAccount,
		Merkle:  testMerkle,
		Chunk:   chunk,
		Bid:     bid,
	}

	_, err = msgSrvr.RequestChunk(ctx, &requestChunkMessage)
	suite.Require().NoError(err)

	_, found := k.GetOracleRequest(suite.ctx, testAccount, testMerkle, chunk)
	suite.Require().True(found)

	fulfillMessage := types.MsgFulfillRequest{
		Creator:   testOracleAccount,
		Requester: testAccount,
		Merkle:    testMerkle,
		Chunk:     chunk,
		Data:      chunkData,
		HashList:  jproof,
	}

	_, err = msgSrvr.FulfillRequest(ctx, &fulfillMessage)
	suite.Require().NoError(err)

	_, found = k.GetOracleRequest(suite.ctx, testAccount, testMerkle, chunk)
	suite.Require().False(found)

	_, found = k.GetOracleEntry(suite.ctx, testAccount, testMerkle, chunk)
	suite.Require().True(found)

	all := k.GetAllOracleEntries(suite.ctx)
	for _, entry := range all {
		suite.T().Log(entry.String())
	}

	tAcc, err := sdk.AccAddressFromBech32(testAccount)
	suite.Require().NoError(err)

	oAcc, err := sdk.AccAddressFromBech32(testOracleAccount)
	suite.Require().NoError(err)

	bal := suite.bankKeeper.GetBalance(suite.ctx, oAcc, "ujkl")
	suite.Require().Equal(int64(500), bal.Amount.Int64())

	bal = suite.bankKeeper.GetBalance(suite.ctx, tAcc, "ujkl")
	suite.Require().Equal(int64(100000000000-500), bal.Amount.Int64())

	suite.reset()
}
