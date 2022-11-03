package keeper_test

import (
	types "github.com/jackalLabs/canine-chain/x/rns/types"
)

func (suite *KeeperTestSuite) TestListMsg () {
	suite.reset()
	keeper := suite.rnsKeeper
	msgSrvr, _, ctx := setupMsgServer(suite)

	cases := map[string]struct {
		preRun func() *types.MsgList
		expErr bool
		expErrMsg string
	}{
		"name_already_listed": {
			preRun: func() *types.MsgList {
				err := suite.setupNames()
				suite.Require().NoError(err)

				name, found := keeper.GetNames(suite.ctx, "name1", "jkl")
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
				}
			},
			expErr: true,
			expErrMsg: "Name already listed.",
		},
	}

	for name, tc := range cases {
		suite.Run(name, func(){
			msg := tc.preRun()

			_, err := msgSrvr.List(ctx, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
