default_schema_code=$1
echo "#############################################"
echo "##                                         ##"
echo "##  Welcome to the menu script             ##"
echo "##                                         ##"
echo "##  Please select an option                ##"
echo "##                                         ##"
echo "##  1. Show Schema                         ##"
echo "##  2. Show NFTs                           ##"
echo "##  3. Mockup Token                        ##"
echo "##  4. Do Action                           ##"
echo "##  5. Set NFT Attribute                   ##"
echo "##  6. Oracle - Create Mint Request        ##"
echo "##  7. Oracle - Get Mint Request           ##"
echo "##  8. Oracle - Submit Mint Response       ##"
echo "##  Your choice:                           ##"
echo "##                                         ##"
echo "#############################################"
read -p "Your choice: " choice
case $choice in
    1) echo "Showing Schema"
        read -p "Enter Schema Code: " schema_code 
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        sixnftd q nftmngr show-nft-schema ${schema_code} --output json | jq .
        ;;
    2) echo "Showing NFT"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token ID: " token_id
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        sixnftd q nftmngr show-nft-data ${schema_code} ${token_id} --output json | jq .
        ;;
    3) echo "Mockup Token"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token ID: " token_id
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        BASE64_META=`cat nft-data.json | sed "s/TOKENID/${token_id}/g"  | sed "s/SCHEMA_CODE/${schema_code}/g" | base64 | tr -d '\n'`
        sixnftd tx nftmngr create-metadata "${schema_code}" ${token_id} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y \
            ${BASE64_META}
        ;;
    4) echo "Do Action"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token ID: " token_id
        read -p "Enter Action: " action
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        sixnftd tx nftmngr perform-action-by-admin ${schema_code} ${token_id} ${action} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y 
        ;;
    5) echo "Set NFT Attribute"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Value (attribute_name=N[value]): " value
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi

        ATTRIBUTE_NAME=`echo $value | cut -d'=' -f1`
        ATTRIBUTE_VALUE_STRING=`echo $value | cut -d'=' -f2`
        # get one character from ATTRIBUTE_VALUE
        ATTRIBUTE_VALUE_CHAR=`echo $ATTRIBUTE_VALUE_STRING | cut -c1`
        # get characters between [] from ATTRIBUTE_VALUE_CHAR
        ATTRIBUTE_VALUE_VALUE=`echo $ATTRIBUTE_VALUE_STRING | cut -d'[' -f2 | cut -d']' -f1`

        if [ "$ATTRIBUTE_VALUE_CHAR" = "N" ]; then
            ATTRIBUTE_VALUE_TYPE="number"
            ATTRIBUTE_VALUE_TYPE_VALUE=${ATTRIBUTE_VALUE_VALUE}
        elif [ "$ATTRIBUTE_VALUE_CHAR" = "S" ]; then
            ATTRIBUTE_VALUE_TYPE="string"
            ATTRIBUTE_VALUE_TYPE_VALUE="\"${ATTRIBUTE_VALUE_VALUE}\""
        elif [ "$ATTRIBUTE_VALUE_CHAR" = "B" ]; then
            ATTRIBUTE_VALUE_TYPE="boolean"
            # check if ATTRIBUTE_VALUE_VALUE is true or false
            if [ "$ATTRIBUTE_VALUE_VALUE" = "true" ]; then
                ATTRIBUTE_VALUE_TYPE_VALUE="true"
            elif [ "$ATTRIBUTE_VALUE_VALUE" = "false" ]; then
                ATTRIBUTE_VALUE_TYPE_VALUE="false"
            else
                echo "Invalid boolean value"
                exit 1
            fi
        elif [ "$ATTRIBUTE_VALUE_CHAR" = "F" ]; then
            ATTRIBUTE_VALUE_TYPE="float"
            ATTRIBUTE_VALUE_TYPE_VALUE=${ATTRIBUTE_VALUE_VALUE}
        fi

        BASE64_ATTR=`cat nft-data.json \
            | sed "s/#ATTRIBUTE_NAME#/${ATTRIBUTE_NAME}/g" \
            | sed "s/#ATTRIBUTE_VALUE_TYPE#/${ATTRIBUTE_VALUE_TYPE}/g" \
            | sed "s/#ATTRIBUTE_VALUE_TYPE_VALUE#/${ATTRIBUTE_VALUE_TYPE_VALUE}/g" \
            | base64 | tr -d '\n'`

        echo "BASE64_ATTR: ${BASE64_ATTR}"

        sixnftd tx nftmngr set-nft-attribute ${schema_code} ${BASE64_ATTR} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y 
        ;;
    6) echo "Oracle - Create Mint Request"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token ID: " token_id
        read -p "Require confirmations: " require_confirmations
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        sixnftd tx nftoracle create-mint-request ${schema_code} ${token_id} ${require_confirmations} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y 
        ;;
    7) echo "Oracle - Get Mint Request"
        read -p "Mint Request ID: " mint_request_id 
        sixnftd q nftoracle show-mint-request ${mint_request_id} --output json | jq .
        ;;
    8) echo "Oracle - Submit Mint Response"
        read -p "Mint Request ID: " mint_request_id
        read -p "Oracle : " oracle_key_name
        BASE64_ORIGINDATA=`cat nft-origin-data.json | base64 | tr -d '\n'`

        sixnftd tx nftoracle submit-mint-response ${mint_request_id} ${BASE64_ORIGINDATA} --from ${oracle_key_name} --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y 
        ;;
    *) echo "Invalid choice"
       ;;
esac