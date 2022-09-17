package keeper

import (
	"context"
	"encoding/base64"

	"sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetNFTAttribute(goCtx context.Context, msg *types.MsgSetNFTAttribute) (*types.MsgSetNFTAttributeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	attributeValueJson, err := base64.StdEncoding.DecodeString(msg.Base64NftAttributeValue)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}
	attributeValue := types.NftAttributeValue{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(attributeValueJson, &attributeValue)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingAttributeValue, err.Error())
	}
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	// Check if creator is owner of schema
	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	// Create map of schema.OnchainData.NftAttributes to check if attribute exists
	mapNftAttributes := make(map[string]types.AttributeDefinition)
	for _, attr := range schema.OnchainData.NftAttributes {
		mapNftAttributes[attr.Name] = *attr
	}
	// Check if attribute exists
	_, found = mapNftAttributes[attributeValue.Name]
	if !found {
		return nil, sdkerrors.Wrap(types.ErrAttributeDoesNotExists, attributeValue.Name)
	}

	foundIndex := -1
	for i, attr := range schema.OnchainData.NftAttributesValue {
		if attr.Name == attributeValue.Name {
			foundIndex = i
			break
		}
	}
	if foundIndex != -1 {
		schema.OnchainData.NftAttributesValue[foundIndex] = &attributeValue
	} else {
		schema.OnchainData.NftAttributesValue = append(schema.OnchainData.NftAttributesValue, &attributeValue)
	}

	k.Keeper.SetNFTSchema(ctx, schema)

	return &types.MsgSetNFTAttributeResponse{}, nil
}
