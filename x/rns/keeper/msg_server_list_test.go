package keeper_test

import (
	types "github.com/jackalLabs/canine-chain/x/rns/types"
)

func (suite *KeeperTestSuite) TestListMsg () {
	suite.reset()
	keeper := suite.rnsKeeper
	msgSrvr, _, ctx := setupMsgServer(suite)
	err := suite.setupNames()
	suite.Require().NoError(err)

	cases := map[string]struct {
		preRun func() *types.MsgList
		postRun func() error
		expErr bool
		expErrMsg string
	}{
		"name_already_listed": {
			preRun: func() *types.MsgList {
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
			postRun: func() error {
				// Clean up forsale name
				name, found := keeper.GetNames(suite.ctx, "name1", "jkl")
				suite.Require().True(found)
				keeper.RemoveForsale(suite.ctx, name.Name)
				_, found = keeper.GetForsale(suite.ctx, name.Name)
				suite.Require().False(found)
				return nil
			},
			expErr: true,
			expErrMsg: "Name already listed.",
		},

		"name_not_found": {
			preRun: func() *types.MsgList {
				return &types.MsgList{
					Creator: "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
					Name: "nonexistentjkl",
					Price: "100ujkl",
				}
			},
			expErr: true,
			expErrMsg: "Name does not exist or has expired.",
		},

		"wrong_owner": {
			preRun: func() *types.MsgList {
				blockHeight := suite.ctx.BlockHeight()
				name := types.Names {
					Name: "free_name_",
					Locked: blockHeight + 1,
					Expires: blockHeight + 1,
					Value: "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
					Data: "{}",
					Tld: "jkl",
				}
				keeper.SetNames(suite.ctx, name)
				_, found := keeper.GetNames(suite.ctx, name.Name, name.Tld)
				suite.Require().True(found)
				return &types.MsgList{
					Creator: "wrong_account",
					Name: "free_name_jkl", 
					Price: "100ujkl",
 				}
			},
			postRun: func() error {
				// Clean up name
				name, found := keeper.GetNames(suite.ctx, "free_name_", "jkl")
				suite.Require().True(found)
				keeper.RemoveNames(suite.ctx, name.Name, name.Tld)
				_, found = keeper.GetNames(suite.ctx, name.Name, name.Tld)
				suite.Require().False(found)
				return nil
			},
			expErr: true,
			expErrMsg: "You do not own this name.",
		},

		"cannot_transfer_free_name": {
			preRun: func() *types.MsgList {
				blockHeight := suite.ctx.BlockHeight()
				freeName := types.Names {
					Name: "free_name",
					Locked: blockHeight + 1,
					Expires: blockHeight + 1,
					Value: "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
					Data: "_",
					Tld: "jkl",
				}
				keeper.SetNames(suite.ctx, freeName)
				return &types.MsgList{
					Creator: "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
					Name: "free_name_jkl",
					Price: "100ujkl",
				}
			},
			expErr: true,
			expErrMsg: "cannot transfer free name",
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

			if tc.postRun != nil {
				err = tc.postRun()
				suite.Require().NoError(err)
			}
		})
	}
}
