package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

// SetActionSigner set a specific actionSigner in the store from its index
func (k Keeper) SetActionSigner(ctx sdk.Context, actionSigner types.ActionSigner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerKeyPrefix))
	b := k.cdc.MustMarshal(&actionSigner)
	store.Set(types.ActionSignerKey(
		actionSigner.ActorAddress,
	), b)
}

// GetActionSigner returns a actionSigner from its index
func (k Keeper) GetActionSigner(
	ctx sdk.Context,
	actorAddress string,

) (val types.ActionSigner, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerKeyPrefix))

	b := store.Get(types.ActionSignerKey(
		actorAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActionSigner removes a actionSigner from the store
func (k Keeper) RemoveActionSigner(
	ctx sdk.Context,
	actorAddress string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerKeyPrefix))
	store.Delete(types.ActionSignerKey(
		actorAddress,
	))
}

// GetAllActionSigner returns all actionSigner
func (k Keeper) GetAllActionSigner(ctx sdk.Context) (list []types.ActionSigner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionSigner
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}


func GetActionSignerOwnerAddressBytes(owner_address string) []byte {
	bz := []byte(owner_address)
	return bz
}

// GetCollectionOwnerRequestIDFromBytes returns ID in uint64 format from a byte array
func GetActionSignerOwnerAddressFromBytes(bz []byte) string {
	return string(bz)
}


func (k Keeper) InsertActiveSignerQueue(ctx sdk.Context, owner_address string, endTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	bz := GetActionSignerOwnerAddressBytes(owner_address)
	store.Set(ActiveSignerQueueKey(owner_address, endTime), bz)
}

func (k Keeper) RemoveFromActiveSignerQueue(ctx sdk.Context, owner_address string, endTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(ActiveSignerQueueKey(owner_address, endTime))
}

func (k Keeper) IterateActiveSignersQueue(ctx sdk.Context, endTime time.Time, cb func(mintRequest types.ActionSigner) (stop bool)) {
	iterator := k.ActiveSignerQueueIterator(ctx, endTime)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		owner_address, _ := SplitActiveSignerQueueKey(iterator.Key())
		verifyRequest, found := k.GetActionSigner(ctx, owner_address)
		if !found {
			panic(fmt.Sprintf("verifyRequest %v does not exist", owner_address))
		}

		if cb(verifyRequest) {
			break
		}
	}
}

func (k Keeper) ActiveSignerQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(ActiveVerifyCollectionRequestQueuePrefix, sdk.PrefixEndBytes(ActiveSignerByTimeKey(endTime)))
}

func ActiveSignerQueueKey(owner_address string, endTime time.Time) []byte {
	return append(ActiveSignerByTimeKey(endTime), GetActionSignerOwnerAddressBytes(owner_address)...)
}
func ActiveSignerByTimeKey(endTime time.Time) []byte {
	return append(ActiveVerifyCollectionRequestQueuePrefix, sdk.FormatTimeBytes(endTime)...)
}

func SplitActiveSignerQueueKey(key []byte) (owner_address string, endTime time.Time) {
	return actionsplitKeyWithTime(key)
}

func actionsplitKeyWithTime(key []byte) (owner_address string, endTime time.Time) {
	kv.AssertKeyLength(key[1:], 8+lenTime)

	endTime, err := sdk.ParseTimeBytes(key[1 : 1+lenTime])
	if err != nil {
		panic(err)
	}

	owner_address = GetActionSignerOwnerAddressFromBytes(key[1+lenTime:])

	return
}
