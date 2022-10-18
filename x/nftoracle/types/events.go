package types

const (
	EventTypeMintRequestCreated = "mint_request_created"
	AttributeKeyMintRequestID   = "mint_request_id"
	AttributeKeyNftSchemaCode   = "nft_schema_code"
	AttributeKeyTokenID         = "token_id"
	AttributeKeyRequiredConfirm = "required_confirm"

	EventTypeMintRequest          = "mint_request"
	AttributeKeyMintRequestStatus = "status"

	EventTypeMintRequestConfirmed = "mint_request_confirmed"

	EventTypeActionRequestCreated = "action_request_created"
	AttributeKeyActionRequestID   = "action_request_id"

	EventTypeActionRequest          = "action_request"
	AttributeKeyActionRequestStatus = "status"

	EventMessage                    = "action"
	EventTypeRunAction              = "run_action"
	AttributeKeyRunActionResult     = "run_action_result"
	AttributeKeyRunActionChangeList = "run_action_changelist"

	EventTypeVerificationRequestCreated   = "verify_request_created"
	AttributeKeyVerifyRequestID           = "verify_request_id"
	AttributeKeyVerificationRequestStatus = "status"

	EventTypeVerificationRequest       = "verify_request"
	EventTypeVerificationRequestStatus = "status"

	AttributeRequestorAddress = "requestor_address"

	EventTypeSetMinimumConfirmation = "set_minimum_confirmation"
	AttributeKeyMinimumConfirmation = "minimum_confirmation"
)
