package keeper

import (
	"sixnft/x/nftmngr/types"
)

// Check Duplicate Attributes in array
func HasDuplicateAttributes(attributes []*types.AttributeDefinition) (bool, string) {
	mapAttributes := map[string]*types.AttributeDefinition{}
	mapCount := map[string]int{}
	for _, attriDef := range attributes {
		mapAttributes[attriDef.Name] = attriDef
	}
	for _, attriDef := range attributes {
		if mapCount[attriDef.Name] > 1 {
			return true, attriDef.Name
		}
		if _, ok := mapCount[attriDef.Name]; ok {
			return true, attriDef.Name
		}
		mapCount[attriDef.Name] = 1
	}
	return false, ""
}

func HasDuplicateNftAttributesValue(attributes []*types.NftAttributeValue) (bool, string) {
	mapAttributes := map[string]*types.NftAttributeValue{}
	mapCount := map[string]int{}
	for _, attriDef := range attributes {
		mapAttributes[attriDef.Name] = attriDef
	}
	for _, attriDef := range attributes {
		if mapCount[attriDef.Name] > 1 {
			return true, attriDef.Name
		}
		if _, ok := mapCount[attriDef.Name]; ok {
			return true, attriDef.Name
		}
		mapCount[attriDef.Name] = 1
	}
	return false, ""
}

func HasSameType(originAttributes []*types.AttributeDefinition, onchainAttributes []*types.AttributeDefinition) (bool, string) {
	mapOriginAttributes := map[string]*types.AttributeDefinition{}

	for _, attriDef := range originAttributes {
		mapOriginAttributes[attriDef.Name] = attriDef
	}
	for _, attriVal := range onchainAttributes {
		attrDef := mapOriginAttributes[attriVal.Name]
		if attrDef == nil {
			continue
		}
		if attrDef.DataType != attriVal.DataType {
			return false, attrDef.Name
		}
	}
	return true, ""
}

func HasSameTypeAsSchema(schemaAttributes []*types.AttributeDefinition, dataAttributes []*types.NftAttributeValue) (bool, string) {
	// If attributes have same name, then they must have same type
	mapSchemaAttributes := map[string]*types.AttributeDefinition{}

	for _, attriDef := range schemaAttributes {
		mapSchemaAttributes[attriDef.Name] = attriDef
	}

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
