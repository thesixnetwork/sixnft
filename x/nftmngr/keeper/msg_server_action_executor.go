package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) CreateActionExecutor(goCtx context.Context, msg *types.MsgCreateActionExecutor) (*types.MsgCreateActionExecutorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetActionExecutor(
		ctx,
		msg.NftSchemaCode,
		msg.ExecutorAddress,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var actionExecutor = types.ActionExecutor{
		Creator:         msg.Creator,
		NftSchemaCode:   msg.NftSchemaCode,
		ExecutorAddress: msg.ExecutorAddress,
	}

	k.SetActionExecutor(
		ctx,
		actionExecutor,
	)
	return &types.MsgCreateActionExecutorResponse{}, nil
}

func (k msgServer) UpdateActionExecutor(goCtx context.Context, msg *types.MsgUpdateActionExecutor) (*types.MsgUpdateActionExecutorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetActionExecutor(
		ctx,
		msg.NftSchemaCode,
		msg.ExecutorAddress,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var actionExecutor = types.ActionExecutor{
		Creator:         msg.Creator,
		NftSchemaCode:   msg.NftSchemaCode,
		ExecutorAddress: msg.ExecutorAddress,
	}

	k.SetActionExecutor(ctx, actionExecutor)

	return &types.MsgUpdateActionExecutorResponse{}, nil
}

func (k msgServer) DeleteActionExecutor(goCtx context.Context, msg *types.MsgDeleteActionExecutor) (*types.MsgDeleteActionExecutorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetActionExecutor(
		ctx,
		msg.NftSchemaCode,
		msg.ExecutorAddress,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveActionExecutor(
		ctx,
		msg.NftSchemaCode,
		msg.ExecutorAddress,
	)

	return &types.MsgDeleteActionExecutorResponse{}, nil
}
