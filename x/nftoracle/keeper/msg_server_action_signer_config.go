package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"

	"github.com/ethereum/go-ethereum/common"
)

func (k msgServer) CreateActionSignerConfig(goCtx context.Context, msg *types.MsgCreateActionSignerConfig) (*types.MsgCreateActionSignerConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate if creator has permission to set fee config
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	granted := k.nftadminKeeper.HasPermission(ctx, types.KeyPermissionAdminSignerConfig, creator)
	if !granted {
		return nil, sdkerrors.Wrap(types.ErrNoPermissionAdminSignerConfig, msg.Creator)
	}

	// Check if the value already exists
	_, isFound := k.GetActionSignerConfig(
		ctx,
		msg.Chain,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// validate ContractOwner address
	if !common.IsHexAddress(msg.ContractOwner) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid contract owner address")
	}

	// validate ContractAddress address
	if !common.IsHexAddress(msg.ContractAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid contract address")
	}

	var actionSignerConfig = types.ActionSignerConfig{
		Creator:         msg.Creator,
		Chain:           msg.Chain,
		RpcEndpoint:     msg.RpcEndpoint,
		ContractAddress: msg.ContractAddress,
		ContractName:    msg.ContractName,
		ContractOwner:   msg.ContractOwner,
	}

	k.SetActionSignerConfig(
		ctx,
		actionSignerConfig,
	)
	return &types.MsgCreateActionSignerConfigResponse{}, nil
}

func (k msgServer) UpdateActionSignerConfig(goCtx context.Context, msg *types.MsgUpdateActionSignerConfig) (*types.MsgUpdateActionSignerConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetActionSignerConfig(
		ctx,
		msg.Chain,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var actionSignerConfig = types.ActionSignerConfig{
		Creator:         msg.Creator,
		Chain:           msg.Chain,
		ContractAddress: msg.ContractAddress,
		ContractName:    msg.ContractName,
		ContractOwner:   msg.ContractOwner,
	}

	k.SetActionSignerConfig(ctx, actionSignerConfig)

	return &types.MsgUpdateActionSignerConfigResponse{}, nil
}

func (k msgServer) DeleteActionSignerConfig(goCtx context.Context, msg *types.MsgDeleteActionSignerConfig) (*types.MsgDeleteActionSignerConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetActionSignerConfig(
		ctx,
		msg.Chain,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveActionSignerConfig(
		ctx,
		msg.Chain,
	)

	return &types.MsgDeleteActionSignerConfigResponse{}, nil
}