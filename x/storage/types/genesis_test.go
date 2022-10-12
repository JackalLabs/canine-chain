package types_test

import (
	"testing"

	"github.com/jackal-dao/canine/x/storage/types"
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

				ContractsList: []types.Contracts{
					{
						Cid: "0",
					},
					{
						Cid: "1",
					},
				},
				ProofsList: []types.Proofs{
					{
						Cid: "0",
					},
					{
						Cid: "1",
					},
				},
				ActiveDealsList: []types.ActiveDeals{
					{
						Cid: "0",
					},
					{
						Cid: "1",
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
				PayBlocksList: []types.PayBlocks{
					{
						Blockid: "0",
					},
					{
						Blockid: "1",
					},
				},
				ClientUsageList: []types.ClientUsage{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				StraysList: []types.Strays{
					{
						Cid: "0",
					},
					{
						Cid: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated contracts",
			genState: &types.GenesisState{
				ContractsList: []types.Contracts{
					{
						Cid: "0",
					},
					{
						Cid: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated proofs",
			genState: &types.GenesisState{
				ProofsList: []types.Proofs{
					{
						Cid: "0",
					},
					{
						Cid: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated activeDeals",
			genState: &types.GenesisState{
				ActiveDealsList: []types.ActiveDeals{
					{
						Cid: "0",
					},
					{
						Cid: "0",
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
		{
			desc: "duplicated payBlocks",
			genState: &types.GenesisState{
				PayBlocksList: []types.PayBlocks{
					{
						Blockid: "0",
					},
					{
						Blockid: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated clientUsage",
			genState: &types.GenesisState{
				ClientUsageList: []types.ClientUsage{
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
			desc: "duplicated strays",
			genState: &types.GenesisState{
				StraysList: []types.Strays{
					{
						Cid: "0",
					},
					{
						Cid: "0",
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
