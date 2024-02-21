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
		Cid:          cid,
	}

	suite.storageKeeper.SetReportForm(suite.ctx, report)

	reportRequest := types.QueryReportRequest{
		Cid: cid,
	}

	res, err := suite.queryClient.Reports(suite.ctx.Context(), &reportRequest)
	suite.Require().NoError(err)
	suite.Require().Equal(report.Cid, res.Report.Cid)
	suite.Require().Equal(report.Attestations, res.Report.Attestations)
}

func (suite *KeeperTestSuite) TestGetReportForm() {
	suite.SetupSuite()

	var att []*types.Attestation
	report := types.ReportForm{
		Attestations: att,
		Cid:          cid,
	}

	suite.storageKeeper.SetReportForm(suite.ctx, report)

	foundReport, found := suite.storageKeeper.GetReportForm(suite.ctx, cid)
	suite.Require().Equal(found, true)
	suite.Require().Equal(foundReport.Cid, report.Cid)
	suite.Require().Equal(foundReport.Attestations, report.Attestations)
}

func (suite *KeeperTestSuite) TestGetAllReportForm() {
	suite.SetupSuite()

	report := types.ReportForm{
		Attestations: []*types.Attestation{},
		Cid:          cid,
	}

	report2 := types.ReportForm{
		Attestations: []*types.Attestation{},
		Cid:          "jklc1321",
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
		Cid:          cid,
	}
	suite.storageKeeper.SetReportForm(suite.ctx, report)

	suite.storageKeeper.RemoveReport(suite.ctx, cid)

	foundReport, found := suite.storageKeeper.GetReportForm(suite.ctx, cid)
	suite.Require().Equal(found, false)

	var atts []*types.Attestation
	ghostReport := types.ReportForm{
		Attestations: atts,
		Cid:          "",
	}

	suite.Require().Equal(foundReport, ghostReport)
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

	deal := types.ActiveDeals{
		Cid:          cid,
		Signee:       "",
		Provider:     addresses[10],
		Startblock:   "",
		Endblock:     "0",
		Filesize:     "",
		LastProof:    0,
		Proofsmissed: "",
		Blocktoprove: "",
		Creator:      "",
		Merkle:       "",
		Fid:          "",
	}

	suite.storageKeeper.SetActiveDeals(suite.ctx, deal) // creating storage deal

	_, err = suite.storageKeeper.RequestReport(suite.ctx, cid)
	suite.NoError(err)

	form, found := suite.storageKeeper.GetReportForm(suite.ctx, cid)
	suite.Equal(true, found)

	for _, attestation := range form.Attestations {
		fmt.Printf("%s %t\n", attestation.Provider, attestation.Complete)
	}

	_ = form
	allReportForm := suite.storageKeeper.GetAllReport(suite.ctx)
	suite.Require().Equal(1, len(allReportForm))

	d, found := suite.storageKeeper.GetActiveDeals(suite.ctx, cid)
	suite.Equal(true, found)
	p := suite.storageKeeper.GetParams(suite.ctx)
	verified := d.IsVerified(suite.ctx.BlockHeight(), p.ProofWindow)
	suite.Equal(false, verified)

	for i, attestation := range form.Attestations {
		err := suite.storageKeeper.Report(suite.ctx, cid, attestation.Provider)
		if i >= int(params.AttestMinToPass) {
			suite.Require().Error(err)
		} else {
			suite.Require().NoError(err)
		}
	}

	_, found = suite.storageKeeper.GetReportForm(suite.ctx, cid)
	suite.Equal(false, found)

	_, found = suite.storageKeeper.GetActiveDeals(suite.ctx, cid)
	suite.Equal(false, found)

	_, found = suite.storageKeeper.GetStrays(suite.ctx, cid)
	suite.Equal(true, found)
}
