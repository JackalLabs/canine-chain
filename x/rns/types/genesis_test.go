package types_test

import (
	"testing"

	"github.com/jackalLabs/canine-chain/v5/x/rns/types"
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
				WhoIsList: []types.Whois{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				NamesList: []types.Names{
					{
						Name: "0",
						Tld:  "o",
					},
					{
						Name: "1",
						Tld:  "o",
					},
				},
				BidsList: []types.Bids{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				ForSaleList: []types.Forsale{
					{
						Name: "0",
					},
					{
						Name: "1",
					},
				},
				InitList: []types.Init{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
			},
			valid: true,
		},
		{
			desc: "duplicated whois",
			genState: &types.GenesisState{
				WhoIsList: []types.Whois{
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
			desc: "duplicated names",
			genState: &types.GenesisState{
				NamesList: []types.Names{
					{
						Name: "0",
						Tld:  "o",
					},
					{
						Name: "0",
						Tld:  "o",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated bids",
			genState: &types.GenesisState{
				BidsList: []types.Bids{
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
			desc: "duplicated forsale",
			genState: &types.GenesisState{
				ForSaleList: []types.Forsale{
					{
						Name: "0",
					},
					{
						Name: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated init",
			genState: &types.GenesisState{
				InitList: []types.Init{
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
