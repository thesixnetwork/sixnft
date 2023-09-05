package keeper

import (
	"context"
	"encoding/base64"
	"strconv"

	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// const (
// 	// AttributeName regular expression
// 	RegxAttributeName = `^[a-z]{1}[a-z0-9_]*[a-z0-9]{1}$`
// 	RegxActionName    = `^[A-Za-z]{1}[A-Za-z0-9_]*[A-Za-z0-9]{1}$`
// 	//regexp.MatchString(`^[a-z]{1}[a-z0-9_]*[a-z0-9]{1}$`, "user_name9")
// )

func (k msgServer) CreateNFTSchema(goCtx context.Context, msg *types.MsgCreateNFTSchema) (*types.MsgCreateNFTSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	jsonSchema, err := base64.StdEncoding.DecodeString(msg.NftSchemaBase64)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}
	schema_input := types.NFTSchemaINPUT{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(jsonSchema, &schema_input)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingSchemaMessage, err.Error())
	}
	schema_input.Owner = msg.Creator
	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	// Validate Schema Message and return error if not valid
	valid, err := ValidateNFTSchema(&schema_input)
	_ = valid
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingNFTSchema, err.Error())
	}
	// if mint_authorization is empty then set system to default
	if len(schema_input.MintAuthorization) == 0 || schema_input.MintAuthorization != types.KeyMintPermissionOnlySystem && schema_input.MintAuthorization != types.KeyMintPermissionAll {
		schema_input.MintAuthorization = types.KeyMintPermissionOnlySystem
	}
	// Check if the schema_input already exists
	_, found := k.Keeper.GetNFTSchema(ctx, schema_input.Code)
	if found {
		return nil, sdkerrors.Wrap(types.ErrSchemaAlreadyExists, schema_input.Code)
	}
	foundOrganization, organizationName := GetOrganizationFromSchemaCode(schema_input.Code)
	// If there is organization in schema_input code, check if the organization exists
	if foundOrganization {
		storedOrganization, found := k.Keeper.GetOrganization(ctx, organizationName)
		if found {
			// Check owner of organization
			if storedOrganization.Owner != msg.Creator {
				return nil, sdkerrors.Wrap(types.ErrOrganizationOwner, msg.Creator)
			}
		} else {
			// Store organization
			k.Keeper.SetOrganization(ctx, types.Organization{
				Owner: msg.Creator,
				Name:  organizationName,
			})
		}

	}
	_ = MergeAllAttributesAndAlterOrderIndex(schema_input.OriginData.OriginAttributes, schema_input.OnchainData.TokenAttributes)

	// parse schema_input to NFTSchema
	schema := types.NFTSchema{
		Code:        schema_input.Code,
		Name:        schema_input.Name,
		Owner:       schema_input.Owner,
		Description: schema_input.Description,
		OriginData:  schema_input.OriginData,
		OnchainData: &types.OnChainData{
			TokenAttributes: schema_input.OnchainData.TokenAttributes,
			Actions:         schema_input.OnchainData.Actions,
			Status:          schema_input.OnchainData.Status,
		},
		IsVerified:        schema_input.IsVerified,
		MintAuthorization: schema_input.MintAuthorization,
	}

	// loop over SchemaAttribute and add to nftmngr/code/name
	for _, scheamDefaultMintAttribute := range schema_input.OnchainData.SchemaAttributes {
		// parse DefaultMintValue to SchemaAttributeValue
		schmaAttributeValue, err := ConvertDefaultMintValueToSchemaAttributeValue(scheamDefaultMintAttribute.DefaultMintValue)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
		}

		k.SetSchemaAttribute(ctx, types.SchemaAttribute{
			NftSchemaCode:       schema_input.Code,
			Name:                scheamDefaultMintAttribute.Name,
			DataType:            scheamDefaultMintAttribute.DataType,
			Required:            scheamDefaultMintAttribute.Required,
			DisplayValueField:   scheamDefaultMintAttribute.DisplayValueField,
			DisplayOption:       scheamDefaultMintAttribute.DisplayOption,
			CurrentValue:        schmaAttributeValue,
			HiddenOveride:       scheamDefaultMintAttribute.HiddenOveride,
			HiddenToMarketplace: scheamDefaultMintAttribute.HiddenToMarketplace,
			Creator:             msg.Creator,
		})
	}

	// Add the schema_input to the store
	k.Keeper.SetNFTSchema(ctx, schema)

	// **** ACTION ****
	// set action
	for i, action := range schema_input.OnchainData.Actions {
		// check if action already exists
		_, isFound := k.Keeper.GetActionOfSchema(ctx, schema_input.Code, action.Name)
		if isFound {
			continue
		}
		k.Keeper.SetActionOfSchema(ctx, types.ActionOfSchema{
			Name:          action.Name,
			NftSchemaCode: schema_input.Code,
			Index:         uint64(i),
		})
	}

	// set action executor
	for _, actionExecutor := range schema_input.SystemActioners {

		//validate address
		_, err := sdk.AccAddressFromBech32(actionExecutor)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, actionExecutor)
		}

		// check if actionExecutor already exists
		_, isFound := k.Keeper.GetActionExecutor(ctx, schema_input.Code, actionExecutor)
		if isFound {
			continue
		}

		val, found := k.GetExecutorOfSchema(ctx, schema_input.Code)
		if !found {
			val = types.ExecutorOfSchema{
				NftSchemaCode: schema_input.Code,
				ExecutorAddress: []string{},
			}
		}

		// set executorOfSchema
		val.ExecutorAddress = append(val.ExecutorAddress, actionExecutor)
		
		k.SetExecutorOfSchema(ctx, types.ExecutorOfSchema{
			NftSchemaCode: schema_input.Code,
			ExecutorAddress: val.ExecutorAddress,
		})

		// set actionExecutor
		k.SetActionExecutor(ctx, types.ActionExecutor{
			Creator:         msg.Creator,
			NftSchemaCode:   schema_input.Code,
			ExecutorAddress: actionExecutor,
		})
	}

	// **** ENHANCEMENT ****

	// Check if schema_input by contract already exists
	nftSchemaByContract, found := k.Keeper.GetNFTSchemaByContract(ctx, schema_input.OriginData.OriginContractAddress)
	if !found {
		nftSchemaByContract = types.NFTSchemaByContract{
			OriginContractAddress: schema_input.OriginData.OriginContractAddress,
			SchemaCodes:           []string{schema_input.Code},
		}
	} else {
		nftSchemaByContract.SchemaCodes = append(nftSchemaByContract.SchemaCodes, schema_input.Code)
	}
	// Add the schema_input code to the list of schema_input codes
	k.Keeper.SetNFTSchemaByContract(ctx, nftSchemaByContract)

	// **** SCHEMA FEE ****
	feeConfig, found := k.Keeper.GetNFTFeeConfig(ctx)
	if found {
		// Get Denom
		amount, err := sdk.ParseCoinNormalized(feeConfig.SchemaFee.FeeAmount)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidFeeAmount, err.Error())
		}
		feeBalances, found := k.Keeper.GetNFTFeeBalance(ctx)
		if !found {

			feeBalances = types.NFTFeeBalance{
				FeeBalances: []string{
					"0" + amount.Denom,
				},
			}
		}

		if len(feeBalances.FeeBalances) > 0 {
			feeBalances.FeeBalances[types.FeeSubject_CREATE_NFT_SCHEMA] = "0" + amount.Denom
		}
		err = k.processFee(ctx, &feeConfig, &feeBalances, types.FeeSubject_CREATE_NFT_SCHEMA, creatorAddress)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrProcessingFee, err.Error())
		}
		// Set Fee Balance
		k.Keeper.SetNFTFeeBalance(ctx, feeBalances)
	}
	// **** END SCHEMA FEE ****

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateSchema,
			sdk.NewAttribute(types.AttributeKeyCreateSchemaCode, schema_input.Code),
			sdk.NewAttribute(types.AttributeKeyCreateSchemaResult, "success"),
		),
	})

	return &types.MsgCreateNFTSchemaResponse{
		Code: schema_input.Code,
	}, nil
}

// Total amount of fee collected from schema_input for each block then distribute to validators // ** In the begin block it will set to 0 again
func (k msgServer) processFee(ctx sdk.Context, feeConfig *types.NFTFeeConfig, feeBalances *types.NFTFeeBalance, feeSubject types.FeeSubject, source sdk.AccAddress) error {
	currentFeeBalance, _ := sdk.ParseCoinNormalized(feeBalances.FeeBalances[int32(feeSubject)])
	feeAmount, _ := sdk.ParseCoinNormalized(feeConfig.SchemaFee.FeeAmount)
	// Plus fee amount to fee balance
	currentFeeBalance = currentFeeBalance.Add(feeAmount)
	feeBalances.FeeBalances[int32(feeSubject)] = strconv.FormatInt(currentFeeBalance.Amount.Int64(), 10) + feeAmount.Denom

	err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, source, types.ModuleName, sdk.NewCoins(feeAmount),
	)
	return err
}
