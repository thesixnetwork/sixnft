package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

func (k msgServer) CreateVerifyCollectionOwnerRequest(goCtx context.Context, msg *types.MsgCreateVerifyCollectionOwnerRequest) (*types.MsgCreateVerifyCollectionOwnerRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateVerifyCollectionOwnerRequestResponse{}, nil
}
