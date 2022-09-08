package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nftmngr module sentinel errors
var (
	// General (0-99)
	ErrParsingBase64          = sdkerrors.Register(ModuleName, 2, "Error parsing base64")
	ErrParsingSchemaMessage   = sdkerrors.Register(ModuleName, 3, "Error parsing schema")
	ErrParsingMetadataMessage = sdkerrors.Register(ModuleName, 5, "Error parsing metadata")

	// Metadata (100-199)
	ErrMetadataAlreadyExists                     = sdkerrors.Register(ModuleName, 100, "Metadata already exists")
	ErrValidatingMetadata                        = sdkerrors.Register(ModuleName, 101, "Error validating metadata")
	ErrSchemaDoesNotExists                       = sdkerrors.Register(ModuleName, 102, "Schema does not exists")
	ErrDuplicateOnchainAttributesValue           = sdkerrors.Register(ModuleName, 103, "Duplicate onchain attributes value")
	ErrDuplicateOriginAttributesValue            = sdkerrors.Register(ModuleName, 104, "Duplicate origin attributes value")
	ErrOriginAttributesNotSameTypeAsSchema       = sdkerrors.Register(ModuleName, 105, "Origin attributes not same type as schema")
	ErrOnchainTokenAttributesNotSameTypeAsSchema = sdkerrors.Register(ModuleName, 106, "Onchain token attributes not same type as schema")
	ErrOnchainNFTAttributesNotSameTypeAsSchema   = sdkerrors.Register(ModuleName, 107, "Onchain NFT attributes not same type as schema")

	// Schema (200-299)
	ErrSchemaAlreadyExists             = sdkerrors.Register(ModuleName, 200, "Schema already exists")
	ErrValidatingNFTSchema             = sdkerrors.Register(ModuleName, 201, "Error validating NFT schema")
	ErrDuplicateOriginAttributes       = sdkerrors.Register(ModuleName, 202, "Duplicate origin attributes")
	ErrDuplicateOnchainNFTAttributes   = sdkerrors.Register(ModuleName, 203, "Duplicate onchain NFT attributes")
	ErrDuplicateOnchainTokenAttributes = sdkerrors.Register(ModuleName, 204, "Duplicate onchain token attributes")
	ErrDuplicateAttributesValue        = sdkerrors.Register(ModuleName, 205, "Duplicate attributes value")
	ErrSameTypeNFTAttributes           = sdkerrors.Register(ModuleName, 206, "Same type NFT attributes")
	ErrSameTypeTokenAttributes         = sdkerrors.Register(ModuleName, 207, "Same type token attributes")
)
