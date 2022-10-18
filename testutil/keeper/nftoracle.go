package keeper

import (
	"testing"

	"github.com/thesixnetwork/sixnft/x/nftoracle/keeper"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"

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

	nftmngrkeeper "github.com/thesixnetwork/sixnft/x/nftmngr/keeper"
	nftmngrtypes "github.com/thesixnetwork/sixnft/x/nftmngr/types"

	adminKeeper "github.com/thesixnetwork/sixnft/x/nftadmin/keeper"
	admintypes "github.com/thesixnetwork/sixnft/x/nftadmin/types"

	evmsupportkeeper "github.com/thesixnetwork/sixnft/x/evmsupport/keeper"
	evmsupporttypes "github.com/thesixnetwork/sixnft/x/evmsupport/types"
)

func NftoracleKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
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
		"NftoracleParams",
	)
	nftmngrStoreKey := sdk.NewKVStoreKey(nftmngrtypes.StoreKey)
	nftmngrMemStoreKey := storetypes.NewMemoryStoreKey(nftmngrtypes.MemStoreKey)
	nftMngrParamsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		nftmngrStoreKey,
		nftmngrMemStoreKey,
		"NftmngrParams",
	)

	evmSupportStoreKey := sdk.NewKVStoreKey(evmsupporttypes.StoreKey)
	evmSupportMemStoreKey := storetypes.NewMemoryStoreKey(evmsupporttypes.MemStoreKey)

	adminsupportstoreKey := sdk.NewKVStoreKey(admintypes.StoreKey)
	adminsupportmemStoreKey := storetypes.NewMemoryStoreKey(admintypes.MemStoreKey)

	nftmngrKeeper := nftmngrkeeper.NewKeeper(
		cdc,
		nftmngrStoreKey,
		nftmngrMemStoreKey,
		nftMngrParamsSubspace,
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
		nil,
		nil,
		nil,
		nil,
	)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		nftmngrKeeper,
		adminKeeper.NewKeeper(
			cdc,
			adminsupportstoreKey,
			adminsupportmemStoreKey,
			typesparams.NewSubspace(cdc,
				types.Amino,
				adminsupportstoreKey,
				adminsupportmemStoreKey,
				"AdminParams",
			),
			nil,
		),
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	// k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
