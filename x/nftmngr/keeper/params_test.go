package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "sixnft/testutil/keeper"
	"sixnft/x/nftmngr/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NftmngrKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
