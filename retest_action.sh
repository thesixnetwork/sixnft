ALICE_ADDRESS=`sixnftd keys show alice --output json | jq ".address" | sed  's/\"//g'`

echo "ALICE_ADDRESS: $ALICE_ADDRESS"

BASE64_SCHEMA=`cat nft-schema-example.json | sed "s/0xNFTOWNER/${ALICE_ADDRESS}/g" | base64 | tr -d '\n'`

sixnftd tx nftmngr create-nft-schema --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y \
    ${BASE64_SCHEMA}

BASE64_META=`cat nft-data-example.json | sed "s/0xTOKENOWNER/${ALICE_ADDRESS}/g" | base64 | tr -d '\n'`

sixnftd tx nftmngr create-metadata buakaw1 1 --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y \
    ${BASE64_META}

echo "Enter to continue..."
read
echo "Value before action >>"
sixnftd query nftmngr show-nft-data buakaw1 1 --output json | jq '.nftData.onchain_attributes[] | select(.name=="percent_discount_10").number_attribute_value.value'

sixnftd tx nftmngr perform-action-by-nftadmin buakaw1 1 use10percentdiscount --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y
echo "Value after action"
sixnftd query nftmngr show-nft-data buakaw1 1 --output json | jq '.nftData.onchain_attributes[] | select(.name=="percent_discount_10").number_attribute_value.value'