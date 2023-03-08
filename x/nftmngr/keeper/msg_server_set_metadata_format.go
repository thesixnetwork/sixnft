package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) SetMetadataFormat(goCtx context.Context, msg *types.MsgSetMetadataFormat) (*types.MsgSetMetadataFormatResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.Keeper.GetNFTSchema(ctx, msg.SchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.SchemaCode)
	}

	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	schema.OriginData.MetadataFormat = msg.NewFormat

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetMetadataFormat,
			sdk.NewAttribute(types.EventTypeSetMetadataFormat, msg.SchemaCode),
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NewFormat),
			sdk.NewAttribute(types.AttributeKeySetRetrivalResult, "success"),
		),
	})

	k.Keeper.SetNFTSchema(ctx, schema)

	return &types.MsgSetMetadataFormatResponse{
		SchemaCode: msg.SchemaCode,
		NewFormat:  msg.NewFormat,
	}, nil
}
