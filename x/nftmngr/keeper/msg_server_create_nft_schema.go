package keeper

import (
	"context"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// AttributeName regular expression
	RegxAttributeName = `^[a-z]{1}[a-z0-9_]*[a-z0-9]{1}$`
	RegxActionName    = `^[A-Za-z]{1}[A-Za-z0-9_]*[A-Za-z0-9]{1}$`
	//regexp.MatchString(`^[a-z]{1}[a-z0-9_]*[a-z0-9]{1}$`, "user_name9")
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
	_ = MergeAllSchemaAttributesAndAlterOrderIndex(schema.OriginData.OriginAttributes, schema.OnchainData.NftAttributes, schema.OnchainData.TokenAttributes)

	// Add the schema to the store
	k.Keeper.SetNFTSchema(ctx, schema)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateSchema,
			sdk.NewAttribute(types.AttributeKeyCreateSchemaCode, schema.Code),
			sdk.NewAttribute(types.AttributeKeyCreateSchemaResult, "success"),
		),
	})

	return &types.MsgCreateNFTSchemaResponse{
		Code: schema.Code,
	}, nil
}

// Validate NFT Schema
func (k msgServer) ValidateNFTSchema(schema *types.NFTSchema) (bool, error) {
	// Origin Data Origin Attributes Map
	mapAttributeOriginDefinition := CreateAttrDefMap(schema.OriginData.OriginAttributes)
	for _, attriDef := range mapAttributeOriginDefinition {
		mapAttributeOriginDefinition[attriDef.Name] = attriDef
	}
	// Check for duplicate origin attributes
	duplicated, errString := HasDuplicateAttributes(schema.OriginData.OriginAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOriginAttributes, fmt.Sprintf("Duplicate attribute name: %s", errString))
	}
	// Validate Origin Data Origin Attributes
	err := ValidateAttributeNames(schema.OriginData.OriginAttributes)
	if err != nil {
		return false, err
	}
	// Validate Onchain Data Onchain Attributes
	err = ValidateAttributeNames(append(schema.OnchainData.NftAttributes, schema.OnchainData.TokenAttributes...))
	if err != nil {
		return false, err
	}
	// Loop over actions to validate action name
	for _, action := range schema.OnchainData.Actions {
		// Validate action name
		err = ValidateActionName(action.Name)
		if err != nil {
			return false, err
		}
	}
	// Create map of action name
	mapActionName := make(map[string]bool)
	// Loop over actions to validate action name
	for _, action := range schema.OnchainData.Actions {
		// Validate duplicate action name
		if _, ok := mapActionName[action.Name]; ok {
			return false, sdkerrors.Wrap(types.ErrDuplicateActionName, action.Name)
		}
		mapActionName[action.Name] = true
	}
	// Validate for duplicate onchain nft attributes
	duplicated, errString = HasDuplicateAttributes(schema.OnchainData.NftAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOnchainNFTAttributes, fmt.Sprintf("Duplicate attribute name: %s", errString))
	}
	// Validate for duplicate onchain token attributes
	duplicated, errString = HasDuplicateAttributes(schema.OnchainData.TokenAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOnchainTokenAttributes, fmt.Sprintf("Duplicate attribute name: %s", errString))
	}
	duplicated, errString = HasDuplicateNftAttributesValue(schema.OnchainData.NftAttributesValue)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateAttributesValue, fmt.Sprintf("Duplicate attribute name: %s", errString))
	}
	// Validate if attributes have the same type
	hasSameType, errString := HasSameType(mapAttributeOriginDefinition, schema.OnchainData.NftAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrSameTypeNFTAttributes, fmt.Sprintf("Attribute type not the same: %s", errString))
	}
	hasSameType, errString = HasSameType(mapAttributeOriginDefinition, schema.OnchainData.TokenAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrSameTypeTokenAttributes, fmt.Sprintf("Attribute type not the same: %s", errString))
	}
	// Validate if default mint value has the same type
	hasSameType, errString = DefaultMintValueHasSameType(schema.OnchainData.TokenAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrNotSameTypeDefaultMintValue, fmt.Sprintf("Attribute type not the same: %s", errString))
	}
	// validate if attribute id is set

	return true, nil
}

func ValidateAttributeNames(attributeDefinitions []*types.AttributeDefinition) error {
	// Loop over definitions and validate
	for _, attrDef := range attributeDefinitions {
		// Check if attribute name is snake case
		matched, _ := regexp.MatchString(RegxAttributeName, attrDef.Name)
		if !matched {
			return sdkerrors.Wrap(types.ErrInvalidAttributeName, attrDef.Name)
		}
	}
	return nil
}

// Validate action name
func ValidateActionName(actionName string) error {
	// Check if action name matches RegxActionName
	matched, _ := regexp.MatchString(RegxActionName, actionName)
	if !matched {
		return sdkerrors.Wrap(types.ErrInvalidActionName, actionName)
	}
	return nil
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

func MergeAllSchemaAttributesAndAlterOrderIndex(originAttributes []*types.AttributeDefinition, onchainNFTAttributes []*types.AttributeDefinition, onchainTokenAttribute []*types.AttributeDefinition) []*types.AttributeDefinition {
	mergedAttributes := make([]*types.AttributeDefinition, 0)
	var index uint64 = 0
	for _, originAttribute := range originAttributes {
		originAttribute.Index = index
		mergedAttributes = append(append(mergedAttributes, originAttribute), onchainNFTAttributes...)
		index++
	}
	for _, _onchainNFTAttribute := range onchainNFTAttributes {
		_onchainNFTAttribute.Index = index
		mergedAttributes = append(append(mergedAttributes, _onchainNFTAttribute), onchainNFTAttributes...)
		index++
	}
	for _, _onchainTokenAttribute := range onchainTokenAttribute {
		_onchainTokenAttribute.Index = index
		mergedAttributes = append(append(mergedAttributes, _onchainTokenAttribute), onchainTokenAttribute...)
		index++
	}
	return mergedAttributes
}
