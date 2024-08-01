package msg_server

import (
	"context"
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/keeper"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msg_server) UpdateSchemaAttribute(goCtx context.Context, msg *types.MsgUpdateSchemaAttribute) (*types.MsgUpdateSchemaAttributeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var update_attribute types.AttributeDefinition
	input_addribute, err := base64.StdEncoding.DecodeString(msg.Base64UpdateAttriuteDefenition)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(input_addribute, &update_attribute)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	// Check if the value exists
	valFound, isFound := k.GetSchemaAttribute(
		ctx,
		msg.NftSchemaCode,
		update_attribute.Name,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Attribute not exists in schema")
	}

	//get existing nft schema
	schema, schemaFound := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !schemaFound {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	err = keeper.ValidateAttributeNames([]*types.AttributeDefinition{&update_attribute})
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	// parse DefaultMintValue to SchemaAttributeValue
	schmaAttributeValue, err := keeper.ConvertDefaultMintValueToSchemaAttributeValue(update_attribute.DefaultMintValue)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	var schemaAttribute = types.SchemaAttribute{
		Creator:       msg.Creator,
		NftSchemaCode: valFound.NftSchemaCode,
		Name:          valFound.Name,
		DataType:      update_attribute.DataType,
		CurrentValue:  schmaAttributeValue,
	}

	k.SetSchemaAttribute(ctx, schemaAttribute)

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateAttribute,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.EventTypeAddAttribute, update_attribute.Name),
		),
	)

	return &types.MsgUpdateSchemaAttributeResponse{
		NftSchemaCode: msg.NftSchemaCode,
		Name:          update_attribute.Name,
		NewAttribute:  &schemaAttribute,
	}, nil
}
