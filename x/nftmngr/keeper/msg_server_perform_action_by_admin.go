package keeper

import (
	"context"

	"sixnft/x/nftmngr/types"

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
		if action.Name == msg.Action {
			mapAction = *action
			break
		}
	}
	meta := types.NewMetadata(&tokenData, schema.OriginData.AttributeOverriding)
	err := ProcessAction(meta, &mapAction)
	if err != nil {
		return nil, err
	}

	k.Keeper.SetNftData(ctx, tokenData)

	// Compare the new metadata with the old one and get the difference

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.EventMessage, types.EventTypeRunAction),
			sdk.NewAttribute(types.AttributeKeyRunActionResult, "success"),
			sdk.NewAttribute(types.AttributeKeyOldValue, ""),
			sdk.NewAttribute(types.AttributeKeyNewValue, ""),
		),
	)

	return &types.MsgPerformActionByAdminResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}
