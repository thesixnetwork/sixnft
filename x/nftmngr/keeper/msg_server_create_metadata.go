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
	fmt.Println("#######################CreateMetadata 1")
	metadata, err := base64.StdEncoding.DecodeString(msg.Base64NFTData)
	fmt.Println("#######################CreateMetadata 2", metadata)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}
	data := types.NftData{}
	fmt.Println("#######################CreateMetadata 3", data)
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(metadata, &data)
	fmt.Println("#######################CreateMetadata 4", data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	fmt.Println("#######################", data.NftSchemaCode)

	// Check if the data already exists
	_, found := k.Keeper.GetNftData(ctx, data.NftSchemaCode, data.TokenId)
	if found {
		return nil, sdkerrors.Wrap(types.ErrMetadataAlreadyExists, data.NftSchemaCode)
	}
	// Add the data to the store
	k.Keeper.SetNftData(ctx, data)

	return &types.MsgCreateMetadataResponse{
		NftSchemaCode: data.NftSchemaCode,
		TokenId:       data.TokenId,
	}, nil
}
