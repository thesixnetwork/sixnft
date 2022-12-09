package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgChageOrgOwner = "chage_org_owner"

var _ sdk.Msg = &MsgChageOrgOwner{}

func NewMsgChageOrgOwner(creator string, orgName string, toNewOwner string) *MsgChageOrgOwner {
	return &MsgChageOrgOwner{
		Creator:    creator,
		OrgName:    orgName,
		ToNewOwner: toNewOwner,
	}
}

func (msg *MsgChageOrgOwner) Route() string {
	return RouterKey
}

func (msg *MsgChageOrgOwner) Type() string {
	return TypeMsgChageOrgOwner
}

func (msg *MsgChageOrgOwner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChageOrgOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChageOrgOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
