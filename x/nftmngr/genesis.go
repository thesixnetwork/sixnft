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
	genesis.NftCollectionList = k.GetAllNftCollection(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
