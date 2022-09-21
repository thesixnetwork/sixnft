package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nftoracle module sentinel errors
var (
	ErrNFTSchemaNotFound                   = sdkerrors.Register(ModuleName, 1, "NFTSchema not found")
	ErrMetadataAlreadyExists               = sdkerrors.Register(ModuleName, 2, "Metadata already exists")
	ErrNoOraclePermission                  = sdkerrors.Register(ModuleName, 3, "No oracle permission")
	ErrMintRequestNotFound                 = sdkerrors.Register(ModuleName, 4, "Mint request not found")
	ErrMintRequestNotPending               = sdkerrors.Register(ModuleName, 5, "Mint request not pending")
	ErrMintRequestConfirmedAlreadyComplete = sdkerrors.Register(ModuleName, 6, "Mint request confirmed already complete")
	ErrParsingBase64                       = sdkerrors.Register(ModuleName, 7, "Error parsing base64")
	ErrParsingOriginData                   = sdkerrors.Register(ModuleName, 8, "Error parsing origin data")
	ErrDataHashMismatch                    = sdkerrors.Register(ModuleName, 9, "Data hash mismatch")
	ErrOracleConfirmedAlready              = sdkerrors.Register(ModuleName, 10, "Oracle confirmed already")
	ErrUnsupportedMetadataFormat           = sdkerrors.Register(ModuleName, 11, "Unsupported metadata format")
	ErrInvalidBase64                       = sdkerrors.Register(ModuleName, 12, "Invalid base64")
	ErrParsingActionSignature              = sdkerrors.Register(ModuleName, 13, "Error parsing action signature")
	ErrParsingActionParam                  = sdkerrors.Register(ModuleName, 14, "Error parsing action params")
	ErrVerifyingSignature                  = sdkerrors.Register(ModuleName, 15, "Error verify signature")
	ErrMetadataNotExists                   = sdkerrors.Register(ModuleName, 16, "Metadata not exists")

	ErrActionRequestNotFound                 = sdkerrors.Register(ModuleName, 17, "Action request not found")
	ErrActionRequestNotPending               = sdkerrors.Register(ModuleName, 18, "Action request not pending")
	ErrActionRequestConfirmedAlreadyComplete = sdkerrors.Register(ModuleName, 19, "Action request confirmed already complete")
	ErrMetaDataNotFound                      = sdkerrors.Register(ModuleName, 20, "Metadata not found")
	ErrUnauthorizedCaller                    = sdkerrors.Register(ModuleName, 21, "Unauthorized caller")
	ErrEmptyChangeList                       = sdkerrors.Register(ModuleName, 22, "There's no changes from action")
	ErrRefIdAlreadyExists                    = sdkerrors.Register(ModuleName, 23, "RefId already exists")
)
