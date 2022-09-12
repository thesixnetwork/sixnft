package keeper

import (
	"context"
	"strconv"

	"sixnft/x/nftoracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMintRequest(goCtx context.Context, msg *types.MsgCreateMintRequest) (*types.MsgCreateMintRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if nft_schema_code exists
	_, found := k.nftmngrKeeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNFTSchemaNotFound, msg.NftSchemaCode)
	}
	// Check if the token is already minted
	_, found = k.nftmngrKeeper.GetNftData(ctx, msg.NftSchemaCode, msg.TokenId)
	if found {
		return nil, sdkerrors.Wrap(types.ErrMetadataAlreadyExists, msg.NftSchemaCode)
	}

	id_ := k.Keeper.AppendMintRequest(ctx, types.MintRequest{
		NftSchemaCode:   msg.NftSchemaCode,
		TokenId:         msg.TokenId,
		RequiredConfirm: msg.RequiredConfirm,
	})

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMintRequestCreated,
			sdk.NewAttribute(types.AttributeKeyMintRequestID, strconv.FormatUint(id_, 10)),
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyTokenID, msg.TokenId),
			sdk.NewAttribute(types.AttributeKeyRequiredConfirm, strconv.FormatUint(msg.RequiredConfirm, 10)),
		),
	})

	return &types.MsgCreateMintRequestResponse{
		Id:            id_,
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}
