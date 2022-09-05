package nftmngr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sixnft/x/nftmngr/keeper"
	"sixnft/x/nftmngr/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the nFTSchema
	for _, elem := range genState.NFTSchemaList {
		k.SetNFTSchema(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.NFTSchemaList = k.GetAllNFTSchema(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
