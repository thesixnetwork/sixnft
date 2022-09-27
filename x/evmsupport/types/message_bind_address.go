package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBindAddress = "bind_address"

var _ sdk.Msg = &MsgBindAddress{}

func NewMsgBindAddress(creator string, ethAddress string, signature string, signedMessage string) *MsgBindAddress {
	return &MsgBindAddress{
		Creator:       creator,
		EthAddress:    ethAddress,
		Signature:     signature,
		SignedMessage: signedMessage,
	}
}

func (msg *MsgBindAddress) Route() string {
	return RouterKey
}

func (msg *MsgBindAddress) Type() string {
	return TypeMsgBindAddress
}

func (msg *MsgBindAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBindAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBindAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
