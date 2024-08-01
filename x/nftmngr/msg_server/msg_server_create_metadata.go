package msg_server

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/thesixnetwork/sixnft/x/nftmngr/keeper"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msg_server) CreateMetadata(goCtx context.Context, msg *types.MsgCreateMetadata) (*types.MsgCreateMetadataResponse, error) {
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

	mapOfMinters, userMintfound := k.Keeper.GetMetadataCreator(ctx, data.NftSchemaCode)
	// Check mint authorization
	switch schema.MintAuthorization {
	case types.KeyMintPermissionOnlySystem:
		// Check if creator is the schema owner
		if msg.Creator != schema.Owner {
			return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
		}
	case types.KeyMintPermissionAll:
		// Add creator to minters list
		if !userMintfound {
			mapOfMinters = types.MetadataCreator{
				NftSchemaCode:    schema.Code,
				MetadataMintedBy: make([]*types.MapTokenToMinter, 0),
			}
			mapOfMinters.MetadataMintedBy = append(mapOfMinters.MetadataMintedBy, &types.MapTokenToMinter{
				TokenId: data.TokenId,
				Minter:  msg.Creator,
			})
		} else {
			mapOfMinters.MetadataMintedBy = append(mapOfMinters.MetadataMintedBy, &types.MapTokenToMinter{
				TokenId: data.TokenId,
				Minter:  msg.Creator,
			})
		}
	}

	// Validate Schema Message and return error if not valid
	valid, err := ValidateNFTData(&data, &schema)
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
					data.OnchainAttributes = append(data.OnchainAttributes, keeper.NewNFTAttributeValueFromDefaultValue(attr.Name, attr.DefaultMintValue))
				}
			}
		}
	}

	// Check if the data already exists
	_, dataFound := k.Keeper.GetNftData(ctx, data.NftSchemaCode, data.TokenId)
	if dataFound {
		return nil, sdkerrors.Wrap(types.ErrMetadataAlreadyExists, data.NftSchemaCode)
	}

	if !schema.OnchainData.GetStatusByKey(types.KeyNFTStatusFirstMintComplete) {
		schema.OnchainData.SetStatusByKey(types.KeyNFTStatusFirstMintComplete, true)
		k.Keeper.SetNFTSchema(ctx, schema)
	}

	// Add minter to minters list
	k.Keeper.SetMetadataCreator(ctx, mapOfMinters)

	// Add the data to the store
	k.Keeper.SetNftData(ctx, data)

	// Add the minted of any schema to collection
	k.Keeper.AddMetadataToCollection(ctx, &data)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateMetadata,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyCreateMetaDataTokenID, msg.TokenId),
			sdk.NewAttribute(types.AttributeKeyCreateMetaDataResult, "success"),
		),
	})

	return &types.MsgCreateMetadataResponse{
		NftSchemaCode: data.NftSchemaCode,
		TokenId:       data.TokenId,
	}, nil
}

// Validate NFT Data
func ValidateNFTData(data *types.NftData, schema *types.NFTSchema) (bool, error) {
	// Origin Data Origin Attributes Map
	mapAttributeDefinition := keeper.CreateAttrDefMap(schema.OriginData.OriginAttributes)

	// Merge Origin Attributes and Onchain Token Attributes together
	// NOTE: No need to merge schema attributes because we will retrieve the attributes value from its directly
	// WE WILL NOT CHANGE SCHEMA VALUE IN THIS FUNCTION
	mergedAttributes := keeper.MergeNFTDataAttributes(schema.OriginData.OriginAttributes, schema.OnchainData.TokenAttributes)
	mergedMap := keeper.CreateAttrDefMap(mergedAttributes)

	// Check if attributes exist in schema
	attributesExistsInSchema, err := keeper.NFTDataAttributesExistInSchema(mergedMap, data.OnchainAttributes)
	if !attributesExistsInSchema {
		return false, sdkerrors.Wrap(types.ErrOnchainAttributesNotExistsInSchema, fmt.Sprintf("Attribute does not exist in schema: %s", err))
	}

	// Check if origin attributes exist in schema
	attributesOriginExistsInSchema, err := keeper.NFTDataAttributesExistInSchema(mapAttributeDefinition, data.OriginAttributes)
	if !attributesOriginExistsInSchema {
		return false, sdkerrors.Wrap(types.ErrOnchainAttributesNotExistsInSchema, fmt.Sprintf("Attribute does not exist in schema: %s", err))
	}

	// Validate required attributes
	validated, requiredAttributeName := keeper.ValidateRequiredAttributes(schema.OnchainData.TokenAttributes, keeper.CreateNftAttrValueMap(data.OnchainAttributes))
	if !validated {
		return false, sdkerrors.Wrap(types.ErrRequiredAttributeMissing, requiredAttributeName)
	}

	// Validate Onchain Attributes Value
	duplicated, err := keeper.HasDuplicateNftAttributesValue(data.OnchainAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOnchainAttributesValue, fmt.Sprintf("Duplicate attribute name: %s", err))
	}

	// Validate Origin Attributes Value
	duplicated, err = keeper.HasDuplicateNftAttributesValue(data.OriginAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOriginAttributesValue, fmt.Sprintf("Duplicate attribute name: %s", err))
	}

	// Validate Origin Attributes Exist in Schema
	hasSameType, err := keeper.HasSameTypeAsSchema(mapAttributeDefinition, data.OriginAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrOriginAttributesNotSameTypeAsSchema, fmt.Sprintf("Does not have same type as schema: %s", err))
	}

	// Validate Onchain Attributes Exist in Schema
	hasSameType, err = keeper.HasSameTypeAsSchema(mergedMap, data.OnchainAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrOnchainTokenAttributesNotSameTypeAsSchema, fmt.Sprintf("Does not have same type as schema: %s", err))
	}
	return true, nil
}
