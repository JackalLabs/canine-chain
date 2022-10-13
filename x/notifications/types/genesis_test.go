package types_test

import (
	"testing"

	"github.com/jackal-dao/canine/x/notifications/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated notifications",
			genState: &types.GenesisState{
				NotificationsList: []types.Notifications{
					{
						Count: 0,
					},
					{
						Count: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated notiCounter",
			genState: &types.GenesisState{
				NotiCounterList: []types.NotiCounter{
					{
						Address: "0",
					},
					{
						Address: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
