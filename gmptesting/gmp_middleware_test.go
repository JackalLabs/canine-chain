package gmp_testing

import (
	"testing"

	ibctesting "github.com/cosmos/ibc-go/v4/testing"
	"github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
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
	suite.Setup(suite.T())

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

func (suite *GMPTestSuite) TestOnRecvPacket() {
	var (
		trace    transfertypes.DenomTrace
		amount   sdk.Int
		receiver string
		//  status   testutils.Status don't think we need for now
	)

	// need this later

	suite.SetupTest() // reset

	path := NewTransferPath(suite.chainA, suite.chainB)
	suite.coordinator.Setup(path)
	receiver = suite.chainB.SenderAccount.GetAddress().String() // looks like this is auto generated
	// status = testutils.Status{} don't think we need a status for now

	amount = sdk.NewInt(100)
	seq := uint64(1)

	trace = transfertypes.ParseDenomTrace(sdk.DefaultBondDenom)

	// do we need to send coins first?
	// send coin from chainA to chainB
	transferMsg := transfertypes.NewMsgTransfer(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, sdk.NewCoin(trace.IBCDenom(), amount), suite.chainA.SenderAccount.GetAddress().String(), receiver, clienttypes.NewHeight(1, 110), 0)
	_, err := suite.chainA.SendMsgs(transferMsg)
	suite.Require().NoError(err) // message committed

	data := transfertypes.NewFungibleTokenPacketData(trace.GetFullDenomPath(), amount.String(), suite.chainA.SenderAccount.GetAddress().String(), receiver)
	data.Memo = "placeholder memo"
	packet := channeltypes.NewPacket(data.GetBytes(), seq, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, clienttypes.NewHeight(1, 100), 0)

	// we expect a returned acknowledgement
	logger, logFile := testutil.CreateLogger()
	logger.Println("*********************")
	logger.Printf("%+v\n", suite.chainB.GetContext())
	logger.Printf("%+v\n", packet)
	logger.Printf("%+v\n", suite.chainA.SenderAccount.GetAddress())
	ack := suite.chainB.GetJackalApp().GmpStack.OnRecvPacket(suite.chainB.GetContext(), packet, suite.chainA.SenderAccount.GetAddress())
	logger.Printf("The ack is %+v\n", ack)

	logFile.Close()

	// suite.Require().True(ack.Success())
}

func (suite *GMPTestSuite) TestRecvTransferWithMetadata() {
	// Setup contract
	suite.chainA.StoreContractCode(&suite.Suite, "./bytecode/echo.wasm")
}
