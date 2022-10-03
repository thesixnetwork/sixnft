package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/0.20220523154235-2921a1c3c918/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NftCollection(c context.Context, req *types.QueryGetNftCollectionRequest) (*types.QueryGetNftCollectionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nftCollections []types.NftCollection
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	NftcollectionStore := prefix.NewStore(store, types.KeyPrefix(types.NftCollectionKeyPrefix))

	pageRes, err := query.Paginate(NftcollectionStore, req.Pagination, func(key []byte, value []byte) error {
		var nftCollection types.NftCollection
		if err := k.cdc.Unmarshal(value, &nftCollection); err != nil {
			return err
		}

		nftCollections = append(nftCollections, nftCollection)
		return nil
	})

	val, found := k.GetNftCollection(
		ctx,
		req.NftSchemaCode,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNftCollectionResponse{NftCollection: val, Pagination: pageRes}, nil
}
