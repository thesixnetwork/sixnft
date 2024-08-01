package msg_server

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msg_server) ToggleAction(goCtx context.Context, msg *types.MsgToggleAction) (*types.MsgToggleActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.GetNFTSchema(ctx, msg.Code)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.Code)
	}
	// Check if creator is owner of schema
	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	// Update is_active in schema
	for i, action := range schema.OnchainData.Actions {
		if action.Name == msg.Action {
			schema.OnchainData.Actions[i].Disable = msg.Status
		}
	}

	k.SetNFTSchema(ctx, schema)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeToggleNFTAction,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.Code),
			sdk.NewAttribute(types.AttributeKeyToggleNFTAction, msg.Action),
			sdk.NewAttribute(types.AttributeKeyToggleNFTActionResult, strconv.FormatBool(msg.Status)),
		),
	})

	return &types.MsgToggleActionResponse{
		Code:   msg.Code,
		Name:   msg.Action,
		Status: msg.Status,
	}, nil
}
