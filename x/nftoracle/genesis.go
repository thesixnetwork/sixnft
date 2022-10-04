package nftoracle

import (
	"fmt"

	"github.com/thesixnetwork/sixnft/x/nftoracle/keeper"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the mintRequest
	for _, elem := range genState.MintRequestList {
		k.SetMintRequest(ctx, elem)
	}

	fmt.Println("############################### mint active duration", genState.Params.MintRequestActiveDuration)

	// Set mintRequest count
	k.SetMintRequestCount(ctx, genState.MintRequestCount)
	// Set all the actionRequest
	for _, elem := range genState.ActionRequestList {
		k.SetActionRequest(ctx, elem)
	}

	// Set actionRequest count
	k.SetActionRequestCount(ctx, genState.ActionRequestCount)
	// Set all the collectionOwnerRequest
	for _, elem := range genState.CollectionOwnerRequestList {
		k.SetCollectionOwnerRequest(ctx, elem)
	}

	// Set collectionOwnerRequest count
	k.SetCollectionOwnerRequestCount(ctx, genState.CollectionOwnerRequestCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.MintRequestList = k.GetAllMintRequest(ctx)
	genesis.MintRequestCount = k.GetMintRequestCount(ctx)
	genesis.ActionRequestList = k.GetAllActionRequest(ctx)
	genesis.ActionRequestCount = k.GetActionRequestCount(ctx)
	genesis.CollectionOwnerRequestList = k.GetAllCollectionOwnerRequest(ctx)
	genesis.CollectionOwnerRequestCount = k.GetCollectionOwnerRequestCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
