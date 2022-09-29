package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) SetBaseUri(goCtx context.Context, msg *types.MsgSetBaseUri) (*types.MsgSetBaseUriResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.Keeper.GetNFTSchema(ctx, msg.Code)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.Code)
	}

	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	schema.OriginData.OriginBaseUri = msg.NewBaseUri

	k.Keeper.SetNFTSchema(ctx, schema)

	return &types.MsgSetBaseUriResponse{
		Code: msg.Code,
		Uri:  msg.NewBaseUri,
	}, nil
}
