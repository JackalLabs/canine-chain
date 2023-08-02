package gmp_middleware

import (
	"encoding/json"
	"fmt"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v4/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"

	"github.com/jackalLabs/canine-chain/v3/testutil"
)

type IBCMiddleware struct {
	channel porttypes.ICS4Wrapper
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

	// isIcs20, data := isIcs20Packet(packet)

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

func (im IBCMiddleware) SendPacket(capabilitykeeper capabilitykeeper.ScopedKeeper, ibcKeeper ibckeeper.Keeper, ctx sdk.Context, chanCap *capabilitytypes.Capability, packet ibcexported.PacketI) error {
	concretePacket, ok := packet.(channeltypes.Packet)
	if !ok {
		return im.channel.SendPacket(ctx, chanCap, packet) // send packet will return an error
	}

	isIcs20, data := isIcs20Packet(concretePacket)
	if !isIcs20 {
		return im.channel.SendPacket(ctx, chanCap, packet) // send packet will return an error
	}

	isCallbackRouted, metadata := jsonStringHasKey(data.GetMemo(), IBCCallbackKey) // metadata was here
	fmt.Println(isCallbackRouted)
	/*
		to do: what's the purpose for this call back?
		older versions of IBC did not have a memo but will we ever interact with chains that are still running very old versions of IBC?

		if !isCallbackRouted {
			return im.channel.SendPacket(ctx, chanCap, packet) // send packet will return an error
		}
		// We remove the callback metadata from the memo as it has already been processed.

		// If the only available key in the memo is the callback, we should remove the memo
		// from the data completely so the packet is sent without it.
		// This way receiver chains that are on old versions of IBC will be able to process the packet

		callbackRaw := metadata[IBCCallbackKey] // This will be used later.
		delete(metadata, IBCCallbackKey)
	*/
	bzMetadata, err := json.Marshal(metadata)
	if err != nil {
		return errorsmod.Wrap(err, "Send packet with callback error")
	}

	stringMetadata := string(bzMetadata)
	if stringMetadata == "{}" {
		data.Memo = ""
	} else {
		data.Memo = stringMetadata
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return errorsmod.Wrap(err, "Send packet with callback error")
	}

	packetWithoutCallbackMemo := channeltypes.Packet{
		Sequence:           concretePacket.Sequence,
		SourcePort:         concretePacket.SourcePort,
		SourceChannel:      concretePacket.SourceChannel,
		DestinationPort:    concretePacket.DestinationPort,
		DestinationChannel: concretePacket.DestinationChannel,
		Data:               dataBytes,
		TimeoutTimestamp:   concretePacket.TimeoutTimestamp,
		TimeoutHeight:      concretePacket.TimeoutHeight,
	}

	ibcChannelKeeper := ibcKeeper.ChannelKeeper
	logger, logFile := testutil.CreateLogger()

	// Can the keeper retrieve the channel from our packet?
	channel, found := ibcChannelKeeper.GetChannel(ctx, packetWithoutCallbackMemo.GetSourcePort(), packetWithoutCallbackMemo.GetSourceChannel())
	logger.Printf("Channel found? %t. Channel: %#v\n", found, channel)
	logger.Printf("Channel: %#v\n", channel)
	logger.Printf("Channel state: %#v\n", channel.State)

	// Does caller own capability for the channel?
	isCapable := capabilitykeeper.AuthenticateCapability(ctx, chanCap, host.ChannelCapabilityPath(packet.GetSourcePort(), packet.GetSourceChannel()))
	logger.Printf("isCapable? %t", isCapable)

	// packet destination port matches the counterparty's port?
	portMatched := packet.GetDestPort() == channel.Counterparty.PortId
	logger.Printf("ports match? %t", portMatched)

	// packet destination channel match the counterparty's channel?
	destinationMatch := packet.GetDestChannel() == channel.Counterparty.ChannelId
	logger.Printf("channels match? %t", destinationMatch)

	// has a connection even been established?
	connectionEnd, found := ibcKeeper.ConnectionKeeper.GetConnection(ctx, channel.ConnectionHops[0])
	logger.Printf("connection established? %t", found)
	logger.Printf("connectionEnd: %#v\n", connectionEnd)

	// bypass the ICS4 wrapper and call the channel keeper directly works O.o
	err = ibcChannelKeeper.SendPacket(ctx, chanCap, packet)

	// err = SafeSendPacket(im.channel, ctx, chanCap, packetWithoutCallbackMemo)
	// if err != nil {
	// 	logger.Println(err)
	// 	return err
	// }

	// // Make sure the callback contract is a string and a valid bech32 addr. If it isn't, ignore this packet
	// contract, ok := callbackRaw.(string)
	// if !ok {
	// 	return nil
	// }
	// _, err = sdk.AccAddressFromBech32(contract)
	// if err != nil {
	// 	return nil
	// }

	// h.ibcHooksKeeper.StorePacketCallback(ctx, packet.GetSourceChannel(), packet.GetSequence(), contract)
	logFile.Close()

	return nil
}

// func SafeSendPacket(channel porttypes.ICS4Wrapper, ctx sdk.Context, chanCap *capabilitytypes.Capability, packet ibcexported.PacketI) (err error) {
// 	logger, logFile := testutil.CreateLogger()

// 	defer func() {
// 		if r := recover(); r != nil {
// 			err = fmt.Errorf("SendPacket panic: %v", r)
// 			logger.Println(err)
// 		}
// 	}()

// 	err = channel.SendPacket(ctx, chanCap, packet)
// 	logFile.Close()
// 	return err
// }

// jsonStringHasKey parses the memo as a json object and checks if it contains the key.
func jsonStringHasKey(memo, key string) (found bool, jsonObject map[string]interface{}) {
	jsonObject = make(map[string]interface{})

	// If there is no memo, the packet was either sent with an earlier version of IBC, or the memo was
	// intentionally left blank. Nothing to do here. Ignore the packet and pass it down the stack.
	if len(memo) == 0 {
		return false, jsonObject
	}

	// the jsonObject must be a valid JSON object
	err := json.Unmarshal([]byte(memo), &jsonObject)
	if err != nil {
		return false, jsonObject
	}

	// If the key doesn't exist, there's nothing to do on this hook. Continue by passing the packet
	// down the stack
	_, ok := jsonObject[key]
	if !ok {
		return false, jsonObject
	}

	return true, jsonObject
}
