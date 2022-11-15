package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

//testing: contracts.go active_deals.go...

func (suite *KeeperTestSuite) TestSetActiveDeals() {
	suite.SetupSuite()
	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	provider, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	deal := types.ActiveDeals{
		Cid:           "549",
		Signee:        user.String(),
		Provider:      provider.String(),
		Startblock:    "100",
		Endblock:      "1000",
		Filesize:      "5000",
		Proofverified: "true",
		Proofsmissed:  "0",
		Blocktoprove:  "150",
		Creator:       user.String(),
		Merkle:        "",
		Fid:           "5789",
	}

	suite.storageKeeper.SetActiveDeals(suite.ctx, deal)
	suite.Require().NoError(err)

	dealRequest := types.QueryGetActiveDealsRequest{
		Cid: "549",
	}

	res, err := suite.queryClient.ActiveDeals(suite.ctx.Context(), &dealRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(res.ActiveDeals.Cid, deal.Cid)
	suite.Require().Equal(res.ActiveDeals.Signee, deal.Signee)

}

func (suite *KeeperTestSuite) TestGetActiveDeals() {
	suite.SetupSuite()
	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	provider, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	deal := types.ActiveDeals{
		Cid:           "549",
		Signee:        user.String(),
		Provider:      provider.String(),
		Startblock:    "100",
		Endblock:      "1000",
		Filesize:      "5000",
		Proofverified: "true",
		Proofsmissed:  "0",
		Blocktoprove:  "150",
		Creator:       user.String(),
		Merkle:        "",
		Fid:           "5789",
	}

	suite.storageKeeper.SetActiveDeals(suite.ctx, deal)
	suite.Require().NoError(err)

	foundDeal, found := suite.storageKeeper.GetActiveDeals(suite.ctx, "549")

	suite.Require().NoError(err)
	suite.Require().Equal(found, true)
	suite.Require().Equal(foundDeal.Cid, deal.Cid)
	suite.Require().Equal(foundDeal.Signee, deal.Signee)

}

func (suite *KeeperTestSuite) TestRemoveActiveDeals() {
	suite.SetupSuite()
	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	provider, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	deal := types.ActiveDeals{
		Cid:           "549",
		Signee:        user.String(),
		Provider:      provider.String(),
		Startblock:    "100",
		Endblock:      "1000",
		Filesize:      "5000",
		Proofverified: "true",
		Proofsmissed:  "0",
		Blocktoprove:  "150",
		Creator:       user.String(),
		Merkle:        "",
		Fid:           "5789",
	}

	suite.storageKeeper.SetActiveDeals(suite.ctx, deal)
	suite.Require().NoError(err)

	suite.storageKeeper.RemoveActiveDeals(suite.ctx, "549")

	foundDeal, found := suite.storageKeeper.GetActiveDeals(suite.ctx, "549")

	ghostDeal := types.ActiveDeals{}

	suite.Require().NoError(err)
	suite.Require().Equal(found, false)
	suite.Require().Equal(foundDeal, ghostDeal)

}

func (suite *KeeperTestSuite) TestGetAllActiveDeals() {
	suite.SetupSuite()

	provider, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	alice, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	charlie, err := sdk.AccAddressFromBech32("cosmos1xetrp5dwjplsn4lev5r2cu8en5qsq824vza9nu")
	suite.Require().NoError(err)

	deal := types.ActiveDeals{
		Cid:           "549",
		Signee:        alice.String(),
		Provider:      provider.String(),
		Startblock:    "100",
		Endblock:      "1000",
		Filesize:      "5000",
		Proofverified: "true",
		Proofsmissed:  "0",
		Blocktoprove:  "150",
		Creator:       alice.String(),
		Merkle:        "",
		Fid:           "5789",
	}

	suite.storageKeeper.SetActiveDeals(suite.ctx, deal)
	suite.Require().NoError(err)

	deal1 := types.ActiveDeals{
		Cid:           "1458",
		Signee:        charlie.String(),
		Provider:      provider.String(),
		Startblock:    "200",
		Endblock:      "2000",
		Filesize:      "10000",
		Proofverified: "true",
		Proofsmissed:  "0",
		Blocktoprove:  "200",
		Creator:       charlie.String(),
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
