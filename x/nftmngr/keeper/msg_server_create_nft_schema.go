package keeper

import (
	"context"
	"encoding/base64"
	"fmt"

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
	// Add the schema to the store
	k.Keeper.SetNFTSchema(ctx, schema)

	return &types.MsgCreateNFTSchemaResponse{
		Code: schema.Code,
	}, nil
}

// Validate NFT Schema
func (k msgServer) ValidateNFTSchema(schema *types.NFTSchema) (bool, error) {
	// Validate Origin Attributes
	duplicated, err := HasDuplicateAttributes(schema.OriginData.OriginAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOriginAttributes, fmt.Sprintf("Duplicate attribute name: %s", err))
	}
	// Validate Onchain NFT Attributes
	duplicated, err = HasDuplicateAttributes(schema.OnchainData.NftAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOnchainNFTAttributes, fmt.Sprintf("Duplicate attribute name: %s", err))
	}
	// Validate Onchain Token Attributes
	duplicated, err = HasDuplicateAttributes(schema.OnchainData.TokenAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOnchainTokenAttributes, fmt.Sprintf("Duplicate attribute name: %s", err))
	}
	// Validate NFT Attributes Value
	duplicated, err = HasDuplicateNftAttributesValue(schema.OnchainData.NftAttributesValue)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateAttributesValue, fmt.Sprintf("Duplicate attribute value: %s", err))
	}
	// Validate if attributes have the same type for NFT Attributes
	hasSameType, err := HasSameType(schema.OriginData.OriginAttributes, schema.OnchainData.NftAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrSameTypeNFTAttributes, fmt.Sprintf("NFT Attributes and Origin Attributes have different types: %s", err))
	}
	// Validate if attributes have the same type for Token Attributes
	hasSameType, err = HasSameType(schema.OriginData.OriginAttributes, schema.OnchainData.TokenAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrSameTypeTokenAttributes, fmt.Sprintf("Token Attributes and Origin Attributes have different types: %s", err))
	}
	return true, nil
}
