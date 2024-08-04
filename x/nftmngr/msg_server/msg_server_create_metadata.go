package msg_server

import (
	"context"
	"encoding/base64"

	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msg_server) CreateMetadata(goCtx context.Context, msg *types.MsgCreateMetadata) (*types.MsgCreateMetadataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	newMetadata, err := base64.StdEncoding.DecodeString(msg.Base64NFTData)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}

	metadata := types.NftData{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(newMetadata, &metadata)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	err = k.Keeper.CreateNewMetadataKeeper(ctx, from, msg.NftSchemaCode, msg.TokenId, metadata)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateMetadata,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyCreateMetaDataTokenID, msg.TokenId),
			sdk.NewAttribute(types.AttributeKeyCreateMetaDataResult, "success"),
		),
	})

	return &types.MsgCreateMetadataResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}