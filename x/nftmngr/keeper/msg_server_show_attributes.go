package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

type ReadAttribute struct {
	AttributeName              string
	OnchainAttributeDefinition *types.AttributeDefinition
	AttributeLocation          types.AttributeLocation
	AttrributeIndex            int
}

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

	// Create map of nftattributes and tokenattributes ni ReadAttribute struct
	mapReadAttribute := make(map[string]ReadAttribute)

	for i, tokenAttribute := range schema.OnchainData.TokenAttributes {
		mapReadAttribute[tokenAttribute.Name] = ReadAttribute{
			AttributeName:              tokenAttribute.Name,
			OnchainAttributeDefinition: tokenAttribute,
			AttributeLocation:          types.AttributeLocation_ATTRIBUTE_OF_TOKEN,
			AttrributeIndex:            i,
		}
	}

	// loop over msg.AttributeNames
	for _, attributeName := range msg.AttributeNames {
		// find attribute in schema/attributes
		valFound, fount := k.GetSchemaAttribute(ctx, msg.NftSchemaCode, attributeName)
		switch fount {
		case true:
			k.SetSchemaAttribute(ctx, types.SchemaAttribute{
				NftSchemaCode:       valFound.NftSchemaCode,
				Name:                valFound.Name,
				DataType:            valFound.DataType,
				Required:            valFound.Required,
				DisplayOption:       valFound.DisplayOption,
				Creator:             valFound.Creator,
				HiddenToMarketplace: !msg.Show,
				DisplayValueField:   valFound.DisplayValueField,
				HiddenOveride:       valFound.HiddenOveride,
				CurrentValue:        valFound.CurrentValue,
			})
			// emit events
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeShowAttribute,
					sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
					sdk.NewAttribute(types.AttributeKeyShowAttributeResult, "success"),
				),
			)
		case false:
			// check if attribute is exist in mapReadAttribute
			if _, ok := mapReadAttribute[attributeName]; !ok {
				return nil, sdkerrors.Wrap(types.ErrAttributeDoesNotExists, attributeName)
			}
			readAttributeDef := mapReadAttribute[attributeName]
			schema.OnchainData.TokenAttributes[readAttributeDef.AttrributeIndex].HiddenToMarketplace = !msg.Show

			k.Keeper.SetNFTSchema(ctx, schema)
			// emit events
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeShowAttribute,
					sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
					sdk.NewAttribute(types.AttributeKeyShowAttributeResult, "success"),
				),
			)

		default:
			return nil, sdkerrors.Wrap(types.ErrAttributeDoesNotExists, attributeName)
		}
	}

	return &types.MsgShowAttributesResponse{NftSchema: schema.Code}, nil
}
