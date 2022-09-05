package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nftmngr module sentinel errors
var (
	ErrParsingBase64        = sdkerrors.Register(ModuleName, 2, "Error parsing base64")
	ErrParsingSchemaMessage = sdkerrors.Register(ModuleName, 3, "Error parsing schema")
	ErrSchemaAlreadyExists  = sdkerrors.Register(ModuleName, 4, "Schema already exists")
)
