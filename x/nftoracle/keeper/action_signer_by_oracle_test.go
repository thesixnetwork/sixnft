package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/sixnft/testutil/keeper"
	"github.com/thesixnetwork/sixnft/testutil/nullify"
	"github.com/thesixnetwork/sixnft/x/nftoracle/keeper"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

func createNActionSignerByOracle(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ActionSignerByOracle {
	items := make([]types.ActionSignerByOracle, n)
	for i := range items {
		items[i].Id = keeper.AppendActionSignerByOracle(ctx, items[i])
	}
	return items
}

func TestActionSignerByOracleGet(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSignerByOracle(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetActionSignerByOracle(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestActionSignerByOracleRemove(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSignerByOracle(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveActionSignerByOracle(ctx, item.Id)
		_, found := keeper.GetActionSignerByOracle(ctx, item.Id)
		require.False(t, found)
	}
}

func TestActionSignerByOracleGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSignerByOracle(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllActionSignerByOracle(ctx)),
	)
}

func TestActionSignerByOracleCount(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSignerByOracle(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetActionSignerByOracleCount(ctx))
}
