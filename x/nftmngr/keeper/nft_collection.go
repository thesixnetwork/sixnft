package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

// SetNftCollection set a specific nftCollection in the store from its index
func (k Keeper) SetNftCollection(ctx sdk.Context, nftCollection types.NftCollection) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftCollectionKeyPrefix))
	b := k.cdc.MustMarshal(&nftCollection)
	store.Set(types.NftCollectionKey(
		nftCollection.NftSchemaCode,
	), b)
}

// GetNftCollection returns a nftCollection from its index
func (k Keeper) GetNftCollection(
	ctx sdk.Context,
	nftSchemaCode string,

) (val types.NftCollection, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftCollectionKeyPrefix))

	b := store.Get(types.NftCollectionKey(
		nftSchemaCode,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNftCollection removes a nftCollection from the store
func (k Keeper) RemoveNftCollection(
	ctx sdk.Context,
	nftSchemaCode string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftCollectionKeyPrefix))
	store.Delete(types.NftCollectionKey(
		nftSchemaCode,
	))
}

// GetAllNftCollection returns all nftCollection
func (k Keeper) GetAllNftCollection(ctx sdk.Context) (list []types.NftCollection) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftCollectionKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NftCollection
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
