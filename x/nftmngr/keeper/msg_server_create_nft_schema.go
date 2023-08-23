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
	schema := types.NFTSchema{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(jsonSchema, &schema)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingSchemaMessage, err.Error())
	}
	schema.Owner = msg.Creator
	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	// Validate Schema Message and return error if not valid
	valid, err := ValidateNFTSchema(&schema)
	_ = valid
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingNFTSchema, err.Error())
	}
	// if mint_authorization is empty then set system to default
	if len(schema.MintAuthorization) == 0 || schema.MintAuthorization != types.KeyMintPermissionOnlySystem && schema.MintAuthorization != types.KeyMintPermissionAll {
		schema.MintAuthorization = types.KeyMintPermissionOnlySystem
	}
	// Check if the schema already exists
	_, found := k.Keeper.GetNFTSchema(ctx, schema.Code)
	if found {
		return nil, sdkerrors.Wrap(types.ErrSchemaAlreadyExists, schema.Code)
	}
	foundOrganization, organizationName := GetOrganizationFromSchemaCode(schema.Code)
	// If there is organization in schema code, check if the organization exists
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
	_ = MergeAllSchemaAttributesAndAlterOrderIndex(schema.OriginData.OriginAttributes, schema.OnchainData.SchemaAttributes, schema.OnchainData.TokenAttributes)

	// Add the schema to the store
	k.Keeper.SetNFTSchema(ctx, schema)


	// **** ENHANCEMENT ****

	// Check if schema by contract already exists
	nftSchemaByContract, found := k.Keeper.GetNFTSchemaByContract(ctx, schema.OriginData.OriginContractAddress)
	if !found {
		nftSchemaByContract = types.NFTSchemaByContract{
			OriginContractAddress: schema.OriginData.OriginContractAddress,
			SchemaCodes:           []string{schema.Code},
		}
	} else {
		nftSchemaByContract.SchemaCodes = append(nftSchemaByContract.SchemaCodes, schema.Code)
	}
	// Add the schema code to the list of schema codes
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
			sdk.NewAttribute(types.AttributeKeyCreateSchemaCode, schema.Code),
			sdk.NewAttribute(types.AttributeKeyCreateSchemaResult, "success"),
		),
	})

	return &types.MsgCreateNFTSchemaResponse{
		Code: schema.Code,
	}, nil
}

// Total amount of fee collected from schema for each block then distribute to validators // ** In the begin block it will set to 0 again
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
