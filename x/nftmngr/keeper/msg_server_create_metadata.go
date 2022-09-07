package keeper

import (
	"context"
	"encoding/base64"

	"sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMetadata(goCtx context.Context, msg *types.MsgCreateMetadata) (*types.MsgCreateMetadataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	metadata, err := base64.StdEncoding.DecodeString(msg.Base64NFTData)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}
	data := types.NftData{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(metadata, &data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}
	// Get nft schema from store
	schema, schemaFound := k.Keeper.GetNFTSchema(ctx, data.NftSchemaCode)
	// Check if the schema already exists
	if !schemaFound {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, data.NftSchemaCode)
	}
	// Validate Schema Message and return error if not valid
	valid, err := k.ValidateNFTData(&data, &schema)
	_ = valid
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}
	// Add attributes from schema to metadata onchain attributes
	for _, attribute := range schema.OnchainData.NftAttributesValue {
		data.OnchainAttributes = append(data.OnchainAttributes, attribute)
	}

	// Check if the data already exists
	_, dataFound := k.Keeper.GetNftData(ctx, data.NftSchemaCode, data.TokenId)
	if dataFound {
		return nil, sdkerrors.Wrap(types.ErrMetadataAlreadyExists, data.NftSchemaCode)
	}
	// Add the data to the store
	k.Keeper.SetNftData(ctx, data)

	return &types.MsgCreateMetadataResponse{
		NftSchemaCode: data.NftSchemaCode,
		TokenId:       data.TokenId,
	}, nil
}

// Validate NFT Data
func (k msgServer) ValidateNFTData(data *types.NftData, schema *types.NFTSchema) (bool, error) {
	// Validate Onchain Attributes Value
	_, err := HasDuplicateNftAttributesValue(data.OnchainAttributes)
	if err != nil {
		return false, err
	}
	// Validate Origin Attributes Value
	_, err = HasDuplicateNftAttributesValue(data.OriginAttributes)
	if err != nil {
		return false, err
	}
	// Validate if attributes have the same type for NFT Attributes
	_, err = HasSameType(schema.OriginData.OriginAttributes, schema.OnchainData.NftAttributes)
	if err != nil {
		return false, err
	}
	// Validate if attributes have the same type for Token Attributes
	_, err = HasSameType(schema.OriginData.OriginAttributes, schema.OnchainData.TokenAttributes)
	if err != nil {
		return false, err
	}
	// Validate if origin attributes have the same type as the schema
	_, err = HasSameTypeAsSchema(schema.OriginData.OriginAttributes, data.OriginAttributes)
	if err != nil {
		return false, err
	}
	// Validate if onchain token attributes have the same type as the schema
	_, err = HasSameTypeAsSchema(schema.OnchainData.TokenAttributes, data.OnchainAttributes)
	if err != nil {
		return false, err
	}
	// Validate if onchain nft attributes have the same type as the schema
	_, err = HasSameTypeAsSchema(schema.OnchainData.NftAttributes, data.OnchainAttributes)
	if err != nil {
		return false, err
	}
	return true, nil
}
