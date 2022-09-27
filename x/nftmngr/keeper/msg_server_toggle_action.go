package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	// Toggle action
	if mapAction.IsActive {
		mapAction.IsActive = false
	} else {
		mapAction.IsActive = true
	}
	// Update is_active in schema
	for i, action := range schema.OnchainData.Actions {
		if action.Name == msg.Action {
			schema.OnchainData.Actions[i] = &mapAction
			break
		}
	}


	k.Keeper.SetNFTSchema(ctx, schema)
	

	return &types.MsgToggleActionResponse{
		Code:  msg.Code,
		Name:  mapAction.Name,
		OnchainDataAction: schema.OnchainData,
	}, nil
}
