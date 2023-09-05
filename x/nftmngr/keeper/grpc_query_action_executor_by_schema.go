package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ActionExecutorbySchema(c context.Context, req *types.QueryActionExecutorbySchemaRequest) (*types.QueryActionExecutorbySchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionExecutors []string
	ctx := sdk.UnwrapSDKContext(c)

	_, found := k.GetNFTSchema(ctx, req.NftSchemaCode)
	if !found {
		return nil, status.Error(codes.NotFound, "Schema Not Found")
	}
	
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var actionExecutor types.ActionExecutor
		k.cdc.MustUnmarshal(iterator.Value(), &actionExecutor)
		if actionExecutor.NftSchemaCode == req.NftSchemaCode {
			actionExecutors = append(actionExecutors, actionExecutor.ExecutorAddress)
		}
	}

	return &types.QueryActionExecutorbySchemaResponse{NftSchemaCode: req.NftSchemaCode, ExecutorAddress: actionExecutors}, nil
}
