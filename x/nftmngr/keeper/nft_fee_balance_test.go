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

func createTestNFTFeeBalance(keeper *keeper.Keeper, ctx sdk.Context) types.NFTFeeBalance {
	item := types.NFTFeeBalance{}
	keeper.SetNFTFeeBalance(ctx, item)
	return item
}

func TestNFTFeeBalanceGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	item := createTestNFTFeeBalance(keeper, ctx)
	rst, found := keeper.GetNFTFeeBalance(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestNFTFeeBalanceRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	createTestNFTFeeBalance(keeper, ctx)
	keeper.RemoveNFTFeeBalance(ctx)
	_, found := keeper.GetNFTFeeBalance(ctx)
	require.False(t, found)
}
