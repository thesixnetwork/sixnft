package keeper

import (
	"context"

	"sixnft/x/evmsupport/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RemoveBinding(goCtx context.Context, msg *types.MsgRemoveBinding) (*types.MsgRemoveBindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valid, err := k.ValidateEVM(msg.SignedMessage, msg.Signature, msg.EthAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrVerifyingSignature, err.Error())
	}

	if !valid {
		return nil, sdkerrors.Wrap(types.ErrInvalidSignature, "Invalid Signature")
	}

	binding, found := k.Keeper.GetAddressBinding(ctx, msg.GetEthAddress())
	if !found {
		return nil, sdkerrors.Wrap(types.ErrAddressNotBound, msg.GetEthAddress())
	}

	k.Keeper.RemoveAddressBinding(ctx, msg.EthAddress)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.EventMessage, types.EventTypeDeleteBinding),
			sdk.NewAttribute(types.AttributeValueBinder, msg.GetCreator()),
			sdk.NewAttribute(types.AttributeValueEthAddress, msg.GetEthAddress()),
			sdk.NewAttribute(types.AttributeKeyDeleteBindingResult, "success"),
		),
	)

	return &types.MsgRemoveBindingResponse{
		EthAddress:    msg.EthAddress,
		NativeAddress: binding.NativeAddress,
	}, nil
}
