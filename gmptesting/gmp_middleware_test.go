package gmp_testing

import (
	"testing"

	ibctesting "github.com/cosmos/ibc-go/v4/testing"
	"github.com/stretchr/testify/suite"

	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
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
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 2)
	suite.chainA = TestChain{
		TestChain: suite.coordinator.GetChain(ibctesting.GetChainID(1)),
	}
	suite.chainB = TestChain{
		TestChain: suite.coordinator.GetChain(ibctesting.GetChainID(2)),
	}

	suite.pathAB = NewTransferPath(suite.chainA, suite.chainB)
	suite.coordinator.Setup(suite.pathAB)
	suite.pathBA = NewTransferPath(suite.chainB, suite.chainA)
	suite.coordinator.Setup(suite.pathBA)
}

func NewTransferPath(chainA, chainB TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA.TestChain, chainB.TestChain)
	path.EndpointA.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointA.ChannelConfig.Version = transfertypes.Version
	path.EndpointB.ChannelConfig.Version = transfertypes.Version

	return path
}
