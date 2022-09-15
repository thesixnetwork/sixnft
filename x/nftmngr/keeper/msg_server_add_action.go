package keeper

import (
	"context"
	"encoding/base64"

	"sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddAction(goCtx context.Context, msg *types.MsgAddAction) (*types.MsgAddActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// structure for new action
	var new_action types.Action
	
	//decode base64 string to bytes
	input_action, err := base64.StdEncoding.DecodeString(msg.Base64NewAction)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}

	//unmarshal bytes to Action structure
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(input_action, &new_action)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	//get existing action in schema
	schema, schemaFound := k.Keeper.GetNFTSchema(ctx, msg.GetCode())
	if !schemaFound {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.GetCode())
	}

	//validate Action data
	err = k.ValidateAction(&new_action, &schema)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	// append new action
	schema.OnchainData.Actions = append(schema.OnchainData.Actions, &new_action)

	// save schema
	k.Keeper.SetNFTSchema(ctx, schema)


	return &types.MsgAddActionResponse{
		Code:        msg.GetCode(),
		Name:        schema.Name,
		OnchainData: schema.OnchainData,
	}, nil
}

//validate Action data
func (k Keeper) ValidateAction(action *types.Action, schema *types.NFTSchema) error {
	// Onchain Data Nft Actions Map
	mapNftActions := CreateActionMap(schema.OnchainData.Actions)
	// check if action name is unique
	if _, found := mapNftActions[action.Name]; found {
		return sdkerrors.Wrap(types.ErrActionAlreadyExists, action.Name)
	}

	//validate action struct
	if action.Name == "" {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action name is empty")
	}
	if action.Desc == "" {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action description is empty")
	}
	if action.When == "" {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action type is empty")
	}
	//validate array of action.Then is not empty
	if len(action.Then) == 0 {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action.Then is empty")
	}

	return nil
}