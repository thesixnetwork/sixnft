grantOracle()
{
    echo "Grant 'oracle' to $1"
    sixnftd tx nftadmin grant-permission oracle $1 --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y \
        --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}
}

RPC_ENDPOINT=http://localhost:26657
CHAIN_ID=sixnft
BASE64_SCHEMA=`cat nft-schema.json | base64 | tr -d '\n'`

# sixnftd tx nftadmin grant-permission oracle_admin $(sixnftd keys show alice -a) --from alice -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}
# sixnftd tx nftoracle set-minimum-confirmation 1 --from alice --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} -y
sixnftd tx nftmngr create-nft-schema ${BASE64_SCHEMA} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y \
    --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} 

# grantOracle $(sixnftd keys show oracle1 -a)
# grantOracle $(sixnftd keys show oracle2 -a)
# grantOracle $(sixnftd keys show oracle3 -a)
# grantOracle $(sixnftd keys show oracle4 -a)

# sixnftd tx nftadmin grant-permission binder $(sixnftd keys show alice -a) --from alice --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} -y

# sixnftd q nftadmin show-authorization --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}
