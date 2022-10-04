package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) RemoveSystemActioner(goCtx context.Context, msg *types.MsgRemoveSystemActioner) (*types.MsgRemoveSystemActionerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieve the schema
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}
	// Check if the creator is the owner of the schema
	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	// Create map of system actioners
	mapSystemActioners := make(map[string]bool)
	for _, systemActioner := range schema.SystemActioners {
		mapSystemActioners[systemActioner] = true
	}
	// Check if system actioner exists
	if _, ok := mapSystemActioners[msg.Actioner]; !ok {
		return nil, sdkerrors.Wrap(types.ErrSystemActionerDoesNotExists, msg.Actioner)
	}
	// Remove system actioner from schema
	for i, systemActioner := range schema.SystemActioners {
		if systemActioner == msg.Actioner {
			schema.SystemActioners = append(schema.SystemActioners[:i], schema.SystemActioners[i+1:]...)
			break
		}
	}
	// Set schema
	k.Keeper.SetNFTSchema(ctx, schema)

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRemoveSystemActioner,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyActioner, msg.Actioner),
		),
	)

	return &types.MsgRemoveSystemActionerResponse{
		NftSchemaCode: msg.NftSchemaCode,
		Actioner:      msg.Actioner,
	}, nil
}
