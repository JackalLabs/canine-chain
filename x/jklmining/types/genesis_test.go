package types_test

import (
	"testing"

	"github.com/jackal-dao/canine/x/jklmining/types"
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

				SaveRequestsList: []types.SaveRequests{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				MinersList: []types.Miners{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				MinedList: []types.Mined{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				MinedCount: 2,
				MinerClaimsList: []types.MinerClaims{
					{
						Hash: "0",
					},
					{
						Hash: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated saveRequests",
			genState: &types.GenesisState{
				SaveRequestsList: []types.SaveRequests{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated miners",
			genState: &types.GenesisState{
				MinersList: []types.Miners{
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
		{
			desc: "duplicated mined",
			genState: &types.GenesisState{
				MinedList: []types.Mined{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid mined count",
			genState: &types.GenesisState{
				MinedList: []types.Mined{
					{
						Id: 1,
					},
				},
				MinedCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated minerClaims",
			genState: &types.GenesisState{
				MinerClaimsList: []types.MinerClaims{
					{
						Hash: "0",
					},
					{
						Hash: "0",
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
