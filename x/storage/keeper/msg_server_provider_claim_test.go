package keeper_test

import (
	testutil "github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (suite *KeeperTestSuite) TestAddProviderClaimer() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)
	msgSrvr, k, context := setupMsgServer(suite)

	alice := testAddresses[0]
	claimAddress := testAddresses[1]

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
				ClaimAddress: claimAddress,
			},
			expErr: false,
		},
		{
			name: "add the same claim addr",
			msg: types.MsgAddClaimer{
				Creator:      alice,
				ClaimAddress: claimAddress,
			},
			expErr: true,
			errMsg: "cannot add the same claimer twice: conflict",
		},
		{
			name: "add claimer to a non-existing provider",
			msg: types.MsgAddClaimer{
				Creator:      "non-existing provider",
				ClaimAddress: claimAddress,
			},
			expErr: true,
			errMsg: "Provider not found. Please init your provider.",
		},
	}
	for _, tcs := range cases {
		tc := tcs
		suite.Run(tc.name, func() {
			_, err := msgSrvr.AddProviderClaimer(context, &tc.msg)
			if tc.expErr {
				suite.Require().EqualError(err, tc.errMsg)
			} else {
				provider, _ := k.GetProviders(suite.ctx, alice)
				suite.Require().Equal(provider.AuthClaimers[0], claimAddress)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestRemoveProviderClaimer() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)
	msgSrvr, k, context := setupMsgServer(suite)

	alice := testAddresses[0]

	provider := types.Providers{
		Address:         alice,
		Ip:              "192.158.1.38",
		Totalspace:      "1280000",
		BurnedContracts: "0",
		Creator:         alice,
		AuthClaimers: []string{
			"claimer_1",
			"claimer_2",
			"claimer_3",
			"claimer_4",
		},
	}

	suite.storageKeeper.SetProviders(suite.ctx, provider)
	suite.Require().NoError(err)

	providerTwo := types.Providers{
		Address:         "provider_two",
		Ip:              "192.188.1.1",
		Totalspace:      "1280000",
		BurnedContracts: "0",
		Creator:         "provider_two",
	}
	suite.storageKeeper.SetProviders(suite.ctx, providerTwo)
	suite.Require().NoError(err)

	cases := []struct {
		name   string
		msg    types.MsgRemoveClaimer
		expErr bool
		errMsg string
	}{
		{
			name: "remove claimer addr",
			msg: types.MsgRemoveClaimer{
				Creator:      alice,
				ClaimAddress: "claimer_4",
			},
			expErr: false,
		},
		{
			name: "remove non-existing claimer addr",
			msg: types.MsgRemoveClaimer{
				Creator:      alice,
				ClaimAddress: "non-existing_claimer",
			},
			expErr: true,
			errMsg: "this address is not a claimer: key not found",
		},
		{
			name: "remove from provider with no claimer",
			msg: types.MsgRemoveClaimer{
				Creator:      "provider_two",
				ClaimAddress: "claimer_address",
			},
			expErr: true,
			errMsg: "Provider has no claimer addresses: conflict",
		},
	}
	for _, tcs := range cases {
		tc := tcs
		suite.Run(tc.name, func() {
			_, err := msgSrvr.RemoveProviderClaimer(context, &tc.msg)
			if tc.expErr {
				suite.Require().EqualError(err, tc.errMsg)
			} else {
				provider, _ := k.GetProviders(suite.ctx, alice)
				suite.Require().Equal(3, len(provider.AuthClaimers))
			}
		})
	}
}
