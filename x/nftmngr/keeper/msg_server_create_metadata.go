package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sixnft/x/nftmngr/types"
)

func (k msgServer) CreateMetadata(goCtx context.Context, msg *types.MsgCreateMetadata) (*types.MsgCreateMetadataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateMetadataResponse{}, nil
}
