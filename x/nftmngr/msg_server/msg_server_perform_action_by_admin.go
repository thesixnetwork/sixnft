package msg_server

import (
	"context"
	"fmt"

	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msg_server) PerformActionByAdmin(goCtx context.Context, msg *types.MsgPerformActionByAdmin) (*types.MsgPerformActionByAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
  
  fmt.Printf("ADDRESS SDK AccAddress %v \n", from)
  fmt.Printf("ADDRESS String %v \n", from.String())

	// Emit events on metadata change
	changeList, err := k.ActionByAdmin(ctx, from, msg.NftSchemaCode, msg.TokenId, msg.Action, msg.RefId, msg.Parameters)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.EventMessage, types.EventTypeRunAction),
			sdk.NewAttribute(types.AttributeKeyRunActionResult, "success"),
			// Emit change list attributes
			sdk.NewAttribute(types.AttributeKeyRunActionChangeList, string(changeList)),
		),
	)

	return &types.MsgPerformActionByAdminResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}
