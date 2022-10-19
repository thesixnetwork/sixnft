EVMSIGN=./evmsign
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
echo "##  9. Oracle - Create Action Request      ##"
echo "##  10. Oracle - Get Action Request        ##"
echo "##  11. Oracle - Submit Action Response    ##"
echo "##  12. Oracle - Create Verfify Request    ##"
echo "##  13. Oracle - Get Verify Request        ##"
echo "##  14. Oracle - Submit Verify Response    ##"
echo "##  15. Add Attribute                      ##"
echo "##  16. Add Action                         ##"
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
        sixnftd tx nftmngr perform-action-by-nftadmin ${schema_code} ${token_id} ${action} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y
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
    9) echo "Oracle - Create Action Request"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token ID: " token_id
        read -p "Enter Action: " action
        read -p "Require confirmations: " require_confirmations
        read -p "Reference ID: " reference_id
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi

        BASE64JSON=`cat action-param.json | sed "s/ACTION/${action}/g" | sed "s/TOKEN_ID/${token_id}/g" | sed "s/SCHEMA_CODE/${schema_code}/g" | sed "s/REFID/${reference_id}/g"`
        # echo "BASE64JSON: ${BASE64JSON}"
        BASE64_MESSAGE=`echo -n $BASE64JSON | base64 | tr -d '\n'`
        # echo "BASE64_MESSAGE: ${BASE64_MESSAGE}"
        MESSAGE_SIG=`echo -n ${BASE64_MESSAGE} | $EVMSIGN ./.secret`
        # echo "MESSAGE_SIG: ${MESSAGE_SIG}"

        BASE64_ACTION_SIG=`cat action-signature.json | sed "s/SIGNATURE/${MESSAGE_SIG}/g" | sed "s/MESSAGE/${BASE64_MESSAGE}/g" | base64 | tr -d '\n'`

        # echo -n ${BASE64_MESSAGE} | $EVMSIGN ./.secret 1
        # echo -n ${BASE64_MESSAGE} | wc -c
        sixnftd tx nftoracle create-action-request ethereum ${BASE64_ACTION_SIG} ${require_confirmations} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y 
        ;;
    10) echo "Oracle - Get Action Request"
        read -p "Action Request ID: " action_request_id 
        sixnftd q nftoracle show-action-request ${action_request_id} --output json | jq .
        ;;
    11) echo "Oracle - Submit Action Response"
        read -p "Action Request ID: " action_request_id
        read -p "Oracle : " oracle_key_name
        BASE64_ORIGINDATA=`cat nft-origin-data.json | base64 | tr -d '\n'`

        sixnftd tx nftoracle submit-action-response ${action_request_id} ${BASE64_ORIGINDATA} --from ${oracle_key_name} --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y
        ;;
    12) echo "Oracle - Create Verify Schema Request"
        read -p "Enter Schema Code: " schema_code
        read -p "Require confirmations: " require_confirmations
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi

        BASE64JSON=`cat verify-collection-owner.json`
        # echo "BASE64JSON: ${BASE64JSON}"
        BASE64_MESSAGE=`echo -n $BASE64JSON | base64 | tr -d '\n'`
        # echo "BASE64_MESSAGE: ${BASE64_MESSAGE}"
        MESSAGE_SIG=`echo -n ${BASE64_MESSAGE} | $EVMSIGN ./.secret`
        # echo "MESSAGE_SIG: ${MESSAGE_SIG}"

        BASE64_VERIFY_SIG=`cat verify-signature.json | sed "s/SIGNATURE/${MESSAGE_SIG}/g" | sed "s/MESSAGE/${BASE64_MESSAGE}/g" | base64 | tr -d '\n'`

        sixnftd tx nftoracle create-verify-collection-owner-request ${schema_code} ${BASE64_VERIFY_SIG} ${require_confirmations} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y
        ;;
    13) echo "Oracle - Get Verify Request"
        read -p "Verify Request ID: " verfiry_request_id 
        sixnftd q nftoracle show-collection-owner-request ${verfiry_request_id} --output json | jq .
        ;;
    14) echo "Oracle - Submit Verify Response"
        read -p "Enter Schema Code: " schema_code
        read -p "Verify Request ID: " verfiry_request_id
        read -p "Oracle : " oracle_key_name
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        BASE64_ORIGINDATA=`cat verify-collection-owner.json | base64 | tr -d '\n'`

        sixnftd tx nftoracle submit-verify-collection-owner ${verfiry_request_id} ${schema_code} ${BASE64_ORIGINDATA} --from ${oracle_key_name} --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y
        ;;
     15) echo "Add Attribute"
        read -p "Enter Schema Code: " schema_code 
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        read -p "Location of attribute (0 or 1): " location
        BASE64_ATTRIBUTE=`cat new-attribute.json | base64 | tr -d '\n'`
        sixnftd tx nftmngr add-attribute ${schema_code} ${location} ${BASE64_ATTRIBUTE} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y \
            --chain-id sixnft
        ;;
     16) echo "Add Action"
        read -p "Enter Schema Code: " schema_code 
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        BASE64_ACTION=`cat new-action.json | base64 | tr -d '\n'`
        sixnftd tx nftmngr add-action ${schema_code} ${BASE64_ACTION} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake -y \
            --chain-id sixnft
        ;;
    *) echo "Invalid choice"
       ;;
esac