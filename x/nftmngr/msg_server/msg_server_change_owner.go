package msg_server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msg_server) ChangeOrgOwner(goCtx context.Context, msg *types.MsgChangeOrgOwner) (*types.MsgChangeOrgOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//get the organization
	organization, found := k.GetOrganization(ctx, msg.OrgName)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOrganizationNotFound, msg.OrgName)
	}

	if organization.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(types.ErrOrganizationOwner, msg.Creator)
	}

	// validate newOwner address
	_, err := sdk.AccAddressFromBech32(msg.ToNewOwner)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ToNewOwner)
	}

	//change the owner
	organization.Owner = msg.ToNewOwner

	//save the organization
	k.SetOrganization(ctx, organization)

	//emit event
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
