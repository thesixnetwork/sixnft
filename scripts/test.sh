echo "Deploy Schema"
read -p "Press enter to continue"
RPC_ENDPOINT=http://localhost:26657
CHAIN_ID=sixnft
BASE64_SCHEMA=`cat nft-schema.json | base64 | tr -d '\n'`
sixnftd tx nftmngr create-nft-schema ${BASE64_SCHEMA} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y \
    --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} 
echo "Done"

echo "Create NFT"
read -p "Press enter to continue"
BASE64_META=`cat nft-data.json | sed "s/TOKENID/1/g"  | sed "s/SCHEMA_CODE/sixnetwork.develop_v073/g" | base64 | tr -d '\n'`
sixnftd tx nftmngr create-metadata sixnetwork.develop_v073 1 --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y \
    ${BASE64_META} --chain-id sixnft
echo "Done"

echo "start_mission"
read -p "Press enter to continue"
sixnftd tx nftmngr perform-action-by-nftadmin sixnetwork.develop_v073 1 start_mission start "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake -y \
    --chain-id sixnft
echo "Done"
echo ""

echo "Test Read NFT"
read -p "Press enter to continue"
sixnftd tx nftmngr perform-action-by-nftadmin sixnetwork.develop_v073 1 test_read_nft readnft "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake -y \
    --chain-id sixnft
echo "Done 🖕"
echo ""

echo "Test Split"
read -p "Press enter to continue"
sixnftd tx nftmngr perform-action-by-nftadmin sixnetwork.develop_v073 1 test_split split "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake -y \
    --chain-id sixnft
echo "Done 🖕"
echo ""

echo "Test Lowercase"
read -p "Press enter to continue"
sixnftd tx nftmngr perform-action-by-nftadmin sixnetwork.develop_v073 1 test_lowercase lowercase "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake -y \
    --chain-id sixnft
echo "Done 🖕"
echo ""

echo "Test Uppercase"
read -p "Press enter to continue"
sixnftd tx nftmngr perform-action-by-nftadmin sixnetwork.develop_v073 1 test_uppercase uppercase "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake -y \
    --chain-id sixnft
echo "Done 🖕"
echo ""

echo "Test Hide Fail"
read -p "Press enter to continue"
sixnftd tx nftmngr perform-action-by-nftadmin sixnetwork.develop_v073 1 test_hidden hidden_fail "[{\"name\":\"attribute_name\",\"value\":\"hide_fail\"},{\"name\":\"show\",\"value\":\"false\"}]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake -y --chain-id sixnft
echo "Done 🖕"
echo ""

echo "Test Hide Fail"
read -p "Press enter to continue"
sixnftd tx nftmngr perform-action-by-nftadmin sixnetwork.develop_v073 1 test_hidden hidden_pass "[{\"name\":\"attribute_name\",\"value\":\"hide_pass\"},{\"name\":\"show\",\"value\":\"true\"}]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake -y --chain-id sixnft
echo "Done 🖕"
echo ""

## test new msg
sixnftd tx nftmngr set-metadata-format sixnetwork.develop_v073 xclusive --from alice -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake | grep -q 'msg_index: 0';
sixnftd tx nftmngr set-uri-retrieval-method sixnetwork.develop_v073 1 --from alice -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake | grep -q 'msg_index: 0';
sixnftd tx nftmngr set-origin-chain sixnetwork.develop_v073 sixchain --from alice -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake | grep -q 'msg_index: 0';
sixnftd tx nftmngr set-origin-contract sixnetwork.develop_v073 0xe2762507764fE0d20bfb490Fc6Eb3EB837042c35 --from alice --from alice -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake | grep -q 'msg_index: 0';