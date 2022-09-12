package keeper

import (
	"context"

	"sixnft/x/nftoracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateMintRequest(goCtx context.Context, msg *types.MsgCreateMintRequest) (*types.MsgCreateMintRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateMintRequestResponse{}, nil
}
