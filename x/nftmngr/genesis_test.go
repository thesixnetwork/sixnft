package nftmngr_test

import (
	"testing"

	keepertest "github.com/thesixnetwork/sixnft/testutil/keeper"

	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/sixnft/testutil/nullify"
	"github.com/thesixnetwork/sixnft/x/nftmngr"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

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
		NFTSchemaByContractList: []types.NFTSchemaByContract{
			{
				OriginContractAddress: "0",
			},
			{
				OriginContractAddress: "1",
			},
		},
		NftFeeConfig:  &types.NFTFeeConfig{},
		NFTFeeBalance: &types.NFTFeeBalance{},
		MetadataCreatorList: []types.MetadataCreator{
			{
				NftSchemaCode: "0",
			},
			{
				NftSchemaCode: "1",
			},
		},
		NftCollectionList: []types.NftCollection{
			{
				NftSchemaCode: "0",
			},
			{
				NftSchemaCode: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NftmngrKeeper(t)
	nftmngr.InitGenesis(ctx, *k, genesisState)
	got := nftmngr.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.NFTSchemaList, got.NFTSchemaList)
	require.ElementsMatch(t, genesisState.NftDataList, got.NftDataList)
	require.ElementsMatch(t, genesisState.ActionByRefIdList, got.ActionByRefIdList)
	require.ElementsMatch(t, genesisState.OrganizationList, got.OrganizationList)
	require.ElementsMatch(t, genesisState.NFTSchemaByContractList, got.NFTSchemaByContractList)
	require.Equal(t, genesisState.NftFeeConfig, got.NftFeeConfig)
	require.Equal(t, genesisState.NFTFeeBalance, got.NFTFeeBalance)
	require.ElementsMatch(t, genesisState.MetadataCreatorList, got.MetadataCreatorList)
	require.ElementsMatch(t, genesisState.NftCollectionList, got.NftCollectionList)
	// this line is used by starport scaffolding # genesis/test/assert
}
