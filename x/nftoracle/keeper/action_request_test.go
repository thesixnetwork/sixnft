package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "sixnft/testutil/keeper"
	"sixnft/testutil/nullify"
	"sixnft/x/nftoracle/keeper"
	"sixnft/x/nftoracle/types"
)

func createNActionRequest(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ActionRequest {
	items := make([]types.ActionRequest, n)
	for i := range items {
		items[i].Id = keeper.AppendActionRequest(ctx, items[i])
	}
	return items
}

func TestActionRequestGet(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionRequest(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetActionRequest(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestActionRequestRemove(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionRequest(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveActionRequest(ctx, item.Id)
		_, found := keeper.GetActionRequest(ctx, item.Id)
		require.False(t, found)
	}
}

func TestActionRequestGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionRequest(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllActionRequest(ctx)),
	)
}

func TestActionRequestCount(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionRequest(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetActionRequestCount(ctx))
}
