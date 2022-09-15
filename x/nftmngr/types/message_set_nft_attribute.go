package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetNFTAttribute = "set_nft_attribute"

var _ sdk.Msg = &MsgSetNFTAttribute{}

func NewMsgSetNFTAttribute(creator string, nftSchemaCode string, attributeValue string) *MsgSetNFTAttribute {
	return &MsgSetNFTAttribute{
		Creator:                 creator,
		NftSchemaCode:           nftSchemaCode,
		Base64NftAttributeValue: attributeValue,
	}
}

func (msg *MsgSetNFTAttribute) Route() string {
	return RouterKey
}

func (msg *MsgSetNFTAttribute) Type() string {
	return TypeMsgSetNFTAttribute
}

func (msg *MsgSetNFTAttribute) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetNFTAttribute) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetNFTAttribute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
