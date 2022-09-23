package keeper

import (
	"github.com/thesixnetwork/sixnft/x/admin/types"
)

var _ types.QueryServer = Keeper{}
