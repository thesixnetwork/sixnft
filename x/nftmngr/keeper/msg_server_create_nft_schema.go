package keeper

import (
	"context"
	"encoding/base64"

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
	_, err := HasDuplicateAttributes(schema.OriginData.OriginAttributes)
	if err != nil {
		return false, err
	}
	// Validate Onchain NFT Attributes
	_, err = HasDuplicateAttributes(schema.OnchainData.NftAttributes)
	if err != nil {
		return false, err
	}
	// Validate Onchain Token Attributes
	_, err = HasDuplicateAttributes(schema.OnchainData.TokenAttributes)
	if err != nil {
		return false, err
	}
	// Validate NFT Attributes Value
	_, err = HasDuplicateNftAttributesValue(schema.OnchainData.NftAttributesValue)
	if err != nil {
		return false, err
	}
	return true, nil
}
