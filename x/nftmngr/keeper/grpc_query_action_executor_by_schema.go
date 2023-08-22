package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ActionExecutorbySchema(c context.Context, req *types.QueryActionExecutorbySchemaRequest) (*types.QueryActionExecutorbySchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionExecutors []types.ActionExecutor
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	actionExecutorStore := prefix.NewStore(store, types.KeyPrefix(types.ActionExecutorKeyPrefix))

	_, err := query.Paginate(actionExecutorStore, req.Pagination, func(key []byte, value []byte) error {
		var actionExecutor types.ActionExecutor
		if err := k.cdc.Unmarshal(value, &actionExecutor); err != nil {
			return err
		}

		actionExecutors = append(actionExecutors, actionExecutor)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}


	var executorAddresses []string 
	// find all actionExecutors by schema
	for _, actionExecutor := range actionExecutors {
		if actionExecutor.NftSchemaCode == req.NftSchemaCode {
			executorAddresses = append(executorAddresses, actionExecutor.ExecutorAddress)
		}
	}


	return &types.QueryActionExecutorbySchemaResponse{NftSchemaCode: req.NftSchemaCode,ExecutorAddress: executorAddresses}, nil
}