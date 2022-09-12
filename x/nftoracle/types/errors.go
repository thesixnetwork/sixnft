package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nftoracle module sentinel errors
var (
	ErrNFTSchemaNotFound     = sdkerrors.Register(ModuleName, 1, "NFTSchema not found")
	ErrMetadataAlreadyExists = sdkerrors.Register(ModuleName, 2, "Metadata already exists")
)
