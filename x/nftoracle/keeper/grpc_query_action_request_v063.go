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

func (k Keeper) ActionRequestAllV063(c context.Context, req *types.QueryAllActionRequestRequest) (*types.QueryAllActionRequestResponseV063, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionRequests []types.ActionRequestV063
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	actionRequestStore := prefix.NewStore(store, types.KeyPrefix(types.ActionRequestKey))

	pageRes, err := query.Paginate(actionRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var actionRequest types.ActionRequestV063
		if err := k.cdc.Unmarshal(value, &actionRequest); err != nil {
			return err
		}

		actionRequests = append(actionRequests, actionRequest)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActionRequestResponseV063{ActionRequestV063: actionRequests, Pagination: pageRes}, nil
}

func (k Keeper) ActionOracleRequestV063(c context.Context, req *types.QueryGetActionRequestRequest) (*types.QueryGetActionRequestResponseV063, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	actionRequest, found := k.GetActionRequestV063(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetActionRequestResponseV063{ActionRequestV063: actionRequest}, nil
}
