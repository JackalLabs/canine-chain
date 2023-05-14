package keeper_test

import (
	testutil "github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// testing: contracts.go active_deals.go...
func (suite *KeeperTestSuite) TestSetContracts() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("jkl", 2)
	suite.Require().NoError(err)

	user := testAddresses[0]
	provider := testAddresses[1]

	contract := types.ContractV2{
		Cid:      "549",
		Merkle:   "",
		Signer:   user,
		Duration: "1000",
		FileSize: 5000,
		Fid:      "5789",
		Creator:  provider,
	}

	suite.storageKeeper.SetContracts(suite.ctx, contract)
	suite.Require().NoError(err)

	contractRequest := types.QueryContractRequest{
		Cid: "549",
	}

	res, err := suite.queryClient.Contracts(suite.ctx.Context(), &contractRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(res.Contracts.Cid, contract.Cid)
	suite.Require().Equal(res.Contracts.Signer, contract.Signer)
}

func (suite *KeeperTestSuite) TestGetContracts() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("jkl", 2)
	suite.Require().NoError(err)

	user := testAddresses[0]
	provider := testAddresses[1]

	contract := types.ContractV2{
		Cid:      "549",
		Merkle:   "",
		Signer:   user,
		Duration: "1000",
		FileSize: 5000,
		Fid:      "5789",
		Creator:  provider,
	}

	suite.storageKeeper.SetContracts(suite.ctx, contract)
	suite.Require().NoError(err)

	foundContract, found := suite.storageKeeper.GetContracts(suite.ctx, "549")

	suite.Require().NoError(err)
	suite.Require().Equal(found, true)
	suite.Require().Equal(foundContract.Cid, contract.Cid)
	suite.Require().Equal(foundContract.Signer, contract.Signer)
}

func (suite *KeeperTestSuite) TestGetAllContracts() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("jkl", 3)
	suite.Require().NoError(err)

	provider := testAddresses[0]
	alice := testAddresses[1]
	charlie := testAddresses[2]

	contract := types.ContractV2{
		Cid:      "549",
		Merkle:   "",
		Signer:   alice,
		Duration: "1000",
		FileSize: 5000,
		Fid:      "5789",
		Creator:  provider,
	}

	suite.storageKeeper.SetContracts(suite.ctx, contract)
	suite.Require().NoError(err)

	contract1 := types.ContractV2{
		Cid:      "649",
		Merkle:   "",
		Signer:   charlie,
		Duration: "2000",
		FileSize: 10000,
		Fid:      "4587",
		Creator:  provider,
	}

	suite.storageKeeper.SetContracts(suite.ctx, contract1)
	suite.Require().NoError(err)

	allContracts := suite.storageKeeper.GetAllContracts(suite.ctx)
	aliceContract := allContracts[0]
	charlieContract := allContracts[1]

	suite.Require().NoError(err)
	suite.Require().Equal(aliceContract, contract)
	suite.Require().Equal(charlieContract, contract1)
}

func (suite *KeeperTestSuite) TestRemoveContracts() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("jkl", 2)
	suite.Require().NoError(err)

	user := testAddresses[0]
	provider := testAddresses[1]

	contract := types.ContractV2{
		Cid:      "549",
		Merkle:   "",
		Signer:   user,
		Duration: "1000",
		FileSize: 5000,
		Fid:      "5789",
		Creator:  provider,
	}

	suite.storageKeeper.SetContracts(suite.ctx, contract)
	suite.Require().NoError(err)

	suite.storageKeeper.RemoveContracts(suite.ctx, "549")

	foundContract, found := suite.storageKeeper.GetContracts(suite.ctx, "549")

	ghostContract := types.ContractV2{}

	suite.Require().NoError(err)
	suite.Require().Equal(found, false)
	suite.Require().Equal(foundContract, ghostContract)
}

func (suite *KeeperTestSuite) TestSetActiveDeals() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("jkl", 2)
	suite.Require().NoError(err)

	user := testAddresses[0]
	provider := testAddresses[1]

	deal := types.ActiveDealsV2{
		Cid:           "549",
		Signer:        user,
		Provider:      provider,
		StartBlock:    100,
		EndBlock:      1000,
		FileSize:      5000,
		ProofVerified: true,
		ProofsMissed:  0,
		BlockToProve:  150,
		Creator:       user,
		Merkle:        "",
		Fid:           "5789",
	}

	suite.storageKeeper.SetActiveDeals(suite.ctx, deal)
	suite.Require().NoError(err)

	dealRequest := types.QueryActiveDealRequest{
		Cid: "549",
	}

	res, err := suite.queryClient.ActiveDeals(suite.ctx.Context(), &dealRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(res.ActiveDeals.Cid, deal.Cid)
	suite.Require().Equal(res.ActiveDeals.Signer, deal.Signer)
}

func (suite *KeeperTestSuite) TestGetActiveDeals() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("jkl", 2)
	suite.Require().NoError(err)

	user := testAddresses[0]
	provider := testAddresses[1]

	deal := types.ActiveDealsV2{
		Cid:           "549",
		Signer:        user,
		Provider:      provider,
		StartBlock:    100,
		EndBlock:      1000,
		FileSize:      5000,
		ProofVerified: true,
		ProofsMissed:  0,
		BlockToProve:  150,
		Creator:       user,
		Merkle:        "",
		Fid:           "5789",
	}

	suite.storageKeeper.SetActiveDeals(suite.ctx, deal)
	suite.Require().NoError(err)

	foundDeal, found := suite.storageKeeper.GetActiveDeals(suite.ctx, "549")

	suite.Require().NoError(err)
	suite.Require().Equal(found, true)
	suite.Require().Equal(foundDeal.Cid, deal.Cid)
	suite.Require().Equal(foundDeal.Signer, deal.Signer)
}

func (suite *KeeperTestSuite) TestGetAllActiveDeals() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("jkl", 3)
	suite.Require().NoError(err)

	provider := testAddresses[0]
	alice := testAddresses[1]
	charlie := testAddresses[2]

	deal := types.ActiveDealsV2{
		Cid:           "549",
		Signer:        alice,
		Provider:      provider,
		StartBlock:    100,
		EndBlock:      1000,
		FileSize:      5000,
		ProofVerified: true,
		ProofsMissed:  0,
		BlockToProve:  150,
		Creator:       alice,
		Merkle:        "",
		Fid:           "5789",
	}

	suite.storageKeeper.SetActiveDeals(suite.ctx, deal)
	suite.Require().NoError(err)

	deal1 := types.ActiveDealsV2{
		Cid:           "1458",
		Signer:        charlie,
		Provider:      provider,
		StartBlock:    200,
		EndBlock:      2000,
		FileSize:      10000,
		ProofVerified: true,
		ProofsMissed:  0,
		BlockToProve:  200,
		Creator:       charlie,
		Merkle:        "",
		Fid:           "4589",
	}

	suite.storageKeeper.SetActiveDeals(suite.ctx, deal1)
	suite.Require().NoError(err)

	allDeals := suite.storageKeeper.GetAllActiveDeals(suite.ctx)
	aliceDeal := allDeals[1]
	charleDeal := allDeals[0]

	suite.Require().NoError(err)
	suite.Require().Equal(aliceDeal, deal)
	suite.Require().Equal(charleDeal, deal1)
}

func (suite *KeeperTestSuite) TestRemoveActiveDeals() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("jkl", 2)
	suite.Require().NoError(err)

	user := testAddresses[0]
	provider := testAddresses[1]

	deal := types.ActiveDealsV2{
		Cid:           "549",
		Signer:        user,
		Provider:      provider,
		StartBlock:    100,
		EndBlock:      1000,
		FileSize:      5000,
		ProofVerified: true,
		ProofsMissed:  0,
		BlockToProve:  150,
		Creator:       user,
		Merkle:        "",
		Fid:           "5789",
	}

	suite.storageKeeper.SetActiveDeals(suite.ctx, deal)
	suite.Require().NoError(err)

	suite.storageKeeper.RemoveActiveDeals(suite.ctx, "549")

	foundDeal, found := suite.storageKeeper.GetActiveDeals(suite.ctx, "549")

	ghostDeal := types.ActiveDealsV2{}

	suite.Require().NoError(err)
	suite.Require().Equal(found, false)
	suite.Require().Equal(foundDeal, ghostDeal)
}

func (suite *KeeperTestSuite) TestSetStrays() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("jkl", 1)
	suite.Require().NoError(err)

	user := testAddresses[0]

	stray := types.StrayV2{
		Cid:      "549",
		Fid:      "5789",
		Signer:   user,
		FileSize: 1000,
		Merkle:   "",
		DealType: 0,
	}
	suite.storageKeeper.SetStrays(suite.ctx, stray)
	suite.Require().NoError(err)

	strayRequest := types.QueryStrayRequest{
		Cid: "549",
	}

	res, err := suite.queryClient.Strays(suite.ctx.Context(), &strayRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(res.Strays.Cid, stray.Cid)
	suite.Require().Equal(res.Strays.Fid, stray.Fid)
}

func (suite *KeeperTestSuite) TestGetStrays() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("jkl", 1)
	suite.Require().NoError(err)

	user := testAddresses[0]

	stray := types.StrayV2{
		Cid:      "549",
		Fid:      "5789",
		Signer:   user,
		FileSize: 1000,
		Merkle:   "",
	}
	suite.storageKeeper.SetStrays(suite.ctx, stray)
	suite.Require().NoError(err)

	foundStray, found := suite.storageKeeper.GetStrays(suite.ctx, "549")

	suite.Require().NoError(err)
	suite.Require().Equal(found, true)
	suite.Require().Equal(foundStray.Cid, stray.Cid)
	suite.Require().Equal(foundStray.Fid, stray.Fid)
}

func (suite *KeeperTestSuite) TestGetAllStrays() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("jkl", 1)
	suite.Require().NoError(err)

	user := testAddresses[0]

	stray := types.StrayV2{
		Cid:      "549",
		Fid:      "5789",
		Signer:   user,
		FileSize: 1000,
		Merkle:   "",
	}
	suite.storageKeeper.SetStrays(suite.ctx, stray)
	suite.Require().NoError(err)

	stray1 := types.StrayV2{
		Cid:      "649",
		Fid:      "5789",
		Signer:   user,
		FileSize: 1000,
		Merkle:   "",
	}
	suite.storageKeeper.SetStrays(suite.ctx, stray1)
	suite.Require().NoError(err)

	foundStrays := suite.storageKeeper.GetAllStrays(suite.ctx)
	foundStayZero := foundStrays[0]
	foundStayOne := foundStrays[1]

	suite.Require().NoError(err)
	suite.Require().Equal(foundStayZero, stray)
	suite.Require().Equal(foundStayOne, stray1)
}

func (suite *KeeperTestSuite) TestRemoveStrays() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("jkl", 1)
	suite.Require().NoError(err)

	user := testAddresses[0]

	stray := types.StrayV2{
		Cid:      "549",
		Fid:      "5789",
		Signer:   user,
		FileSize: 1000,
		Merkle:   "",
	}
	suite.storageKeeper.SetStrays(suite.ctx, stray)
	suite.Require().NoError(err)

	suite.storageKeeper.RemoveStrays(suite.ctx, "549")

	foundStray, found := suite.storageKeeper.GetStrays(suite.ctx, "549")
	ghostStray := types.StrayV2{}

	suite.Require().NoError(err)
	suite.Require().Equal(found, false)
	suite.Require().Equal(foundStray, ghostStray)
}

func (suite *KeeperTestSuite) TestSetFidCid() {
	suite.SetupSuite()

	FidCid := types.FidCid{Fid: "549", Cids: "['628', '629', '630']"}

	suite.storageKeeper.SetFidCid(suite.ctx, FidCid)

	FidCidRequest := types.QueryFidCidRequest{
		Fid: "549",
	}

	res, err := suite.queryClient.FidCid(suite.ctx.Context(), &FidCidRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(res.FidCid.Fid, FidCid.Fid)
}

func (suite *KeeperTestSuite) TestGetFidCid() {
	suite.SetupSuite()

	FidCid := types.FidCid{Fid: "549", Cids: "['628', '629', '630']"}

	suite.storageKeeper.SetFidCid(suite.ctx, FidCid)

	foundFidCid, found := suite.storageKeeper.GetFidCid(suite.ctx, "549")

	suite.Require().Equal(found, true)
	suite.Require().Equal(foundFidCid.Fid, FidCid.Fid)
}

func (suite *KeeperTestSuite) TestGetAllFidCid() {
	suite.SetupSuite()

	FidCid := types.FidCid{Fid: "549", Cids: "['628', '629', '630']"}

	suite.storageKeeper.SetFidCid(suite.ctx, FidCid)

	FidCid1 := types.FidCid{Fid: "649", Cids: "['728', '729', '730']"}

	suite.storageKeeper.SetFidCid(suite.ctx, FidCid1)

	foundAllFidCid := suite.storageKeeper.GetAllFidCid(suite.ctx)

	suite.Require().Equal(foundAllFidCid[0], FidCid)
	suite.Require().Equal(foundAllFidCid[1], FidCid1)
}

func (suite *KeeperTestSuite) TestRemoveFidCid() {
	suite.SetupSuite()

	FidCid := types.FidCid{Fid: "549", Cids: "['628', '629', '630']"}
	suite.storageKeeper.SetFidCid(suite.ctx, FidCid)

	suite.storageKeeper.RemoveFidCid(suite.ctx, "549")

	foundFidCid, found := suite.storageKeeper.GetFidCid(suite.ctx, "549")
	ghostFidCid := types.FidCid{}
	suite.Require().Equal(found, false)
	suite.Require().Equal(foundFidCid, ghostFidCid)
}
