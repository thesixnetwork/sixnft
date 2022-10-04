package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveSystemActioner = "remove_system_actioner"

var _ sdk.Msg = &MsgRemoveSystemActioner{}

func NewMsgRemoveSystemActioner(creator string, nftSchemaCode string, actioner string) *MsgRemoveSystemActioner {
	return &MsgRemoveSystemActioner{
		Creator:       creator,
		NftSchemaCode: nftSchemaCode,
		Actioner:      actioner,
	}
}

func (msg *MsgRemoveSystemActioner) Route() string {
	return RouterKey
}

func (msg *MsgRemoveSystemActioner) Type() string {
	return TypeMsgRemoveSystemActioner
}

func (msg *MsgRemoveSystemActioner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveSystemActioner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveSystemActioner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
