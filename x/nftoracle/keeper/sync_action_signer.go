package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

// GetSyncActionSignerCount get the total number of syncActionSigner
func (k Keeper) GetSyncActionSignerCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SyncActionSignerCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSyncActionSignerCount set the total number of syncActionSigner
func (k Keeper) SetSyncActionSignerCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SyncActionSignerCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSyncActionSigner appends a syncActionSigner in the store with a new id and update the count
func (k Keeper) AppendSyncActionSigner(
	ctx sdk.Context,
	syncActionSigner types.SyncActionSigner,
) uint64 {
	// Create the syncActionSigner
	count := k.GetSyncActionSignerCount(ctx)

	// Set the ID of the appended value
	syncActionSigner.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SyncActionSignerKey))
	appendedValue := k.cdc.MustMarshal(&syncActionSigner)
	store.Set(GetSyncActionSignerIDBytes(syncActionSigner.Id), appendedValue)

	// Update syncActionSigner count
	k.SetSyncActionSignerCount(ctx, count+1)

	return count
}

// SetSyncActionSigner set a specific syncActionSigner in the store
func (k Keeper) SetSyncActionSigner(ctx sdk.Context, syncActionSigner types.SyncActionSigner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SyncActionSignerKey))
	b := k.cdc.MustMarshal(&syncActionSigner)
	store.Set(GetSyncActionSignerIDBytes(syncActionSigner.Id), b)
}

// GetSyncActionSigner returns a syncActionSigner from its id
func (k Keeper) GetSyncActionSigner(ctx sdk.Context, id uint64) (val types.SyncActionSigner, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SyncActionSignerKey))
	b := store.Get(GetSyncActionSignerIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSyncActionSigner removes a syncActionSigner from the store
func (k Keeper) RemoveSyncActionSigner(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SyncActionSignerKey))
	store.Delete(GetSyncActionSignerIDBytes(id))
}

// GetAllSyncActionSigner returns all syncActionSigner
func (k Keeper) GetAllSyncActionSigner(ctx sdk.Context) (list []types.SyncActionSigner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SyncActionSignerKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SyncActionSigner
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSyncActionSignerIDBytes returns the byte representation of the ID
func GetSyncActionSignerIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSyncActionSignerIDFromBytes returns ID in uint64 format from a byte array
func GetSyncActionSignerIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
