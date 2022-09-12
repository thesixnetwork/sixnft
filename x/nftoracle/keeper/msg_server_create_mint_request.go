package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sixnft/x/nftoracle/types"
)

func (k msgServer) CreateMintRequest(goCtx context.Context, msg *types.MsgCreateMintRequest) (*types.MsgCreateMintRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateMintRequestResponse{}, nil
}
