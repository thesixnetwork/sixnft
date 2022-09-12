package nftoracle_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "sixnft/testutil/keeper"
	"sixnft/testutil/nullify"
	"sixnft/x/nftoracle"
	"sixnft/x/nftoracle/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MintRequestList: []types.MintRequest{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		MintRequestCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NftoracleKeeper(t)
	nftoracle.InitGenesis(ctx, *k, genesisState)
	got := nftoracle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MintRequestList, got.MintRequestList)
	require.Equal(t, genesisState.MintRequestCount, got.MintRequestCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
