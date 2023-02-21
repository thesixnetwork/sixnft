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

func (k Keeper) SyncActionSignerAll(c context.Context, req *types.QueryAllSyncActionSignerRequest) (*types.QueryAllSyncActionSignerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var syncActionSigners []types.SyncActionSigner
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	syncActionSignerStore := prefix.NewStore(store, types.KeyPrefix(types.SyncActionSignerKey))

	pageRes, err := query.Paginate(syncActionSignerStore, req.Pagination, func(key []byte, value []byte) error {
		var syncActionSigner types.SyncActionSigner
		if err := k.cdc.Unmarshal(value, &syncActionSigner); err != nil {
			return err
		}

		syncActionSigners = append(syncActionSigners, syncActionSigner)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSyncActionSignerResponse{SyncActionSigner: syncActionSigners, Pagination: pageRes}, nil
}

func (k Keeper) SyncActionSigner(c context.Context, req *types.QueryGetSyncActionSignerRequest) (*types.QueryGetSyncActionSignerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	syncActionSigner, found := k.GetSyncActionSigner(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSyncActionSignerResponse{SyncActionSigner: syncActionSigner}, nil
}
