package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sixnft/x/nftoracle/types"
)

// GetMintRequestCount get the total number of mintRequest
func (k Keeper) GetMintRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MintRequestCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMintRequestCount set the total number of mintRequest
func (k Keeper) SetMintRequestCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MintRequestCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMintRequest appends a mintRequest in the store with a new id and update the count
func (k Keeper) AppendMintRequest(
	ctx sdk.Context,
	mintRequest types.MintRequest,
) uint64 {
	// Create the mintRequest
	count := k.GetMintRequestCount(ctx)

	// Set the ID of the appended value
	mintRequest.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintRequestKey))
	appendedValue := k.cdc.MustMarshal(&mintRequest)
	store.Set(GetMintRequestIDBytes(mintRequest.Id), appendedValue)

	// Update mintRequest count
	k.SetMintRequestCount(ctx, count+1)

	return count
}

// SetMintRequest set a specific mintRequest in the store
func (k Keeper) SetMintRequest(ctx sdk.Context, mintRequest types.MintRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintRequestKey))
	b := k.cdc.MustMarshal(&mintRequest)
	store.Set(GetMintRequestIDBytes(mintRequest.Id), b)
}

// GetMintRequest returns a mintRequest from its id
func (k Keeper) GetMintRequest(ctx sdk.Context, id uint64) (val types.MintRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintRequestKey))
	b := store.Get(GetMintRequestIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMintRequest removes a mintRequest from the store
func (k Keeper) RemoveMintRequest(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintRequestKey))
	store.Delete(GetMintRequestIDBytes(id))
}

// GetAllMintRequest returns all mintRequest
func (k Keeper) GetAllMintRequest(ctx sdk.Context) (list []types.MintRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintRequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MintRequest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMintRequestIDBytes returns the byte representation of the ID
func GetMintRequestIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMintRequestIDFromBytes returns ID in uint64 format from a byte array
func GetMintRequestIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
