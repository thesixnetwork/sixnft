package keeper

import (
	"github.com/thesixnetwork/sixnft/x/nftadmin/types"
)

var _ types.QueryServer = Keeper{}
