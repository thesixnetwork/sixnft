package keeper

import (
	"context"
	"encoding/base64"

	// "strconv"

	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddTokenAttribute(goCtx context.Context, msg *types.MsgAddTokenAttribute) (*types.MsgAddTokenAttributeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var new_add_token_attribute types.AttributeDefinition
	input_token_addribute, err := base64.StdEncoding.DecodeString(msg.Base64NewTokenAttriute)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(input_token_addribute, &new_add_token_attribute)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}
	//get existing nft schema
	schema, schemaFound := k.Keeper.GetNFTSchema(ctx, msg.GetCode())
	if !schemaFound {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.GetCode())
	}

	//validate AttributeDefinition data
	err = k.ValidateTokenAttribute(&new_add_token_attribute, &schema)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	// append new nft_attributes to array of OnchainData.TokenAttributes
	schema.OnchainData.TokenAttributes = append(schema.OnchainData.TokenAttributes, &new_add_token_attribute)

	// set schema
	k.Keeper.SetNFTSchema(ctx, schema)

	return &types.MsgAddTokenAttributeResponse{
		Code:        msg.GetCode(),
		Name:        schema.Name,
		OnchainData: schema.OnchainData,
	}, nil
}

// validate AttributeDefinition data
func (k Keeper) ValidateTokenAttribute(attribute *types.AttributeDefinition, schema *types.NFTSchema) error {
	// Onchain Data Nft Attributes Map
	mapNftOnchainAttributes := CreateAttrDefMap(schema.OnchainData.NftAttributes)
	mapNftTokenAttributes := CreateAttrDefMap(schema.OnchainData.TokenAttributes)
	mapNFTOriginAttributes := CreateAttrDefMap(schema.OriginData.OriginAttributes)
	// check if attribute name is unique
	if _, found := mapNftOnchainAttributes[attribute.Name]; found {
		return sdkerrors.Wrap(types.ErrAttributeAlreadyExists, attribute.Name)
	}
	if _, found := mapNftTokenAttributes[attribute.Name]; found {
		return sdkerrors.Wrap(types.ErrAttributeAlreadyExists, attribute.Name)
	}
	if _, found := mapNFTOriginAttributes[attribute.Name]; found {
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
