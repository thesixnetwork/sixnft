package keeper

import (
	"context"
	"encoding/json"
	"strconv"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) PerformActionByAdmin(goCtx context.Context, msg *types.MsgPerformActionByAdmin) (*types.MsgPerformActionByAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	tokenData, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, msg.TokenId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, msg.NftSchemaCode)
	}

	// Create map of existing attribute in nftdata
	mapExistingAttributes := make(map[string]bool)
	for _, attribute := range tokenData.OnchainAttributes {
		mapExistingAttributes[attribute.Name] = true
	}

	// Loop over schema.TokenAttributes to check if exists in nftdata
	for _, attribute := range schema.OnchainData.TokenAttributes {
		if _, ok := mapExistingAttributes[attribute.Name]; !ok {
			if attribute.DefaultMintValue == nil {
				return nil, sdkerrors.Wrap(types.ErrNoDefaultValue, attribute.Name)
			}
			// Add attribute to nftdata with default value
			tokenData.OnchainAttributes = append(tokenData.OnchainAttributes,
				NewNFTAttributeValueFromDefaultValue(attribute.Name, attribute.DefaultMintValue))
		}
	}

	// Map system actioners
	mapSystemActioners := make(map[string]bool)
	for _, systemActioner := range schema.SystemActioners {
		mapSystemActioners[systemActioner] = true
	}

	// Check if Creator is one of system actioners
	if _, ok := mapSystemActioners[msg.Creator]; !ok {
		if msg.Creator != schema.Owner {
			return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
		}
	}
	mapAction := types.Action{}
	for _, action := range schema.OnchainData.Actions {
		if action.Disable {
			return nil, sdkerrors.Wrap(types.ErrActionIsDisabled, action.Name)
		}
		if action.Name == msg.Action {
			mapAction = *action
			break
		}
	}
	// Check if AllowedAction is for system
	if mapAction.GetAllowedActioner() == types.AllowedActioner_ALLOWED_ACTIONER_USER_ONLY {
		return nil, sdkerrors.Wrap(types.ErrActionIsForUserOnly, msg.Action)
	}

	// Check if action requires parameters
	param := mapAction.GetParams()
	// if len(param) != len(msg.Parameters) {
	// 	return nil, sdkerrors.Wrap(types.ErrInvalidParameter, msg.Action)
	// }

	// for i, p := range param {
	// 	if p.Required {
	// 		if params := msg.Parameters[i];
	// 			params.Name != p.Name || params.Value == "" {
	// 			return nil, sdkerrors.Wrap(types.ErrInvalidParameter, p.Name)
	// 		}
	// 	}
	// }


	// array of *ActionParams
	required_param := make([]*types.ActionParams, 0)

	for _, p := range param {
		if p.Required {
			required_param = append(required_param, p)
		}
	}

	if len(required_param) > len(msg.Parameters) {
		return nil, sdkerrors.Wrap(types.ErrInvalidParameter, "Input parameters length is not equal to required parameters length")
	}

	for i := 0; i < len(required_param); i++ {
		if msg.Parameters[i].Name != required_param[i].Name {
			return nil, sdkerrors.Wrap(types.ErrInvalidParameter, "input paramter name is not match to "+required_param[i].Name)
		}
		if msg.Parameters[i].Value == "" {
			msg.Parameters[i].Value = strconv.Itoa(int(required_param[i].DefaultValue))
		}
	}

	meta := types.NewMetadata(&schema, &tokenData, schema.OriginData.AttributeOverriding)

	err := ProcessAction(meta, &mapAction, msg.Parameters)
	if err != nil {
		return nil, err
	}
	// Check if ChangeList is empty, error if empty
	if len(meta.ChangeList) == 0 {
		return nil, sdkerrors.Wrap(types.ErrEmptyChangeList, msg.Action)
	}

	k.Keeper.SetNftData(ctx, tokenData)

	// Check action with reference exists
	if msg.RefId != "" {

		_, found := k.Keeper.GetActionByRefId(ctx, msg.RefId)
		if found {
			return nil, sdkerrors.Wrap(types.ErrRefIdAlreadyExists, msg.RefId)
		}

		k.Keeper.SetActionByRefId(ctx, types.ActionByRefId{
			RefId:         msg.RefId,
			Creator:       msg.Creator,
			NftSchemaCode: msg.NftSchemaCode,
			TokenId:       msg.TokenId,
			Action:        mapAction.Name,
		})
	}
	// Emit events on metadata change
	changeList, _ := json.Marshal(meta.ChangeList)
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
