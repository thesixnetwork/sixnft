package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AttributeOfSchema(c context.Context, req *types.QueryGetAttributeOfSchemaRequest) (*types.QueryGetAttributeOfSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAttributeOfSchema(
		ctx,
		req.NftSchemaCode,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	var schemaAttributes []types.SchemaAttribute
	for _, v := range val.SchemaAttributes {
		schemaAttributes = append(schemaAttributes, *v)
	}

	return &types.QueryGetAttributeOfSchemaResponse{NftSchemaCode: val.NftSchemaCode, SchemaAttribute: schemaAttributes}, nil
}
