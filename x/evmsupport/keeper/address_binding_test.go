package keeper_test

import (
	"strconv"
	"testing"

	keepertest "sixnft/testutil/keeper"
	"sixnft/testutil/nullify"
	"sixnft/x/evmsupport/keeper"
	"sixnft/x/evmsupport/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAddressBinding(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AddressBinding {
	items := make([]types.AddressBinding, n)
	for i := range items {
		items[i].EthAddress = strconv.Itoa(i)
		items[i].NativeAddress = strconv.Itoa(i)

		keeper.SetAddressBinding(ctx, items[i])
	}
	return items
}

func TestAddressBindingGet(t *testing.T) {
	keeper, ctx := keepertest.EvmsupportKeeper(t)
	items := createNAddressBinding(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAddressBinding(ctx,
			item.EthAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAddressBindingRemove(t *testing.T) {
	keeper, ctx := keepertest.EvmsupportKeeper(t)
	items := createNAddressBinding(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAddressBinding(ctx,
			item.EthAddress,
			item.NativeAddress,
		)
		_, found := keeper.GetAddressBinding(ctx,
			item.EthAddress,
		)
		require.False(t, found)
	}
}

func TestAddressBindingGetAll(t *testing.T) {
	keeper, ctx := keepertest.EvmsupportKeeper(t)
	items := createNAddressBinding(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAddressBinding(ctx)),
	)
}
