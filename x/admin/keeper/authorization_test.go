package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "sixnft/testutil/keeper"
	"sixnft/testutil/nullify"
	"sixnft/x/admin/keeper"
	"sixnft/x/admin/types"
)

func createTestAuthorization(keeper *keeper.Keeper, ctx sdk.Context) types.Authorization {
	item := types.Authorization{}
	keeper.SetAuthorization(ctx, item)
	return item
}

func TestAuthorizationGet(t *testing.T) {
	keeper, ctx := keepertest.AdminKeeper(t)
	item := createTestAuthorization(keeper, ctx)
	rst, found := keeper.GetAuthorization(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestAuthorizationRemove(t *testing.T) {
	keeper, ctx := keepertest.AdminKeeper(t)
	createTestAuthorization(keeper, ctx)
	keeper.RemoveAuthorization(ctx)
	_, found := keeper.GetAuthorization(ctx)
	require.False(t, found)
}
