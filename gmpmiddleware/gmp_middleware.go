package gmp_middleware

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v4/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	"github.com/jackalLabs/canine-chain/v3/testutil"
)

type IBCMiddleware struct {
	app     porttypes.IBCModule
	handler GeneralMessageHandler
}

func NewIBCMiddleware(app porttypes.IBCModule, handler GeneralMessageHandler) IBCMiddleware {
	return IBCMiddleware{
		app:     app,
		handler: handler,
	}
}

// OnChanOpenInit implements the IBCModule interface
func (im IBCMiddleware) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	// call underlying callback
	return im.app.OnChanOpenInit(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, version)
}

// OnChanOpenTry implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	return im.app.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, channelCap, counterparty, counterpartyVersion)
}

// OnChanOpenAck implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	return im.app.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
}

// OnChanOpenConfirm implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return im.app.OnChanOpenConfirm(ctx, portID, channelID)
}

// OnChanCloseInit implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return im.app.OnChanCloseInit(ctx, portID, channelID)
}

// OnChanCloseConfirm implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return im.app.OnChanCloseConfirm(ctx, portID, channelID)
}

// OnRecvPacket implements the IBCMiddleware interface
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	// Comment out call to underlying OnRecvPacket for now, we apply our own logic
	// ack := im.app.OnRecvPacket(ctx, packet, relayer)

	// declare an ack that we always return so we can dissect the code without erroring
	var ack ibcexported.Acknowledgement
	// if !ack.Success() {
	// 	return ack
	// }

	isIcs20, data := isIcs20Packet(packet)
	logger, logFile := testutil.CreateLogger()
	logger.Println("*********************")
	logger.Printf("Is it ICS20? %t\n", isIcs20)
	logger.Printf("%+v\n", data)

	logFile.Close()

	// The below will always make our tests fail unless we hard code that the sender is the AxelarGMPAcc
	// // authenticate the message with packet sender + channel-id
	// // TODO: authenticate the message with channel-id
	// if data.Sender != AxelarGMPAcc {
	// 	return ack
	// }

	// var msg Message
	// var err error

	// if err = json.Unmarshal([]byte(data.GetMemo()), &msg); err != nil {
	// 	return channeltypes.NewErrorAcknowledgement(fmt.Errorf("cannot unmarshal memo"))
	// }

	// switch msg.Type {
	// case TypeGeneralMessage:
	// 	// implement the handler
	// 	err = im.handler.HandleGeneralMessage(ctx, msg.SourceChain, msg.SourceAddress, data.Receiver, msg.Payload)
	// default:
	// 	err = fmt.Errorf("unrecognized mesasge type: %d", msg.Type)
	// }

	// if err != nil {
	// 	return channeltypes.NewErrorAcknowledgement(err)
	// }

	return ack
}

// OnAcknowledgementPacket implements the IBCMiddleware interface
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	return im.app.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

// OnTimeoutPacket implements the IBCMiddleware interface
func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	return im.app.OnTimeoutPacket(ctx, packet, relayer)
}
