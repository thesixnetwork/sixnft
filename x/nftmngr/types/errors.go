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
	ErrDuplicateAttributes    = sdkerrors.Register(ModuleName, 8, "Duplicate attributes")
	ErrValidatingNFTSchema    = sdkerrors.Register(ModuleName, 9, "Error validating NFT schema")
	ErrValidatingMetadata     = sdkerrors.Register(ModuleName, 10, "Error validating metadata")
)
