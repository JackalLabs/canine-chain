package keeper_test

import (
	testutil "github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (suite *KeeperTestSuite) TestAddProviderClaimer() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)
	msgSrvr, k, context := setupMsgServer(suite)

	alice := testAddresses[0]
	claim_address := testAddresses[1]

	provider := types.Providers{
		Address:         alice,
		Ip:              "192.158.1.38",
		Totalspace:      "1280000",
		BurnedContracts: "0",
		Creator:         alice,
	}

	suite.storageKeeper.SetProviders(suite.ctx, provider)
	suite.Require().NoError(err)

	cases := []struct {
		name   string
		msg    types.MsgAddClaimer
		expErr bool
		errMsg string
	}{
		{
			name: "add a new claim addr",
			msg: types.MsgAddClaimer{
				Creator:      alice,
				ClaimAddress: claim_address,
			},
			expErr: false,
		},
		{
			name: "add the same claim addr",
			msg: types.MsgAddClaimer{
				Creator:      alice,
				ClaimAddress: claim_address,
			},
			expErr: true,
			errMsg: "cannot add the same claimer twice: conflict",
		},
		{
			name: "add claimer to a non-existing provider",
			msg: types.MsgAddClaimer{
				Creator:      "non-existing provider",
				ClaimAddress: claim_address,
			},
			expErr: true,
			errMsg: "Provider not found. Please init your provider.",
		},
	}
	for _, tc := range cases {
		suite.Run(tc.name, func() {
			_, err := msgSrvr.AddProviderClaimer(context, &tc.msg)
			if tc.expErr {
				suite.Require().EqualError(err, tc.errMsg)
			} else {
				provider, _ := k.GetProviders(suite.ctx, alice)
				suite.Require().Equal(provider.AuthClaimers[0], claim_address)
			}
		})
	}
}
