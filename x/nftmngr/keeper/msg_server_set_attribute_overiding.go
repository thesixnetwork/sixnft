package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) SetAttributeOveriding(goCtx context.Context, msg *types.MsgSetAttributeOveriding) (*types.MsgSetAttributeOveridingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.Keeper.GetNFTSchema(ctx, msg.SchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.SchemaCode)
	}

	if msg.Creator != schema.Owner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	// int32 to string
	NewOveridingType := strconv.Itoa(int(msg.NewOveridingType))
	
	switch msg.NewOveridingType {
	case 0:
		schema.OriginData.AttributeOverriding = types.AttributeOverriding_ORIGIN
	case 1:
		schema.OriginData.AttributeOverriding = types.AttributeOverriding_ORIGIN
	default:
		return nil, sdkerrors.Wrap(types.ErrAttributeOptionDoesNotExists, NewOveridingType)
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetAttributeOveriding,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.SchemaCode),
			sdk.NewAttribute(types.AttributeKeySetAttributeOverride, schema.OriginData.AttributeOverriding.String()),
			sdk.NewAttribute(types.AttributeKeySetAttributeOverrideResult, "success"),
		),
	})
	
	k.Keeper.SetNFTSchema(ctx, schema)

	return &types.MsgSetAttributeOveridingResponse{}, nil
}
