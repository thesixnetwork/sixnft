# Create new chain
ignite scaffold chain sixnft --no-module --address-prefix 6nft

# Create module nftmngr
ignite scaffold module nftmngr

# Create NFT Data type fot NFT Schema
ignite scaffold type OpenseaDisplayOption display_type trait_type max_value \
    --no-message \
    --no-simulation \
    --module nftmngr

# maxValue is a string, so we modify the type to be a uint64 and also change format of the fields to be under_score_case
# message OpenseaDisplayOption {
  
#   string display_type = 1; 
#   string trait_type = 2; 
#   uint64 max_value = 3; 
# }
ignite scaffold type DisplayOption opensea \
    --no-message \
    --no-simulation \
    --module nftmngr
# modify proto file to be like this

# syntax = "proto3";
# package sixnft.nftmngr;
# import "nftmngr/opensea_display_option.proto";
# option go_package = "sixnft/x/nftmngr/types";

# message DisplayOption {
  
#   OpenseaDisplayOption opensea = 1; 
# }    

# Attribute Definition
ignite scaffold type AttributeDefinition name data_type display_value_field display_option default_mint_value  \
    --no-message \
    --no-simulation \
    --module nftmngr
# Modify proto file to be like this
# syntax = "proto3";
# package sixnft.nftmngr;
# import "nftmngr/display_option.proto";
# option go_package = "sixnft/x/nftmngr/types";

# message AttributeDefinition {
  
#   string name = 1; 
#   string data_type = 2; 
#   string display_value_field = 3; 
#   DisplayOption display_option = 4; 
#   string default_mint_value = 5; 
# }

ignite scaffold type OriginData origin_chain origin_contract_address origin_base_uri attribute_overriding metadata_format origin_attributes \
    --no-message \
    --no-simulation \
    --module nftmngr
# Modify proto file to be like this
# syntax = "proto3";
# package sixnft.nftmngr;
# import "nftmngr/attribute_definition.proto";
# option go_package = "sixnft/x/nftmngr/types";

# enum AttributeOverriding {
#   ORIGIN = 0;
#   CHAIN = 1;
# }

# message OriginData {
  
#   string origin_chain = 1; 
#   string origin_contract_address = 2; 
#   string origin_base_uri = 3; 
#   AttributeOverriding  attribute_overriding = 4; 
#   string metadata_format = 5; 
#   repeated AttributeDefinition origin_attributes = 6; 
# }

ignite scaffold type Action name desc when then \
    --no-message \
    --no-simulation \
    --module nftmngr
# Modify proto file to be like this
# syntax = "proto3";
# package sixnft.nftmngr;

# option go_package = "sixnft/x/nftmngr/types";

# message Action {
  
#   string name = 1; 
#   string desc = 2; 
#   string when = 3; 
#   repeated string then = 4; 
# }
ignite scaffold type OnChainData nft_attributes token_attributes actions then \
    --no-message \
    --no-simulation \
    --module nftmngr
# Modify proto file to be like this
# syntax = "proto3";
# package sixnft.nftmngr;
# import "nftmngr/attribute_definition.proto";
# import "nftmngr/action.proto";
# option go_package = "sixnft/x/nftmngr/types";

# message OnChainData {
  
#   repeated AttributeDefinition nft_attributes = 1; 
#   repeated AttributeDefinition token_attributes = 2; 
#   repeated Action actions = 3; 
# }
ignite scaffold map NFTSchema name owner origin_data onchain_data \
    --index code \
    --no-message \
    --no-simulation \
    --module nftmngr
# Modify proto file to be like this
# syntax = "proto3";
# package sixnft.nftmngr;


# import "nftmngr/on_chain_data.proto";
# import "nftmngr/origin_data.proto";

# option go_package = "sixnft/x/nftmngr/types";

# message NFTSchema {
#   string code = 1; 
#   string name = 2; 
#   string owner = 3; 
#   OriginData origin_data = 4; 
#   OnChainData onchain_data = 5; 
  
# }
ignite scaffold message createNFTSchema nft_schema_base64 \
    --desc "To create NFT schema from JSON in Base64 format" \
    --response code \
    --no-simulation \
    --module nftmngr

# Unmarshal JSON in msg_server_xxxxx
# func (k msgServer) CreateNFTSchema(goCtx context.Context, msg *types.MsgCreateNFTSchema) (*types.MsgCreateNFTSchemaResponse, error) {
# 	ctx := sdk.UnwrapSDKContext(goCtx)

# 	jsonSchema, err := base64.StdEncoding.DecodeString(msg.NftSchemaBase64)
# 	if err != nil {
# 		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
# 	}
# 	schema := types.NFTSchema{}
# 	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(jsonSchema, &schema)
# 	if err != nil {
# 		return nil, sdkerrors.Wrap(types.ErrParsingSchemaMessage, err.Error())
# 	}

# 	fmt.Println("####################### ", schema.Code)

# 	// TODO: Handling the message
# 	_ = ctx

# 	return &types.MsgCreateNFTSchemaResponse{}, nil
# }

# Test create NFTSchema
sixnftd tx nftmngr create-nft-schema --from alice \
    <BASE64 schema>

sixnftd query nftmngr show-nft-schema \
    <NFT schema code>


ignite scaffold type Status first_mint_complete \
    --no-message \
    --no-simulation \
    --module nftmngr

ignite scaffold type OnOffSwitch active \
    --no-message \
    --no-simulation \
    --module nftmngr


# Create NFT Attributes Value
ignite scaffold type NftAttributeValue name value \
    --no-message \
    --no-simulation \
    --module nftmngr

# Create NFT Data
ignite scaffold map NftData token_owner origin_image origin_attributes onchain_attributes \
    --index nft_schema_code,token_id \
    --no-message \
    --no-simulation \
    --module nftmngr
# Modify proto file
# syntax = "proto3";
# package sixnft.nftmngr;

# import "nftmngr/nft_attribute_value.proto";
# option go_package = "sixnft/x/nftmngr/types";

# message NftData {
#   string nftSchemaCode = 1; 
#   uint64 tokenId = 2; 
#   string tokenOwner = 3; 
#   string originImage = 4; 
#   repeated NftAttributeValue originAttributes = 5; 
#   repeated NftAttributeValue onchainAttributes = 6; 
  
# }

ignite scaffold message createMetadata nft_schema_code token_id base64NFTData \
    --desc "To create NFT Metadata from Base64" \
    --response nft_schema_code,token_id \
    --no-simulation \
    --module nftmngr

#Test create metadata
sixnftd tx nftmngr create-metadata buakaw1 1 --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake \
    <BASE64 NFT Data>


# Create Action
ignite scaffold message performActionByAdmin nft_schema_code token_id action \
    --desc "To do action" \
    --response nft_schema_code,token_id \
    --no-simulation \
    --module nftmngr


# Test create action
sixnftd tx nftmngr perform-action-by-admin buakaw1 1 use10percentdiscount --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1stake

# Create evm_support module

ignite scaffold module evm_support 

ignite scaffold map AddressBinding \
    --index eth_address:string,native_address:string \
    --no-message \
    --no-simulation \
    --module evm_support

ignite scaffold message bindAddress eth_address:string signature:string signed_message:string \
    --desc "To bind EVM address with native address" \
    --response eth_address,native_address \
    --no-simulation \
    --module evm_support

ignite scaffold message removeBinding eth_address:string signature:string signed_message:string \
    --desc "To remove binding address" \
    --response eth_address \
    --no-simulation \
    --module evm_support