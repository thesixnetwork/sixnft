package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) ToggleAction(goCtx context.Context, msg *types.MsgToggleAction) (*types.MsgToggleActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.Keeper.GetNFTSchema(ctx, msg.Code)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.Code)
	}
	// Check if creator is owner of schema
	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}
	mapAction := types.Action{}
	for _, action := range schema.OnchainData.Actions {
		if action.Name == msg.Action {
			mapAction = *action
			break
		}
	}
	// Update is_active in schema
	for i, action := range schema.OnchainData.Actions {
		if action.Name == msg.Action && action.Disable != msg.Disable {
			schema.OnchainData.Actions[i].Disable = msg.Disable
		}else if action.Name == msg.Action && action.Disable == msg.Disable {
			return nil, sdkerrors.Wrap(types.ErrActionAlreadyDisabled, msg.Action)
		}
	}

	k.Keeper.SetNFTSchema(ctx, schema)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeToggleNFTAction,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.Code),
			sdk.NewAttribute(types.AttributeKeyToggleNFTAction, msg.Action),
			sdk.NewAttribute(types.AttributeKeyToggleNFTActionResult, "success"),
		),
	})

	return &types.MsgToggleActionResponse{
		Code:              msg.Code,
		Name:              mapAction.Name,
		OnchainDataAction: schema.OnchainData,
	}, nil
}
