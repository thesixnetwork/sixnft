package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "sixnft/testutil/keeper"
	"sixnft/x/nftoracle/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NftoracleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
