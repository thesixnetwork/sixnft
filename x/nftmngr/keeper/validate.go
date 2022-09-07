package keeper

import (
	"fmt"

	"sixnft/x/nftmngr/types"
)

// Check Duplicate Attributes in array
func HasDuplicateAttributes(attributes []*types.AttributeDefinition) (bool, error) {
	attributesLen := len(attributes)
	for i := 0; i < attributesLen; i++ {
		for j := i + 1; j < attributesLen; j++ {
			if attributes[i].Name == attributes[j].Name {
				return false, fmt.Errorf("Duplicate attribute name: %s", attributes[i].Name)
			}
		}
	}
	return true, nil
}

func HasDuplicateNftAttributesValue(attributes []*types.NftAttributeValue) (bool, error) {
	attributesLen := len(attributes)
	for i := 0; i < attributesLen; i++ {
		for j := i + 1; j < attributesLen; j++ {
			if attributes[i].Name == attributes[j].Name {
				return false, fmt.Errorf("Duplicate attribute value name: %s", attributes[i].Name)
			}
		}
	}
	return true, nil
}

func HasSameType(attributes_1 []*types.AttributeDefinition, attributes_2 []*types.AttributeDefinition) (bool, error) {
	// If attributes have same name, then they must have same type
	for i := 0; i < len(attributes_1); i++ {
		for j := i + 1; j < len(attributes_2); j++ {
			if attributes_1[i].Name == attributes_2[j].Name {
				if attributes_1[i].DataType != attributes_2[j].DataType {
					return false, fmt.Errorf("Attribute %s has different type", attributes_1[i].Name)
				}
				return true, nil
			}
		}
	}
	return true, nil
}
