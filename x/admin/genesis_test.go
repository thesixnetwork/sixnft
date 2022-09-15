package admin_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "sixnft/testutil/keeper"
	"sixnft/testutil/nullify"
	"sixnft/x/admin"
	"sixnft/x/admin/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Authorization: &types.Authorization{
			RootAdmin:   "12",
			Permissions: new(types.Permissions),
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AdminKeeper(t)
	admin.InitGenesis(ctx, *k, genesisState)
	got := admin.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Authorization, got.Authorization)
	// this line is used by starport scaffolding # genesis/test/assert
}
