package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgHidddenAttributes = "hiddden_attributes"

var _ sdk.Msg = &MsgHidddenAttributes{}

func NewMsgHidddenAttributes(creator string, nftSchemaCode string, attibuteName string, attributeName string) *MsgHidddenAttributes {
	return &MsgHidddenAttributes{
		Creator:       creator,
		NftSchemaCode: nftSchemaCode,
		AttibuteName:  attibuteName,
		AttributeName: attributeName,
	}
}

func (msg *MsgHidddenAttributes) Route() string {
	return RouterKey
}

func (msg *MsgHidddenAttributes) Type() string {
	return TypeMsgHidddenAttributes
}

func (msg *MsgHidddenAttributes) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgHidddenAttributes) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgHidddenAttributes) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
