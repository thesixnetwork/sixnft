package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

// GetActionSignerByOracleCount get the total number of actionSignerByOracle
func (k Keeper) GetActionSignerByOracleCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ActionSignerByOracleCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetActionSignerByOracleCount set the total number of actionSignerByOracle
func (k Keeper) SetActionSignerByOracleCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ActionSignerByOracleCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendActionSignerByOracle appends a actionSignerByOracle in the store with a new id and update the count
func (k Keeper) AppendActionSignerByOracle(
	ctx sdk.Context,
	actionSignerByOracle types.ActionSignerByOracle,
) uint64 {
	// Create the actionSignerByOracle
	count := k.GetActionSignerByOracleCount(ctx)

	// Set the ID of the appended value
	actionSignerByOracle.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerByOracleKey))
	appendedValue := k.cdc.MustMarshal(&actionSignerByOracle)
	store.Set(GetActionSignerByOracleIDBytes(actionSignerByOracle.Id), appendedValue)

	// Update actionSignerByOracle count
	k.SetActionSignerByOracleCount(ctx, count+1)

	return count
}

// SetActionSignerByOracle set a specific actionSignerByOracle in the store
func (k Keeper) SetActionSignerByOracle(ctx sdk.Context, actionSignerByOracle types.ActionSignerByOracle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerByOracleKey))
	b := k.cdc.MustMarshal(&actionSignerByOracle)
	store.Set(GetActionSignerByOracleIDBytes(actionSignerByOracle.Id), b)
}

// GetActionSignerByOracle returns a actionSignerByOracle from its id
func (k Keeper) GetActionSignerByOracle(ctx sdk.Context, id uint64) (val types.ActionSignerByOracle, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerByOracleKey))
	b := store.Get(GetActionSignerByOracleIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActionSignerByOracle removes a actionSignerByOracle from the store
func (k Keeper) RemoveActionSignerByOracle(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerByOracleKey))
	store.Delete(GetActionSignerByOracleIDBytes(id))
}

// GetAllActionSignerByOracle returns all actionSignerByOracle
func (k Keeper) GetAllActionSignerByOracle(ctx sdk.Context) (list []types.ActionSignerByOracle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerByOracleKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionSignerByOracle
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetActionSignerByOracleIDBytes returns the byte representation of the ID
func GetActionSignerByOracleIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetActionSignerByOracleIDFromBytes returns ID in uint64 format from a byte array
func GetActionSignerByOracleIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
