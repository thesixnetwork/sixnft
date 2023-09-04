package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

// SetSchemaAttribute set a specific schemaAttribute in the store from its index
func (k Keeper) SetSchemaAttribute(ctx sdk.Context, schemaAttribute types.SchemaAttribute) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SchemaAttributeKeyPrefix))
	b := k.cdc.MustMarshal(&schemaAttribute)
	store.Set(types.SchemaAttributeKey(
		schemaAttribute.NftSchemaCode,
		schemaAttribute.Name,
	), b)
}

// GetSchemaAttribute returns a schemaAttribute from its index
func (k Keeper) GetSchemaAttribute(
	ctx sdk.Context,
	nftSchemaCode string,
	name string,

) (val types.SchemaAttribute, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SchemaAttributeKeyPrefix))

	b := store.Get(types.SchemaAttributeKey(
		nftSchemaCode,
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSchemaAttribute removes a schemaAttribute from the store
func (k Keeper) RemoveSchemaAttribute(
	ctx sdk.Context,
	nftSchemaCode string,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SchemaAttributeKeyPrefix))
	store.Delete(types.SchemaAttributeKey(
		nftSchemaCode,
		name,
	))
}

// GetAllSchemaAttribute returns all schemaAttribute
func (k Keeper) GetAllSchemaAttribute(ctx sdk.Context) (list []types.SchemaAttribute) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SchemaAttributeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SchemaAttribute
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
