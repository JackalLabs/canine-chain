package keeper_test

import (
	"fmt"

	"github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

// testing attestations.go file
func (suite *KeeperTestSuite) TestSetAttestationForm() {
	suite.SetupSuite()

	var att []*types.Attestation
	attestation := types.AttestationForm{
		Attestations: att,
		Prover:       "prover",
		Merkle:       []byte("merkle"),
		Owner:        "owner",
		Start:        0,
	}

	suite.storageKeeper.SetAttestationForm(suite.ctx, attestation)

	attestationRequest := types.QueryAttestation{
		Prover: "prover",
		Merkle: []byte("merkle"),
		Owner:  "owner",
		Start:  0,
	}

	res, err := suite.queryClient.Attestation(suite.ctx.Context(), &attestationRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(attestation.Prover, res.Attestation.Prover)
	suite.Require().Equal(attestation.Owner, res.Attestation.Owner)
	suite.Require().Equal(attestation.Start, res.Attestation.Start)
	suite.Require().Equal(attestation.Merkle, res.Attestation.Merkle)
	suite.Require().Equal(attestation.Attestations, res.Attestation.Attestations)
}

func (suite *KeeperTestSuite) TestGetAttestationForm() {
	suite.SetupSuite()

	var att []*types.Attestation
	attestation := types.AttestationForm{
		Attestations: att,
		Prover:       "prover",
		Merkle:       []byte{},
		Owner:        "owner",
		Start:        0,
	}

	suite.storageKeeper.SetAttestationForm(suite.ctx, attestation)

	foundAttestation, found := suite.storageKeeper.GetAttestationForm(suite.ctx, "prover", []byte{}, "owner", 0)
	suite.Require().Equal(found, true)
	suite.Require().Equal(foundAttestation.Prover, attestation.Prover)
	suite.Require().Equal(foundAttestation.Attestations, attestation.Attestations)
}

func (suite *KeeperTestSuite) TestGetAllAttestationForm() {
	suite.SetupSuite()

	attestation := types.AttestationForm{
		Attestations: []*types.Attestation{},
		Prover:       "prover",
		Merkle:       []byte("merkle"),
		Owner:        "owner",
		Start:        0,
	}

	attestation2 := types.AttestationForm{
		Attestations: []*types.Attestation{},
		Prover:       "prover2",
		Merkle:       []byte("merkle2"),
		Owner:        "owner2",
		Start:        0,
	}

	allAttestationFormbefore := suite.storageKeeper.GetAllAttestation(suite.ctx)
	suite.Require().Equal(0, len(allAttestationFormbefore))

	suite.storageKeeper.SetAttestationForm(suite.ctx, attestation)
	suite.storageKeeper.SetAttestationForm(suite.ctx, attestation2)

	allAttestationForm := suite.storageKeeper.GetAllAttestation(suite.ctx)
	suite.Require().Equal(2, len(allAttestationForm))
}

func (suite *KeeperTestSuite) TestRemoveAttestationForm() {
	suite.SetupSuite()

	attestation := types.AttestationForm{
		Attestations: []*types.Attestation{},
		Prover:       "prover",
		Merkle:       []byte("merkle"),
		Owner:        "owner",
		Start:        0,
	}
	suite.storageKeeper.SetAttestationForm(suite.ctx, attestation)

	suite.storageKeeper.RemoveAttestation(suite.ctx, "prover", []byte{}, "owner", 0)

	_, found := suite.storageKeeper.GetAttestationForm(suite.ctx, "prover", []byte{}, "owner", 0)
	suite.Require().Equal(found, false)
}

func (suite *KeeperTestSuite) TestMakeAttestation() {
	suite.SetupSuite()
	params := suite.storageKeeper.GetParams(suite.ctx)

	addresses, err := testutil.CreateTestAddresses("jkl", 50)
	suite.NoError(err)

	for i, address := range addresses {
		realProvider := types.Providers{
			Address: address,
			Ip:      fmt.Sprintf("https://test%d.com", i),
		}

		suite.storageKeeper.SetProviders(suite.ctx, realProvider)
		suite.storageKeeper.SetActiveProviders(suite.ctx, types.ActiveProviders{
			Address: address,
		})
	}

	file := types.UnifiedFile{
		Merkle:        []byte("merkle"),
		Owner:         "owner",
		Start:         0,
		Expires:       0,
		FileSize:      100,
		ProofInterval: 100,
		ProofType:     0,
		Proofs:        make([]string, 0),
		MaxProofs:     3,
		Note:          "test",
	}
	suite.storageKeeper.SetFile(suite.ctx, file) // creating storage deal

	file.AddProver(suite.ctx, suite.storageKeeper, addresses[10])

	_, err = suite.storageKeeper.RequestAttestation(suite.ctx, []byte("merkle"), "owner", 0, addresses[10])
	suite.NoError(err)

	form, found := suite.storageKeeper.GetAttestationForm(suite.ctx, addresses[10], []byte("merkle"), "owner", 0)
	suite.Equal(true, found)

	for _, attestation := range form.Attestations {
		fmt.Printf("%s %t\n", attestation.Provider, attestation.Complete)
	}

	_ = form
	allAttestationForm := suite.storageKeeper.GetAllAttestation(suite.ctx)
	suite.Require().Equal(1, len(allAttestationForm))

	d, found := suite.storageKeeper.GetProof(suite.ctx, addresses[10], []byte("merkle"), "owner", 0)
	suite.Equal(true, found)
	suite.Equal(int64(0), d.LastProven)

	for i, attestation := range form.Attestations {
		err := suite.storageKeeper.Attest(suite.ctx, addresses[10], []byte("merkle"), "owner", 0, attestation.Provider)
		if i >= int(params.AttestMinToPass) {
			suite.Require().Error(err)
		} else {
			suite.Require().NoError(err)
		}
	}

	_, found = suite.storageKeeper.GetAttestationForm(suite.ctx, addresses[10], []byte("merkle"), "owner", 0)
	suite.Equal(false, found)

	d, found = suite.storageKeeper.GetProof(suite.ctx, addresses[10], []byte("merkle"), "owner", 0)
	suite.Equal(true, found)
	suite.Equal(int64(0), d.LastProven)
}
