package keeper

import (
	"sixnft/x/nftmngr/types"
)

var _ types.QueryServer = Keeper{}
