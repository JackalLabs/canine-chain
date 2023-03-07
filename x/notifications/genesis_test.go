package notifications_test

import (
	"testing"

	"github.com/jackal-dao/canine/testutil/nullify"
	"github.com/jackal-dao/canine/x/notifications"
	"github.com/jackal-dao/canine/x/notifications/types"
	keepertest "github.com/jackalLabs/canine-chain/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		NotificationsList: []types.Notifications{
			{
				Count: 0,
			},
			{
				Count: 1,
			},
		},
		NotiCounterList: []types.NotiCounter{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NotificationsKeeper(t)
	notifications.InitGenesis(ctx, *k, genesisState)
	got := notifications.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.NotificationsList, got.NotificationsList)
	require.ElementsMatch(t, genesisState.NotiCounterList, got.NotiCounterList)
	// this line is used by starport scaffolding # genesis/test/assert
}
