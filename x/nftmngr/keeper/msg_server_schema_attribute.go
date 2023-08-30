package keeper

import (
	"context"
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) CreateSchemaAttribute(goCtx context.Context, msg *types.MsgCreateSchemaAttribute) (*types.MsgCreateSchemaAttributeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetSchemaAttribute(
		ctx,
		msg.NftSchemaCode,
		msg.Name,
	)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var new_add_attribute types.AttributeDefinition
	input_addribute, err := base64.StdEncoding.DecodeString(msg.Base64NewAttriuteDefenition)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(input_addribute, &new_add_attribute)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}
	//get existing nft schema
	schema, schemaFound := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !schemaFound {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	//validate AttributeDefinition data
	err = k.ValidateAttributeDefinition(ctx, &new_add_attribute, &schema)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	err = ValidateAttributeNames([]*types.AttributeDefinition{&new_add_attribute})
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	// parse DefaultMintValue to SchemaAttributeValue
	schmaAttributeValue, err := ConvertDefaultMintValueToSchemaAttributeValue(new_add_attribute.DefaultMintValue)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	var schemaAttribute = types.SchemaAttribute{
		Creator:             msg.Creator,
		NftSchemaCode:       msg.NftSchemaCode,
		Name:                msg.Name,
		DataType:            new_add_attribute.DataType,
		Required:            new_add_attribute.Required,
		DisplayValueField:   new_add_attribute.DisplayValueField,
		DisplayOption:       new_add_attribute.DisplayOption,
		CurrentValue:        schmaAttributeValue,
		HiddenOveride:       new_add_attribute.HiddenOveride,
		HiddenToMarketplace: new_add_attribute.HiddenToMarketplace,
	}

	k.SetSchemaAttribute(
		ctx,
		schemaAttribute,
	)

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddAttribute,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.EventTypeAddAttribute, msg.Name),
		),
	)

	return &types.MsgCreateSchemaAttributeResponse{
		NftSchemaCode: msg.NftSchemaCode,
		Name:          msg.Name,
		NewAttribute:  &schemaAttribute,
	}, nil
}

func (k msgServer) UpdateSchemaAttribute(goCtx context.Context, msg *types.MsgUpdateSchemaAttribute) (*types.MsgUpdateSchemaAttributeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSchemaAttribute(
		ctx,
		msg.NftSchemaCode,
		msg.Name,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	var new_add_attribute types.AttributeDefinition
	input_addribute, err := base64.StdEncoding.DecodeString(msg.Base64NewAttriuteDefenition)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(input_addribute, &new_add_attribute)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}
	//get existing nft schema
	schema, schemaFound := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !schemaFound {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	//validate AttributeDefinition data
	err = k.ValidateAttributeDefinition(ctx, &new_add_attribute, &schema)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	err = ValidateAttributeNames([]*types.AttributeDefinition{&new_add_attribute})
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	// parse DefaultMintValue to SchemaAttributeValue
	schmaAttributeValue, err := ConvertDefaultMintValueToSchemaAttributeValue(new_add_attribute.DefaultMintValue)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	var schemaAttribute = types.SchemaAttribute{
		Creator:             msg.Creator,
		NftSchemaCode:       valFound.NftSchemaCode,
		Name:                valFound.Name,
		DataType:            new_add_attribute.DataType,
		Required:            new_add_attribute.Required,
		DisplayValueField:   new_add_attribute.DisplayValueField,
		DisplayOption:       new_add_attribute.DisplayOption,
		CurrentValue:        schmaAttributeValue,
		HiddenOveride:       new_add_attribute.HiddenOveride,
		HiddenToMarketplace: new_add_attribute.HiddenToMarketplace,
	}

	k.SetSchemaAttribute(ctx, schemaAttribute)

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateAttribute,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.EventTypeAddAttribute, msg.Name),
		),
	)

	return &types.MsgUpdateSchemaAttributeResponse{
		NftSchemaCode: msg.NftSchemaCode,
		Name:          msg.Name,
		NewAttribute:  &schemaAttribute,
	}, nil
}

func (k msgServer) DeleteSchemaAttribute(goCtx context.Context, msg *types.MsgDeleteSchemaAttribute) (*types.MsgDeleteSchemaAttributeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSchemaAttribute(
		ctx,
		msg.NftSchemaCode,
		msg.Name,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSchemaAttribute(
		ctx,
		msg.NftSchemaCode,
		msg.Name,
	)

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRemoveAttribute,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.EventTypeAddAttribute, msg.Name),
		),
	)

	return &types.MsgDeleteSchemaAttributeResponse{
		NftSchemaCode: msg.NftSchemaCode,
		Name:          msg.Name,
	}, nil
}
