grantOracle()
{
    echo "Grant 'oracle' to $1"
    sixd tx nftadmin grant-permission oracle $1 --from super-admin --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y \
        --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}
}

RPC_ENDPOINT=http://localhost:26657
CHAIN_ID=testnet
BASE64_SCHEMA=`cat nft-schema.json | base64 | tr -d '\n'`

sixd tx nftadmin grant-permission oracle_admin $(sixd keys show super-admin -a) --from super-admin -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}
sixd tx nftoracle set-minimum-confirmation 1 --from super-admin --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} -y
sixd tx nftmngr create-nft-schema ${BASE64_SCHEMA} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y \
    --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} 

grantOracle $(sixd keys show oracle1 -a)
grantOracle $(sixd keys show oracle2 -a)
grantOracle $(sixd keys show oracle3 -a)
grantOracle $(sixd keys show oracle4 -a)

sixd tx nftadmin grant-permission binder $(sixd keys show super-admin -a) --from super-admin --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} -y

sixd q nftadmin show-authorization --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}