package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sixnft/x/nftoracle/types"
)

func (k msgServer) SubmitMintResponse(goCtx context.Context, msg *types.MsgSubmitMintResponse) (*types.MsgSubmitMintResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitMintResponseResponse{}, nil
}
