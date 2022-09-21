package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sixnft/x/nftoracle/types"
)

// GetActionRequestCount get the total number of actionRequest
func (k Keeper) GetActionRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ActionRequestCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetActionRequestCount set the total number of actionRequest
func (k Keeper) SetActionRequestCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ActionRequestCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendActionRequest appends a actionRequest in the store with a new id and update the count
func (k Keeper) AppendActionRequest(
	ctx sdk.Context,
	actionRequest types.ActionRequest,
) uint64 {
	// Create the actionRequest
	count := k.GetActionRequestCount(ctx)

	// Set the ID of the appended value
	actionRequest.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionRequestKey))
	appendedValue := k.cdc.MustMarshal(&actionRequest)
	store.Set(GetActionRequestIDBytes(actionRequest.Id), appendedValue)

	// Update actionRequest count
	k.SetActionRequestCount(ctx, count+1)

	return count
}

// SetActionRequest set a specific actionRequest in the store
func (k Keeper) SetActionRequest(ctx sdk.Context, actionRequest types.ActionRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionRequestKey))
	b := k.cdc.MustMarshal(&actionRequest)
	store.Set(GetActionRequestIDBytes(actionRequest.Id), b)
}

// GetActionRequest returns a actionRequest from its id
func (k Keeper) GetActionRequest(ctx sdk.Context, id uint64) (val types.ActionRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionRequestKey))
	b := store.Get(GetActionRequestIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActionRequest removes a actionRequest from the store
func (k Keeper) RemoveActionRequest(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionRequestKey))
	store.Delete(GetActionRequestIDBytes(id))
}

// GetAllActionRequest returns all actionRequest
func (k Keeper) GetAllActionRequest(ctx sdk.Context) (list []types.ActionRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionRequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionRequest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetActionRequestIDBytes returns the byte representation of the ID
func GetActionRequestIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetActionRequestIDFromBytes returns ID in uint64 format from a byte array
func GetActionRequestIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
