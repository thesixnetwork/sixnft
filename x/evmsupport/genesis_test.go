package evmsupport_test

import (
	"testing"

	keepertest "github.com/thesixnetwork/sixnft/testutil/keeper"

	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/sixnft/testutil/nullify"
	"github.com/thesixnetwork/sixnft/x/evmsupport"
	"github.com/thesixnetwork/sixnft/x/evmsupport/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EvmsupportKeeper(t)
	evmsupport.InitGenesis(ctx, *k, genesisState)
	got := evmsupport.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AddressBindingList, got.AddressBindingList)
	// this line is used by starport scaffolding # genesis/test/assert
}
