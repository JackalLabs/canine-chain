package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (suite *KeeperTestSuite) TestCreateContracts () {
	suite.SetupSuite()
	msgSrvr, storageKeeper, goCtx := setupMsgServer(suite)
	creator, err := sdk.AccAddressFromBech32("cosmos17j2hkm7n9fz9dpntyj2kxgxy5pthzd289nvlfl")
	suite.Require().NoError(err)

	cases := map[string]struct{
		preRun func() *types.MsgCreateContracts
		postRun func()
		expErr bool
		expErrMsg string
	}{
		"contract_already_exists": {
			preRun: func() *types.MsgCreateContracts {
				c := types.Contracts{
					Creator: creator.String(),
					Cid: "1",
					Priceamt: "1",
					Pricedenom: "1",
					Merkle: "1",
					Signee: "1",
					Duration: "1",
					Filesize: "1",
					Fid: "1",
				}
				storageKeeper.SetContracts(suite.ctx, c)
				return &types.MsgCreateContracts{
					Creator: c.Creator,
					Cid: c.Cid,
					Priceamt: c.Priceamt,
					Pricedenom: c.Pricedenom,
					Merkle: c.Merkle,
					Signee: c.Signee,
					Duration: c.Duration,
					Filesize: c.Filesize,
					Fid: c.Fid,
				}
			},
			postRun: func() {
				c, found := storageKeeper.GetContracts(suite.ctx, "1")
				suite.Require().True(found)
				storageKeeper.RemoveContracts(suite.ctx, c.Cid)
				_, found = storageKeeper.GetContracts(suite.ctx, c.Cid)
				suite.Require().True(found)
			
			},
			expErr: true,
			expErrMsg: "index already set",
		},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
			// preRun must be defined to get MsgCreateContracts
			suite.Require().NotNil(tc.preRun())
			c := tc.preRun()
			_, err := msgSrvr.CreateContracts(goCtx, c)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}

			if tc.postRun != nil {
				tc.postRun()
			}
		})
	}
}
