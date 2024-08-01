package msg_server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/keeper"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msg_server) ResyncAttributes(goCtx context.Context, msg *types.MsgResyncAttributes) (*types.MsgResyncAttributesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retreive schema
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}
	// Check if creator is owner of schema
	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	// Retrieve NFT Data
	nftData, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, msg.TokenId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNftDataDoesNotExists, msg.TokenId)
	}

	// Create map of existing attribute in nftdata
	mapExistingAttributes := make(map[string]bool)
	for _, attribute := range nftData.OnchainAttributes {
		mapExistingAttributes[attribute.Name] = true
	}

	// Loop over schema.TokenAttributes to check if exists in nftdata
	for _, attribute := range schema.OnchainData.TokenAttributes {
		if _, ok := mapExistingAttributes[attribute.Name]; !ok {
			if attribute.DefaultMintValue == nil {
				return nil, sdkerrors.Wrap(types.ErrNoDefaultValue, attribute.Name)
			}
			// Add attribute to nftdata with default value
			nftData.OnchainAttributes = append(nftData.OnchainAttributes,
				keeper.NewNFTAttributeValueFromDefaultValue(attribute.Name, attribute.DefaultMintValue))
		}
	}

	// Emit Event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeResyncAttributes,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyTokenId, msg.TokenId),
		),
	)

	// Set nftdata
	k.Keeper.SetNftData(ctx, nftData)

	return &types.MsgResyncAttributesResponse{}, nil
}
