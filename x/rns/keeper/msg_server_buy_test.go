package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/jackalLabs/canine-chain/x/rns/types"
)

func (suite *KeeperTestSuite) TestBuyMsg() {
	suite.SetupSuite()
	msgSrvr, _, ctx := setupMsgServer(suite)
}
