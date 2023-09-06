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

func createNAttributeOfSchema(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AttributeOfSchema {
	items := make([]types.AttributeOfSchema, n)
	for i := range items {
		items[i].NftSchemaCode = strconv.Itoa(i)

		keeper.SetAttributeOfSchema(ctx, items[i])
	}
	return items
}

func TestAttributeOfSchemaGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNAttributeOfSchema(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAttributeOfSchema(ctx,
			item.NftSchemaCode,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAttributeOfSchemaRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNAttributeOfSchema(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAttributeOfSchema(ctx,
			item.NftSchemaCode,
		)
		_, found := keeper.GetAttributeOfSchema(ctx,
			item.NftSchemaCode,
		)
		require.False(t, found)
	}
}

func TestAttributeOfSchemaGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNAttributeOfSchema(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAttributeOfSchema(ctx)),
	)
}
