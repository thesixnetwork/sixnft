package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) HidddenAttributes(goCtx context.Context, msg *types.MsgHidddenAttributes) (*types.MsgHidddenAttributesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retreive schema
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	// Check if creator is the owner of the schema
	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	// Check if list of hidden attributes already exists
	if schema.HiddenAttributes == nil {
		schema.HiddenAttributes = make([]string, 0)
	}

	// check if attribute already hidden
	for _, attribute := range schema.HiddenAttributes {
		if attribute == msg.AttibuteName {
			return nil, sdkerrors.Wrap(types.ErrAttributeAlreadyHidden, msg.AttibuteName)
		}
	}

	for _, hiddenAttribute := range schema.HiddenAttributes {
		schema.HiddenAttributes = append(schema.HiddenAttributes, hiddenAttribute)
	}

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeHiddenAttribute,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyHiddenAttributeResult, msg.AttibuteName),
			sdk.NewAttribute(types.AttributeKeyHiddenAttributeResult, "success"),
		),
	)

	k.Keeper.SetNFTSchema(ctx, schema)

	return &types.MsgHidddenAttributesResponse{}, nil
}
