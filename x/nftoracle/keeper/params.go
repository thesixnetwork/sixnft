package keeper

import (
	"sixnft/x/nftoracle/types"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MintRequestActiveDuration(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) MintRequestActiveDuration(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyMintRequestActiveDuration, &res)
	return
}
