package keeper

import (
	"context"
	"encoding/base64"
	"fmt"

	"sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMetadata(goCtx context.Context, msg *types.MsgCreateMetadata) (*types.MsgCreateMetadataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	metadata, err := base64.StdEncoding.DecodeString(msg.Base64NFTData)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}
	data := types.NftData{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(metadata, &data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}
	// Get nft schema from store
	schema, schemaFound := k.Keeper.GetNFTSchema(ctx, data.NftSchemaCode)
	// Check if the schema already exists
	if !schemaFound {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, data.NftSchemaCode)
	}
	// Validate Schema Message and return error if not valid
	valid, err := k.ValidateNFTData(&data, &schema)
	_ = valid
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	// Append Attribute with default value to NFT Data if not exist in NFT Data yet
	mapOfTokenAttributeValues := map[string]*types.NftAttributeValue{}
	for _, attr := range data.OnchainAttributes {
		mapOfTokenAttributeValues[attr.Name] = attr
	}
	for _, attr := range schema.OnchainData.TokenAttributes {
		if attr.Required {
			if _, ok := mapOfTokenAttributeValues[attr.Name]; !ok {
				if attr.DefaultMintValue != nil {
					data.OnchainAttributes = append(data.OnchainAttributes, NewNFTAttributeValueFromDefaultValue(attr.Name, attr.DefaultMintValue))
				}
			}
		}
	}

	// Add attributes from schema to metadata onchain attributes
	for _, attribute := range schema.OnchainData.NftAttributesValue {
		data.OnchainAttributes = append(append(data.OnchainAttributes, attribute), data.OnchainAttributes...)
	}
	// Check if the data already exists
	_, dataFound := k.Keeper.GetNftData(ctx, data.NftSchemaCode, data.TokenId)
	if dataFound {
		return nil, sdkerrors.Wrap(types.ErrMetadataAlreadyExists, data.NftSchemaCode)
	}
	// Set first mint complete as true
	if !schema.OnchainData.Status[types.KeyNFTStatusFirstMintComplete] {
		schema.OnchainData.Status[types.KeyNFTStatusFirstMintComplete] = true
		k.Keeper.SetNFTSchema(ctx, schema)
	}
	// Add the data to the store
	k.Keeper.SetNftData(ctx, data)

	return &types.MsgCreateMetadataResponse{
		NftSchemaCode: data.NftSchemaCode,
		TokenId:       data.TokenId,
	}, nil
}

// Validate NFT Data
func (k msgServer) ValidateNFTData(data *types.NftData, schema *types.NFTSchema) (bool, error) {
	// Origin Data Origin Attributes Map
	mapAttributeDefinition := CreateAttrDefMap(schema.OriginData.OriginAttributes)
	// Validate required attributes
	validated, requiredAttributeName := ValidateRequiredAttributes(schema.OnchainData.TokenAttributes, CreateNftAttrValueMap(data.OnchainAttributes))
	if !validated {
		return false, sdkerrors.Wrap(types.ErrRequiredAttributeMissing, requiredAttributeName)
	}
	// Validate Onchain Attributes Value
	duplicated, err := HasDuplicateNftAttributesValue(data.OnchainAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOnchainAttributesValue, fmt.Sprintf("Duplicate attribute name: %s", err))
	}
	// Validate Origin Attributes Value
	duplicated, err = HasDuplicateNftAttributesValue(data.OriginAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOriginAttributesValue, fmt.Sprintf("Duplicate attribute name: %s", err))
	}
	// Validate Origin Attributes Exist in Schema
	hasSameType, err := HasSameTypeAsSchema(mapAttributeDefinition, data.OriginAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrOriginAttributesNotSameTypeAsSchema, fmt.Sprintf("Does not have same type as schema: %s", err))
	}
	// Merge Origin Attributes and Onchain Attributes together
	mergedAttributes := MergeNFTDataAttributes(schema.OriginData.OriginAttributes, schema.OnchainData.TokenAttributes)
	mergedMap := CreateAttrDefMap(mergedAttributes)
	// Validate Onchain Attributes Exist in Schema
	hasSameType, err = HasSameTypeAsSchema(mergedMap, data.OnchainAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrOnchainTokenAttributesNotSameTypeAsSchema, fmt.Sprintf("Does not have same type as schema: %s", err))
	}
	return true, nil
}

func NewNFTAttributeValueFromDefaultValue(name string, defaultValue *types.DefaultMintValue) *types.NftAttributeValue {
	nftAttributeValue := &types.NftAttributeValue{
		Name: name,
	}
	switch defaultValue.Value.(type) {
	case *types.DefaultMintValue_NumberAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_NumberAttributeValue{NumberAttributeValue: defaultValue.GetNumberAttributeValue()}
	case *types.DefaultMintValue_StringAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_StringAttributeValue{StringAttributeValue: defaultValue.GetStringAttributeValue()}
	case *types.DefaultMintValue_BooleanAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_BooleanAttributeValue{BooleanAttributeValue: defaultValue.GetBooleanAttributeValue()}
	case *types.DefaultMintValue_FloatAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_FloatAttributeValue{FloatAttributeValue: defaultValue.GetFloatAttributeValue()}
	default:
		return nil
	}
	return nftAttributeValue
}
