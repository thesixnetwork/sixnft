package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nftmngr module sentinel errors
var (
	ErrParsingBase64          = sdkerrors.Register(ModuleName, 2, "Error parsing base64")
	ErrParsingSchemaMessage   = sdkerrors.Register(ModuleName, 3, "Error parsing schema")
	ErrSchemaAlreadyExists    = sdkerrors.Register(ModuleName, 4, "Schema already exists")
	ErrParsingMetadataMessage = sdkerrors.Register(ModuleName, 5, "Error parsing metadata")
	ErrMetadataAlreadyExists  = sdkerrors.Register(ModuleName, 6, "Metadata already exists")
	ErrSchemaDoesNotExists    = sdkerrors.Register(ModuleName, 7, "Schema does not exists")
	ErrCreatorDoesNotMatch    = sdkerrors.Register(ModuleName, 8, "Only creator can run this action")
	ErrMetadataDoesNotExists  = sdkerrors.Register(ModuleName, 9, "Metadata of the toekn does not exists")

	ErrAttributeOverriding   = sdkerrors.Register(ModuleName, 100, "Attribute overriding is not allowed")
	ErrAttributeTypeNotMatch = sdkerrors.Register(ModuleName, 101, "Attribute type does not match")
)
