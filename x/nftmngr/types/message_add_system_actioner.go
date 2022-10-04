package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddSystemActioner = "add_system_actioner"

var _ sdk.Msg = &MsgAddSystemActioner{}

func NewMsgAddSystemActioner(creator string, nftSchemaCode string, actioner string) *MsgAddSystemActioner {
	return &MsgAddSystemActioner{
		Creator:       creator,
		NftSchemaCode: nftSchemaCode,
		Actioner:      actioner,
	}
}

func (msg *MsgAddSystemActioner) Route() string {
	return RouterKey
}

func (msg *MsgAddSystemActioner) Type() string {
	return TypeMsgAddSystemActioner
}

func (msg *MsgAddSystemActioner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddSystemActioner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddSystemActioner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
