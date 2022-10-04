package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) AddSystemActioner(goCtx context.Context, msg *types.MsgAddSystemActioner) (*types.MsgAddSystemActionerResponse, error) {
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
	// Create a map of syste actioners
	mapSystemActioners := make(map[string]bool)
	for _, systemActioner := range schema.SystemActioners {
		mapSystemActioners[systemActioner] = true
	}
	// Check if system actioner already exists
	if _, ok := mapSystemActioners[msg.Actioner]; ok {
		return nil, sdkerrors.Wrap(types.ErrSystemActionerAlreadyExists, msg.Actioner)
	}
	// Add system actioner to schema
	schema.SystemActioners = append(schema.SystemActioners, msg.Actioner)
	// Set schema
	k.Keeper.SetNFTSchema(ctx, schema)

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddSystemActioner,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyActioner, msg.Actioner),
		),
	)

	return &types.MsgAddSystemActionerResponse{
		NftSchemaCode: msg.NftSchemaCode,
		Actioner:      msg.Actioner,
	}, nil
}
