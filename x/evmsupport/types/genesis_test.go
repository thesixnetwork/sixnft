package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/sixnft/x/evmsupport/types"
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

				AddressBindingList: []types.AddressBinding{
					{
						EthAddress:    "0",
						NativeAddress: "0",
					},
					{
						EthAddress:    "1",
						NativeAddress: "1",
					},
				},
				ActionSignerList: []types.ActionSigner{
					{
						ActorAddress: "0",
					},
					{
						ActorAddress: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated addressBinding",
			genState: &types.GenesisState{
				AddressBindingList: []types.AddressBinding{
					{
						EthAddress:    "0",
						NativeAddress: "0",
					},
					{
						EthAddress:    "0",
						NativeAddress: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated actionSigner",
			genState: &types.GenesisState{
				ActionSignerList: []types.ActionSigner{
					{
						ActorAddress: "0",
					},
					{
						ActorAddress: "0",
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
