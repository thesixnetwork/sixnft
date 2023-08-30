package keeper

import (
	"context"
	"encoding/json"
	"strconv"

	// "strconv"
	"time"

	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) PerformActionByAdmin(goCtx context.Context, msg *types.MsgPerformActionByAdmin) (*types.MsgPerformActionByAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// ** SCHEMA LAYER **
	// check if schema exists
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	tokenData, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, msg.TokenId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, "Schema: "+msg.NftSchemaCode+" TokenID: "+msg.TokenId)
	}

	// check if executor is authorized to perform action
	var isOwner bool
	if msg.Creator == schema.Owner {
		isOwner = true
	}

	// if not owner, check if executor is authorized to perform action
	if !isOwner {

		_, isFound := k.GetActionExecutor(
			ctx,
			msg.NftSchemaCode,
			msg.Creator,
		)

		if !isFound {
			return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
		}
	}

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

	// Check if AllowedAction is for system // if yes whick means only oracle request can perform this action
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

	var list_schema_attributes_ []*types.SchemaAttribute
	var map_converted_schema_attributes []*types.NftAttributeValue

	// get schema attributes and convert to NFtAttributeValue
	all_schema_attributes := k.GetAllSchemaAttribute(ctx)

	attributeMap := make(map[string]bool)

	for _, schema_attribute := range all_schema_attributes {
		if schema_attribute.NftSchemaCode != msg.NftSchemaCode {
			continue
		}
		// Check if the attribute has already been added
		if attributeMap[schema_attribute.Name] {
			continue
		}
		// Add the attribute to the list of schema attributes
		list_schema_attributes_ = append(list_schema_attributes_, &schema_attribute)

		// Add the attribute to the map
		attributeMap[schema_attribute.Name] = true

		nftAttributeValue_ := ConverSchemaAttributeToNFTAttributeValue(&schema_attribute)
		map_converted_schema_attributes = append(map_converted_schema_attributes, nftAttributeValue_)
	}

	// ** META path ../types/meta.go **
	meta := types.NewMetadata(&schema, &tokenData, schema.OriginData.AttributeOverriding, map_converted_schema_attributes)
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

	for _, change := range meta.ChangeList {
		val, found := k.Keeper.GetSchemaAttribute(ctx, msg.NftSchemaCode, change.Key)
		if found {
			switch val.DataType {
			case "string":
				val.CurrentValue.Value = &types.SchemaAttributeValue_StringAttributeValue{
					StringAttributeValue: &types.StringAttributeValue{
						Value: change.NewValue,
					},
				}
			case "boolean":
				boolValue, err := strconv.ParseBool(change.NewValue)
				if err != nil {
					return nil, err
				}
				val.CurrentValue.Value = &types.SchemaAttributeValue_BooleanAttributeValue{
					BooleanAttributeValue: &types.BooleanAttributeValue{
						Value: boolValue,
					},
				}
			case "number":
				uintValue, err := strconv.ParseUint(change.NewValue, 10, 64)
				if err != nil {
					return nil, err
				}
				val.CurrentValue.Value = &types.SchemaAttributeValue_NumberAttributeValue{
					NumberAttributeValue: &types.NumberAttributeValue{
						Value: uintValue,
					},
				}
			case "float":
				floatValue, err := strconv.ParseFloat(change.NewValue, 64)
				if err != nil {
					return nil, err
				}
				val.CurrentValue.Value = &types.SchemaAttributeValue_FloatAttributeValue{
					FloatAttributeValue: &types.FloatAttributeValue{
						Value: floatValue,
					},
				}
			default:
				return nil, sdkerrors.Wrap(types.ErrParsingAttributeValue, val.DataType)
			}

			k.SetSchemaAttribute(ctx, val)
		}
	}

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
