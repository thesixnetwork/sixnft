package msg_server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msg_server) SetOriginContract(goCtx context.Context, msg *types.MsgSetOriginContract) (*types.MsgSetOriginContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.Keeper.GetNFTSchema(ctx, msg.SchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.SchemaCode)
	}

	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	schema.OriginData.OriginContractAddress = msg.NewContractAddress

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetOriginContract,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.SchemaCode),
			sdk.NewAttribute(types.AttributeKeySetOriginContract, msg.NewContractAddress),
			sdk.NewAttribute(types.AttributeKeySetOriginContractResult, "success"),
		),
	})

	k.Keeper.SetNFTSchema(ctx, schema)
	return &types.MsgSetOriginContractResponse{
		SchemaCode:         msg.SchemaCode,
		NewContractAddress: msg.NewContractAddress,
	}, nil
}
