package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

func (k msgServer) SubmitVerifyCollectionOwner(goCtx context.Context, msg *types.MsgSubmitVerifyCollectionOwner) (*types.MsgSubmitVerifyCollectionOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitVerifyCollectionOwnerResponse{}, nil
}
