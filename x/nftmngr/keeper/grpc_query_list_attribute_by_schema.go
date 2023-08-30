package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListAttributeBySchema(goCtx context.Context, req *types.QueryListAttributeBySchemaRequest) (*types.QueryListAttributeBySchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var schemaAttributes []types.SchemaAttribute
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	store := ctx.KVStore(k.storeKey)
	schemaAttributesStore := prefix.NewStore(store, types.KeyPrefix(types.SchemaAttributeKeyPrefix))

	pagination, err := query.Paginate(schemaAttributesStore, req.Pagination, func(key []byte, value []byte) error {
		var schemaAttribute types.SchemaAttribute
		if err := k.cdc.Unmarshal(value, &schemaAttribute); err != nil {
			return err
		}

		schemaAttributes = append(schemaAttributes, schemaAttribute)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var listOfSchemaAttibutes []types.SchemaAttribute
	// find all actionExecutors by schema
	for _, attribute := range schemaAttributes {
		if attribute.NftSchemaCode == req.NftSchemaCode {
			listOfSchemaAttibutes = append(listOfSchemaAttibutes, attribute)
		}
	}

	pagination.Total = uint64(len(listOfSchemaAttibutes))

	return &types.QueryListAttributeBySchemaResponse{SchemaAttribute: listOfSchemaAttibutes,Pagination: pagination,}, nil
}





// func (k Keeper) ListAttributeBySchemaUsingCtx(ctx sdk.Context, req Pagination *query.PageRequest) ([]*types.SchemaAttribute, error) {
// 	var schemaAttributes []types.SchemaAttribute
	
// 	store := ctx.KVStore(k.storeKey)
// 	schemaAttributesStore := prefix.NewStore(store, types.KeyPrefix(types.SchemaAttributeKeyPrefix))

// 	// set default pagination
// 	req := &types.QueryListAttributeBySchemaRequest{
// 		NftSchemaCode: schemacode,
// 		Pagination: &query.PageRequest{
// 			Key:        []byte{},
// 			Offset:     0,
// 			Limit:      0,
// 			CountTotal: true,
// 		},
// 	}

// 	_, err := query.Paginate(schemaAttributesStore, req.Pagination, func(key []byte, value []byte) error {
// 		var schemaAttribute types.SchemaAttribute
// 		if err := k.cdc.Unmarshal(value, &schemaAttribute); err != nil {
// 			return err
// 		}

// 		schemaAttributes = append(schemaAttributes, schemaAttribute)
// 		return nil
// 	})

// 	if err != nil {
// 		return nil, status.Error(codes.Internal, err.Error())
// 	}

// 	var listOfSchemaAttibutes []*types.SchemaAttribute
// 	// find all actionExecutors by schema
// 	for _, attribute := range schemaAttributes {
// 		if attribute.NftSchemaCode == req.NftSchemaCode {
// 			listOfSchemaAttibutes = append(listOfSchemaAttibutes, &attribute)
// 		}
// 	}
// 	return listOfSchemaAttibutes, nil
// }