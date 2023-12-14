package keeper_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	testutil "github.com/jackalLabs/canine-chain/v3/testutil"
	testkeeper "github.com/jackalLabs/canine-chain/v3/testutil/interchaintxs/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"

	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
)

func TestHandleChanOpenAck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	icak, ctx := testkeeper.InterchainTxsKeeper(t, nil, nil, nil)
	testAddress, err := testutil.CreateTestAddresses("cosmos", 1)
	require.NoError(t, err)

	portID := icatypes.PortPrefix + testAddress[0] + ".ica0"
	contractAddress := sdk.MustAccAddressFromBech32(testAddress[0])
	channelID := "channel-0"
	counterpartyChannelID := "channel-1"

	error := icak.HandleChanOpenAck(ctx, "", channelID, counterpartyChannelID, "1")
	require.ErrorContains(t, error, "failed to get ica owner from port")

	error = icak.HandleChanOpenAck(ctx, portID, channelID, counterpartyChannelID, "1")
	require.NoError(t, err)

	fmt.Println(contractAddress)
	// NOTE: could we call the bindings contract directly and bypass the wasmbindings interface?
}

func TestHandleOnRecvPacket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	icak, ctx := testkeeper.InterchainTxsKeeper(t, nil, nil, nil)
	testAddress, err := testutil.CreateTestAddresses("cosmos", 3)
	require.NoError(t, err)

	sender := testAddress[0]
	receiver := testAddress[1]
	relayer := testAddress[2]

	packet := makeMockPacket(receiver, "placeholder memo", 0, sender)

	ack := icak.HandleOnRecvPacket(ctx, packet, sdk.AccAddress(relayer))
	fmt.Println(ack)
}

// NOTE: Always make sure this resembles osmosis' mock packet
func makeMockPacket(receiver, memo string, prevSequence uint64, sender string) channeltypes.Packet {
	packetData := transfertypes.FungibleTokenPacketData{
		Denom:    sdk.DefaultBondDenom,
		Amount:   "1",
		Sender:   sender,
		Receiver: receiver,
		Memo:     memo, // attempted removing memo but packet still won't send. Nil pointer de-reference error remains the same.
	}

	return channeltypes.NewPacket(
		packetData.GetBytes(),
		prevSequence+1,
		icatypes.PortPrefix+sender+".ica0",
		"channel-0",
		icatypes.PortID+sender+".ica1",
		"channel-1",
		clienttypes.NewHeight(0, 100),
		0,
	)
}
