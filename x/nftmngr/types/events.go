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
)
