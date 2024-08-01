package msg_server

import (
	"context"
	"encoding/base64"

	// "strconv"

	"github.com/thesixnetwork/sixnft/x/nftmngr/keeper"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msg_server) AddAttribute(goCtx context.Context, msg *types.MsgAddAttribute) (*types.MsgAddAttributeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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
	schema, schemaFound := k.Keeper.GetNFTSchema(ctx, msg.GetCode())
	if !schemaFound {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.GetCode())
	}

	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	//validate AttributeDefinition data
	err = k.ValidateAttributeDefinition(ctx, &new_add_attribute, &schema)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	err = keeper.ValidateAttributeNames([]*types.AttributeDefinition{&new_add_attribute})
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	// Swith location of attribute
	switch msg.Location {
	case types.AttributeLocation_NFT_ATTRIBUTE:
		_defaultMintValue, err := keeper.ConvertDefaultMintValueToSchemaAttributeValue(new_add_attribute.DefaultMintValue)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
		}
		// this case will use Msg.CreateSchemaAtribute
		k.SetSchemaAttribute(ctx, types.SchemaAttribute{
			NftSchemaCode: msg.Code,
			Name:          new_add_attribute.Name,
			CurrentValue:  _defaultMintValue,
			DataType:      new_add_attribute.DataType,
			Creator:       msg.Creator,
		})
		schema.OnchainData.NftAttributes = append(schema.OnchainData.NftAttributes, &new_add_attribute)
	case types.AttributeLocation_TOKEN_ATTRIBUTE:
		// append new token_attributes to array of OnchainData.TokenAttributes
		schema.OnchainData.TokenAttributes = append(schema.OnchainData.TokenAttributes, &new_add_attribute)
		// end the case
	}
	// count the index of new attribute
	count := MergeAndCountAllAttributes(schema.OriginData.OriginAttributes, schema.OnchainData.NftAttributes, schema.OnchainData.TokenAttributes)
	// set new index to new attribute
	new_add_attribute.Index = uint64(count - 1)

	// set schema
	k.Keeper.SetNFTSchema(ctx, schema)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddAttribute,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.Code),
			sdk.NewAttribute(types.AttributeKeyAddAttributeName, new_add_attribute.Name),
			sdk.NewAttribute(types.AttributeKeyAddAttributeLocation, types.AttributeLocation.String(msg.Location)),
			sdk.NewAttribute(types.AttributeKeyAddAttributeResult, "success"),
		),
	})

	return &types.MsgAddAttributeResponse{
		Code:        msg.GetCode(),
		Name:        schema.Name,
		OnchainData: schema.OnchainData,
	}, nil

}

// validate AttributeDefinition data
func (k msg_server) ValidateAttributeDefinition(ctx sdk.Context, attribute *types.AttributeDefinition, schema *types.NFTSchema) error {

	valFound, found := k.GetSchemaAttribute(ctx, schema.Code, attribute.Name)
	if found {
		return sdkerrors.Wrap(types.ErrAttributeAlreadyExists, valFound.Name)
	}
	// Onchain Data Nft Attributes Map
	mapOriginAttributes := keeper.CreateAttrDefMap(schema.OriginData.OriginAttributes)
	mapTokenAttributes := keeper.CreateAttrDefMap(schema.OnchainData.TokenAttributes)
	// check if attribute name is unique
	if _, found := mapOriginAttributes[attribute.Name]; found {
		return sdkerrors.Wrap(types.ErrAttributeAlreadyExists, attribute.Name)
	}
	if _, found := mapTokenAttributes[attribute.Name]; found {
		return sdkerrors.Wrap(types.ErrAttributeAlreadyExists, attribute.Name)
	}

	// // validate struct data
	// if attribute.Name == "" {
	// 	return sdkerrors.Wrap(types.ErrInvalidAttribute, "Attribute name is empty")
	// }
	// if attribute.DataType == "" {
	// 	return sdkerrors.Wrap(types.ErrInvalidAttribute, "Attribute type is empty")
	// }
	// if strconv.FormatBool(attribute.Required) == "" {
	// 	return sdkerrors.Wrap(types.ErrInvalidAttribute, "Attribute required is false")
	// }
	// if attribute.DisplayOption.BoolFalseValue == "" {
	// 	return sdkerrors.Wrap(types.ErrInvalidAttribute, "Attribute BoolFalseValue is empty")
	// }
	// if attribute.DisplayOption.BoolTrueValue == "" {
	// 	return sdkerrors.Wrap(types.ErrInvalidAttribute, "Attribute BoolTrueValue is empty")
	// }
	// if attribute.DisplayOption.Opensea.DisplayType == "" {
	// 	return sdkerrors.Wrap(types.ErrInvalidAttribute, "Attribute DisplayType is empty")
	// }
	// if attribute.DisplayOption.Opensea.TraitType == "" {
	// 	return sdkerrors.Wrap(types.ErrInvalidAttribute, "Attribute TraitType is empty")
	// }
	// if strconv.FormatUint(attribute.DisplayOption.Opensea.MaxValue, 10) == "" {
	// 	return sdkerrors.Wrap(types.ErrInvalidAttribute, "Attribute MaxValue is empty")
	// }
	// if attribute.DefaultMintValue == "" {
	// 	return sdkerrors.Wrap(types.ErrInvalidAttribute, "Attribute DefaultMintValue is empty")
	// }
	// if strconv.FormatBool(attribute.HiddenToMarketplace) == "" {
	// 	return sdkerrors.Wrap(types.ErrInvalidAttribute, "Attribute HiddenToMarketplace is false")
	// }
	return nil
}

// merge all attributes and count the index
func MergeAndCountAllAttributes(originAttributes []*types.AttributeDefinition, onchainNFTAttributes []*types.AttributeDefinition, onchainTokenAttribute []*types.AttributeDefinition) int {
	// length or originAttributes
	length_originAttributes := len(originAttributes)
	// length or onchainNFTAttributes
	length_onchainNFTAttributes := len(onchainNFTAttributes)
	// length or onchainTokenAttribute
	length_onchainTokenAttribute := len(onchainTokenAttribute)

	// length of all attributes
	length_allAttributes := length_originAttributes + length_onchainNFTAttributes + length_onchainTokenAttribute
	return length_allAttributes
}
