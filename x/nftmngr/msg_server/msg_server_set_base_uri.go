package msg_server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msg_server) SetBaseUri(goCtx context.Context, msg *types.MsgSetBaseUri) (*types.MsgSetBaseUriResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	k.SetBaseURIKeeper(ctx, from, msg.Code, msg.NewBaseUri)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetBaseURI,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.Code),
			sdk.NewAttribute(types.AttributeKeySetBaseURI, msg.NewBaseUri),
		),
	})

	return &types.MsgSetBaseUriResponse{
		Code: msg.Code,
		Uri:  msg.NewBaseUri,
	}, nil
}
