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
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"
	filetreekeeper "github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
	filetreetypes "github.com/jackalLabs/canine-chain/v3/x/filetree/types"

	"github.com/jackalLabs/canine-chain/v3/gmpmiddleware/types"
	"github.com/jackalLabs/canine-chain/v3/gmpmiddleware/utils"

	"github.com/jackalLabs/canine-chain/v3/testutil"
)

type IBCMiddleware struct {
	channel        porttypes.ICS4Wrapper
	app            porttypes.IBCModule
	handler        types.GeneralMessageHandler
	filetreekeeper filetreekeeper.Keeper
}

func NewIBCMiddleware(channel porttypes.ICS4Wrapper, app porttypes.IBCModule, handler types.GeneralMessageHandler, filetreekeeper filetreekeeper.Keeper) IBCMiddleware {
	return IBCMiddleware{
		channel:        channel,
		app:            app,
		handler:        handler,
		filetreekeeper: filetreekeeper,
	}
}

// GetChannel returns the 'channel' field of the IBCMiddleware struct.
func (m IBCMiddleware) GetChannel() porttypes.ICS4Wrapper {
	return m.channel
}

// GetApp returns the 'app' field of the IBCMiddleware struct.
func (m IBCMiddleware) GetApp() porttypes.IBCModule {
	return m.app
}

// GetHandler returns the 'handler' field of the IBCMiddleware struct.
func (m IBCMiddleware) GetHandler() types.GeneralMessageHandler {
	return m.handler
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
	// im.ProperlyConfigured()? // consider adding this function if there are keepers involved
	logger, logFile := testutil.CreateLogger()

	ack := im.app.OnRecvPacket(ctx, packet, relayer)

	isIcs20, data := isIcs20Packet(packet)
	if !isIcs20 {
		// make sure this properly calls the underlying IBCModule
		return im.app.OnRecvPacket(ctx, packet, relayer)
	}

	// Validate the memo
	isWasmRouted, contractAddr, msgBytes, err := ValidateAndParseMemo(data.GetMemo(), data.Receiver)
	if !isWasmRouted {
		return im.app.OnRecvPacket(ctx, packet, relayer)
	}
	if err != nil {
		return utils.NewEmitErrorAcknowledgement(ctx, types.ErrMsgValidation, err.Error())
	}
	fmt.Println(contractAddr)
	fmt.Println(msgBytes)

	// The below will always make our tests fail unless we hard code that the sender is the AxelarGMPAcc
	// // authenticate the message with packet sender + channel-id
	// // TODO: authenticate the message with channel-id
	// if data.Sender != AxelarGMPAcc {
	// 	return ack
	// }

	var msg types.Message

	// var err error

	// if err = json.Unmarshal([]byte(data.GetMemo()), &msg); err != nil {
	// 	return channeltypes.NewErrorAcknowledgement(fmt.Errorf("cannot unmarshal memo"))
	// }

	logger.Println(isIcs20)
	logger.Printf("The message is %v", msg)
	logger.Printf("The ack is %v", ack.Success())

	filetreeMsgServer := filetreekeeper.NewMsgServerImpl(im.filetreekeeper)

	msgMakeRoot := filetreetypes.NewMsgMakeRootV2(
		msg.SourceAddress,
		msg.SourceChain,
		string(msg.Payload),
		"placeholder",
	)

	if err := msgMakeRoot.ValidateBasic(); err != nil {
		return ack
	}

	// Post files
	_, error := filetreeMsgServer.MakeRootV2(
		sdk.WrapSDKContext(ctx),
		msgMakeRoot,
	)
	if error != nil {
		return ack
	}

	logFile.Close()

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

	isCallbackRouted, metadata := jsonStringHasKey(data.GetMemo(), types.IBCCallbackKey) // metadata was here
	if !isCallbackRouted {
		return im.channel.SendPacket(ctx, chanCap, packet) // continue
	}

	// We remove the callback metadata from the memo as it has already been processed.

	// If the only available key in the memo is the callback, we should remove the memo
	// from the data completely so the packet is sent without it.
	// This way receiver chains that are on old versions of IBC will be able to process the packet

	callbackRaw := metadata[types.IBCCallbackKey]
	delete(metadata, types.IBCCallbackKey)

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

	// What's the consequence of not using our own hooks for ICS4?
	err = im.channel.SendPacket(ctx, chanCap, packetWithoutCallbackMemo)
	if err != nil {
		return err
	}

	// Make sure the callback contract is a string and a valid bech32 addr. If it isn't, ignore this packet
	contract, ok := callbackRaw.(string)
	if !ok {
		return nil
	}
	_, err = sdk.AccAddressFromBech32(contract)
	if err != nil {
		return nil
	}

	// Do we need 'StorePacketCallback' to store
	// which contract will be listening for the ack or timeout of a packet

	return nil
}

// do we really want to store the call back?...
/*
// StorePacketCallback stores which contract will be listening for the ack or timeout of a packet
func (k Keeper) StorePacketCallback(ctx sdk.Context, channel string, packetSequence uint64, contract string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(GetPacketKey(channel, packetSequence), []byte(contract))
}
*/
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

func ValidateAndParseMemo(memo string, receiver string) (isWasmRouted bool, contractAddr sdk.AccAddress, msgBytes []byte, err error) {
	isWasmRouted, metadata := jsonStringHasKey(memo, "wasm")
	if !isWasmRouted {
		return isWasmRouted, sdk.AccAddress{}, nil, nil
	}

	wasmRaw := metadata["wasm"]

	// Make sure the wasm key is a map. If it isn't, ignore this packet
	wasm, ok := wasmRaw.(map[string]interface{})
	if !ok {
		return isWasmRouted, sdk.AccAddress{}, nil,
			fmt.Errorf(types.ErrBadMetadataFormatMsg, memo, "wasm metadata is not a valid JSON map object")
	}

	// Get the contract
	contract, ok := wasm["contract"].(string)
	if !ok {
		// The tokens will be returned
		return isWasmRouted, sdk.AccAddress{}, nil,
			fmt.Errorf(types.ErrBadMetadataFormatMsg, memo, `Could not find key wasm["contract"]`)
	}

	contractAddr, err = sdk.AccAddressFromBech32(contract)
	if err != nil {
		return isWasmRouted, sdk.AccAddress{}, nil,
			fmt.Errorf(types.ErrBadMetadataFormatMsg, memo, `wasm["contract"] is not a valid bech32 address`)
	}

	// The contract and the receiver should be the same for the packet to be valid
	if contract != receiver {
		return isWasmRouted, sdk.AccAddress{}, nil,
			fmt.Errorf(types.ErrBadMetadataFormatMsg, memo, `wasm["contract"] should be the same as the receiver of the packet`)
	}

	// Ensure the message key is provided
	if wasm["msg"] == nil {
		return isWasmRouted, sdk.AccAddress{}, nil,
			fmt.Errorf(types.ErrBadMetadataFormatMsg, memo, `Could not find key wasm["msg"]`)
	}

	// Make sure the msg key is a map. If it isn't, return an error
	_, ok = wasm["msg"].(map[string]interface{})
	if !ok {
		return isWasmRouted, sdk.AccAddress{}, nil,
			fmt.Errorf(types.ErrBadMetadataFormatMsg, memo, `wasm["msg"] is not a map object`)
	}

	// Get the message string by serializing the map
	msgBytes, err = json.Marshal(wasm["msg"])
	if err != nil {
		// The tokens will be returned
		return isWasmRouted, sdk.AccAddress{}, nil,
			fmt.Errorf(types.ErrBadMetadataFormatMsg, memo, err.Error())
	}

	return isWasmRouted, contractAddr, msgBytes, nil
}
