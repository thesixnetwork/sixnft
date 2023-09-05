package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListActionBySchema(goCtx context.Context, req *types.QueryListActionBySchemaRequest) (*types.QueryListActionBySchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionOfSchemas []*types.ActionOfSchema

	ctx := sdk.UnwrapSDKContext(goCtx)
	_, found := k.GetNFTSchema(ctx, req.NftSchemaCode)
	if !found {
		return nil, status.Error(codes.NotFound, "Schema Not Found")
	}

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var actionOfSchema types.ActionOfSchema
		k.cdc.MustUnmarshal(iterator.Value(), &actionOfSchema)
		if actionOfSchema.NftSchemaCode == req.NftSchemaCode {
			actionOfSchemas = append(actionOfSchemas, &actionOfSchema)
		}
	}

	return &types.QueryListActionBySchemaResponse{NftSchemaCode: req.GetNftSchemaCode(), ActionOfSchema: actionOfSchemas}, nil

}
