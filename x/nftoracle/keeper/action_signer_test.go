package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/sixnft/testutil/keeper"
	"github.com/thesixnetwork/sixnft/testutil/nullify"
	"github.com/thesixnetwork/sixnft/x/nftoracle/keeper"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNActionSigner(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ActionSigner {
	items := make([]types.ActionSigner, n)
	for i := range items {
		items[i].ActorAddress = strconv.Itoa(i)

		keeper.SetActionSigner(ctx, items[i])
	}
	return items
}

func TestActionSignerGet(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSigner(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetActionSigner(ctx,
			item.ActorAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestActionSignerRemove(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSigner(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveActionSigner(ctx,
			item.ActorAddress,
		)
		_, found := keeper.GetActionSigner(ctx,
			item.ActorAddress,
		)
		require.False(t, found)
	}
}

func TestActionSignerGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSigner(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllActionSigner(ctx)),
	)
}
