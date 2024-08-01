package msg_server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msg_server) SetBaseUri(goCtx context.Context, msg *types.MsgSetBaseUri) (*types.MsgSetBaseUriResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.Keeper.GetNFTSchema(ctx, msg.Code)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.Code)
	}

	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	schema.OriginData.OriginBaseUri = msg.NewBaseUri

	k.Keeper.SetNFTSchema(ctx, schema)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetBaseURI,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.Code),
			sdk.NewAttribute(types.AttributeKeySetBaseURI, msg.NewBaseUri),
			sdk.NewAttribute(types.AttributeKeySetBaseURIResult, "success"),
		),
	})

	return &types.MsgSetBaseUriResponse{
		Code: msg.Code,
		Uri:  msg.NewBaseUri,
	}, nil
}
