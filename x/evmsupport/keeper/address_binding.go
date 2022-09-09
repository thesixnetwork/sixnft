package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sixnft/x/evmsupport/types"
)

// SetAddressBinding set a specific addressBinding in the store from its index
func (k Keeper) SetAddressBinding(ctx sdk.Context, addressBinding types.AddressBinding) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressBindingKeyPrefix))
	b := k.cdc.MustMarshal(&addressBinding)
	store.Set(types.AddressBindingKey(
		addressBinding.EthAddress,
		addressBinding.NativeAddress,
	), b)
}

// GetAddressBinding returns a addressBinding from its index
func (k Keeper) GetAddressBinding(
	ctx sdk.Context,
	ethAddress string,
	nativeAddress string,

) (val types.AddressBinding, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressBindingKeyPrefix))

	b := store.Get(types.AddressBindingKey(
		ethAddress,
		nativeAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAddressBinding removes a addressBinding from the store
func (k Keeper) RemoveAddressBinding(
	ctx sdk.Context,
	ethAddress string,
	nativeAddress string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressBindingKeyPrefix))
	store.Delete(types.AddressBindingKey(
		ethAddress,
		nativeAddress,
	))
}

// GetAllAddressBinding returns all addressBinding
func (k Keeper) GetAllAddressBinding(ctx sdk.Context) (list []types.AddressBinding) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressBindingKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AddressBinding
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
