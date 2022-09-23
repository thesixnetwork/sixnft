package keeper

import (
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

var _ types.QueryServer = Keeper{}
