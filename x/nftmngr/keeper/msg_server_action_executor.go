package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) CreateActionExecutor(goCtx context.Context, msg *types.MsgCreateActionExecutor) (*types.MsgCreateActionExecutorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieve the schema
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}
	// Check if the creator is the owner of the schema
	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	//validate that the actioner is a valid address
	// validate grantee format as 6x0000000000000000 or not
	_, err := sdk.AccAddressFromBech32(msg.ExecutorAddress)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}

	// Check if the value already exists
	_, isFound := k.GetActionExecutor(
		ctx,
		msg.NftSchemaCode,
		msg.ExecutorAddress,
	)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Action Executor already exists")
	}

	var actionExecutor = types.ActionExecutor{
		Creator:         msg.Creator,
		NftSchemaCode:   msg.NftSchemaCode,
		ExecutorAddress: msg.ExecutorAddress,
	}

	val, found := k.GetExecutorOfSchema(ctx, msg.NftSchemaCode)
	if !found {
		val = types.ExecutorOfSchema{
			NftSchemaCode: msg.NftSchemaCode,
			ExecutorAddress: []string{},
		}
	}

	// set executorOfSchema
	val.ExecutorAddress = append(val.ExecutorAddress, msg.ExecutorAddress)
	
	k.SetExecutorOfSchema(ctx, types.ExecutorOfSchema{
		NftSchemaCode: msg.NftSchemaCode,
		ExecutorAddress: val.ExecutorAddress,
	})

	k.SetActionExecutor(
		ctx,
		actionExecutor,
	)

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddActionExecutor,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyActionExecutor, msg.ExecutorAddress),
		),
	)

	return &types.MsgCreateActionExecutorResponse{
		NftSchemaCode:   msg.NftSchemaCode,
		ExecutorAddress: msg.ExecutorAddress,
	}, nil
}

func (k msgServer) UpdateActionExecutor(goCtx context.Context, msg *types.MsgUpdateActionExecutor) (*types.MsgUpdateActionExecutorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieve the schema
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}
	// Check if the creator is the owner of the schema
	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Check if the value exists
	_, isFound := k.GetActionExecutor(
		ctx,
		msg.NftSchemaCode,
		msg.ExecutorAddress,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// ** Validate by creator is schema owner instead
	// // Checks if the the msg creator is the same as the current owner
	// if msg.Creator != valFound.Creator {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }

	var actionExecutor = types.ActionExecutor{
		Creator:         msg.Creator,
		NftSchemaCode:   msg.NftSchemaCode,
		ExecutorAddress: msg.ExecutorAddress,
	}
	
	val, found := k.GetExecutorOfSchema(ctx, msg.NftSchemaCode)
	if !found {
		val = types.ExecutorOfSchema{
			NftSchemaCode: msg.NftSchemaCode,
			ExecutorAddress: []string{},
		}
	}

	// set executorOfSchema
	val.ExecutorAddress = append(val.ExecutorAddress, msg.ExecutorAddress)
	
	k.SetExecutorOfSchema(ctx, types.ExecutorOfSchema{
		NftSchemaCode: msg.NftSchemaCode,
		ExecutorAddress: val.ExecutorAddress,
	})

	k.SetActionExecutor(ctx, actionExecutor)

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateActionExecutor,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyActionExecutor, msg.ExecutorAddress),
		),
	)

	return &types.MsgUpdateActionExecutorResponse{}, nil
}

func (k msgServer) DeleteActionExecutor(goCtx context.Context, msg *types.MsgDeleteActionExecutor) (*types.MsgDeleteActionExecutorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}
	// Check if the creator is the owner of the schema
	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Check if the value exists
	_, isFound := k.GetActionExecutor(
		ctx,
		msg.NftSchemaCode,
		msg.ExecutorAddress,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	k.RemoveActionExecutor(
		ctx,
		msg.NftSchemaCode,
		msg.ExecutorAddress,
	)

	val, found := k.GetExecutorOfSchema(ctx, msg.NftSchemaCode)
	if !found {
		val = types.ExecutorOfSchema{
			NftSchemaCode: msg.NftSchemaCode,
			ExecutorAddress: []string{},
		}
	}

	// remove executorOfSchema
	for i, executor := range val.ExecutorAddress {
		if executor == msg.ExecutorAddress {
			val.ExecutorAddress = append(val.ExecutorAddress[:i], val.ExecutorAddress[i+1:]...)
			break
		}
	}
	
	k.SetExecutorOfSchema(ctx, types.ExecutorOfSchema{
		NftSchemaCode: msg.NftSchemaCode,
		ExecutorAddress: val.ExecutorAddress,
	})

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRemoveActionExecutor,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyActionExecutor, msg.ExecutorAddress),
		),
	)

	return &types.MsgDeleteActionExecutorResponse{}, nil
}
