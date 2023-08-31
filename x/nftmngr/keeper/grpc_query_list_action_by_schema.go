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

func (k Keeper) ListActionBySchema(goCtx context.Context, req *types.QueryListActionBySchemaRequest) (*types.QueryListActionBySchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var actionOfSchemas []*types.ActionOfSchema

	store := ctx.KVStore(k.storeKey)
	actionOfSchemaStore := prefix.NewStore(store, types.KeyPrefix(types.ActionOfSchemaKeyPrefix))

	pagination, err := query.Paginate(actionOfSchemaStore, req.Pagination, func(key []byte, value []byte) error {
		var actionSchema types.ActionOfSchema
		if err := k.cdc.Unmarshal(value, &actionSchema); err != nil {
			return err
		}
		if actionSchema.NftSchemaCode == req.NftSchemaCode {
			actionOfSchemas = append(actionOfSchemas, &actionSchema)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pagination.Total = uint64(len(actionOfSchemas))

	return &types.QueryListActionBySchemaResponse{NftSchemaCode: req.GetNftSchemaCode(), ActionOfSchema: actionOfSchemas, Pagination: pagination}, nil

}
