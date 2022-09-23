package keeper_test

import (
	"testing"

	testkeeper "github.com/thesixnetwork/sixnft/testutil/keeper"

	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/sixnft/x/evmsupport/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EvmsupportKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
