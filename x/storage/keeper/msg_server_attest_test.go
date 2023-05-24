package keeper_test

import (
	testutil "github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (suite *KeeperTestSuite) TestAttest() {
	addresses, err := testutil.CreateTestAddresses("cosmos", keeper.FormSize+2)
	suite.Require().NoError(err)

	validCid := "cid1"
	noActiveDealCid := "no_active_deal_cid"

	cases := map[string]struct {
		cid     string
		creator string
		expErr  bool
	}{
		"attestation form not found": {
			cid:     "I do not exist",
			creator: addresses[keeper.FormSize],
			expErr:  true,
		},
		"not requested provider": {
			cid:     validCid,
			creator: "not requested provider",
			expErr:  true,
		},
		"active deal not found": {
			cid:     noActiveDealCid,
			creator: addresses[keeper.FormSize],
			expErr:  true,
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
				Cid:          validCid,
			}

			noActiveDealAttestForm := types.AttestationForm{
				Attestations: attestations,
				Cid:          noActiveDealCid,
			}

			suite.storageKeeper.SetAttestationForm(suite.ctx, attestForm)
			suite.storageKeeper.SetAttestationForm(suite.ctx, noActiveDealAttestForm)

			activeDeal := types.ActiveDeals{
				Cid: validCid,
			}

			suite.storageKeeper.SetActiveDeals(suite.ctx, activeDeal)

			err = suite.storageKeeper.Attest(suite.ctx, tc.cid, tc.creator)

			if tc.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
			}
		})
	}
}
