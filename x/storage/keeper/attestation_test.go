package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/testutil"
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

// testing attestations.go file
func (suite *KeeperTestSuite) TestRewardsAttestationForm() {
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

	addresses, err := testutil.CreateTestAddresses("cosmos", 50)
	suite.NoError(err)

	res, err := suite.queryClient.Attestation(suite.ctx.Context(), &attestationRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(attestation.Cid, res.Attestation.Cid)
	suite.Require().Equal(attestation.Attestations, res.Attestation.Attestations)

	address, err := sdk.AccAddressFromBech32(addresses[0])
	suite.Require().NoError(err)

	err = suite.storageKeeper.InternalRewards(suite.ctx, make([]types.ActiveDeals, 0), address)
	suite.Require().NoError(err)

	_, found := suite.storageKeeper.GetAttestationForm(suite.ctx, cid)
	suite.Require().Equal(false, found)
}

func (suite *KeeperTestSuite) TestGetAttestationForm() {
	suite.SetupSuite()

	var att []*types.Attestation
	attestation := types.AttestationForm{
		Attestations: att,
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

func (suite *KeeperTestSuite) TestMakeAttestation() {
	suite.SetupSuite()

	addresses, err := testutil.CreateTestAddresses("jkl", 50)
	suite.NoError(err)

	for _, address := range addresses {
		suite.storageKeeper.SetActiveProviders(suite.ctx, types.ActiveProviders{
			Address: address,
		})
	}

	deal := types.ActiveDeals{
		Cid:           cid,
		Signee:        "",
		Provider:      addresses[10],
		Startblock:    "",
		Endblock:      "",
		Filesize:      "",
		Proofverified: "false",
		Proofsmissed:  "",
		Blocktoprove:  "",
		Creator:       "",
		Merkle:        "",
		Fid:           "",
	}

	suite.storageKeeper.SetActiveDeals(suite.ctx, deal) // creating storage deal

	_, err = suite.storageKeeper.RequestAttestation(suite.ctx, cid, addresses[10])
	suite.NoError(err)

	form, found := suite.storageKeeper.GetAttestationForm(suite.ctx, cid)
	suite.Equal(true, found)

	for _, attestation := range form.Attestations {
		fmt.Printf("%s %t\n", attestation.Provider, attestation.Complete)
	}

	_ = form
	allAttestationForm := suite.storageKeeper.GetAllAttestation(suite.ctx)
	suite.Require().Equal(1, len(allAttestationForm))

	d, found := suite.storageKeeper.GetActiveDeals(suite.ctx, cid)
	suite.Equal(true, found)
	suite.Equal("false", d.Proofverified)

	for _, attestation := range form.Attestations {
		err := suite.storageKeeper.Attest(suite.ctx, cid, attestation.Provider)
		suite.Require().NoError(err)
	}

	_, found = suite.storageKeeper.GetAttestationForm(suite.ctx, cid)
	suite.Equal(false, found)

	d, found = suite.storageKeeper.GetActiveDeals(suite.ctx, cid)
	suite.Equal(true, found)
	suite.Equal("true", d.Proofverified)
}
