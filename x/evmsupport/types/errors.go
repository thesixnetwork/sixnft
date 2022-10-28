package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/evmsupport module sentinel errors
var (
	ErrAddressAlreadyBound   = sdkerrors.Register(ModuleName, 1, "Address Already Bound")
	ErrVerifyingSignature    = sdkerrors.Register(ModuleName, 2, "Verifying Signature")
	ErrInvalidSignature      = sdkerrors.Register(ModuleName, 3, "Invalid Signature")
	ErrAddressNotBound       = sdkerrors.Register(ModuleName, 4, "Binding Not Found")
	ErrInvalidBase64         = sdkerrors.Register(ModuleName, 5, "Invalid Base64")
	ErrInvalidHex            = sdkerrors.Register(ModuleName, 6, "Invalid Hex")
	ErrParsingSetSignerParam = sdkerrors.Register(ModuleName, 7, "Parsing Action Params")
)
