package keeper

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/thesixnetwork/sixnft/x/nftoracle/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	ActiveActionRequestQueuePrefix = []byte{0x02}
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
	actionRequest types.ActionOracleRequest,
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
func (k Keeper) SetActionRequest(ctx sdk.Context, actionRequest types.ActionOracleRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionRequestKey))
	b := k.cdc.MustMarshal(&actionRequest)
	store.Set(GetActionRequestIDBytes(actionRequest.Id), b)
}

// GetActionRequest returns a actionRequest from its id
func (k Keeper) GetActionRequest(ctx sdk.Context, id uint64) (val types.ActionOracleRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionRequestKey))
	b := store.Get(GetActionRequestIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetActionRequest returns a actionRequest from its id
func (k Keeper) GetActionRequestV063(ctx sdk.Context, id uint64) (val types.ActionRequestV063, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionRequestKey))
	b := store.Get(GetActionRequestIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// IsActionRequestOldVersion
func (k Keeper) IsActionRequestOldVersion(ctx sdk.Context, id uint64) (val types.ActionOracleRequest, isOldVersion bool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionRequestKey))
	b := store.Get(GetActionRequestIDBytes(id))
	if b == nil {
		return val, false, false
	}
	err := k.cdc.Unmarshal(b, &val)
	if err != nil {
		return val, true, false
	}
	return val, false, true
}

// RemoveActionRequest removes a actionRequest from the store
func (k Keeper) RemoveActionRequest(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionRequestKey))
	store.Delete(GetActionRequestIDBytes(id))
}

// GetAllActionRequest returns all actionRequest
func (k Keeper) GetAllActionRequest(ctx sdk.Context) (list []types.ActionOracleRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionRequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionOracleRequest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllActionRequestV603 returns all ActionRequestV063
func (k Keeper) GetAllActionRequestV603(ctx sdk.Context) (list []types.ActionRequestV063, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionRequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionRequestV063
		err := k.cdc.Unmarshal(iterator.Value(), &val)
		if err != nil {
			return nil, err
		}
		list = append(list, val)
	}

	return list, nil
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

func (k Keeper) InsertActiveActionRequestQueue(ctx sdk.Context, requestID uint64, endTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	bz := GetActionRequestIDBytes(requestID)
	store.Set(ActiveActionRequestQueueKey(requestID, endTime), bz)
}

func (k Keeper) RemoveFromActiveActionRequestQueue(ctx sdk.Context, requestID uint64, endTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(ActiveActionRequestQueueKey(requestID, endTime))
}

func (k Keeper) IterateActiveActionRequestsQueue(ctx sdk.Context, endTime time.Time, cb func(ActionOracleRequest types.ActionOracleRequest) (stop bool)) {
	iterator := k.ActiveActionRequestQueueIterator(ctx, endTime)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		requestID, _ := SplitActiveActionRequestQueueKey(iterator.Key())

		ActionOracleRequest, isOldVersion, found := k.IsActionRequestOldVersion(ctx, requestID)
		if !found {
			(fmt.Printf("ActionOracleRequest %d does not exist", requestID))
			continue
		}

		if isOldVersion {
			ActionOracleRequestV063, found := k.GetActionRequestV063(ctx, requestID)
			if !found {
				(fmt.Printf("ActionOracleRequest %d does not exist", requestID))
				continue
			}
			ActionOracleRequest = types.ActionOracleRequest{
				Id:                    ActionOracleRequest.Id,
				NftSchemaCode:         ActionOracleRequestV063.NftSchemaCode,
				TokenId:               ActionOracleRequestV063.TokenId,
				Action:                ActionOracleRequestV063.Action,
				Params:                make([]*types.ActionParameter, 0),
				Caller:                ActionOracleRequestV063.Caller,
				RefId:                 ActionOracleRequestV063.RefId,
				RequiredConfirm:       ActionOracleRequestV063.RequiredConfirm,
				Status:                ActionOracleRequestV063.Status,
				CurrentConfirm:        ActionOracleRequestV063.CurrentConfirm,
				Confirmers:            ActionOracleRequestV063.Confirmers,
				CreatedAt:             ActionOracleRequestV063.CreatedAt,
				ValidUntil:            ActionOracleRequestV063.ValidUntil,
				DataHashes:            ActionOracleRequestV063.DataHashes,
				ExpiredHeight:         ActionOracleRequestV063.ExpiredHeight,
				ExecutionErrorMessage: ActionOracleRequestV063.ExecutionErrorMessage,
			}
			/// migrate to new version
			// Migrate
			k.SetActionRequest(ctx, ActionOracleRequest)
		}

		if cb(ActionOracleRequest) {
			break
		}
	}
}

func (k Keeper) ActiveActionRequestQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(ActiveActionRequestQueuePrefix, sdk.PrefixEndBytes(ActiveActionRequestByTimeKey(endTime)))
}

func ActiveActionRequestQueueKey(requestID uint64, endTime time.Time) []byte {
	return append(ActiveActionRequestByTimeKey(endTime), GetActionRequestIDBytes(requestID)...)
}
func ActiveActionRequestByTimeKey(endTime time.Time) []byte {
	return append(ActiveActionRequestQueuePrefix, sdk.FormatTimeBytes(endTime)...)
}

func SplitActiveActionRequestQueueKey(key []byte) (requestID uint64, endTime time.Time) {
	return splitKeyWithTime(key)
}
