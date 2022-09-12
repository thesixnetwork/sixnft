package nftoracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sixnft/x/nftoracle/keeper"
	"sixnft/x/nftoracle/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the mintRequest
	for _, elem := range genState.MintRequestList {
		k.SetMintRequest(ctx, elem)
	}

	// Set mintRequest count
	k.SetMintRequestCount(ctx, genState.MintRequestCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.MintRequestList = k.GetAllMintRequest(ctx)
	genesis.MintRequestCount = k.GetMintRequestCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
