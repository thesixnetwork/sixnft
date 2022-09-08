package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPerformActionByAdmin = "perform_action_by_admin"

var _ sdk.Msg = &MsgPerformActionByAdmin{}

func NewMsgPerformActionByAdmin(creator string, nftSchemaCode string, tokenId string, action string) *MsgPerformActionByAdmin {
	return &MsgPerformActionByAdmin{
		Creator:       creator,
		NftSchemaCode: nftSchemaCode,
		TokenId:       tokenId,
		Action:        action,
	}
}

func (msg *MsgPerformActionByAdmin) Route() string {
	return RouterKey
}

func (msg *MsgPerformActionByAdmin) Type() string {
	return TypeMsgPerformActionByAdmin
}

func (msg *MsgPerformActionByAdmin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPerformActionByAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPerformActionByAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
