package keeper

import (
	"context"

	"sixnft/x/evmsupport/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BindAddress(goCtx context.Context, msg *types.MsgBindAddress) (*types.MsgBindAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valid, err := k.ValidateEVM(msg.SignedMessage, msg.Signature, msg.EthAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrVerifyingSignature, err.Error())
	}

	if !valid {
		return nil, sdkerrors.Wrap(types.ErrInvalidSignature, "Invalid Signature")
	}

	_, found := k.Keeper.GetAddressBinding(ctx, msg.GetEthAddress())
	if found {
		return nil, sdkerrors.Wrap(types.ErrAddressAlreadyBound, msg.GetEthAddress()+"|"+msg.GetCreator())
	}

	k.Keeper.SetAddressBinding(ctx, types.AddressBinding{
		EthAddress:    msg.GetEthAddress(),
		NativeAddress: msg.GetCreator(),
	})


	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.EventMessage, types.EventTypeDeleteBinding),
			sdk.NewAttribute(types.AttributeValueBinder, msg.GetCreator()),
			sdk.NewAttribute(types.AttributeValueEthAddress, msg.GetEthAddress()),
			sdk.NewAttribute(types.AttributeKeyCreateBindingResult, "success"),
		),
	)

	return &types.MsgBindAddressResponse{
		EthAddress:    msg.GetEthAddress(),
		NativeAddress: msg.GetCreator(),
	}, nil
}
