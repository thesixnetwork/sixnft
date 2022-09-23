package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/thesixnetwork/sixnft/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/evmsupport/keeper"
	"github.com/thesixnetwork/sixnft/x/evmsupport/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EvmsupportKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
