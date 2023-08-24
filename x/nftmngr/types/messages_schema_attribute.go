package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateSchemaAttribute = "create_schema_attribute"
	TypeMsgUpdateSchemaAttribute = "update_schema_attribute"
	TypeMsgDeleteSchemaAttribute = "delete_schema_attribute"
)

var _ sdk.Msg = &MsgCreateSchemaAttribute{}

func NewMsgCreateSchemaAttribute(
	creator string,
	nftSchemaCode string,
	name string,
	base64NewAttriuteDefenition string,

) *MsgCreateSchemaAttribute {
	return &MsgCreateSchemaAttribute{
		Creator:             creator,
		NftSchemaCode:       nftSchemaCode,
		Name:                name,
		Base64NewAttriuteDefenition: base64NewAttriuteDefenition,
	}
}

func (msg *MsgCreateSchemaAttribute) Route() string {
	return RouterKey
}

func (msg *MsgCreateSchemaAttribute) Type() string {
	return TypeMsgCreateSchemaAttribute
}

func (msg *MsgCreateSchemaAttribute) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSchemaAttribute) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSchemaAttribute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSchemaAttribute{}

func NewMsgUpdateSchemaAttribute(
	creator string,
	nftSchemaCode string,
	name string,
	base64NewAttriuteDefenition string,

) *MsgUpdateSchemaAttribute {
	return &MsgUpdateSchemaAttribute{
		Creator:             creator,
		NftSchemaCode:       nftSchemaCode,
		Name:                name,
		Base64NewAttriuteDefenition: base64NewAttriuteDefenition,
	}
}

func (msg *MsgUpdateSchemaAttribute) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSchemaAttribute) Type() string {
	return TypeMsgUpdateSchemaAttribute
}

func (msg *MsgUpdateSchemaAttribute) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSchemaAttribute) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSchemaAttribute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSchemaAttribute{}

func NewMsgDeleteSchemaAttribute(
	creator string,
	nftSchemaCode string,
	name string,

) *MsgDeleteSchemaAttribute {
	return &MsgDeleteSchemaAttribute{
		Creator:       creator,
		NftSchemaCode: nftSchemaCode,
		Name:          name,
	}
}
func (msg *MsgDeleteSchemaAttribute) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSchemaAttribute) Type() string {
	return TypeMsgDeleteSchemaAttribute
}

func (msg *MsgDeleteSchemaAttribute) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSchemaAttribute) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSchemaAttribute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
