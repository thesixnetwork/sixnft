package msg_server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msg_server) ChangeOrgOwner(goCtx context.Context, msg *types.MsgChangeOrgOwner) (*types.MsgChangeOrgOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	to, err := sdk.AccAddressFromBech32(msg.ToNewOwner)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ToNewOwner)
	}

	err = k.Keeper.ChangeOrgOwner(ctx, from, to, msg.OrgName)
	if err != nil {
		return nil, err
	}

	// emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeChangeOrgOwner,
			sdk.NewAttribute(types.AttributeKeyOrgName, msg.OrgName),
			sdk.NewAttribute(types.AttributeKeyOldOwner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyNewOwner, msg.ToNewOwner),
		),
	)

	return &types.MsgChangeOrgOwnerResponse{
		OrgName:  msg.OrgName,
		OldOwner: msg.Creator,
		NewOwner: msg.ToNewOwner,
	}, nil
}
