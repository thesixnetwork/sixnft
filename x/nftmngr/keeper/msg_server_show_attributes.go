package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) ShowAttributes(goCtx context.Context, msg *types.MsgShowAttributes) (*types.MsgShowAttributesResponse, error) {
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

	// Check if list of hidden attributes already exists and remove attribute from list
	if schema.HiddenAttributes == nil {
		schema.HiddenAttributes = make([]string, 0)
	} else {
		for i, attribute := range schema.HiddenAttributes {
			if attribute == msg.AttributeName {
				schema.HiddenAttributes = append(schema.HiddenAttributes[:i], schema.HiddenAttributes[i+1:]...)
			}
		}
	}
	
	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeShowAttribute,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyShowAttributeResult, msg.AttributeName),
			sdk.NewAttribute(types.AttributeKeyShowAttributeResult, "success"),
		),
	)

	k.Keeper.SetNFTSchema(ctx, schema)


	return &types.MsgShowAttributesResponse{
		NftSchema: schema.Code,
		ShowedAttributeName: msg.AttributeName,
	}, nil
}
