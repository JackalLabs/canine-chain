package keeper_test

import (
//	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil "github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
)

func (suite *KeeperTestSuite) TestAttest() {
	addresses, err := testutil.CreateTestAddresses("cosmos", keeper.FormSize+2)
	suite.Require().NoError(err)

	validCid := "cid1"

	cases := map[string]struct {
		cid string
		creator string
		expErr bool
	}{
		"attestation form not found": {
			cid: "I do not exist",
			creator: addresses[keeper.FormSize],
			expErr: true,
		},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
			suite.SetupSuite()
			
			attestations := make([]*types.Attestation, keeper.FormSize)

			providerAddresses := make([]string, keeper.FormSize)

			for i := 0; i < keeper.FormSize; i++ {
				providerAddresses[i] = addresses[i]

				attestations[i] = &types.Attestation{
					Provider: addresses[i],
					Complete: false,
				}
			}

			attestForm := types.AttestationForm{
				Attestations: attestations,
				Cid: validCid,
			}

			suite.storageKeeper.SetAttestationForm(suite.ctx, attestForm)

			err := suite.storageKeeper.Attest(suite.ctx, tc.cid, tc.creator)

			if !tc.expErr {
				suite.NoError(err)
			}
		})
	}
}
