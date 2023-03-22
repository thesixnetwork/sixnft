package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ActionSignerConfigAll(c context.Context, req *types.QueryAllActionSignerConfigRequest) (*types.QueryAllActionSignerConfigResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionSignerConfigs []types.ActionSignerConfig
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	actionSignerConfigStore := prefix.NewStore(store, types.KeyPrefix(types.ActionSignerConfigKeyPrefix))

	pageRes, err := query.Paginate(actionSignerConfigStore, req.Pagination, func(key []byte, value []byte) error {
		var actionSignerConfig types.ActionSignerConfig
		if err := k.cdc.Unmarshal(value, &actionSignerConfig); err != nil {
			return err
		}

		actionSignerConfigs = append(actionSignerConfigs, actionSignerConfig)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActionSignerConfigResponse{ActionSignerConfig: actionSignerConfigs, Pagination: pageRes}, nil
}

func (k Keeper) ActionSignerConfig(c context.Context, req *types.QueryGetActionSignerConfigRequest) (*types.QueryGetActionSignerConfigResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetActionSignerConfig(
		ctx,
		req.Chain,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetActionSignerConfigResponse{ActionSignerConfig: val}, nil
}
