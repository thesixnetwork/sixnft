package keeper

import (
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

var _ types.QueryServer = Keeper{}
