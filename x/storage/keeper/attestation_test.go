package keeper_test

import (
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

const cid = "jklc123"

// testing attestations.go file
func (suite *KeeperTestSuite) TestSetAttestationForm() {
	suite.SetupSuite()

	var att []*types.Attestation
	attestation := types.AttestationForm{
		Attestations: att,
		Cid:          cid,
	}

	suite.storageKeeper.SetAttestationForm(suite.ctx, attestation)

	attestationRequest := types.QueryAttestationRequest{
		Cid: cid,
	}

	res, err := suite.queryClient.Attestation(suite.ctx.Context(), &attestationRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(attestation.Cid, res.Attestation.Cid)
	suite.Require().Equal(attestation.Attestations, res.Attestation.Attestations)
}

func (suite *KeeperTestSuite) TestGetAttestationForm() {
	suite.SetupSuite()

	attestation := types.AttestationForm{
		Attestations: []*types.Attestation{},
		Cid:          cid,
	}

	suite.storageKeeper.SetAttestationForm(suite.ctx, attestation)

	foundAttestation, found := suite.storageKeeper.GetAttestationForm(suite.ctx, cid)
	suite.Require().Equal(found, true)
	suite.Require().Equal(foundAttestation.Cid, attestation.Cid)
	suite.Require().Equal(foundAttestation.Attestations, attestation.Attestations)
}

func (suite *KeeperTestSuite) TestGetAllAttestationForm() {
	suite.SetupSuite()

	attestation := types.AttestationForm{
		Attestations: []*types.Attestation{},
		Cid:          cid,
	}

	attestation2 := types.AttestationForm{
		Attestations: []*types.Attestation{},
		Cid:          "jklc1321",
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
		Cid:          cid,
	}
	suite.storageKeeper.SetAttestationForm(suite.ctx, attestation)

	suite.storageKeeper.RemoveAttestation(suite.ctx, cid)

	foundAttestation, found := suite.storageKeeper.GetAttestationForm(suite.ctx, cid)
	suite.Require().Equal(found, false)

	var atts []*types.Attestation
	ghostAttestation := types.AttestationForm{
		Attestations: atts,
		Cid:          "",
	}

	suite.Require().Equal(foundAttestation, ghostAttestation)
}
