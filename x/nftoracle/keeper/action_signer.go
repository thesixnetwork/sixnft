package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

var (
	ActiveActionSignerQueuePrefix = []byte{0x04}
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

// GetActionSignerOwnerAddressBytes
func (k Keeper) GetActionSignerOwnerAddressBytes(ctx sdk.Context,actorAddress string) []byte {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ActionSignerKey(
		actorAddress,
	))
	return bz
}

// GetActionSignerOwnerAddressFromBytes returns address in string format from a byte array
func GetActionSignerOwnerAddressFromBytes(bz []byte) string {
	return string(bz)
}


func (k Keeper) InsertActiveSignerQueue(ctx sdk.Context, actorAddress string, endTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ActionSignerKey(
		actorAddress,
	))
	store.Set(ActiveSignerQueueKey(actorAddress, endTime), bz)
}

func (k Keeper) RemoveFromActiveSignerQueue(ctx sdk.Context, actorAddress string, endTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(ActiveSignerQueueKey(actorAddress, endTime))
}

func (k Keeper) IterateActiveSignersQueue(ctx sdk.Context, endTime time.Time, cb func(actionSigner types.ActionSigner) (stop bool)) {
	iterator := k.ActiveSignerQueueIterator(ctx, endTime)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		actorAddress, _ := SplitActiveSignerQueueKey(iterator.Key())
		singer, found := k.GetActionSigner(ctx, actorAddress)
		if !found {
			panic(fmt.Sprintf("singer %v does not exist", actorAddress))
		}

		if cb(singer) {
			break
		}
	}
}

func (k Keeper) ActiveSignerQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(ActiveActionSignerQueuePrefix, sdk.PrefixEndBytes(ActiveSignerByTimeKey(endTime)))
}

func ActiveSignerQueueKey(actorAddress string, endTime time.Time) []byte {
	return append(ActiveSignerByTimeKey(endTime), types.ActionSignerKey(
		actorAddress,
	)...)
}
func ActiveSignerByTimeKey(endTime time.Time) []byte {
	return append(ActiveActionSignerQueuePrefix, sdk.FormatTimeBytes(endTime)...)
}

func SplitActiveSignerQueueKey(key []byte) (actorAddess string, endTime time.Time) {
	return actionsplitKeyWithTime(key)
}

func actionsplitKeyWithTime(key []byte) (actorAddess string, endTime time.Time) {
	kv.AssertKeyLength(key[1:], 8+lenTime)

	endTime, err := sdk.ParseTimeBytes(key[1 : 1+lenTime])
	if err != nil {
		panic(err)
	}

	actorAddess = GetActionSignerOwnerAddressFromBytes(key[1+lenTime:])

	return
}
