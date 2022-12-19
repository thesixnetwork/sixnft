package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ActionRequestAll(c context.Context, req *types.QueryAllActionRequestRequest) (*types.QueryAllActionRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionRequests []types.ActionOracleRequest
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	actionRequestStore := prefix.NewStore(store, types.KeyPrefix(types.ActionRequestKey))

	pageRes, err := query.Paginate(actionRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var actionRequest types.ActionOracleRequest
		if err := k.cdc.Unmarshal(value, &actionRequest); err != nil {
			return err
		}

		actionRequests = append(actionRequests, actionRequest)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActionRequestResponse{ActionOracleRequest: actionRequests, Pagination: pageRes}, nil
}

func (k Keeper) ActionOracleRequest(c context.Context, req *types.QueryGetActionRequestRequest) (*types.QueryGetActionRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	actionRequest, found := k.GetActionRequest(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetActionRequestResponse{ActionOracleRequest: actionRequest}, nil
}
