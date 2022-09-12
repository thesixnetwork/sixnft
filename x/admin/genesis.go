package admin

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sixnft/x/admin/keeper"
	"sixnft/x/admin/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.Authorization != nil {
		k.SetAuthorization(ctx, *genState.Authorization)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all authorization
	authorization, found := k.GetAuthorization(ctx)
	if found {
		genesis.Authorization = &authorization
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
