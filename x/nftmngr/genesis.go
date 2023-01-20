package nftmngr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftmngr/keeper"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the nFTSchema
	for _, elem := range genState.NFTSchemaList {
		k.SetNFTSchema(ctx, elem)
	}
	// Set all the nftData
	for _, elem := range genState.NftDataList {
		k.SetNftData(ctx, elem)
	}
	// Set all the actionByRefId
	for _, elem := range genState.ActionByRefIdList {
		k.SetActionByRefId(ctx, elem)
	}
	// Set all the organization
	for _, elem := range genState.OrganizationList {
		k.SetOrganization(ctx, elem)
	}
	// Set all the nFTSchemaByContract
	for _, elem := range genState.NFTSchemaByContractList {
		k.SetNFTSchemaByContract(ctx, elem)
	}
	// Set if defined
	if genState.NftFeeConfig != nil {
		err := k.ValidateFeeConfig(genState.NftFeeConfig)
		if err != nil {
			panic(err)
		}
		k.SetNFTFeeConfig(ctx, *genState.NftFeeConfig)
	}
	// Set if defined
	if genState.NFTFeeBalance != nil {
		k.SetNFTFeeBalance(ctx, *genState.NFTFeeBalance)
	}
	// Set all the metadataCreator
	for _, elem := range genState.MetadataCreatorList {
		k.SetMetadataCreator(ctx, elem)
	}
	// Set all the nftCollection
	for _, elem := range genState.NftCollectionList {
		k.SetNftCollection(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.NFTSchemaList = k.GetAllNFTSchema(ctx)
	genesis.NftDataList = k.GetAllNftData(ctx)
	genesis.ActionByRefIdList = k.GetAllActionByRefId(ctx)
	genesis.OrganizationList = k.GetAllOrganization(ctx)
	genesis.NFTSchemaByContractList = k.GetAllNFTSchemaByContract(ctx)
	// Get all nFTFeeConfig
	nFTFeeConfig, found := k.GetNFTFeeConfig(ctx)
	if found {
		genesis.NftFeeConfig = &nFTFeeConfig
	}
	// Get all nFTFeeBalance
	nFTFeeBalance, found := k.GetNFTFeeBalance(ctx)
	if found {
		genesis.NFTFeeBalance = &nFTFeeBalance
	}
	genesis.MetadataCreatorList = k.GetAllMetadataCreator(ctx)
	genesis.NftCollectionList = k.GetAllNftCollection(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
