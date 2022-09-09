package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveBinding = "remove_binding"

var _ sdk.Msg = &MsgRemoveBinding{}

func NewMsgRemoveBinding(creator string, ethAddress string, signature string, signedMessage string) *MsgRemoveBinding {
	return &MsgRemoveBinding{
		Creator:       creator,
		EthAddress:    ethAddress,
		Signature:     signature,
		SignedMessage: signedMessage,
	}
}

func (msg *MsgRemoveBinding) Route() string {
	return RouterKey
}

func (msg *MsgRemoveBinding) Type() string {
	return TypeMsgRemoveBinding
}

func (msg *MsgRemoveBinding) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveBinding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveBinding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
