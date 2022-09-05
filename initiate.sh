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
    ewogICAgImNvZGUiOiAiYnVha2F3MSIsCiAgICAibmFtZSI6ICJCdWFrYXcxIiwKICAgICJvd25lciI6ICI2bmZ0MTVzOHlldjJjdjJ2ZTdsemVtNXl1Zmp5bm5mZTB1MnQ5c2xwd3ZuIiwKICAgICJvcmlnaW5fZGF0YSI6IHsKICAgICAgICAib3JpZ2luX2Jhc2VfdXJpIjogImh0dHBzOi8vYmsxbmZ0LnNpeG5ldHdvcmsuaW8vaXBmcy9RbWNvdkVjWnhpTTFNOWdmMzUzNWpVQ3BicUZaU2pjMmhIdVluWkViV05tdDVhIiwKICAgICAgICAib3JpZ2luX2NoYWluIjogImV0aGVyZXVtIiwKICAgICAgICAib3JpZ2luX2NvbnRyYWN0X2FkZHJlc3MiOiAiMHg5RjFDQzcwYjExZjQxMjlkMDQyZDAwMzdjMjA2NmQxMkUxNmQ5YTUyIiwKICAgICAgICAiYXR0cmlidXRlX292ZXJyaWRpbmciOiAiT1JJR0lOIiwKICAgICAgICAibWV0YWRhdGFfZm9ybWF0IjogIm9wZW5zZWEiLAogICAgICAgICJvcmlnaW5fYXR0cmlidXRlcyI6IFsKICAgICAgICAgICAgewogICAgICAgICAgICAgICAgIm5hbWUiOiAiYmFja2dyb3VuZF9sIiwKICAgICAgICAgICAgICAgICJkYXRhX3R5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICAgICAgICJkaXNwbGF5X3ZhbHVlX2ZpZWxkIjogInZhbHVlIiwKICAgICAgICAgICAgICAgICJkaXNwbGF5X29wdGlvbiI6IHsKICAgICAgICAgICAgICAgICAgICAib3BlbnNlYSI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgInRyYWl0X3R5cGUiOiAiQmFja2dyb3VuZCBMIgogICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgfSwKICAgICAgICAgICAgewogICAgICAgICAgICAgICAgIm5hbWUiOiAiYmFja2dyb3VuZF9yIiwKICAgICAgICAgICAgICAgICJkYXRhX3R5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICAgICAgICJkaXNwbGF5X3ZhbHVlX2ZpZWxkIjogInZhbHVlIiwKICAgICAgICAgICAgICAgICJkaXNwbGF5X29wdGlvbiI6IHsKICAgICAgICAgICAgICAgICAgICAib3BlbnNlYSI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgInRyYWl0X3R5cGUiOiAiQmFja2dyb3VuZCBSIgogICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgfSwKICAgICAgICAgICAgewogICAgICAgICAgICAgICAgIm5hbWUiOiAicGxhdGVfbCIsCiAgICAgICAgICAgICAgICAiZGF0YV90eXBlIjogInN0cmluZyIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV92YWx1ZV9maWVsZCI6ICJ2YWx1ZSIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV9vcHRpb24iOiB7CiAgICAgICAgICAgICAgICAgICAgIm9wZW5zZWEiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICJ0cmFpdF90eXBlIjogIlBsYXRlIEwiCiAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgfQogICAgICAgICAgICB9LAogICAgICAgICAgICB7CiAgICAgICAgICAgICAgICAibmFtZSI6ICJwbGF0ZV9yIiwKICAgICAgICAgICAgICAgICJkYXRhX3R5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICAgICAgICJkaXNwbGF5X3ZhbHVlX2ZpZWxkIjogInZhbHVlIiwKICAgICAgICAgICAgICAgICJkaXNwbGF5X29wdGlvbiI6IHsKICAgICAgICAgICAgICAgICAgICAib3BlbnNlYSI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgInRyYWl0X3R5cGUiOiAiUGxhdGUgUiIKICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgIH0sCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICJuYW1lIjogImJvZHlfbCIsCiAgICAgICAgICAgICAgICAiZGF0YV90eXBlIjogInN0cmluZyIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV92YWx1ZV9maWVsZCI6ICJ2YWx1ZSIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV9vcHRpb24iOiB7CiAgICAgICAgICAgICAgICAgICAgIm9wZW5zZWEiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICJ0cmFpdF90eXBlIjogIkJvZHkgTCIKICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgIH0sCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICJuYW1lIjogImJvZHlfciIsCiAgICAgICAgICAgICAgICAiZGF0YV90eXBlIjogInN0cmluZyIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV92YWx1ZV9maWVsZCI6ICJ2YWx1ZSIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV9vcHRpb24iOiB7CiAgICAgICAgICAgICAgICAgICAgIm9wZW5zZWEiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICJ0cmFpdF90eXBlIjogIkJvZHkgUiIKICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgIH0sCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICJuYW1lIjogImhlYWRfbCIsCiAgICAgICAgICAgICAgICAiZGF0YV90eXBlIjogInN0cmluZyIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV92YWx1ZV9maWVsZCI6ICJ2YWx1ZSIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV9vcHRpb24iOiB7CiAgICAgICAgICAgICAgICAgICAgIm9wZW5zZWEiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICJ0cmFpdF90eXBlIjogIkhlYWQgTCIKICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgIH0sCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICJuYW1lIjogImhlYWRfciIsCiAgICAgICAgICAgICAgICAiZGF0YV90eXBlIjogInN0cmluZyIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV92YWx1ZV9maWVsZCI6ICJ2YWx1ZSIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV9vcHRpb24iOiB7CiAgICAgICAgICAgICAgICAgICAgIm9wZW5zZWEiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICJ0cmFpdF90eXBlIjogIkhlYWQgUiIKICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgIH0sCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICJuYW1lIjogImNsb3RoZXNfbCIsCiAgICAgICAgICAgICAgICAiZGF0YV90eXBlIjogInN0cmluZyIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV92YWx1ZV9maWVsZCI6ICJ2YWx1ZSIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV9vcHRpb24iOiB7CiAgICAgICAgICAgICAgICAgICAgIm9wZW5zZWEiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICJ0cmFpdF90eXBlIjogIkNsb3RoZXMgTCIKICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgIH0sCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICJuYW1lIjogImNsb3RoZXNfciIsCiAgICAgICAgICAgICAgICAiZGF0YV90eXBlIjogInN0cmluZyIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV92YWx1ZV9maWVsZCI6ICJ2YWx1ZSIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV9vcHRpb24iOiB7CiAgICAgICAgICAgICAgICAgICAgIm9wZW5zZWEiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICJ0cmFpdF90eXBlIjogIkNsb3RoZXMgUiIKICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgIH0sCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICJuYW1lIjogImV4dHJhX2wiLAogICAgICAgICAgICAgICAgImRhdGFfdHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgICAgICAgImRpc3BsYXlfdmFsdWVfZmllbGQiOiAidmFsdWUiLAogICAgICAgICAgICAgICAgImRpc3BsYXlfb3B0aW9uIjogewogICAgICAgICAgICAgICAgICAgICJvcGVuc2VhIjogewogICAgICAgICAgICAgICAgICAgICAgICAidHJhaXRfdHlwZSI6ICJFeHRyYSBMIgogICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgfSwKICAgICAgICAgICAgewogICAgICAgICAgICAgICAgIm5hbWUiOiAiZXh0cmFfciIsCiAgICAgICAgICAgICAgICAiZGF0YV90eXBlIjogInN0cmluZyIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV92YWx1ZV9maWVsZCI6ICJ2YWx1ZSIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV9vcHRpb24iOiB7CiAgICAgICAgICAgICAgICAgICAgIm9wZW5zZWEiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICJ0cmFpdF90eXBlIjogIkV4dHJhIFIiCiAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgfQogICAgICAgICAgICB9LAogICAgICAgICAgICB7CiAgICAgICAgICAgICAgICAibmFtZSI6ICJoYW5kX2wiLAogICAgICAgICAgICAgICAgImRhdGFfdHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgICAgICAgImRpc3BsYXlfdmFsdWVfZmllbGQiOiAidmFsdWUiLAogICAgICAgICAgICAgICAgImRpc3BsYXlfb3B0aW9uIjogewogICAgICAgICAgICAgICAgICAgICJvcGVuc2VhIjogewogICAgICAgICAgICAgICAgICAgICAgICAidHJhaXRfdHlwZSI6ICJIYW5kIEwiCiAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgfQogICAgICAgICAgICB9LAogICAgICAgICAgICB7CiAgICAgICAgICAgICAgICAibmFtZSI6ICJoYW5kX3IiLAogICAgICAgICAgICAgICAgImRhdGFfdHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgICAgICAgImRpc3BsYXlfdmFsdWVfZmllbGQiOiAidmFsdWUiLAogICAgICAgICAgICAgICAgImRpc3BsYXlfb3B0aW9uIjogewogICAgICAgICAgICAgICAgICAgICJvcGVuc2VhIjogewogICAgICAgICAgICAgICAgICAgICAgICAidHJhaXRfdHlwZSI6ICJIYW5kIFIiCiAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgfQogICAgICAgICAgICB9LAogICAgICAgICAgICB7CiAgICAgICAgICAgICAgICAibmFtZSI6ICJpbmZsdWVuY2VyIiwKICAgICAgICAgICAgICAgICJkYXRhX3R5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICAgICAgICJkaXNwbGF5X3ZhbHVlX2ZpZWxkIjogInZhbHVlIiwKICAgICAgICAgICAgICAgICJkaXNwbGF5X29wdGlvbiI6IHsKICAgICAgICAgICAgICAgICAgICAib3BlbnNlYSI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgInRyYWl0X3R5cGUiOiAiSW5mbHVlbmNlciIKICAgICAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgIH0KICAgICAgICBdCiAgICB9LAogICAgIm9uY2hhaW5fZGF0YSI6IHsKICAgICAgICAibmZ0X2F0dHJpYnV0ZXMiOiBbCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICJuYW1lIjogImV4cGlyZV9kYXRlIiwKICAgICAgICAgICAgICAgICJkYXRhX3R5cGUiOiAibG9uZyIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV92YWx1ZV9maWVsZCI6ICJ2YWx1ZSIsCiAgICAgICAgICAgICAgICAiZGlzcGxheV9vcHRpb24iOiB7CiAgICAgICAgICAgICAgICAgICAgIm9wZW5zZWEiOiB7CiAgICAgICAgICAgICAgICAgICAgICAgICJkaXNwbGF5X3R5cGUiOiAiZGF0ZSIsCiAgICAgICAgICAgICAgICAgICAgICAgICJ0cmFpdF90eXBlIjogImJpcnRoZGF5IgogICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgfQogICAgICAgIF0sCiAgICAgICAgInRva2VuX2F0dHJpYnV0ZXMiOiBbCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICJuYW1lIjogImV4Y2x1c2l2ZV9wYXJ0eV9hY2Nlc3MiLAogICAgICAgICAgICAgICAgImRhdGFfdHlwZSI6ICJib29sZWFuIiwKICAgICAgICAgICAgICAgICJoaWRkZW5fdG9fbWFya2V0cGxhY2UiOiB0cnVlCiAgICAgICAgICAgIH0sCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICJuYW1lIjogImZpcnN0X2Rpc2NvdW50X3VzZWQiLAogICAgICAgICAgICAgICAgImRhdGFfdHlwZSI6ICJib29sZWFuIiwKICAgICAgICAgICAgICAgICJoaWRkZW5fdG9fbWFya2V0cGxhY2UiOiB0cnVlCiAgICAgICAgICAgIH0sCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICJuYW1lIjogInBlcmNlbnRfZGlzY291bnRfMTAiLAogICAgICAgICAgICAgICAgImRlZmF1bHRfbWludF92YWx1ZSI6ICIxMCIsCiAgICAgICAgICAgICAgICAiZGF0YV90eXBlIjogImludGVyZ2VyIiwKICAgICAgICAgICAgICAgICJkaXNwbGF5X3ZhbHVlX2ZpZWxkIjogInZhbHVlIiwKICAgICAgICAgICAgICAgICJkaXNwbGF5X29wdGlvbiI6IHsKICAgICAgICAgICAgICAgICAgICAib3BlbnNlYSI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgInRyYWl0X3R5cGUiOiAiQXZhaWxhYmxlIDEwIFBlcmNlbnQgRGlzY291bnQiLAogICAgICAgICAgICAgICAgICAgICAgICAibWF4X3ZhbHVlIjogNSwKICAgICAgICAgICAgICAgICAgICAgICAgImRpc3BsYXlfdHlwZSI6ICJudW1iZXIiCiAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgfQogICAgICAgICAgICB9LAogICAgICAgICAgICB7CiAgICAgICAgICAgICAgICAibmFtZSI6ICJkaXNjb3VudF9wZXJjZW50YWdlIiwKICAgICAgICAgICAgICAgICJkZWZhdWx0X21pbnRfdmFsdWUiOiAiMCIsCiAgICAgICAgICAgICAgICAiZGF0YV90eXBlIjogImludGVyZ2VyIiwKICAgICAgICAgICAgICAgICJkaXNwbGF5X3ZhbHVlX2ZpZWxkIjogInZhbHVlIiwKICAgICAgICAgICAgICAgICJkaXNwbGF5X29wdGlvbiI6IHsKICAgICAgICAgICAgICAgICAgICAib3BlbnNlYSI6IHsKICAgICAgICAgICAgICAgICAgICAgICAgInRyYWl0X3R5cGUiOiAiRGlzY291bnQiLAogICAgICAgICAgICAgICAgICAgICAgICAiZGlzcGxheV90eXBlIjogImJvb3N0X3BlcmNlbnRhZ2UiCiAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgfQogICAgICAgICAgICB9LAogICAgICAgICAgICB7CiAgICAgICAgICAgICAgICAibmFtZSI6ICJkaXNjb3VudF9hbW91bnQiLAogICAgICAgICAgICAgICAgImRlZmF1bHRfbWludF92YWx1ZSI6ICIwIiwKICAgICAgICAgICAgICAgICJkYXRhX3R5cGUiOiAiaW50ZXJnZXIiLAogICAgICAgICAgICAgICAgImRpc3BsYXlfdmFsdWVfZmllbGQiOiAidmFsdWUiLAogICAgICAgICAgICAgICAgImRpc3BsYXlfb3B0aW9uIjogewogICAgICAgICAgICAgICAgICAgICJvcGVuc2VhIjogewogICAgICAgICAgICAgICAgICAgICAgICAidHJhaXRfdHlwZSI6ICJEaXNjb3VudCIsCiAgICAgICAgICAgICAgICAgICAgICAgICJkaXNwbGF5X3R5cGUiOiAiYm9vc3RfbnVtYmVyIgogICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgfQogICAgICAgIF0sCiAgICAgICAgImFjdGlvbnMiOiBbCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAgICJuYW1lIjogInVzZTEwcGVyY2VudGRpc2NvdW50IiwKICAgICAgICAgICAgICAgICJkZXNjIjogIlVzZSAxMCUgZGlzY291bnQgb24gcHVyY2hhc2UiLAogICAgICAgICAgICAgICAgIndoZW4iOiAibWV0YS5HZXROdW1iZXIoJ3BlcmNlbnRfZGlzY291bnRfMTAnKSA


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

