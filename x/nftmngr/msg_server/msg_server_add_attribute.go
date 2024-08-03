package msg_server

import (
	"context"
	"encoding/base64"

	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msg_server) AddAttribute(goCtx context.Context, msg *types.MsgAddAttribute) (*types.MsgAddAttributeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	var new_add_attribute types.AttributeDefinition

	input_addribute, err := base64.StdEncoding.DecodeString(msg.Base64NewAttriuteDefenition)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}

	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(input_addribute, &new_add_attribute)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	err = k.AddAttributeKeeper(ctx, from, msg.Code, new_add_attribute, msg.Location)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddAttribute,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.Code),
			sdk.NewAttribute(types.AttributeKeyAddAttributeName, new_add_attribute.Name),
			sdk.NewAttribute(types.AttributeKeyAddAttributeLocation, types.AttributeLocation.String(msg.Location)),
		),
	})

	return &types.MsgAddAttributeResponse{
		Code: msg.Code,
		Name: new_add_attribute.Name,
	}, nil
}
