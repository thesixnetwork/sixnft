package keeper

import (
	"sixnft/x/evmsupport/types"
)

var _ types.QueryServer = Keeper{}
