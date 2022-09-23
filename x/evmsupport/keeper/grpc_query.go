package keeper

import (
	"github.com/thesixnetwork/sixnft/x/evmsupport/types"
)

var _ types.QueryServer = Keeper{}
