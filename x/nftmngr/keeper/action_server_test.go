package keeper_test

import (
	"fmt"
	"testing"

	"github.com/thesixnetwork/sixnft/x/nftmngr/keeper"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/stretchr/testify/require"
)

var (
	schemaJSON = `
	{
		"code": "buakaw1",
		"name": "Buakaw1",
		"owner": "6nft1hf3qa8kmysg9cl7fjm44n2n6p8cj5gzatags65",
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
						"meta.SetNumber('percent_discount_10',meta.GetNumber('percent_discount_10') - 1);",
						"meta.SetBoolean('first_discount_used', true);"
					]
				}
			],
			"status": [{
				"status_name": "first_mint_complete",
				"status_value": false
			}],
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
	metaJSON = `
	{
		"nft_schema_code": "buakaw1",
		"token_id": "1",
		"token_owner": "6nft1hf3qa8kmysg9cl7fjm44n2n6p8cj5gzatags65",
		"origin_image": "https://bk1nft.sixnetwork.io/ipfs/QmaBUdqTctVteajpoMVEpo55XEXe2YVwWDgeNi1UT47UKU/1.png",
		"origin_attributes": [
			{
				"name": "background_l",
				"string_attribute_value": {
					"value": "L BG Ringside Golden"
				}
			},
			{
				"name": "background_r",
				"string_attribute_value": {
					"value": "R Tiger"
				}
			},
			{
				"name": "plate_l",
				"string_attribute_value": {
					"value": "L Banchamek Black"
				}
			},
			{
				"name": "plate_r",
				"string_attribute_value": {
					"value": "R Banchamek Golden"
				}
			},
			{
				"name": "body_l",
				"string_attribute_value": {
					"value": "L Body Normal Gentleman"
				}
			},
			{
				"name": "head_l",
				"string_attribute_value": {
					"value": "L NON HEAD"
				}
			},
			{
				"name": "clothes_l",
				"string_attribute_value": {
					"value": "L Robe Golden"
				}
			},
			{
				"name": "extra_l",
				"string_attribute_value": {
					"value": "L NO EXTRA"
				}
			},
			{
				"name": "hand_l",
				"string_attribute_value": {
					"value": "L Glove Black BV"
				}
			},
			{
				"name": "influencer",
				"string_attribute_value": {
					"value": "Teerawat Yioyim"
				}
			}
		],
		"onchain_attributes": [
			{
				"name": "exclusive_party_access",
				"boolean_attribute_value": {
					"value": true
				}
			},
			{
				"name": "percent_discount_10",
				"number_attribute_value": {
					"value": 10
				}
			},
			{
				"name": "first_discount_used",
				"boolean_attribute_value": {
					"value": false
				}
			},
			{
				"name": "influencer",
				"string_attribute_value": {
					"value": "Teerawat Yioyim"
				}
			}
		]
	}`
)

// func TestMetadata(t *testing.T) {

// 	data := types.NftData{}
// 	// err := k.GetCodec().(*codec.ProtoCodec).UnmarshalJSON([]byte(metadata), &data)
// 	err := jsonpb.UnmarshalString(metaJSON, &data)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	meta := types.NewMetadata(&data, types.AttributeOverriding_ORIGIN)
// 	_, bv := meta.GetBool("first_discount_used")
// 	require.False(t, bv)
// 	_, bv = meta.GetBool("exclusive_party_access")
// 	require.True(t, bv)

// 	_, sv := meta.GetString("influencer")
// 	require.Equal(t, "Teerawat Yioyim", sv)

// 	sv = meta.GetImage()
// 	require.Equal(t, "https://bk1nft.sixnetwork.io/ipfs/QmaBUdqTctVteajpoMVEpo55XEXe2YVwWDgeNi1UT47UKU/1.png", sv)

// 	err = meta.SetString("influencer", "Teerawat")
// 	require.NotNil(t, err)

// 	// reset overide to be CHAIN
// 	meta = types.NewMetadata(&data, types.AttributeOverriding_CHAIN)
// 	err = meta.SetString("influencer", "Teerawat")
// 	require.Nil(t, err)

// 	_, inf := meta.GetString("influencer")
// 	require.Equal(t, "Teerawat", inf)

// }

func TestAction(t *testing.T) {
	data := types.NftData{}
	// err := k.GetCodec().(*codec.ProtoCodec).UnmarshalJSON([]byte(metadata), &data)
	err := jsonpb.UnmarshalString(metaJSON, &data)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	schema := types.NFTSchema{}
	err = jsonpb.UnmarshalString(schemaJSON, &schema)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	selectAction := "use10percentdiscount"

	fmt.Println("Before start action")

	meta := types.NewMetadata(&schema, &data, types.AttributeOverriding_CHAIN)
	actionParams_ := []*types.ActionParameter{}
	for _, action := range schema.OnchainData.Actions {
		if action.Name == selectAction {
			keeper.ProcessAction(meta, action,actionParams_)
			break
		}
	}

	newAvailableDiscount, err := meta.MustGetNumber("percent_discount_10")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	require.True(t, newAvailableDiscount == 9)

	fmt.Println("newAvailableDiscount", newAvailableDiscount)

	newFirstDiscountUsed, err := meta.MustGetBool("first_discount_used")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	require.True(t, newFirstDiscountUsed == true)
}
