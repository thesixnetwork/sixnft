read -p "Enter Schema Code: " schema_code 
read -p "Enter Token IDs: " token_id
read -p "Enter Action: " action
read -p "Enter Ref ID: " ref_id
if [ -z "$schema_code" ]; then
    schema_code=$default_schema_code
fi
# array from action
arrAction=(${action//,/ })
all_required_params=()
# iterate through array using for loop
for i in "${arrAction[@]}"
do
    echo "$i"
    read -p "Enter number of Required Params of $i: " num_params
    required_params=()
    # check if required_params is empty
    if [[ -z "$num_params" || "$num_params" -eq 0 ]]; then
        required_params="[]"
    else
        for ((j=1; j<=num_params; j++)); do
            read -p "Enter name of param $j: " param_name
            read -p "Enter value of >> $param_name << : " param_value
            required_params+=( "{\"name\":\"$param_name\",\"value\":\"$param_value\"}" )
        done
        required_params="["$(echo ${required_params[@]} | tr ' ' ',')"]"
        echo $required_params
    fi
    all_required_params+=($required_params)
done
echo all_required_params="["$(echo ${all_required_params[@]} | tr ' ' ',')"]"
# # check if required_params is empty
# if [[ -z "$num_params" || "$num_params" -eq 0 ]]; then
#     required_params="[]"
# else
#     for ((i=1; i<=num_params; i++)); do
#         read -p "Enter name of param $i: " param_name
#         read -p "Enter value of >> $param_name << : " param_value
#         required_params+=( "{\"name\":\"$param_name\",\"value\":\"$param_value\"}" )
#     done
#     required_params=$(echo ${required_params[@]} | tr ' ' ',')
#     required_params="["$required_params"]"
#     echo $required_params
# fi

# sixnftd tx nftmngr perform-multi-token-action ${schema_code} ${token_id} ${action} ${ref_id} ${required_params} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake -y \
#     --chain-id sixnft
# ;;