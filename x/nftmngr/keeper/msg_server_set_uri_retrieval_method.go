package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) SetUriRetrievalMethod(goCtx context.Context, msg *types.MsgSetUriRetrievalMethod) (*types.MsgSetUriRetrievalMethodResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.Keeper.GetNFTSchema(ctx, msg.SchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.SchemaCode)
	}

	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	if msg.NewMethod == 0 {
		schema.OriginData.UriRetrievalMethod = types.URIRetrievalMethod_BASE
	} else if msg.NewMethod == 1 {
		schema.OriginData.UriRetrievalMethod = types.URIRetrievalMethod_TOKEN
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetBaseURI,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.SchemaCode),
			sdk.NewAttribute(types.AttributeKeySetRetrievalMethod, schema.OriginData.UriRetrievalMethod.String()),
			sdk.NewAttribute(types.AttributeKeySetRetrivalResult, "success"),
		),
	})

	k.Keeper.SetNFTSchema(ctx, schema)

	return &types.MsgSetUriRetrievalMethodResponse{
		SchemaCode: msg.SchemaCode,
		NewMethod:  schema.OriginData.UriRetrievalMethod.String(),
	}, nil
}
