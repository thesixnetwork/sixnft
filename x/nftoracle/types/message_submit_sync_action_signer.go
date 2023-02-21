package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitSyncActionSigner = "submit_sync_action_signer"

var _ sdk.Msg = &MsgSubmitSyncActionSigner{}

func NewMsgSubmitSyncActionSigner(creator string, syncId uint64, actorAddress string, ownerAddress string, expireAt string) *MsgSubmitSyncActionSigner {
	return &MsgSubmitSyncActionSigner{
		Creator:      creator,
		SyncId:       syncId,
		ActorAddress: actorAddress,
		OwnerAddress: ownerAddress,
		ExpireAt:     expireAt,
	}
}

func (msg *MsgSubmitSyncActionSigner) Route() string {
	return RouterKey
}

func (msg *MsgSubmitSyncActionSigner) Type() string {
	return TypeMsgSubmitSyncActionSigner
}

func (msg *MsgSubmitSyncActionSigner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitSyncActionSigner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitSyncActionSigner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
