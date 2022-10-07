package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

type ReadAttribute struct {
	AttributeName       string
	AttributeDefinition *types.AttributeDefinition
	AttributeLocation   types.AttributeLocation
	AttrributeIndex     int
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
	for i, nftAttribute := range schema.OnchainData.NftAttributes {
		mapReadAttribute[nftAttribute.Name] = ReadAttribute{
			AttributeName:       nftAttribute.Name,
			AttributeDefinition: nftAttribute,
			AttributeLocation:   types.AttributeLocation_NFT_ATTRIBUTE,
			AttrributeIndex:     i,
		}
	}

	for i, tokenAttribute := range schema.OnchainData.TokenAttributes {
		mapReadAttribute[tokenAttribute.Name] = ReadAttribute{
			AttributeName:       tokenAttribute.Name,
			AttributeDefinition: tokenAttribute,
			AttributeLocation:   types.AttributeLocation_TOKEN_ATTRIBUTE,
			AttrributeIndex:     i,
		}
	}

	// loop over msg.AttributeNames
	for _, attributeName := range msg.AttributeNames {
		// check if attribute is exist in mapReadAttribute
		if _, ok := mapReadAttribute[attributeName]; !ok {
			return nil, sdkerrors.Wrap(types.ErrAttributeDoesNotExists, attributeName)
		}
		readAttributeDef := mapReadAttribute[attributeName]
		if readAttributeDef.AttributeLocation == types.AttributeLocation_NFT_ATTRIBUTE {
			schema.OnchainData.NftAttributes[readAttributeDef.AttrributeIndex].HiddenToMarketplace = !msg.Show
		} else {
			schema.OnchainData.TokenAttributes[readAttributeDef.AttrributeIndex].HiddenToMarketplace = !msg.Show
		}
	}

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeShowAttribute,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyShowAttributeResult, "success"),
		),
	)

	k.Keeper.SetNFTSchema(ctx, schema)

	return &types.MsgShowAttributesResponse{
		NftSchema:           schema.Code,
	}, nil
}
