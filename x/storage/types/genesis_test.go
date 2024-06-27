package types_test

import (
	"testing"

	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
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
				FileList: []types.UnifiedFile{
					{
						Merkle: []byte("0"),
						Owner:  "0",
						Start:  0,
					},
					{
						Merkle: []byte("1"),
						Owner:  "1",
						Start:  0,
					},
				},
				ProvidersList: []types.Providers{
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
			desc: "duplicated activeDeals",
			genState: &types.GenesisState{
				FileList: []types.UnifiedFile{
					{
						Merkle: []byte("0"),
						Owner:  "0",
						Start:  0,
					},
					{
						Merkle: []byte("0"),
						Owner:  "0",
						Start:  0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated providers",
			genState: &types.GenesisState{
				ProvidersList: []types.Providers{
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
