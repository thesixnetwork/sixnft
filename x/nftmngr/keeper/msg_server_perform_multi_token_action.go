package keeper

import (
	"context"
	"encoding/json"
	"time"

	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) PerformMultiTokenAction(goCtx context.Context, msg *types.MsgPerformMultiTokenAction) (*types.MsgPerformMultiTokenActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	//check if id in msg.TokenId is duplicate
	mapOfTokenId := make(map[string]bool)
	for _, tokenId := range msg.TokenId {
		if _, ok := mapOfTokenId[tokenId]; ok {
			return nil, sdkerrors.Wrap(types.ErrDuplicateInputTokenID, tokenId)
		}
		mapOfTokenId[tokenId] = true
	}

	// ** SCHEMA LAYER **
	// check if schema exists
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	//check if all token exists
	for _, tokenId := range msg.TokenId {
		_, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, "Schema: "+msg.NftSchemaCode+" TokenID: "+tokenId)
		}
	}
	
	// ** This might be different from PerformActionByAdmin but to prevent time consuming process, we will use the same code out of iteration process **
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

	// switch case for action
	// case 1 len(action) == 1 && len(tokenid) == 1
	// case 2 len(action) == 1 && len(tokenid) > 1
	// case 3 len(action) > 1 && len(tokenid) == 1
	// default len(action) > 1 && len(tokenid) > 1
	switch {
	case len(msg.Action) == 1 && len(msg.TokenId) == 1:
		msg_ := &types.MsgPerformActionByAdmin{
			Creator:       msg.Creator,
			NftSchemaCode: msg.NftSchemaCode,
			TokenId:       msg.TokenId[0],
			Action:        msg.Action[0],
			RefId:         msg.RefId,
			Parameters:    msg.Parameters[0].Parameters,
		}
		k.PerformActionByAdmin(goCtx, msg_)
	case len(msg.Action) == 1 && len(msg.TokenId) > 1:
		msg_ := &types.MsgPerformMultiTokenOneAction{
			Creator:       msg.Creator,
			NftSchemaCode: msg.NftSchemaCode,
			TokenId:       msg.TokenId,
			Action:        msg.Action[0],
			RefId:         msg.RefId,
			Parameters:    msg.Parameters[0].Parameters,
		}
		k.PerformMultiTokenOneAction(goCtx, msg_)
	case len(msg.Action) > 1 && len(msg.TokenId) == 1:
		msg_ := &types.MsgPerformMultiTokenMultiAction{
			Creator:       msg.Creator,
			NftSchemaCode: msg.NftSchemaCode,
			TokenId:       msg.TokenId,
			Action:        msg.Action,
			RefId:         msg.RefId,
			Parameters:    msg.Parameters,
		}

		k.PerformMultiTokenMultiAction(goCtx, msg_)
	}

	return &types.MsgPerformMultiTokenActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
		Action:        msg.Action,
	}, nil
}



func (k msgServer) PerformMultiTokenOneAction(goCtx context.Context, msg *types.MsgPerformMultiTokenOneAction) (*types.MsgPerformMultiTokenOneActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	//check if id in msg.TokenId is duplicate
	mapOfTokenId := make(map[string]bool)
	for _, tokenId := range msg.TokenId {
		if _, ok := mapOfTokenId[tokenId]; ok {
			return nil, sdkerrors.Wrap(types.ErrDuplicateInputTokenID, tokenId)
		}
		mapOfTokenId[tokenId] = true
	}

	// ** SCHEMA LAYER **
	// check if schema exists
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	//check if all token exists
	for _, tokenId := range msg.TokenId {
		_, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, "Schema: "+msg.NftSchemaCode+" TokenID: "+tokenId)
		}
	}
	
	// ** This might be different from PerformActionByAdmin but to prevent time consuming process, we will use the same code out of iteration process **
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

	// check if action is disabled
	mapAction := types.Action{}
	for _, action := range schema.OnchainData.Actions {
		if action.Name == msg.Action && action.Disable {
			return nil, sdkerrors.Wrap(types.ErrActionIsDisabled, action.Name)
		}
		if action.Name == msg.Action {
			mapAction = *action
			break
		}
	}

	// Check if action exists
	if mapAction.Name == "" {
		return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, msg.Action)
	}
	
	// Check if AllowedAction is for system
	if mapAction.GetAllowedActioner() == types.AllowedActioner_ALLOWED_ACTIONER_USER_ONLY {
		return nil, sdkerrors.Wrap(types.ErrActionIsForUserOnly, msg.Action)
	}

	// Check if action requires parameters
	param := mapAction.GetParams()
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
			msg.Parameters[i].Value = required_param[i].DefaultValue
		}
	}

	// ** TOKEN DATA LAYER **
	// iterate over token ids
	for _, tokenId := range msg.TokenId {
		tokenData, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, msg.NftSchemaCode) //! This should not happen since we already check if all token exists
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

		// ** META path ../types/meta.go **
		meta := types.NewMetadata(&schema, &tokenData, schema.OriginData.AttributeOverriding)
		meta.SetGetNFTFunction(func(tokenId string) (*types.NftData, error) {
			tokenData, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
			if !found {
				return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, msg.NftSchemaCode)
			}
			return &tokenData, nil
		})

		// utils function
		meta.SetGetBlockTimeFunction(func() time.Time {
			return ctx.BlockTime()
		})
		// utils function
		meta.SetGetBlockHeightFunction(func() int64 {
			return ctx.BlockHeight()
		})

		err := ProcessAction(meta, &mapAction, msg.Parameters)
		if err != nil {
			return nil, err
		}
		// Check if ChangeList is empty, error if empty
		if len(meta.ChangeList) == 0 {
			return nil, sdkerrors.Wrap(types.ErrEmptyChangeList, msg.Action)
		}

		// Update back to nftdata
		k.Keeper.SetNftData(ctx, tokenData)

		// Udpate to target
		// loop over meta.OtherUpdatedTokenDatas
		for _, otherTokenData := range meta.OtherUpdatedTokenDatas {
			k.Keeper.SetNftData(ctx, *otherTokenData)
		}

		// Emit events on metadata change
		// Check action with reference exists
		refId := msg.RefId + "token_id_" + tokenId
		if msg.RefId != "" {

			_, found := k.Keeper.GetActionByRefId(ctx, refId)
			if found {
				return nil, sdkerrors.Wrap(types.ErrRefIdAlreadyExists, refId)
			}

			k.Keeper.SetActionByRefId(ctx, types.ActionByRefId{
				RefId:         refId,
				Creator:       msg.Creator,
				NftSchemaCode: msg.NftSchemaCode,
				TokenId:       tokenId,
				Action:        mapAction.Name,
			})
		}

		// Emit events on metadata change
		// appply change list to nftdata
		changeList, _ := json.Marshal(meta.ChangeList)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(types.EventMessage, types.EventTypeRunAction),
				sdk.NewAttribute(types.AttributeKeyRunActionResult, "success"),
				sdk.NewAttribute(types.AttributeKeyTokenId, tokenId),
				sdk.NewAttribute(types.AttributeKeyRunActionRefId, refId),
				// Emit change list attributes
				sdk.NewAttribute(types.AttributeKeyRunActionChangeList, string(changeList)),
			),
		)

	}

	return &types.MsgPerformMultiTokenOneActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}


func (k msgServer) PerformMultiTokenMultiAction(goCtx context.Context, msg *types.MsgPerformMultiTokenMultiAction) (*types.MsgPerformMultiTokenMultiActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	//check if id in msg.TokenId is duplicate
	mapOfTokenId := make(map[string]bool)
	for _, tokenId := range msg.TokenId {
		if _, ok := mapOfTokenId[tokenId]; ok {
			return nil, sdkerrors.Wrap(types.ErrDuplicateInputTokenID, tokenId)
		}
		mapOfTokenId[tokenId] = true
	}

	// ** SCHEMA LAYER **
	// check if schema exists
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	//check if all token exists
	for _, tokenId := range msg.TokenId {
		_, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, "Schema: "+msg.NftSchemaCode+" TokenID: "+tokenId)
		}
	}
	
	// ** This might be different from PerformActionByAdmin but to prevent time consuming process, we will use the same code out of iteration process **
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

	mapAction := []types.Action{}
	for index, action_ := range msg.Action {
		// check if action is disabled
		for _, _action := range schema.OnchainData.Actions {
			if _action.Name == action_ && _action.Disable {
				return nil, sdkerrors.Wrap(types.ErrActionIsDisabled, _action.Name)
			}
			if _action.Name == action_ {
				mapAction[index] = *_action
				break
			}
		}

		// Check if action exists
		if mapAction[index].Name == "" {
			return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, action_)
		}
		
		// Check if AllowedAction is for system
		if mapAction[index].GetAllowedActioner() == types.AllowedActioner_ALLOWED_ACTIONER_USER_ONLY {
			return nil, sdkerrors.Wrap(types.ErrActionIsForUserOnly, action_)
		}
	}

	for index, params_ := range msg.Parameters {
		// Check if action requires parameters
		param := mapAction[index].GetParams()
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
			if params_.Parameters[i].Name != required_param[i].Name {
				return nil, sdkerrors.Wrap(types.ErrInvalidParameter, "input paramter name is not match to "+required_param[i].Name)
			}
			if params_.Parameters[i].Value == "" {
				params_.Parameters[i].Value = required_param[i].DefaultValue
			}
		}

	}

	// ** TOKEN DATA LAYER **
	// iterate over token ids
	for index, tokenId := range msg.TokenId {
		tokenData, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, msg.NftSchemaCode) //! This should not happen since we already check if all token exists
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

		// ** META path ../types/meta.go **
		meta := types.NewMetadata(&schema, &tokenData, schema.OriginData.AttributeOverriding)
		meta.SetGetNFTFunction(func(tokenId string) (*types.NftData, error) {
			tokenData, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
			if !found {
				return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, msg.NftSchemaCode)
			}
			return &tokenData, nil
		})

		// utils function
		meta.SetGetBlockTimeFunction(func() time.Time {
			return ctx.BlockTime()
		})
		// utils function
		meta.SetGetBlockHeightFunction(func() int64 {
			return ctx.BlockHeight()
		})

		err := ProcessAction(meta, &mapAction[index], msg.Parameters[index].Parameters)
		if err != nil {
			return nil, err
		}
		// Check if ChangeList is empty, error if empty
		if len(meta.ChangeList) == 0 {
			return nil, sdkerrors.Wrap(types.ErrEmptyChangeList, msg.Action[index])
		}

		// Update back to nftdata
		k.Keeper.SetNftData(ctx, tokenData)

		// Udpate to target
		// loop over meta.OtherUpdatedTokenDatas
		for _, otherTokenData := range meta.OtherUpdatedTokenDatas {
			k.Keeper.SetNftData(ctx, *otherTokenData)
		}

		// Emit events on metadata change
		// Check action with reference exists
		refId := msg.RefId + "token_id_" + tokenId
		if msg.RefId != "" {

			_, found := k.Keeper.GetActionByRefId(ctx, refId)
			if found {
				return nil, sdkerrors.Wrap(types.ErrRefIdAlreadyExists, refId)
			}

			k.Keeper.SetActionByRefId(ctx, types.ActionByRefId{
				RefId:         refId,
				Creator:       msg.Creator,
				NftSchemaCode: msg.NftSchemaCode,
				TokenId:       tokenId,
				Action:        mapAction[index].Name,
			})
		}

		// Emit events on metadata change
		// appply change list to nftdata
		changeList, _ := json.Marshal(meta.ChangeList)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(types.EventMessage, types.EventTypeRunAction),
				sdk.NewAttribute(types.AttributeKeyRunActionResult, "success"),
				sdk.NewAttribute(types.AttributeKeyTokenId, tokenId),
				sdk.NewAttribute(types.AttributeKeyRunActionRefId, refId),
				// Emit change list attributes
				sdk.NewAttribute(types.AttributeKeyRunActionChangeList, string(changeList)),
			),
		)

	}

	return &types.MsgPerformMultiTokenMultiActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}
