package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "sixnft/testutil/keeper"
	"sixnft/x/evmsupport/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EvmsupportKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
