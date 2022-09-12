package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/admin module sentinel errors
var (
	ErrAuthorizationNotFound = sdkerrors.Register(ModuleName, 1, "authorization not found")
	ErrUnauthorized          = sdkerrors.Register(ModuleName, 2, "unauthorized")
)
