package types_test

import (
	"testing"

	"github.com/jackalLabs/canine-chain/x/oracle/types"
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
			desc:     "valid genesis state",
			genState: &types.GenesisState{},
			valid:    true,
		},
		{
			desc:     "duplicated whois",
			genState: &types.GenesisState{},
			valid:    false,
		},
		{
			desc:     "duplicated names",
			genState: &types.GenesisState{},
			valid:    false,
		},
		{
			desc:     "duplicated bids",
			genState: &types.GenesisState{},
			valid:    false,
		},
		{
			desc:     "duplicated forsale",
			genState: &types.GenesisState{},
			valid:    false,
		},
		{
			desc:     "duplicated init",
			genState: &types.GenesisState{},
			valid:    false,
		},
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
