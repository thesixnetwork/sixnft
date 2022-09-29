package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetSchemaOwner = "set_schema_owner"

var _ sdk.Msg = &MsgSetSchemaOwner{}

func NewMsgSetSchemaOwner(creator string, schemaCode string, newOwner string) *MsgSetSchemaOwner {
	return &MsgSetSchemaOwner{
		Creator:    creator,
		SchemaCode: schemaCode,
		NewOwner:   newOwner,
	}
}

func (msg *MsgSetSchemaOwner) Route() string {
	return RouterKey
}

func (msg *MsgSetSchemaOwner) Type() string {
	return TypeMsgSetSchemaOwner
}

func (msg *MsgSetSchemaOwner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetSchemaOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetSchemaOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
