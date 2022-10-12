package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/sixnft/testutil/keeper"
	"github.com/thesixnetwork/sixnft/testutil/nullify"
	"github.com/thesixnetwork/sixnft/x/nftmngr/keeper"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func createTestNFTFeeConfig(keeper *keeper.Keeper, ctx sdk.Context) types.NFTFeeConfig {
	item := types.NFTFeeConfig{}
	keeper.SetNFTFeeConfig(ctx, item)
	return item
}

func TestNFTFeeConfigGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	item := createTestNFTFeeConfig(keeper, ctx)
	rst, found := keeper.GetNFTFeeConfig(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestNFTFeeConfigRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	createTestNFTFeeConfig(keeper, ctx)
	keeper.RemoveNFTFeeConfig(ctx)
	_, found := keeper.GetNFTFeeConfig(ctx)
	require.False(t, found)
}
