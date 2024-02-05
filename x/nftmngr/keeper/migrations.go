package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	v2types "github.com/thesixnetwork/sixnft/x/nftmngr/migrations/v2/types"
	types "github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
	}
}

func (m Migrator) Migrate2toV3(ctx sdk.Context) error {
	allSchemas := m.keeper.GetAllNFTSchemaLegacy(ctx)

	for _, eachSchema := range allSchemas {

		v2OriginAttribute := make([]*types.AttributeDefinition, 0)
		v2OnchainDataAttribute := make([]*types.AttributeDefinition, 0)
		v2TokenAttribute := make([]*types.AttributeDefinition, 0)
		v2flagStatus := make([]*types.FlagStatus,0)

		for _, attribute := range eachSchema.OriginData.OriginAttributes {

			convertedDefaultMintValue, err := ConvertDefaultMintValueFromV2ToV3(attribute.DefaultMintValue)
			if err != nil {
				// Handle the error appropriately
				return fmt.Errorf("failed to convert DefaultMintValue for attribute '%s': %w", attribute.Name, err)
			}

			v2OriginAttribute = append(v2OriginAttribute, &types.AttributeDefinition{
				Name:              attribute.Name,
				DataType:          attribute.DataType,
				Required:          attribute.Required,
				DisplayValueField: attribute.DisplayValueField,
				DisplayOption: &types.DisplayOption{
					BoolTrueValue:  attribute.DisplayOption.BoolTrueValue,
					BoolFalseValue: attribute.DisplayOption.BoolFalseValue,
					Opensea:        (*types.OpenseaDisplayOption)(attribute.DisplayOption.Opensea),
				},
				DefaultMintValue:    convertedDefaultMintValue,
				HiddenOveride:       attribute.HiddenOveride,
				HiddenToMarketplace: attribute.HiddenOveride,
				Index:               attribute.Index,
			})
		}

		for _ , ocAttribute := range eachSchema.OnchainData.NftAttributes {
			convertedDefaultMintValue, err := ConvertDefaultMintValueFromV2ToV3(ocAttribute.DefaultMintValue)
			if err != nil {
				// Handle the error appropriately
				return fmt.Errorf("failed to convert DefaultMintValue for attribute '%s': %w", ocAttribute.Name, err)
			}

			v2OnchainDataAttribute = append(v2OnchainDataAttribute, &types.AttributeDefinition{
				Name:              ocAttribute.Name,
				DataType:          ocAttribute.DataType,
				Required:          ocAttribute.Required,
				DisplayValueField: ocAttribute.DisplayValueField,
				DisplayOption: &types.DisplayOption{
					BoolTrueValue:  ocAttribute.DisplayOption.BoolTrueValue,
					BoolFalseValue: ocAttribute.DisplayOption.BoolFalseValue,
					Opensea:        (*types.OpenseaDisplayOption)(ocAttribute.DisplayOption.Opensea),
				},
				DefaultMintValue:    convertedDefaultMintValue,
				HiddenOveride:       ocAttribute.HiddenOveride,
				HiddenToMarketplace: ocAttribute.HiddenOveride,
				Index:               ocAttribute.Index,
			})
		}

		for _, tkAttribute := range eachSchema.OnchainData.TokenAttributes {
			convertedDefaultMintValue, err := ConvertDefaultMintValueFromV2ToV3(tkAttribute.DefaultMintValue)
			if err != nil {
				// Handle the error appropriately
				return fmt.Errorf("failed to convert DefaultMintValue for attribute '%s': %w", tkAttribute.Name, err)
			}

			v2TokenAttribute = append(v2TokenAttribute, &types.AttributeDefinition{
				Name:              tkAttribute.Name,
				DataType:          tkAttribute.DataType,
				Required:          tkAttribute.Required,
				DisplayValueField: tkAttribute.DisplayValueField,
				DisplayOption: &types.DisplayOption{
					BoolTrueValue:  tkAttribute.DisplayOption.BoolTrueValue,
					BoolFalseValue: tkAttribute.DisplayOption.BoolFalseValue,
					Opensea:        (*types.OpenseaDisplayOption)(tkAttribute.DisplayOption.Opensea),
				},
				DefaultMintValue:    convertedDefaultMintValue,
				HiddenOveride:       tkAttribute.HiddenOveride,
				HiddenToMarketplace: tkAttribute.HiddenOveride,
				Index:               tkAttribute.Index,
			})
		}

		v2Action := make([]*types.Action, 0)
		

		for _, action := range eachSchema.OnchainData.Actions {
			v2Params := make([]*types.ActionParams, 0)
		
			for _, param := range action.Params {
				// v2Params = append(v2Params, &types.ActionParams{
				// 	Name: param.Name,
				// 	Desc: param.Desc,
				// 	DefaultValue: param.DefaultValue,
				// 	DataType: param.DataType,
				// })
				v2Params = append(v2Params, (*types.ActionParams)(param))
			}

			v2Action = append(v2Action, &types.Action{
				Name: action.Name,
				Desc: action.Desc,
				Disable: action.Disable,
				When: action.When,
				Then: action.Then,
				AllowedActioner: types.AllowedActioner(action.AllowedActioner),
				Params: v2Params,
			})
		}


		for _, status := range eachSchema.OnchainData.Status {
			v2flagStatus = append(v2flagStatus, &types.FlagStatus{
				StatusName: status.StatusName,
				StatusValue: status.StatusValue,
			})
		}

		// migration action_params to ActionParams
		m.keeper.SetNFTSchema(ctx, types.NFTSchema{
			Code:        eachSchema.Code,
			Name:        eachSchema.Name,
			Owner:       eachSchema.Owner,
			Description: eachSchema.Description,
			// OriginData: (*types.OriginData)(eachSchema.OriginData),
			OriginData: &types.OriginData{
				OriginChain:           eachSchema.OriginData.OriginChain,
				OriginContractAddress: eachSchema.OriginData.OriginContractAddress,
				OriginBaseUri:         eachSchema.OriginData.OriginBaseUri,
				AttributeOverriding:   types.AttributeOverriding(eachSchema.OriginData.AttributeOverriding),
				MetadataFormat:        eachSchema.OriginData.MetadataFormat,
				OriginAttributes:      v2OriginAttribute,
				UriRetrievalMethod:    types.URIRetrievalMethod(eachSchema.OriginData.UriRetrievalMethod),
			},
			OnchainData: &types.OnChainData{
				NftAttributes: v2OriginAttribute,
				TokenAttributes: v2TokenAttribute,
				Actions: v2Action,
				Status: v2flagStatus,
			},
			IsVerified: eachSchema.IsVerified,
			MintAuthorization: eachSchema.MintAuthorization,
		})
	}

	return nil
}

func ConvertDefaultMintValueFromV2ToV3(defaultMintValue *v2types.DefaultMintValue) (*types.DefaultMintValue, error) {
	attributeValue := &types.DefaultMintValue{}

	switch value := defaultMintValue.Value.(type) {
	case *v2types.DefaultMintValue_NumberAttributeValue:
		attributeValue.Value = &types.DefaultMintValue_NumberAttributeValue{
			NumberAttributeValue: &types.NumberAttributeValue{
				Value: value.NumberAttributeValue.Value,
			},
		}
	case *v2types.DefaultMintValue_StringAttributeValue:
		attributeValue.Value = &types.DefaultMintValue_StringAttributeValue{
			StringAttributeValue: &types.StringAttributeValue{
				Value: value.StringAttributeValue.Value,
			},
		}
	case *v2types.DefaultMintValue_BooleanAttributeValue:
		attributeValue.Value = &types.DefaultMintValue_BooleanAttributeValue{
			BooleanAttributeValue: &types.BooleanAttributeValue{
				Value: value.BooleanAttributeValue.Value,
			},
		}
	case *v2types.DefaultMintValue_FloatAttributeValue:
		attributeValue.Value = &types.DefaultMintValue_FloatAttributeValue{
			FloatAttributeValue: &types.FloatAttributeValue{
				Value: value.FloatAttributeValue.Value,
			},
		}
	default:
		return nil, fmt.Errorf("unknown value type: %T", value)
	}

	return attributeValue, nil
}
