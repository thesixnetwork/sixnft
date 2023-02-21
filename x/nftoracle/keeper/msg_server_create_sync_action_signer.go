package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

func (k msgServer) CreateSyncActionSigner(goCtx context.Context, msg *types.MsgCreateSyncActionSigner) (*types.MsgCreateSyncActionSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateSyncActionSignerResponse{}, nil
}
