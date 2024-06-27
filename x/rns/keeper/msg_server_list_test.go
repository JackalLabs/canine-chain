package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/testutil"
	types "github.com/jackalLabs/canine-chain/v4/x/rns/types"
)

func (suite *KeeperTestSuite) TestListMsg() {
	accs, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	names := []string{"jackal", "maxi", "free", "expired"}

	tc := map[string]struct {
		Creator string
		Name    string
		Price   sdk.Coin
		expErr  bool
	}{
		"list": {
			Creator: accs[1],
			Name:    names[1],
			Price:   sdk.NewInt64Coin("ujkl", 10000),
			expErr:  false,
		},
		"name already listed": {
			Creator: accs[0],
			Name:    names[0],
			Price:   sdk.NewInt64Coin("ujkl", 10000),
			expErr:  true,
		},
		"name not found": {
			Creator: accs[0],
			Name:    "null",
			Price:   sdk.NewInt64Coin("ujkl", 10000),
			expErr:  true,
		},
		"invalid owner": {
			Creator: accs[0],
			Name:    names[1],
			Price:   sdk.NewInt64Coin("ujkl", 10000),
			expErr:  true,
		},
		"expired": {
			Creator: accs[1],
			Name:    names[3],
			Price:   sdk.NewInt64Coin("ujkl", 10000),
			expErr:  true,
		},
	}

	for name, tc := range tc {
		suite.Run(name, func() {
			suite.SetupSuite()
			suite.ctx = suite.ctx.WithBlockHeight(0)
			msgSrvr, _, ctx := setupMsgServer(suite)

			rns := types.Names{
				Name:       names[0],
				Expires:    suite.ctx.BlockHeight() + 5733818,
				Value:      accs[0],
				Data:       "{}",
				Subdomains: nil,
				Tld:        types.SupportedTLDs[0],
				Locked:     0,
			}
			suite.rnsKeeper.SetNames(suite.ctx, rns)

			rns.Name, rns.Value = names[1], accs[1]
			suite.rnsKeeper.SetNames(suite.ctx, rns)

			rns.Name, rns.Locked = names[2], suite.ctx.BlockHeight()+1
			suite.rnsKeeper.SetNames(suite.ctx, rns)

			rns.Name, rns.Expires, rns.Locked = names[3], -1, 0
			suite.rnsKeeper.SetNames(suite.ctx, rns)

			tld := types.SupportedTLDs[0]

			msg := &types.MsgList{
				Creator: accs[0],
				Name:    names[0] + "." + tld,
				Price:   sdk.NewInt64Coin("ujkl", 10000),
			}
			_, err = msgSrvr.List(ctx, msg)
			suite.Require().NoError(err)

			msg = &types.MsgList{
				Creator: tc.Creator,
				Name:    tc.Name + "." + tld,
				Price:   tc.Price,
			}

			res, err := msgSrvr.List(ctx, msg)
			if !tc.expErr {
				suite.Require().NoError(err)
				suite.NotNil(res)
			}
		})
	}
}
