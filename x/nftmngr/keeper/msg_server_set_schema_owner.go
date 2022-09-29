package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) SetSchemaOwner(goCtx context.Context, msg *types.MsgSetSchemaOwner) (*types.MsgSetSchemaOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.Keeper.GetNFTSchema(ctx, msg.Code)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.Code)
	}

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetSchemaOwnerResponse{}, nil
}
