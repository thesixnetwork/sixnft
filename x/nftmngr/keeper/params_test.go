package keeper_test

import (
	"testing"

	testkeeper "github.com/thesixnetwork/sixnft/testutil/keeper"

	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NftmngrKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
