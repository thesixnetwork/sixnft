package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sixnft/x/admin/types"
)

// SetAuthorization set authorization in the store
func (k Keeper) SetAuthorization(ctx sdk.Context, authorization types.Authorization) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorizationKey))
	b := k.cdc.MustMarshal(&authorization)
	store.Set([]byte{0}, b)
}

// GetAuthorization returns authorization
func (k Keeper) GetAuthorization(ctx sdk.Context) (val types.Authorization, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorizationKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAuthorization removes authorization from the store
func (k Keeper) RemoveAuthorization(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorizationKey))
	store.Delete([]byte{0})
}
