package keeper

import (
	"testing"

	"sixnft/x/nftmngr/keeper"
	"sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	evmsupportkeeper "sixnft/x/evmsupport/keeper"
	evmsupporttypes "sixnft/x/evmsupport/types"
)

func NftmngrKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"NftmngrParams",
	)

	evmSupportStoreKey := sdk.NewKVStoreKey(evmsupporttypes.StoreKey)
	evmSupportMemStoreKey := storetypes.NewMemoryStoreKey(evmsupporttypes.MemStoreKey)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		evmsupportkeeper.NewKeeper(
			cdc,
			evmSupportStoreKey,
			evmSupportMemStoreKey,
			typesparams.NewSubspace(cdc,
				types.Amino,
				evmSupportStoreKey,
				evmSupportMemStoreKey,
				"EvmsupportParams",
			),
		),
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
