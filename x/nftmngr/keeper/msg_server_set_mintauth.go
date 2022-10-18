package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetMintauth(goCtx context.Context, msg *types.MsgSetMintauth) (*types.MsgSetMintauthResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get nft schema from store
	schema, schemaFound := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	// Check if the schema already exists
	if !schemaFound {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	// Swith location of attribute
	switch msg.AuthorizeTo {
	case types.AuthorizeTo_SYSTEM:
		// append new nft_attributes to array of OnchainData.NftAttributes
		schema.MintAuthorization = types.KeyMintPermissionOnlySystem
	case types.AuthorizeTo_ALL:
		// append new token_attributes to array of OnchainData.TokenAttributes
		schema.MintAuthorization = types.KeyMintPermissionAll
		// end the case
	}

	k.Keeper.SetNFTSchema(ctx, schema)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetMintAuth,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeAutorizeTo, schema.MintAuthorization),
			sdk.NewAttribute(types.AttributeKeySetMinAuthResult, "success"),
		),
	})

	return &types.MsgSetMintauthResponse{}, nil
}
