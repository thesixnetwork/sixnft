package msg_server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msg_server) ChangeSchemaOwner(goCtx context.Context, msg *types.MsgChangeSchemaOwner) (*types.MsgChangeSchemaOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retreive schema data
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	// Check if the creator is the same as the current owner
	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	_, err := sdk.AccAddressFromBech32(msg.NewOwner)
	if err != nil {
		return nil, err
	}

	// Change the owner
	schema.Owner = msg.NewOwner

	// Save the schema
	k.Keeper.SetNFTSchema(ctx, schema)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSchemaOwnerChanged,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyNewOwner, msg.NewOwner),
		),
	})

	return &types.MsgChangeSchemaOwnerResponse{
		NftSchemaCode: msg.NftSchemaCode,
		NewOwner:      msg.NewOwner,
	}, nil
}
