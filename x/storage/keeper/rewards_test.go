package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
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

	const signer = "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg"
	const providerOne = "cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl"

	dealOne := types.ActiveDeals{
		Cid:           "cid1test",
		Signee:        signer,
		Provider:      providerOne,
		Startblock:    "0",
		Endblock:      "0",
		Filesize:      "100",
		Proofverified: "true",
		Proofsmissed:  "0",
		Blocktoprove:  "1",
		Creator:       providerOne,
		Merkle:        "nil",
		Fid:           "fid1test",
	}

	suite.storageKeeper.SetActiveDeals(suite.ctx, dealOne)

	suite.storageKeeper.HandleRewardBlock(suite.ctx)
}
