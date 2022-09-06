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
