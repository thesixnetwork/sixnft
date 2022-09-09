package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/evmsupport module sentinel errors
var (
	ErrAddressAlreadyBound = sdkerrors.Register(ModuleName, 1, "Address Already Bound")
	ErrVerifyingSignature  = sdkerrors.Register(ModuleName, 2, "Verifying Signature")
	ErrInvalidSignature    = sdkerrors.Register(ModuleName, 3, "Invalid Signature")
)
