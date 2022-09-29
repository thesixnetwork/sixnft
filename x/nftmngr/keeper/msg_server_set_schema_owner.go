package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) SetSchemaOwner(goCtx context.Context, msg *types.MsgSetSchemaOwner) (*types.MsgSetSchemaOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.Keeper.GetNFTSchema(ctx, msg.SchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.SchemaCode)
	}

	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	new_owner, err := sdk.AccAddressFromBech32(msg.NewOwner)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccdress, msg.Creator)
	}

	schema.Owner = new_owner.String()

	k.Keeper.SetNFTSchema(ctx, schema)

	return &types.MsgSetSchemaOwnerResponse{
		SchemaCode: msg.SchemaCode,
		NewOwner:   msg.NewOwner,
	}, nil
}
