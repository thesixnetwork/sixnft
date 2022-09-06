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
	// Check if the schema already exists
	schema, schemaFound := k.Keeper.GetNFTSchema(ctx, data.NftSchemaCode)
	_ = schema
	if !schemaFound {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, data.NftSchemaCode)
	}

	data.OnchainAttributes = append(schema.OnchainData.NftAttributeValues)

	fmt.Println("data: ", data)

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
