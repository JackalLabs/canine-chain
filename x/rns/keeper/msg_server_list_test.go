package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil "github.com/jackalLabs/canine-chain/testutil"
	types "github.com/jackalLabs/canine-chain/x/rns/types"
)

func (suite *KeeperTestSuite) TestListMsg () {
	suite.reset()
	keeper := suite.rnsKeeper
	names := types.Names {}


	cases := map[string]struct {
		preRun func() (*types.MsgList, error)
		expErr bool
		expErrMsg string
	}{
		"name_already_listed": {
			preRun: func() (*types.MsgList, error) {
				err := suite.setupNames()
				return nil, err

				name, found := keeper.GetNames(suite.ctx, "1", "jkl") 
				suite.Require().True(found)
				newsale := types.Forsale{
					Name: name.Name,
					Price: "100000000ujkl",
					Owner: "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
				}
				keeper.SetForsale(suite.ctx, newsale)
				return &types.MsgList{
					Creator: "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
					Name: name.Name,
					Price: "123123453819283ujkl",
				}, nil
			},
			expErr: true,
			expErrMsg: "Name already listed.",
		},
	}
}
