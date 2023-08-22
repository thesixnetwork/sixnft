package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

// SetActionExecutor set a specific actionExecutor in the store from its index
func (k Keeper) SetActionExecutor(ctx sdk.Context, actionExecutor types.ActionExecutor) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionExecutorKeyPrefix))
	b := k.cdc.MustMarshal(&actionExecutor)
	store.Set(types.ActionExecutorKey(
		actionExecutor.NftSchemaCode,
		actionExecutor.ExecutorAddress,
	), b)
}

// GetActionExecutor returns a actionExecutor from its index
func (k Keeper) GetActionExecutor(
	ctx sdk.Context,
	nftSchemaCode string,
	executorAddress string,

) (val types.ActionExecutor, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionExecutorKeyPrefix))

	b := store.Get(types.ActionExecutorKey(
		nftSchemaCode,
		executorAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActionExecutor removes a actionExecutor from the store
func (k Keeper) RemoveActionExecutor(
	ctx sdk.Context,
	nftSchemaCode string,
	executorAddress string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionExecutorKeyPrefix))
	store.Delete(types.ActionExecutorKey(
		nftSchemaCode,
		executorAddress,
	))
}

// GetAllActionExecutor returns all actionExecutor
func (k Keeper) GetAllActionExecutor(ctx sdk.Context) (list []types.ActionExecutor) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionExecutorKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionExecutor
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
