package keeper

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

// validate Action data
func ValidateAction(action *types.Action, schema *types.NFTSchema) error {
	// Onchain Data Nft Actions Map
	mapNftActions := CreateActionMap(schema.OnchainData.Actions)
	// check if action name is unique
	if _, found := mapNftActions[action.Name]; found {
		return sdkerrors.Wrap(types.ErrActionAlreadyExists, action.Name)
	}

	//validate action struct
	if action.Name == "" || action.Name == " " {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action name is empty")
	}
	if action.Desc == "" || action.Desc == " " {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action description is empty")
	}
	if action.When == "" || action.When == " " {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action type is empty")
	}
	//validate array of action.Then is not empty
	if len(action.Then) == 0 {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action.Then is empty")
	}
	return nil
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


// stringfy tokenId list to string token_id1,token_id2,token_id3
func StringfyTokenIdList(list []string) string {
	var result string
	for i, item := range list {
		if i == 0 {
			result = item
		} else {
			result = result + "," + item
		}
	}
	return result
}
