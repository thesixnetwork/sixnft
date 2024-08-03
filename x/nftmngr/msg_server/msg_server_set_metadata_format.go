package msg_server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msg_server) SetMetadataFormat(goCtx context.Context, msg *types.MsgSetMetadataFormat) (*types.MsgSetMetadataFormatResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	k.SetMetadataFormatKeeper(ctx, from, msg.SchemaCode, msg.NewFormat)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetMetadataFormat,
			sdk.NewAttribute(types.EventTypeSetMetadataFormat, msg.SchemaCode),
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NewFormat),
		),
	})

	return &types.MsgSetMetadataFormatResponse{
		SchemaCode: msg.SchemaCode,
		NewFormat:  msg.NewFormat,
	}, nil
}
