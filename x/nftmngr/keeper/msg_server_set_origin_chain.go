package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) SetOriginChain(goCtx context.Context, msg *types.MsgSetOriginChain) (*types.MsgSetOriginChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.Keeper.GetNFTSchema(ctx, msg.SchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.SchemaCode)
	}

	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	// to uppercase
	chainUpperCase := strings.ToUpper(msg.NewOriginChain)
	schema.OriginData.OriginChain = chainUpperCase

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetBaseURI,
			sdk.NewAttribute(types.EventTypeSetOriginChain, msg.SchemaCode),
			sdk.NewAttribute(types.AttributeKeySetOriginChain, msg.NewOriginChain),
			sdk.NewAttribute(types.AttributeKeySetOriginChainResult, "success"),
		),
	})

	k.Keeper.SetNFTSchema(ctx, schema)

	return &types.MsgSetOriginChainResponse{
		SchemaCode:     msg.SchemaCode,
		NewOriginChain: msg.NewOriginChain,
	}, nil
}
