package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"sixnft/x/nftoracle/types"
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

				MintRequestList: []types.MintRequest{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				MintRequestCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated mintRequest",
			genState: &types.GenesisState{
				MintRequestList: []types.MintRequest{
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
			desc: "invalid mintRequest count",
			genState: &types.GenesisState{
				MintRequestList: []types.MintRequest{
					{
						Id: 1,
					},
				},
				MintRequestCount: 0,
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
