package gmp_testing

import (
	"testing"

	ibctesting "github.com/cosmos/ibc-go/v4/testing"
	"github.com/stretchr/testify/suite"
)

type GMPTestSuite struct {
	GMPTestHelper

	coordinator *ibctesting.Coordinator

	chainA TestChain
	chainB TestChain

	pathAB *ibctesting.Path
	pathBA *ibctesting.Path
}

func TestGMPTestSuite(t *testing.T) {
	suite.Run(t, new(GMPTestSuite))
}

func (suite *GMPTestSuite) SetupTest() {
	suite.Setup()
	ibctesting.DefaultTestingAppInit = SetupTestingApp
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 3)
}
