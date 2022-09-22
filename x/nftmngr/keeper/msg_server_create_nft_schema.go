package keeper

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateNFTSchema(goCtx context.Context, msg *types.MsgCreateNFTSchema) (*types.MsgCreateNFTSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	jsonSchema, err := base64.StdEncoding.DecodeString(msg.NftSchemaBase64)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}
	schema := types.NFTSchema{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(jsonSchema, &schema)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingSchemaMessage, err.Error())
	}
	schema.Owner = msg.Creator
	// Validate Schema Message and return error if not valid
	valid, err := k.ValidateNFTSchema(&schema)
	_ = valid
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingNFTSchema, err.Error())
	}
	// Check if the schema already exists
	_, found := k.Keeper.GetNFTSchema(ctx, schema.Code)
	if found {
		return nil, sdkerrors.Wrap(types.ErrSchemaAlreadyExists, schema.Code)
	}
	foundOrganization, organizationName := GetOrganizationFromSchemaCode(schema.Code)
	// If there is organization in schema code, check if the organization exists
	if foundOrganization {
		storedOrganization, found := k.Keeper.GetOrganization(ctx, organizationName)
		if found {
			// Check owner of organization
			if storedOrganization.Owner != msg.Creator {
				return nil, sdkerrors.Wrap(types.ErrOrganizationOwner, msg.Creator)
			}
		} else {
			// Store organization
			k.Keeper.SetOrganization(ctx, types.Organization{
				Owner: msg.Creator,
				Name:  organizationName,
			})
		}

	}
	// Add the schema to the store
	k.Keeper.SetNFTSchema(ctx, schema)

	return &types.MsgCreateNFTSchemaResponse{
		Code: schema.Code,
	}, nil
}

// Validate NFT Schema
func (k msgServer) ValidateNFTSchema(schema *types.NFTSchema) (bool, error) {
	// Origin Data Origin Attributes Map
	mapAttributeDefinition := CreateAttrDefMap(schema.OriginData.OriginAttributes)
	var i uint64 = 0
	for _, attriDef := range mapAttributeDefinition {
		mapAttributeDefinition[attriDef.Name] = attriDef
		attriDef.AttributeId = i
		i ++
		// isAttributeIdSet := IsAttributeIdSet(attriDef)
		
		// if !isAttributeIdSet {
		// 	attriDef.Id = i
		// 	i ++
		// }
	}
	// Check for duplicate origin attributes
	duplicated, err := HasDuplicateAttributes(schema.OriginData.OriginAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOriginAttributes, fmt.Sprintf("Duplicate attribute name: %s", err))
	}
	// Validate for duplicate onchain nft attributes
	duplicated, err = HasDuplicateAttributes(schema.OnchainData.NftAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOnchainNFTAttributes, fmt.Sprintf("Duplicate attribute name: %s", err))
	}
	// Validate for duplicate onchain token attributes
	duplicated, err = HasDuplicateAttributes(schema.OnchainData.TokenAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOnchainTokenAttributes, fmt.Sprintf("Duplicate attribute name: %s", err))
	}
	duplicated, err = HasDuplicateNftAttributesValue(schema.OnchainData.NftAttributesValue)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateAttributesValue, fmt.Sprintf("Duplicate attribute name: %s", err))
	}
	// Validate if attributes have the same type
	hasSameType, err := HasSameType(mapAttributeDefinition, schema.OnchainData.NftAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrSameTypeNFTAttributes, fmt.Sprintf("Attribute type not the same: %s", err))
	}
	hasSameType, err = HasSameType(mapAttributeDefinition, schema.OnchainData.TokenAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrSameTypeTokenAttributes, fmt.Sprintf("Attribute type not the same: %s", err))
	}
	// Validate if default mint value has the same type
	hasSameType, err = DefaultMintValueHasSameType(schema.OnchainData.TokenAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrNotSameTypeDefaultMintValue, fmt.Sprintf("Attribute type not the same: %s", err))
	}
	// validate if attribute id is set
	

	return true, nil
}

func GetOrganizationFromSchemaCode(nftSchemaCode string) (bool, string) {
	// Get orgationzation from schema code follow by "."
	// Example: "org1.nft1" => "org1"

	// Get the first index of "."
	index := strings.Index(nftSchemaCode, ".")
	if index == -1 {
		return false, ""
	}
	// Get the organization name
	organizationName := nftSchemaCode[:index]
	return true, organizationName
}
