package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

// SetActionExecutor set a specific actionExecutor in the store from its index
func (k Keeper) SetActionExecutor(ctx sdk.Context, actionExecutor types.ActionExecutor) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionExecutorKeyPrefix))
	b := k.cdc.MustMarshal(&actionExecutor)
	store.Set(types.ActionExecutorKey(
		actionExecutor.NftSchemaCode,
		actionExecutor.ExecutorAddress,
	), b)
}

// GetActionExecutor returns a actionExecutor from its index
func (k Keeper) GetActionExecutor(
	ctx sdk.Context,
	nftSchemaCode string,
	executorAddress string,
) (val types.ActionExecutor, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionExecutorKeyPrefix))

	b := store.Get(types.ActionExecutorKey(
		nftSchemaCode,
		executorAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActionExecutor removes a actionExecutor from the store
func (k Keeper) RemoveActionExecutor(
	ctx sdk.Context,
	nftSchemaCode string,
	executorAddress string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionExecutorKeyPrefix))
	store.Delete(types.ActionExecutorKey(
		nftSchemaCode,
		executorAddress,
	))
}

// GetAllActionExecutor returns all actionExecutor
func (k Keeper) GetAllActionExecutor(ctx sdk.Context) (list []types.ActionExecutor) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionExecutorKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionExecutor
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) AddActionExecutor(ctx sdk.Context, signer sdk.AccAddress, nftSchemaName string, executorAddress sdk.AccAddress) error {
	creator := signer.String()
	newExecutorAddress := executorAddress.String()

	// Retrieve the schema
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return sdkerrors.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	// Check if the creator is the owner of the schema
	if creator != schema.Owner {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Check if the value already exists
	_, isFound := k.GetActionExecutor(
		ctx,
		nftSchemaName,
		newExecutorAddress,
	)

	if isFound {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Action Executor already exists")
	}

	actionExecutor := types.ActionExecutor{
		Creator:         creator,
		NftSchemaCode:   nftSchemaName,
		ExecutorAddress: newExecutorAddress,
	}

	val, found := k.GetExecutorOfSchema(ctx, nftSchemaName)
	if !found {
		val = types.ExecutorOfSchema{
			NftSchemaCode:   nftSchemaName,
			ExecutorAddress: []string{},
		}
	}

	// set executorOfSchema
	val.ExecutorAddress = append(val.ExecutorAddress, newExecutorAddress)

	k.SetExecutorOfSchema(ctx, types.ExecutorOfSchema{
		NftSchemaCode:   val.NftSchemaCode,
		ExecutorAddress: val.ExecutorAddress,
	})

	k.SetActionExecutor(ctx, actionExecutor)

	return nil
}

// RemoveActionExecutor removes a actionExecutor from the store
func (k Keeper) DelActionExecutor(ctx sdk.Context, signer sdk.AccAddress, nftSchemaName string, executorAddress sdk.AccAddress) error {
	creator := signer.String()
	toremoveExecutorAddress := executorAddress.String()

	// Retrieve the schema
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return sdkerrors.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	// Check if the creator is the owner of the schema
	if creator != schema.Owner {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Check if the value exists
	_, isFound := k.GetActionExecutor(
		ctx,
		nftSchemaName,
		toremoveExecutorAddress,
	)
	if !isFound {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	k.RemoveActionExecutor(
		ctx,
		nftSchemaName,
		toremoveExecutorAddress,
	)

	val, found := k.GetExecutorOfSchema(ctx, nftSchemaName)
	if !found {
		val = types.ExecutorOfSchema{
			NftSchemaCode:   nftSchemaName,
			ExecutorAddress: []string{},
		}
	}

	// remove executorOfSchema
	for i, executor := range val.ExecutorAddress {
		if executor == toremoveExecutorAddress {
			val.ExecutorAddress = append(val.ExecutorAddress[:i], val.ExecutorAddress[i+1:]...)
			break
		}
	}

	k.SetExecutorOfSchema(ctx, types.ExecutorOfSchema{
		NftSchemaCode:   val.NftSchemaCode,
		ExecutorAddress: val.ExecutorAddress,
	})

	return nil
}
