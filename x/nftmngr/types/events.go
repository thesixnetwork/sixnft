package types

const (
	EventMessage                    = "action"
	EventTypeRunAction              = "run_action"
	AttributeKeyRunActionResult     = "run_action_result"
	AttributeKeyRunActionChangeList = "run_action_changelist"

	EventTypeAddAction          = "add_action"
	AttributeKeyAddActionName   = "add_action_name"
	AttributeKeyAddActionResult = "add_action_result"

	EventTypeAddAttribute          = "add_attribute"
	AttributeKeyAddAttributeName   = "add_attribute_name"
	AttributeKeyAddAttributeResult = "add_attribute_result"

	EventTypeSchemaOwnerChanged = "schema_owner_changed"
	AttributeKeyNftSchemaCode   = "nft_schema_code"
	AttributeKeyNewOwner        = "new_owner"

	EventTypeAddSystemActioner    = "add_system_actioner"
	EventTypeRemoveSystemActioner = "remove_system_actioner"
	AttributeKeyActioner          = "actioner"

	EventTypeCreateSchema          = "craete_schema"
	AttributeKeyCreateSchemaCode   = "create_schema_code"
	AttributeKeyCreateSchemaResult = "create_schema_result"

	EventTypeCreateMetadata = "create_metadata"
	AttributeKeyCreateMetaDataTokenID = "token_id"
	AttributeKeyCreateMetaDataResult = "create_metadata_result"

	//set base uri
	EventTypeSetBaseURI = "set_base_uri"
	AttributeKeySetBaseURI = "base_uri"
	AttributeKeySetBaseURIResult = "set_base_uri_result"

	//set nft schema value
	EventTypeSetNFTSchemaValue = "set_nft_schema_value"
	AttributeKeySetNFTSchemaValue = "nft_schema_new_value"
	AttributeKeySetNFTSchemaValueResult = "set_nft_schema_value_result"

	//toggle nft action
	EventTypeToggleNFTAction = "toggle_nft_action"
	AttributeKeyToggleNFTAction = "nft_action"
	AttributeKeyToggleNFTActionResult = "toggle_nft_action_result"


)
