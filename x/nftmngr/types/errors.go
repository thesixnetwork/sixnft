package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nftmngr module sentinel errors
var (
	// General (0-99)
	ErrParsingBase64          = sdkerrors.Register(ModuleName, 1, "Error parsing base64")
	ErrParsingSchemaMessage   = sdkerrors.Register(ModuleName, 2, "Error parsing schema")
	ErrParsingMetadataMessage = sdkerrors.Register(ModuleName, 3, "Error parsing metadata")
	ErrCreatorDoesNotMatch    = sdkerrors.Register(ModuleName, 4, "Only creator can run this action")
	ErrMetadataDoesNotExists  = sdkerrors.Register(ModuleName, 5, "Metadata of the toekn does not exists")
	ErrParsingAttributeValue  = sdkerrors.Register(ModuleName, 6, "Error parsing attribute value")

	// Metadata (100-199)
	ErrMetadataAlreadyExists                     = sdkerrors.Register(ModuleName, 100, "Metadata already exists")
	ErrValidatingMetadata                        = sdkerrors.Register(ModuleName, 101, "Error validating metadata")
	ErrSchemaDoesNotExists                       = sdkerrors.Register(ModuleName, 102, "Schema does not exists")
	ErrDuplicateOnchainAttributesValue           = sdkerrors.Register(ModuleName, 103, "Duplicate onchain attributes value")
	ErrDuplicateOriginAttributesValue            = sdkerrors.Register(ModuleName, 104, "Duplicate origin attributes value")
	ErrOriginAttributesNotSameTypeAsSchema       = sdkerrors.Register(ModuleName, 105, "Origin attributes not same type as schema")
	ErrOnchainTokenAttributesNotSameTypeAsSchema = sdkerrors.Register(ModuleName, 106, "Onchain token attributes not same type as schema")
	ErrOnchainNFTAttributesNotSameTypeAsSchema   = sdkerrors.Register(ModuleName, 107, "Onchain NFT attributes not same type as schema")
	ErrAttributeOverriding                       = sdkerrors.Register(ModuleName, 108, "Attribute overriding is not allowed")
	ErrAttributeTypeNotMatch                     = sdkerrors.Register(ModuleName, 109, "Attribute type does not match")
	ErrAttributeNotFoundForAction                = sdkerrors.Register(ModuleName, 110, "Attribute not found for action")
	ErrAttributeDefinitionTypesMismatch          = sdkerrors.Register(ModuleName, 111, "Attribute definition type is mismatch")
	ErrAttributeAlreadyExists                    = sdkerrors.Register(ModuleName, 112, "Attribute already exists")
	ErrInvalidAttribute                          = sdkerrors.Register(ModuleName, 113, "Invalid attribute")
	ErrRequiredAttributeMissing                  = sdkerrors.Register(ModuleName, 114, "Required attribute missing")
	ErrOnchainAttributesNotExistsInSchema        = sdkerrors.Register(ModuleName, 115, "Onchain attributes not exists in schema")
	ErrNftDataDoesNotExists                      = sdkerrors.Register(ModuleName, 116, "NFT data does not exists")
	ErrNoDefaultValue                            = sdkerrors.Register(ModuleName, 117, "No default value")
	// Schema (200-299)
	ErrSchemaAlreadyExists             = sdkerrors.Register(ModuleName, 200, "Schema already exists")
	ErrValidatingNFTSchema             = sdkerrors.Register(ModuleName, 201, "Error validating NFT schema")
	ErrDuplicateOriginAttributes       = sdkerrors.Register(ModuleName, 202, "Duplicate origin attributes")
	ErrDuplicateOnchainNFTAttributes   = sdkerrors.Register(ModuleName, 203, "Duplicate onchain NFT attributes")
	ErrDuplicateOnchainTokenAttributes = sdkerrors.Register(ModuleName, 204, "Duplicate onchain token attributes")
	ErrDuplicateAttributesValue        = sdkerrors.Register(ModuleName, 205, "Duplicate attributes value")
	ErrSameTypeNFTAttributes           = sdkerrors.Register(ModuleName, 206, "Same type NFT attributes")
	ErrSameTypeTokenAttributes         = sdkerrors.Register(ModuleName, 207, "Same type token attributes")
	ErrAttributeDoesNotExists          = sdkerrors.Register(ModuleName, 208, "Attribute does not exists")
	ErrNotSameTypeDefaultMintValue     = sdkerrors.Register(ModuleName, 209, "Not same type default mint value")
	ErrInvalidAccdress                 = sdkerrors.Register(ModuleName, 210, "Invalid address")
	ErrAttributeAlreadyHidden          = sdkerrors.Register(ModuleName, 211, "Attribute already hidden")
	ErrDuplicateActionName             = sdkerrors.Register(ModuleName, 212, "Duplicate action name")

	// Action (300-399)
	ErrRefIdAlreadyExists     = sdkerrors.Register(ModuleName, 300, "RefId already exists")
	ErrEmptyChangeList        = sdkerrors.Register(ModuleName, 301, "No changes were updated")
	ErrActionAlreadyExists    = sdkerrors.Register(ModuleName, 302, "Action already exists")
	ErrInvalidActionAttribute = sdkerrors.Register(ModuleName, 303, "Invalid action attribute")
	ErrActionIsDisabled       = sdkerrors.Register(ModuleName, 304, "Action is disabled")
	ErrActionIsForUserOnly    = sdkerrors.Register(ModuleName, 305, "This action can be done by user only")
	ErrActionIsForSystemOnly  = sdkerrors.Register(ModuleName, 306, "This action can be done by system only")
	ErrInvalidParameter	   	  = sdkerrors.Register(ModuleName, 307, "Invalid parameter")

	ErrOrganizationOwner = sdkerrors.Register(ModuleName, 400, "Unauthorized organization owner")

	ErrInvalidAttributeName = sdkerrors.Register(ModuleName, 500, "Invalid attribute name")
	ErrInvalidActionName    = sdkerrors.Register(ModuleName, 501, "Invalid action name")

	ErrSystemActionerAlreadyExists = sdkerrors.Register(ModuleName, 600, "System actioner already exists")
	ErrSystemActionerDoesNotExists = sdkerrors.Register(ModuleName, 601, "System actioner does not exists")

	ErrUnauthorized = sdkerrors.Register(ModuleName, 700, "Unauthorized")

	ErrNoNftFeeAdminPermission = sdkerrors.Register(ModuleName, 800, "No NFT fee admin permission")
	ErrInvalidFeeAmount        = sdkerrors.Register(ModuleName, 801, "Invalid fee amount")
	ErrProcessingFee           = sdkerrors.Register(ModuleName, 802, "Error processing fee")
	ErrInvalidFeePortion       = sdkerrors.Register(ModuleName, 803, "Invalid fee portion")
)
