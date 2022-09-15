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

func CreateNftAttrValueMap(nftAttrValues []*types.NftAttributeValue) map[string]*types.NftAttributeValue {
	nftAttrValueMap := make(map[string]*types.NftAttributeValue)
	for _, nftAttrValue := range nftAttrValues {
		nftAttrValueMap[nftAttrValue.Name] = nftAttrValue
	}
	return nftAttrValueMap
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
		mergedAttributes = append(mergedAttributes, originAttribute)
	}
	for _, onchainAttribute := range onchainAttributes {
		mergedAttributes = append(mergedAttributes, onchainAttribute)
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
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_BooleanAttributeValue_); ok {
		return "boolean"
	}
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_StringAttributeValue_); ok {
		return "string"
	}
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_NumberAttributeValue_); ok {
		return "number"
	}
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_FloatAttributeValue_); ok {
		return "float"
	}
	return "default"
}
