package types

const (
	EventMessage                    = "action"
	EventTypeRunAction              = "run_action"
	AttributeKeyRunActionResult     = "run_action_result"
	AttributeKeyRunActionChangeList = "run_action_changelist"
	AttributeKeyRunActionRefId      = "run_action_ref_id"

	EventTypeAddAction          = "add_action"
	AttributeKeyAddActionName   = "add_action_name"
	AttributeKeyAddActionResult = "add_action_result"

	EventTypeAddAttribute            = "add_attribute"
	AttributeKeyAddAttributeName     = "add_attribute_name"
	AttributeKeyAddAttributeLocation = "add_attribute_location"
	AttributeKeyAddAttributeResult   = "add_attribute_result"

	EventTypeSchemaOwnerChanged = "schema_owner_changed"
	AttributeKeyNftSchemaCode   = "nft_schema_code"
	AttributeKeyNewOwner        = "new_owner"

	EventTypeAddSystemActioner    = "add_system_actioner"
	EventTypeRemoveSystemActioner = "remove_system_actioner"
	AttributeKeyActioner          = "actioner"

	EventTypeCreateSchema          = "craete_schema"
	AttributeKeyCreateSchemaCode   = "create_schema_code"
	AttributeKeyCreateSchemaResult = "create_schema_result"

	EventTypeCreateMetadata           = "create_metadata"
	AttributeKeyCreateMetaDataTokenID = "token_id"
	AttributeKeyCreateMetaDataResult  = "create_metadata_result"

	//set base uri
	EventTypeSetBaseURI          = "set_base_uri"
	AttributeKeySetBaseURI       = "base_uri"
	AttributeKeySetBaseURIResult = "set_base_uri_result"

	//set method retrieval
	EventTypeSetRetrievalMethod    = "set_retrieval_method"
	AttributeKeySetRetrievalMethod = "retrieval_method"
	AttributeKeySetRetrivalResult   = "set_retrieval_method_result"

	//set nft schema value
	EventTypeSetNFTSchemaValue          = "set_nft_schema_value"
	AttributeKeySetNFTSchemaValue       = "nft_schema_new_value"
	AttributeKeySetNFTSchemaValueResult = "set_nft_schema_value_result"

	//toggle nft action
	EventTypeToggleNFTAction          = "toggle_nft_action"
	AttributeKeyToggleNFTAction       = "nft_action"
	AttributeKeyToggleNFTActionResult = "toggle_nft_action_result"

	//Hidden attribute
	EventTypeHiddenAttribute          = "hide_attribute"
	AttributeKeyHiddenAttributeResult = "hide_attribute_result"

	//Show attribute
	EventTypeShowAttribute          = "show_attribute"
	AttributeKeyShowAttributeResult = "show_attribute_result"
	AttributeKeyTokenId             = "token_id"
	EventTypeResyncAttributes       = "resync_attributes"

	EventTypeSetFeeConfig  = "set_fee_config"
	AttributeKeyFeeSubject = "fee_subject"
	AttributeKeyFeeConfig  = "fee_config"

	EventTypeSetMintAuth         = "set_mint_auth"
	AttributeAutorizeTo          = "authorize_to"
	AttributeKeySetMinAuthResult = "add_action_result"

	EventTypeChangeOrgOwner = "change_org_owner"
	AttributeKeyOrgName     = "org_name"
	AttributeKeyOldOwner    = "old_owner"
)
