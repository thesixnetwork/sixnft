package types

// Governance module event types
const (
	EventTypeCreateBinding          = "create_binding"
	EventTypeDeleteBinding          = "delete_binding"
	AttributeKeyCreateBindingResult = "craete_binding_result"
	AttributeKeyDeleteBindingResult = "delete_binding_result"
	AttributeValueBinder            = "native_address"
	AttributeValueEthAddress        = "eth_address"
	// Fail events
	AttributeValidateSignatureFail  = "validate_signature_fail"
	AttributeAlreadyBind  = "duplicate_pair"
)
