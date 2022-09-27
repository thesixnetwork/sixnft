package keeper

import (
	"context"
	"encoding/json"

	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) PerformActionByAdmin(goCtx context.Context, msg *types.MsgPerformActionByAdmin) (*types.MsgPerformActionByAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	tokenData, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, msg.TokenId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, msg.NftSchemaCode)
	}

	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}
	mapAction := types.Action{}
	for _, action := range schema.OnchainData.Actions {
		if !action.IsActive {
			return nil, sdkerrors.Wrap(types.ErrActionIsNotActive, action.Name)
		}
		if action.Name == msg.Action {
			mapAction = *action
			break
		}
	}
	meta := types.NewMetadata(&schema, &tokenData, schema.OriginData.AttributeOverriding)

	err := ProcessAction(meta, &mapAction)
	if err != nil {
		return nil, err
	}
	// Check if ChangeList is empty, error if empty
	if len(meta.ChangeList) == 0 {
		return nil, sdkerrors.Wrap(types.ErrEmptyChangeList, msg.Action)
	}

	k.Keeper.SetNftData(ctx, tokenData)

	// Check action with reference exists
	if msg.RefId != "" {

		_, found := k.Keeper.GetActionByRefId(ctx, msg.RefId)
		if found {
			return nil, sdkerrors.Wrap(types.ErrRefIdAlreadyExists, msg.RefId)
		}

		k.Keeper.SetActionByRefId(ctx, types.ActionByRefId{
			RefId:         msg.RefId,
			Creator:       msg.Creator,
			NftSchemaCode: msg.NftSchemaCode,
			TokenId:       msg.TokenId,
			Action:        mapAction.Name,
		})
	}
	// Emit events on metadata change
	changeList, _ := json.Marshal(meta.ChangeList)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.EventMessage, types.EventTypeRunAction),
			sdk.NewAttribute(types.AttributeKeyRunActionResult, "success"),
			// Emit change list attributes
			sdk.NewAttribute(types.AttributeKeyRunActionChangeList, string(changeList)),
		),
	)

	return &types.MsgPerformActionByAdminResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}
