package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
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

				NFTSchemaList: []types.NFTSchema{
					{
						Code: "0",
					},
					{
						Code: "1",
					},
				},
				NftDataList: []types.NftData{
					{
						NftSchemaCode: "0",
						TokenId:       "0",
					},
					{
						NftSchemaCode: "1",
						TokenId:       "1",
					},
				},
				ActionByRefIdList: []types.ActionByRefId{
					{
						RefId: "0",
					},
					{
						RefId: "1",
					},
				},
				OrganizationList: []types.Organization{
					{
						Name: "0",
					},
					{
						Name: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated nFTSchema",
			genState: &types.GenesisState{
				NFTSchemaList: []types.NFTSchema{
					{
						Code: "0",
					},
					{
						Code: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated nftData",
			genState: &types.GenesisState{
				NftDataList: []types.NftData{
					{
						NftSchemaCode: "0",
						TokenId:       "0",
					},
					{
						NftSchemaCode: "0",
						TokenId:       "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated actionByRefId",
			genState: &types.GenesisState{
				ActionByRefIdList: []types.ActionByRefId{
					{
						RefId: "0",
					},
					{
						RefId: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated organization",
			genState: &types.GenesisState{
				OrganizationList: []types.Organization{
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
