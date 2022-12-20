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

func (k Keeper) NFTSchemaAllV063(c context.Context, req *types.QueryAllNFTSchemaRequest) (*types.QueryAllNFTSchemaResponseV063, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nFTSchemas []types.NFTSchemaV063
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nFTSchemaStore := prefix.NewStore(store, types.KeyPrefix(types.NFTSchemaKeyPrefix))

	pageRes, err := query.Paginate(nFTSchemaStore, req.Pagination, func(key []byte, value []byte) error {
		var nFTSchema types.NFTSchemaV063
		if err := k.cdc.Unmarshal(value, &nFTSchema); err != nil {
			return err
		}

		nFTSchemas = append(nFTSchemas, nFTSchema)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNFTSchemaResponseV063{NFTSchemaV063: nFTSchemas, Pagination: pageRes}, nil
}

func (k Keeper) NFTSchemaV063(c context.Context, req *types.QueryGetNFTSchemaRequest) (*types.QueryGetNFTSchemaResponseV063, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNFTSchemaV063(
		ctx,
		req.Code,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNFTSchemaResponseV063{NFTSchemaV063: val}, nil
}
