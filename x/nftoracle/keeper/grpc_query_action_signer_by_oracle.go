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

func (k Keeper) ActionSignerByOracleAll(c context.Context, req *types.QueryAllActionSignerByOracleRequest) (*types.QueryAllActionSignerByOracleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionSignerByOracles []types.ActionSignerByOracle
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	actionSignerByOracleStore := prefix.NewStore(store, types.KeyPrefix(types.ActionSignerByOracleKey))

	pageRes, err := query.Paginate(actionSignerByOracleStore, req.Pagination, func(key []byte, value []byte) error {
		var actionSignerByOracle types.ActionSignerByOracle
		if err := k.cdc.Unmarshal(value, &actionSignerByOracle); err != nil {
			return err
		}

		actionSignerByOracles = append(actionSignerByOracles, actionSignerByOracle)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActionSignerByOracleResponse{ActionSignerByOracle: actionSignerByOracles, Pagination: pageRes}, nil
}

func (k Keeper) ActionSignerByOracle(c context.Context, req *types.QueryGetActionSignerByOracleRequest) (*types.QueryGetActionSignerByOracleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	actionSignerByOracle, found := k.GetActionSignerByOracle(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetActionSignerByOracleResponse{ActionSignerByOracle: actionSignerByOracle}, nil
}
