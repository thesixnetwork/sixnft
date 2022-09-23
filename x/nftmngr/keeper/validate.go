package keeper

import (
	"fmt"
	"sixnft/x/nftmngr/types"
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
			fmt.Println("Attribute not found: ", attriVal.Name)
			continue
		}
		if attrDef.DataType != attriVal.DataType {
			fmt.Println("attrDef.DataType: ", attrDef.DataType)
			return false, attrDef.Name
		}
	}
	return true, ""
}

func MergeNFTDataAttributes(originAttributes []*types.AttributeDefinition, onchainAttributes []*types.AttributeDefinition) []*types.AttributeDefinition {
	mergedAttributes := make([]*types.AttributeDefinition, 0)
	for _, originAttribute := range originAttributes {
		mergedAttributes = append(append(mergedAttributes, originAttribute), onchainAttributes...)
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