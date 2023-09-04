package keeper

import (
	"context"
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) UpdateAction(goCtx context.Context, msg *types.MsgUpdateAction) (*types.MsgUpdateActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_updateAction, err := base64.StdEncoding.DecodeString(msg.Base64UpdateAction)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}
	updateAction := types.Action{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(_updateAction, &updateAction)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	// get existing action
	actionOfSchema, found := k.Keeper.GetActionOfSchema(ctx, msg.NftSchemaCode, updateAction.Name)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, updateAction.Name)
	}

	//get existing nft schema
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	// updator is valid
	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
	}

	// update action by its index
	schema.OnchainData.Actions[actionOfSchema.Index] = &updateAction

	// update schema
	k.Keeper.SetNFTSchema(ctx, schema)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateAction,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyUpdateActionName, updateAction.Name),
		),
	})

	return &types.MsgUpdateActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		Name:          updateAction.Name,
	}, nil
}
