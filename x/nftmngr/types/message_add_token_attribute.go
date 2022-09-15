package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddTokenAttribute = "add_token_attribute"

var _ sdk.Msg = &MsgAddTokenAttribute{}

func NewMsgAddTokenAttribute(creator string, code string, base64NewAttriuteDefenition string) *MsgAddTokenAttribute {
	return &MsgAddTokenAttribute{
		Creator:                     creator,
		Code:                        code,
		Base64NewAttriuteDefenition: base64NewAttriuteDefenition,
	}
}

func (msg *MsgAddTokenAttribute) Route() string {
	return RouterKey
}

func (msg *MsgAddTokenAttribute) Type() string {
	return TypeMsgAddTokenAttribute
}

func (msg *MsgAddTokenAttribute) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddTokenAttribute) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddTokenAttribute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
