package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackalLabs/canine-chain/x/rns/types"
)

func (suite *KeeperTestSuite) TestMsgInit() {

	suite.SetupSuite()
	err := suite.setupNames()

	suite.Require().NoError(err)
	address, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	initMsg := types.NewMsgInit(address.String())

	_, err1 := suite.msgSrvr.Init(sdk.WrapSDKContext(suite.ctx), initMsg)
	suite.Require().NoError(err1)

	initReq := types.QueryGetInitRequest{
		Address: address.String(),
	}

	_, err2 := suite.queryClient.Init(suite.ctx.Context(), &initReq)
	suite.Require().NoError(err2)

	//init again should fail
	_, err3 := suite.msgSrvr.Init(sdk.WrapSDKContext(suite.ctx), initMsg)
	suite.Require().Error(err3)

}
