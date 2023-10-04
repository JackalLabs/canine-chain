package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"

	"github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/interchaintxs/types"
)

func (k *Keeper) HandleOnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), LabelLabelHandleChanOpenAck)

	logger, logFile := testutil.CreateLogger()

	/*
		Can we call the wasm keeper right here to interact with Jackal cosmwasm bindings?
		We could marshal a cosmwasm msg object into raw bytes and decode it from the Data field.
		Then we could call: func ICAOwnerFromPort(port string) (ICAOwner, error)
		to make sure that only the owner of the ICA is making the contract call.

		First we need a working relayer.
	*/
	logger.Printf("The data is: %s\n", packet.Data)

	logFile.Close()

	// TODO: return an actual acknowledgement
	return nil
}

// For now we're just going to log the data about a successfully created channel.
// Not sure why Neutron used the cosmwasm sudo entry point for returning an acknowledgement
// Perhaps osmosis' acknowledgement design can fit in here

// Notice that in the case of an ICA channel - it is not yet in OPEN state here
// the last step of channel opening(confirm) happens on the host chain.
func (k *Keeper) HandleChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID,
	counterpartyChannelID,
	counterpartyVersion string,
) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), LabelLabelHandleChanOpenAck)

	logger, logFile := testutil.CreateLogger()
	logger.Printf("HandleChanOpenAck, port_id: %s, channel_id: %s, counterparty_channel_id: %s, counterparty_version: %s\n", portID, channelID, counterpartyChannelID, counterpartyVersion)

	icaOwner, err := types.ICAOwnerFromPort(portID)
	logger.Printf("The ica owner is: %s\n", icaOwner)

	if err != nil {

		logger.Printf("HandleChanOpenAck: failed to get ica owner from source port. error: %s\n", err)
		return sdkerrors.Wrap(err, "failed to get ica owner from port")
	}

	logFile.Close()

	return nil
}
