package keeper_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	testutil "github.com/jackalLabs/canine-chain/v3/testutil"
	testkeeper "github.com/jackalLabs/canine-chain/v3/testutil/interchaintxs/keeper"
)

func TestHandleChanOpenAck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	icak, ctx := testkeeper.InterchainTxsKeeper(t, nil, nil, nil)
	testAddress, err := testutil.CreateTestAddresses("jkl", 1)
	require.NoError(t, err)

	portID := icatypes.PortPrefix + testAddress[0] + ".ica0"
	fmt.Println(icak)
	fmt.Println(ctx)
	fmt.Println(portID)
}
