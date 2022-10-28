package evmsupport

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/evmsupport/keeper"
	"github.com/thesixnetwork/sixnft/x/evmsupport/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the addressBinding
	for _, elem := range genState.AddressBindingList {
		k.SetAddressBinding(ctx, elem)
	}
	// Set all the actionSigner
	for _, elem := range genState.ActionSignerList {
		k.SetActionSigner(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AddressBindingList = k.GetAllAddressBinding(ctx)
	genesis.ActionSignerList = k.GetAllActionSigner(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
