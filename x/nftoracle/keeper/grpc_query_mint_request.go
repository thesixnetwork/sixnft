package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sixnft/x/nftoracle/types"
)

func (k Keeper) MintRequestAll(c context.Context, req *types.QueryAllMintRequestRequest) (*types.QueryAllMintRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var mintRequests []types.MintRequest
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	mintRequestStore := prefix.NewStore(store, types.KeyPrefix(types.MintRequestKey))

	pageRes, err := query.Paginate(mintRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var mintRequest types.MintRequest
		if err := k.cdc.Unmarshal(value, &mintRequest); err != nil {
			return err
		}

		mintRequests = append(mintRequests, mintRequest)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMintRequestResponse{MintRequest: mintRequests, Pagination: pageRes}, nil
}

func (k Keeper) MintRequest(c context.Context, req *types.QueryGetMintRequestRequest) (*types.QueryGetMintRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	mintRequest, found := k.GetMintRequest(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMintRequestResponse{MintRequest: mintRequest}, nil
}
