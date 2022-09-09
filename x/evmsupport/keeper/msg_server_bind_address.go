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

	_, found := k.Keeper.GetAddressBinding(ctx, msg.GetEthAddress(), msg.GetCreator())
	if found {
		return nil, sdkerrors.Wrap(types.ErrAddressAlreadyBound, msg.GetEthAddress()+"|"+msg.GetCreator())
	}

	k.Keeper.SetAddressBinding(ctx, types.AddressBinding{
		EthAddress:    msg.GetEthAddress(),
		NativeAddress: msg.GetCreator(),
	})

	return &types.MsgBindAddressResponse{}, nil
}
