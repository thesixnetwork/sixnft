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

func (k Keeper) NFTSchemaAllV072(c context.Context, req *types.QueryAllNFTSchemaRequest) (*types.QueryAllNFTSchemaResponseV072, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nFTSchemas []types.NFTSchemaV072
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nFTSchemaStore := prefix.NewStore(store, types.KeyPrefix(types.NFTSchemaKeyPrefix))

	pageRes, err := query.Paginate(nFTSchemaStore, req.Pagination, func(key []byte, value []byte) error {
		var nFTSchema types.NFTSchemaV072
		if err := k.cdc.Unmarshal(value, &nFTSchema); err != nil {
			return err
		}

		nFTSchemas = append(nFTSchemas, nFTSchema)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNFTSchemaResponseV072{NFTSchemaV072: nFTSchemas, Pagination: pageRes}, nil
}

func (k Keeper) NFTSchemaV072(c context.Context, req *types.QueryGetNFTSchemaRequest) (*types.QueryGetNFTSchemaResponseV072, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNFTSchemaV072(
		ctx,
		req.Code,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNFTSchemaResponseV072{NFTSchemaV072: val}, nil
}
