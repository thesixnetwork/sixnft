package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
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
				ActionRequestList: []types.ActionOracleRequest{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ActionRequestCount: 2,
				CollectionOwnerRequestList: []types.CollectionOwnerRequest{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				CollectionOwnerRequestCount: 2,
				OracleConfig:                &types.OracleConfig{},
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
		{
			desc: "duplicated actionRequest",
			genState: &types.GenesisState{
				ActionRequestList: []types.ActionOracleRequest{
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
			desc: "invalid actionRequest count",
			genState: &types.GenesisState{
				ActionRequestList: []types.ActionOracleRequest{
					{
						Id: 1,
					},
				},
				ActionRequestCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated collectionOwnerRequest",
			genState: &types.GenesisState{
				CollectionOwnerRequestList: []types.CollectionOwnerRequest{
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
			desc: "invalid collectionOwnerRequest count",
			genState: &types.GenesisState{
				CollectionOwnerRequestList: []types.CollectionOwnerRequest{
					{
						Id: 1,
					},
				},
				CollectionOwnerRequestCount: 0,
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
