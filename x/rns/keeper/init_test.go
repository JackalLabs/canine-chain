package keeper_test

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	rns "github.com/jackalLabs/canine-chain/x/rns"
	rnsKeeper "github.com/jackalLabs/canine-chain/x/rns/keeper"
	"github.com/jackalLabs/canine-chain/x/rns/types"
)

// testing the msg server
func (suite *KeeperTestSuite) TestMsgInit() {

	suite.SetupSuite()

	msgSrvr, _, context := setupMsgServer(suite)

	//Need to add mock random addresses
	user, err := sdk.AccAddressFromBech32("cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg")
	suite.Require().NoError(err)

	cases := map[string]struct {
		preRun    func() *types.MsgInit
		expErr    bool
		expErrMsg string
	}{
		"all good": {
			preRun: func() *types.MsgInit {
				return types.NewMsgInit(
					user.String(),
				)
			},
			expErr: false,
		},
		"cannot init twice": {
			preRun: func() *types.MsgInit {
				return types.NewMsgInit(
					user.String(),
				)
			},
			expErr:    true,
			expErrMsg: "cannot initialize more than once: invalid request",
		},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			_, err := msgSrvr.Init(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}

}

func setupMsgServer(suite *KeeperTestSuite) (types.MsgServer, rnsKeeper.Keeper, context.Context) {
	k := suite.rnsKeeper
	rns.InitGenesis(suite.ctx, *k, *types.DefaultGenesis())
	ctx := sdk.WrapSDKContext(suite.ctx)
	return rnsKeeper.NewMsgServerImpl(*k), *k, ctx
}

//suite.Require().EqualValues(suite, types.MsgInitResponse{}, res)
