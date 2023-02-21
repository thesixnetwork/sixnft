package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

func (k msgServer) SubmitSyncActionSigner(goCtx context.Context, msg *types.MsgSubmitSyncActionSigner) (*types.MsgSubmitSyncActionSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitSyncActionSignerResponse{}, nil
}
