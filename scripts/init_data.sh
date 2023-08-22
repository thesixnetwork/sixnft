grantOracle()
{
    echo "Grant 'oracle' to $1"
    sixnftd tx nftadmin grant-permission oracle $1 --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y \
        --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0';
}

RPC_ENDPOINT=http://localhost:26657
CHAIN_ID=sixnft
BASE64_SCHEMA=`cat ./mock-data/nft-schema.json | base64 | tr -d '\n'`

# sixnftd tx nftadmin grant-permission oracle_admin $(sixnftd keys show alice -a) --from alice -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0';
# sixnftd tx nftadmin grant-permission admin_signer_config $(sixnftd keys show alice -a) --from alice -y \
    # --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake | grep -q 'msg_index: 0';

sixnftd tx nftoracle set-minimum-confirmation 1 --from alice --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} -y | grep -q 'msg_index: 0';
sixnftd tx nftmngr create-nft-schema ${BASE64_SCHEMA} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y \
    --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'

# grantOracle $(sixnftd keys show oracle1 -a)
# grantOracle $(sixnftd keys show oracle2 -a)
# grantOracle $(sixnftd keys show oracle3 -a)
# grantOracle $(sixnftd keys show oracle4 -a)

# sixnftd tx nftoracle create-action-signer-config baobab 0xe2762507764fE0d20bfb490Fc6Eb3EB837042c35 --from alice -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake | grep -q 'msg_index: 0';

# sixnftd q nftadmin show-authorization --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}
