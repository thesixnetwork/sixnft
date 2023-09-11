package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

// SetNFTSchema set a specific nFTSchema in the store from its index
func (k Keeper) SetNFTSchema(ctx sdk.Context, nFTSchema types.NFTSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaKeyPrefix))
	b := k.cdc.MustMarshal(&nFTSchema)
	store.Set(types.NFTSchemaKey(
		nFTSchema.Code,
	), b)
}

// GetNFTSchema returns a nFTSchema from its index
func (k Keeper) GetNFTSchema(
	ctx sdk.Context,
	code string,

) (val types.NFTSchema, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaKeyPrefix))

	b := store.Get(types.NFTSchemaKey(
		code,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNFTSchema removes a nFTSchema from the store
func (k Keeper) RemoveNFTSchema(
	ctx sdk.Context,
	code string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaKeyPrefix))
	store.Delete(types.NFTSchemaKey(
		code,
	))
}

// GetAllNFTSchema returns all nFTSchema
func (k Keeper) GetAllNFTSchema(ctx sdk.Context) (list []types.NFTSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NFTSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllNFTSchema returns all nFTSchema
func (k Keeper) GetAllNFTSchemaV2(ctx sdk.Context) (list []types.NFTSchemaV2) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NFTSchemaV2
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// V1 of the NFT Schema

// SetNFTSchema set a specific nFTSchema in the store from its index
func (k Keeper) SetNFTSchemaV1(ctx sdk.Context, nFTSchema types.NFTSchemaV1) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaKeyPrefix))
	b := k.cdc.MustMarshal(&nFTSchema)
	store.Set(types.NFTSchemaKey(
		nFTSchema.Code,
	), b)
}

// GetNFTSchema returns a nFTSchema from its index
func (k Keeper) GetNFTSchemaV1(
	ctx sdk.Context,
	code string,

) (val types.NFTSchemaV1, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaKeyPrefix))

	b := store.Get(types.NFTSchemaKey(
		code,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNFTSchema removes a nFTSchema from the store
func (k Keeper) RemoveNFTSchemaV1(
	ctx sdk.Context,
	code string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaKeyPrefix))
	store.Delete(types.NFTSchemaKey(
		code,
	))
}

// GetAllNFTSchema returns all nFTSchema
func (k Keeper) GetAllNFTSchemaV1(ctx sdk.Context) (list []types.NFTSchemaV1) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NFTSchemaV1
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
