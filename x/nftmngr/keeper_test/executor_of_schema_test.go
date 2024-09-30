package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/sixnft/testutil/keeper"
	"github.com/thesixnetwork/sixnft/testutil/nullify"
	"github.com/thesixnetwork/sixnft/x/nftmngr/keeper"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNExecutorOfSchema(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ExecutorOfSchema {
	items := make([]types.ExecutorOfSchema, n)
	for i := range items {
		items[i].NftSchemaCode = strconv.Itoa(i)

		keeper.SetExecutorOfSchema(ctx, items[i])
	}
	return items
}

func TestExecutorOfSchemaGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNExecutorOfSchema(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetExecutorOfSchema(ctx,
			item.NftSchemaCode,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestExecutorOfSchemaRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNExecutorOfSchema(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveExecutorOfSchema(ctx,
			item.NftSchemaCode,
		)
		_, found := keeper.GetExecutorOfSchema(ctx,
			item.NftSchemaCode,
		)
		require.False(t, found)
	}
}

func TestExecutorOfSchemaGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNExecutorOfSchema(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllExecutorOfSchema(ctx)),
	)
}
