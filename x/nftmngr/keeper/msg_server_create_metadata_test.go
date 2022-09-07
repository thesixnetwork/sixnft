package keeper_test

import (
	"context"
	"fmt"
	"testing"

	keepertest "sixnft/testutil/keeper"
	"sixnft/x/nftmngr/keeper"
	"sixnft/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/jsonpb"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.NftmngrKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestCreateData(t *testing.T) {
	fmt.Println("Start Test")
	// base64Data := "ewogICAgIm5mdF9zY2hlbWFfY29kZSI6ICJidWFrYXcxIiwKICAgICJ0b2tlbl9pZCI6ICIxIiwKICAgICJ0b2tlbl9vd25lciI6ICI2bmZ0MTlwNXl2d3h0YzJxbmdud2RxbGhmemRjZDBzc25jd3d1bXJ4MmphIiwKICAgICJvcmlnaW5faW1hZ2UiOiAiaHR0cHM6Ly9iazFuZnQuc2l4bmV0d29yay5pby9pcGZzL1FtYUJVZHFUY3RWdGVhanBvTVZFcG81NVhFWGUyWVZ3V0RnZU5pMVVUNDdVS1UvMS5wbmciLAogICAgIm9yaWdpbl9hdHRyaWJ1dGVzIjogWwogICAgICAgIHsKICAgICAgICAgICAgIm5hbWUiOiAiYmFja2dyb3VuZF9sIiwKICAgICAgICAgICAgInZhbHVlIjogIkwgQkcgUmluZ3NpZGUgR29sZGVuIgogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgICAibmFtZSI6ICJiYWNrZ3JvdW5kX3IiLAogICAgICAgICAgICAidmFsdWUiOiAiUiBUaWdlciIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgIm5hbWUiOiAicGxhdGVfbCIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJMIEJhbmNoYW1layBCbGFjayIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgIm5hbWUiOiAicGxhdGVfciIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJSIEJhbmNoYW1layBHb2xkZW4iCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAgICJuYW1lIjogImJvZHlfbCIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJMIEJvZHkgTm9ybWFsIEdlbnRsZW1hbiIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgIm5hbWUiOiAiaGVhZF9sIiwKICAgICAgICAgICAgInZhbHVlIjogIkwgTk9OIEhFQUQiCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAgICJuYW1lIjogImNsb3RoZXNfbCIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJMIFJvYmUgR29sZGVuIgogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgICAibmFtZSI6ICJleHRyYV9sIiwKICAgICAgICAgICAgInZhbHVlIjogIkwgTk8gRVhUUkEiCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAgICJuYW1lIjogImhhbmRfbCIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJMIEdsb3ZlIEJsYWNrIEJWIgogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgICAibmFtZSI6ICJpbmZsdWVuY2VyIiwKICAgICAgICAgICAgInZhbHVlIjogIlRlZXJhd2F0IFlpb3lpbSIKICAgICAgICB9CiAgICBdLAogICAgIm9uY2hhaW5fYXR0cmlidXRlcyI6IFsKICAgICAgICB7CiAgICAgICAgICAgICJuYW1lIjogImV4Y2x1c2l2ZV9wYXJ0eV9hY2Nlc3MiLAogICAgICAgICAgICAidmFsdWUiOiBmYWxzZQogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgICAibmFtZSI6ICJmaXJzdF9kaXNjb3VudF91c2VkIiwKICAgICAgICAgICAgInZhbHVlIjogZmFsc2UKICAgICAgICB9CiAgICBdCn0K"
	// k, _ := keepertest.NftmngrKeeper(t)
	// metadata, err := base64.StdEncoding.DecodeString(base64Data)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	schema := `
	{
		"code": "buakaw1",
		"name": "Buakaw1",
		"owner": "6nft19p5yvwxtc2qngnwdqlhfzdcd0ssncwwumrx2ja",
		"origin_data": {
			"origin_base_uri": "https://bk1nft.sixnetwork.io/ipfs/QmcovEcZxiM1M9gf3535jUCpbqFZSjc2hHuYnZEbWNmt5a",
			"origin_chain": "ethereum",
			"origin_contract_address": "0x9F1CC70b11f4129d042d0037c2066d12E16d9a52",
			"attribute_overriding": "ORIGIN",
			"metadata_format": "opensea",
			"origin_attributes": [
				{
					"name": "background_l",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Background L"
						}
					}
				},
				{
					"name": "background_r",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Background R"
						}
					}
				},
				{
					"name": "plate_l",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Plate L"
						}
					}
				},
				{
					"name": "plate_r",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Plate R"
						}
					}
				},
				{
					"name": "body_l",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Body L"
						}
					}
				},
				{
					"name": "body_r",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Body R"
						}
					}
				},
				{
					"name": "head_l",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Head L"
						}
					}
				},
				{
					"name": "head_r",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Head R"
						}
					}
				},
				{
					"name": "clothes_l",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Clothes L"
						}
					}
				},
				{
					"name": "clothes_r",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Clothes R"
						}
					}
				},
				{
					"name": "extra_l",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Extra L"
						}
					}
				},
				{
					"name": "extra_r",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Extra R"
						}
					}
				},
				{
					"name": "hand_l",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Hand L"
						}
					}
				},
				{
					"name": "hand_r",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Hand R"
						}
					}
				},
				{
					"name": "influencer",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Influencer"
						}
					}
				}
			]
		},
		"onchain_data": {
			"reveal_required": true,
			"reveal_secret": "",
			"nft_attributes": [
				{
					"name": "expire_date",
					"data_type": "long",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"display_type": "date",
							"trait_type": "Expire Date"
						}
					}
				}
			],
			"token_attributes": [
				{
					"name": "exclusive_party_access",
					"data_type": "boolean",
					"required": true,
					"hidden_to_marketplace": true
				},
				{
					"name": "first_discount_used",
					"data_type": "boolean",
					"required": true,
					"hidden_to_marketplace": true
				},
				{
					"name": "percent_discount_10",
					"default_mint_value": "10",
					"data_type": "interger",
					"required": true,
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Available 10 Percent Discount",
							"max_value": 5,
							"display_type": "number"
						}
					}
				},
				{
					"name": "discount_percentage",
					"default_mint_value": "0",
					"data_type": "interger",
					"required": true,
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Discount",
							"display_type": "boost_percentage"
						}
					}
				},
				{
					"name": "discount_amount",
					"default_mint_value": "0",
					"data_type": "interger",
					"required": true,
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Discount",
							"display_type": "boost_number"
						}
					}
				}
			],
			"actions": [
				{
					"name": "use10percentdiscount",
					"desc": "Use 10% discount on purchase",
					"when": "meta.GetNumber('percent_discount_10') > 0",
					"then": [
						"meta.SetNumber(meta.GetNumber('percent_discount_10') - 1)",
						"meta.SetBoolean('first_discount_used', true)",
						"meta.Random(x,y)"
					]
				}
			],
			"on_off_switch": {
				"active": true
			},
			"status": {
				"first_mint_complete": false
			},
			"nft_attributes_value": [
				{
					"name": "expire_date",
					"number_attribute_value": {
						"value": 1665009678
					}
				}
			]
		}
	}
	
	`
	metadata := `
	{
		"nft_schema_code": "buakaw1",
		"token_id": "1",
		"token_owner": "6nft19p5yvwxtc2qngnwdqlhfzdcd0ssncwwumrx2ja",
		"origin_image": "https://bk1nft.sixnetwork.io/ipfs/QmaBUdqTctVteajpoMVEpo55XEXe2YVwWDgeNi1UT47UKU/1.png",
		"origin_attributes": [
			{
				"name": "background_l",
				"string_attribute_value" : {"value": "L BG Ringside Golden"}
			},
			{
				"name": "background_r",
				"string_attribute_value" : {"value": "R Tiger"}
			},
			{
				"name": "plate_l",
				"string_attribute_value" : {"value": "L Banchamek Black"}
			},
			{
				"name": "plate_r",
				"string_attribute_value" : {"value": "R Banchamek Golden"}
			},
			{
				"name": "body_l",
				"string_attribute_value" : {"value": "L Body Normal Gentleman"}
			},
			{
				"name": "head_l",
				"string_attribute_value" : {"value": "L NON HEAD"}
			},
			{
				"name": "clothes_l",
				"string_attribute_value" : {"value": "L Robe Golden"}
			},
			{
				"name": "extra_l",
				"string_attribute_value" : {"value": "L NO EXTRA"}
			},
			{
				"name": "hand_l",
				"string_attribute_value" : {"value": "L Glove Black BV"}
			},
			{
				"name": "influencer",
				"string_attribute_value" : {"value": "Teerawat Yioyim"}
			}
		],
		"onchain_attributes": [
			{
				"name": "exclusive_party_access",
				"boolean_attribute_value": {
					"value": false
				}
			},
			{
				"name": "first_discount_used",
				"boolean_attribute_value": {
					"value": false
				}
			}
		]
	}
	
	`
	data := types.NftData{}
	// err := k.GetCodec().(*codec.ProtoCodec).UnmarshalJSON([]byte(metadata), &data)
	err := jsonpb.UnmarshalString(metadata, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_schema := types.NFTSchema{}
	// err := k.GetCodec().(*codec.ProtoCodec).UnmarshalJSON([]byte(metadata), &data)
	err = jsonpb.UnmarshalString(schema, &_schema)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	valid, err := ValidateNFTData(&data, &_schema)
	fmt.Println("Valid: ", valid)
	fmt.Println("Valid Err: ", err)
	// fmt.Println("DataOutput", data)

	// data := types.NftData{}
	// attribute := types.NftAttributeValue{
	// 	Name: "exclusive_party_access",
	// }
	// attribute.Value = &types.NftAttributeValue_BooleanAttributeValue_{BooleanAttributeValue: &types.NftAttributeValue_BooleanAttributeValue{Value: true}}
	// data.OnchainAttributes = append(data.OnchainAttributes, &attribute)

	// m := jsonpb.Marshaler{}
	// str, _ := m.MarshalToString(&attribute)
	// fmt.Println(str)
}

// Validate NFT Data
func ValidateNFTData(data *types.NftData, schema *types.NFTSchema) (bool, error) {
	// Validate Onchain Attributes Value
	_, err := NFTDataHasDuplicateNftAttributesValue(data.OnchainAttributes)
	if err != nil {
		return false, err
	}
	// Validate Origin Attributes Value
	_, err = NFTDataHasDuplicateNftAttributesValue(data.OriginAttributes)
	if err != nil {
		return false, err
	}
	// Validate Origin Attributes Exist in Schema
	_, err = DataAttributesExistInSchema(schema.OriginData.OriginAttributes, data.OriginAttributes)
	if err != nil {
		return false, err
	}
	// Validate Onchain Attributes Exist in Schema
	_, err = DataAttributesExistInSchema(schema.OnchainData.TokenAttributes, data.OnchainAttributes)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NFTDataHasDuplicateNftAttributesValue(attributes []*types.NftAttributeValue) (bool, error) {
	attributesLen := len(attributes)
	for i := 0; i < attributesLen; i++ {
		for j := i + 1; j < attributesLen; j++ {
			if attributes[i].Name == attributes[j].Name {
				return false, fmt.Errorf("Duplicate attribute value name: %s", attributes[i].Name)
			}
		}
	}
	return true, nil
}

func DataAttributesExistInSchema(schemaAttributes []*types.AttributeDefinition, dataAttributes []*types.NftAttributeValue) (bool, error) {
	for _, dataAttribute := range dataAttributes {
		for _, schemaAttribute := range schemaAttributes {
			if dataAttribute.Name != schemaAttribute.Name {
				return false, fmt.Errorf("Attribute %s does not exist in schema", dataAttribute.Name)
			}
		}
	}
	return true, nil
}
