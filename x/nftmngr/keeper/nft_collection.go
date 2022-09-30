package keeper

import (
	"encoding/binary"

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

// AppendNftCollection appends a nftCollection in the store with a new id and update the count
func (k Keeper) AppendNftCollection(
	ctx sdk.Context,
	nftCollection types.NftCollection,
) uint64 {
	// Create the nftCollection
	count := k.GetNftCollectionDataCount(ctx)
	// Set the ID of the appended value
	nftCollection.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftCollectionKeyPrefix))
	appendedValue := k.cdc.MustMarshal(&nftCollection)
	store.Set(types.NftCollectionKey(count), appendedValue)

	// Update nftCollection count
	k.SetNftCollectionDataCount(ctx, count+1)

	return count
}

// SetNftCollectionDataCount set the total number of nftCollection
func (k Keeper) SetNftCollectionDataCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftCollectionDataCountKey))
	byteKey := []byte(types.NftCollectionDataCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// GetNftCollectionDataCount get the total number of nftCollection
func (k Keeper) GetNftCollectionDataCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftCollectionDataCountKey))
	byteKey := []byte(types.NftCollectionDataCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
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
