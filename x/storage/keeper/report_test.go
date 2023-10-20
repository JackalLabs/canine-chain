package keeper_test

import (
	"fmt"

	"github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

// testing attestations.go file
func (suite *KeeperTestSuite) TestSetReportForm() {
	suite.SetupSuite()

	var att []*types.Attestation
	report := types.ReportForm{
		Attestations: att,
		Prover:       "prover",
		Merkle:       []byte("merkle"),
		Owner:        "owner",
		Start:        0,
	}

	suite.storageKeeper.SetReportForm(suite.ctx, report)

	reportRequest := types.QueryReportRequest{
		Prover: "prover",
		Merkle: []byte("merkle"),
		Owner:  "owner",
		Start:  0,
	}

	res, err := suite.queryClient.Reports(suite.ctx.Context(), &reportRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(report.Prover, res.Report.Prover)
	suite.Require().Equal(report.Attestations, res.Report.Attestations)
}

func (suite *KeeperTestSuite) TestGetReportForm() {
	suite.SetupSuite()

	var att []*types.Attestation
	report := types.ReportForm{
		Attestations: att,
		Prover:       "prover",
		Merkle:       []byte("merkle"),
		Owner:        "owner",
		Start:        0,
	}

	suite.storageKeeper.SetReportForm(suite.ctx, report)

	foundReport, found := suite.storageKeeper.GetReportForm(suite.ctx, "prover", []byte{}, "owner", 0)
	suite.Require().Equal(found, true)
	suite.Require().Equal(foundReport.Merkle, report.Merkle)
	suite.Require().Equal(foundReport.Attestations, report.Attestations)
}

func (suite *KeeperTestSuite) TestGetAllReportForm() {
	suite.SetupSuite()

	report := types.ReportForm{
		Attestations: []*types.Attestation{},
		Prover:       "prover",
		Merkle:       []byte("merkle"),
		Owner:        "owner",
		Start:        0,
	}

	report2 := types.ReportForm{
		Attestations: []*types.Attestation{},
		Prover:       "prover",
		Merkle:       []byte("merkle"),
		Owner:        "owner",
		Start:        0,
	}

	allReportFormsBefore := suite.storageKeeper.GetAllReport(suite.ctx)
	suite.Require().Equal(0, len(allReportFormsBefore))

	suite.storageKeeper.SetReportForm(suite.ctx, report)
	suite.storageKeeper.SetReportForm(suite.ctx, report2)

	allReportForm := suite.storageKeeper.GetAllReport(suite.ctx)
	suite.Require().Equal(2, len(allReportForm))
}

func (suite *KeeperTestSuite) TestRemoveReportForm() {
	suite.SetupSuite()

	report := types.ReportForm{
		Attestations: []*types.Attestation{},
		Prover:       "prover",
		Merkle:       []byte("merkle"),
		Owner:        "owner",
		Start:        0,
	}
	suite.storageKeeper.SetReportForm(suite.ctx, report)

	suite.storageKeeper.RemoveReport(suite.ctx, "prover", []byte{}, "owner", 0)

	_, found := suite.storageKeeper.GetReportForm(suite.ctx, "prover", []byte{}, "owner", 0)
	suite.Require().Equal(found, false)
}

func (suite *KeeperTestSuite) TestMakeReport() {
	suite.SetupSuite()
	params := suite.storageKeeper.GetParams(suite.ctx)

	addresses, err := testutil.CreateTestAddresses("jkl", 50)
	suite.NoError(err)

	for i, address := range addresses {
		realProvider := types.Providers{
			Address:         address,
			Ip:              fmt.Sprintf("https://test%d.com", i),
			BurnedContracts: "0",
		}

		suite.storageKeeper.SetProviders(suite.ctx, realProvider)

		suite.storageKeeper.SetActiveProviders(suite.ctx, types.ActiveProviders{
			Address: address,
		})
	}

	uf := types.UnifiedFile{
		Merkle:        []byte("merkle"),
		Owner:         "owner",
		Start:         0,
		Expires:       0,
		FileSize:      1000,
		ProofInterval: 100,
		ProofType:     0,
		Proofs:        make([]string, 0),
		MaxProofs:     3,
		Note:          "test",
	}
	suite.storageKeeper.SetFile(suite.ctx, uf)
	uf.AddProver(suite.ctx, suite.storageKeeper, addresses[10])
	suite.storageKeeper.SetFile(suite.ctx, uf)

	_, err = suite.storageKeeper.RequestReport(suite.ctx, addresses[10], uf.Merkle, uf.Owner, uf.Start)
	suite.NoError(err)

	form, found := suite.storageKeeper.GetReportForm(suite.ctx, addresses[10], uf.Merkle, uf.Owner, uf.Start)
	suite.Equal(true, found)

	for _, attestation := range form.Attestations {
		fmt.Printf("%s %t\n", attestation.Provider, attestation.Complete)
	}

	allReportForm := suite.storageKeeper.GetAllReport(suite.ctx)
	suite.Require().Equal(1, len(allReportForm))

	_, found = suite.storageKeeper.GetProof(suite.ctx, addresses[10], uf.Merkle, uf.Owner, uf.Start)
	suite.Equal(true, found)

	for i, attestation := range form.Attestations {
		err := suite.storageKeeper.Report(suite.ctx, addresses[10], uf.Merkle, uf.Owner, uf.Start, attestation.Provider)
		if i >= int(params.AttestMinToPass) {
			suite.Require().Error(err)
		} else {
			suite.Require().NoError(err)
		}
	}

	_, found = suite.storageKeeper.GetReportForm(suite.ctx, addresses[10], uf.Merkle, uf.Owner, uf.Start)
	suite.Equal(false, found)
}
