package types_test

import (
	"testing"

	"github.com/jackal-dao/canine/x/dsig/types"
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
				UserUploadsList: []types.UserUploads{
					{
						Fid: "0",
					},
					{
						Fid: "1",
					},
				},
				FormList: []types.Form{
					{
						Ffid: "0",
					},
					{
						Ffid: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated userUploads",
			genState: &types.GenesisState{
				UserUploadsList: []types.UserUploads{
					{
						Fid: "0",
					},
					{
						Fid: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated form",
			genState: &types.GenesisState{
				FormList: []types.Form{
					{
						Ffid: "0",
					},
					{
						Ffid: "0",
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
