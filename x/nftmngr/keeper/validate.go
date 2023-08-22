package keeper

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func CreateAttrDefMap(attrDefs []*types.AttributeDefinition) map[string]*types.AttributeDefinition {
	attrDefMap := make(map[string]*types.AttributeDefinition)
	for _, attrDef := range attrDefs {
		attrDefMap[attrDef.Name] = attrDef
	}
	return attrDefMap
}

// create ActionMap
func CreateActionMap(actions []*types.Action) map[string]*types.Action {
	actionMap := make(map[string]*types.Action)
	for _, action := range actions {
		actionMap[action.Name] = action
	}
	return actionMap
}

func CreateNftAttrValueMap(nftAttrValues []*types.NftAttributeValue) map[string]*types.NftAttributeValue {
	nftAttrValueMap := make(map[string]*types.NftAttributeValue)
	for _, nftAttrValue := range nftAttrValues {
		nftAttrValueMap[nftAttrValue.Name] = nftAttrValue
	}
	return nftAttrValueMap
}

func ValidateRequiredAttributes(schemaAttributes []*types.AttributeDefinition, mapAttributeValues map[string]*types.NftAttributeValue) (bool, string) {
	for _, schemaAttribute := range schemaAttributes {
		if schemaAttribute.Required && (schemaAttribute.DefaultMintValue == nil || schemaAttribute.DefaultMintValue.GetValue() == nil) {
			if _, ok := mapAttributeValues[schemaAttribute.Name]; !ok {
				return false, schemaAttribute.Name
			}
		}
	}
	return true, ""
}

func HasDuplicateAttributes(attributes []*types.AttributeDefinition) (bool, string) {
	mapAttributes := map[string]*types.AttributeDefinition{}
	for _, attriDef := range attributes {
		if _, ok := mapAttributes[attriDef.Name]; ok {
			return true, attriDef.Name
		}
		mapAttributes[attriDef.Name] = attriDef
	}
	return false, ""
}

func HasDuplicateNftAttributesValue(attributes []*types.NftAttributeValue) (bool, string) {
	mapAttributes := map[string]*types.NftAttributeValue{}
	for _, attriDef := range attributes {
		if _, ok := mapAttributes[attriDef.Name]; ok {
			return true, attriDef.Name
		}
		mapAttributes[attriDef.Name] = attriDef
	}
	return false, ""
}

func HasSameType(mapOriginAttributes map[string]*types.AttributeDefinition, onchainAttributes []*types.AttributeDefinition) (bool, string) {
	for _, attriVal := range onchainAttributes {
		attrDef := mapOriginAttributes[attriVal.Name]
		if attrDef == nil {
			// fmt.Println("Attribute not found: ", attriVal.Name)
			continue
		}
		if attrDef.DataType != attriVal.DataType {
			// fmt.Println("attrDef.DataType: ", attrDef.DataType)
			return false, attrDef.Name
		}
	}
	return true, ""
}

func MergeNFTDataAttributes(originAttributes []*types.AttributeDefinition, onchainAttributes []*types.AttributeDefinition) []*types.AttributeDefinition {
	mergedAttributes := make([]*types.AttributeDefinition, 0)
	for _, originAttribute := range originAttributes {
		mergedAttributes = append(append(mergedAttributes, originAttribute), onchainAttributes...)

		// ** What is is actually doing is:
		// ** mergedAttributes = append(mergedAttributes, originAttribute)
		// ** for _, attr := range onchainAttributes {
		// ** mergedAttributes = append(mergedAttributes, attr)
		// ** }
		// ** It is better performance to use append(append(mergedAttributes, originAttribute), onchainAttributes...) instead of the above
		// ** Just in case I forget, I will leave this comment here
	}
	for _, onchainAttribute := range onchainAttributes {
		mergedAttributes = append(append(mergedAttributes, onchainAttribute), originAttributes...)
	}
	return mergedAttributes
}
func HasSameTypeAsSchema(mapSchemaAttributes map[string]*types.AttributeDefinition, dataAttributes []*types.NftAttributeValue) (bool, string) {
	// If attributes have same name, then they must have same type
	for _, attriVal := range dataAttributes {
		attDef := mapSchemaAttributes[attriVal.Name]
		if attDef.DataType != GetTypeFromAttributeValue(attriVal) {
			return false, attriVal.Name
		}
	}
	return true, ""
}
func GetTypeFromAttributeValue(attribute *types.NftAttributeValue) string {
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_BooleanAttributeValue); ok {
		return "boolean"
	}
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_StringAttributeValue); ok {
		return "string"
	}
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_NumberAttributeValue); ok {
		return "number"
	}
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_FloatAttributeValue); ok {
		return "float"
	}
	return "default"
}

func DefaultMintValueHasSameType(attributes []*types.AttributeDefinition) (bool, string) {
	for _, attriDef := range attributes {
		_, attrType := HasDefaultMintValue(*attriDef)
		if attriDef.DataType != attrType {
			return false, attriDef.Name
		}
	}
	return true, ""
}

func HasDefaultMintValue(attribute types.AttributeDefinition) (bool, string) {
	// Check if onchain attribute s value exist for each attribute
	if _, ok := attribute.DefaultMintValue.GetValue().(*types.DefaultMintValue_BooleanAttributeValue); ok {
		return ok, "boolean"
	}
	if _, ok := attribute.DefaultMintValue.GetValue().(*types.DefaultMintValue_StringAttributeValue); ok {
		return ok, "string"
	}
	if _, ok := attribute.DefaultMintValue.GetValue().(*types.DefaultMintValue_NumberAttributeValue); ok {
		return ok, "number"
	}
	if _, ok := attribute.DefaultMintValue.GetValue().(*types.DefaultMintValue_FloatAttributeValue); ok {
		return ok, "float"
	}
	return false, "default"
}

// Check if NFT data attributes exists in schema
func NFTDataAttributesExistInSchema(mapAttributes map[string]*types.AttributeDefinition, dataAttributes []*types.NftAttributeValue) (bool, string) {
	// Check if dataAttributes exist in schema
	for _, attriVal := range dataAttributes {
		if _, ok := mapAttributes[attriVal.Name]; !ok {
			return false, attriVal.Name
		}
	}
	return true, ""
}


// Validate NFT Schema
func ValidateNFTSchema(schema *types.NFTSchema) (bool, error) {
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
