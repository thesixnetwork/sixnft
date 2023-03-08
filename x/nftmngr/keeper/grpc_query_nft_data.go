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

func (k Keeper) NftDataAll(c context.Context, req *types.QueryAllNftDataRequest) (*types.QueryAllNftDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var listNFTData []types.NftData
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nftDataStore := prefix.NewStore(store, types.KeyPrefix(types.NftDataKeyPrefix))

	pageRes, err := query.Paginate(nftDataStore, req.Pagination, func(key []byte, value []byte) error {
		var nftData types.NftData
		if err := k.cdc.Unmarshal(value, &nftData); err != nil {
			return err
		}

		updateddata := k.updateNftDataAttributes(ctx, nftData)
		listNFTData = append(listNFTData, updateddata)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNftDataResponse{NftData: listNFTData, Pagination: pageRes}, nil
}

func (k Keeper) NftData(c context.Context, req *types.QueryGetNftDataRequest) (*types.QueryGetNftDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNftData(
		ctx,
		req.NftSchemaCode,
		req.TokenId,
	)
	
	updateddata := k.updateNftDataAttributes(ctx, val)

	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNftDataResponse{NftData: updateddata}, nil
}


// function updateNftDataAttributes updates the nft data attributes from the schema
func (k Keeper) updateNftDataAttributes(ctx sdk.Context, dataOnToken types.NftData) (updatedData types.NftData) {
	nftSchema, _ := k.GetNFTSchema(ctx, dataOnToken.NftSchemaCode)

	// createa map of nftattributes from schema with its default value
	mapFromSchemaAttributes := make(map[string]*types.DefaultMintValue)
	for _, fromSchemaNftAttribute := range nftSchema.OnchainData.NftAttributes {
		mapFromSchemaAttributes[fromSchemaNftAttribute.Name] = fromSchemaNftAttribute.DefaultMintValue
	}

	// loop over nftdata attributes and check if exists in mapFromSchemaAttributes
	// if exists, then update the value
	for _, attribute := range dataOnToken.OnchainAttributes {
		if schemaUpdated, ok := mapFromSchemaAttributes[attribute.Name]; ok {
			switch schemaUpdated.Value.(type) {
			case *types.DefaultMintValue_NumberAttributeValue:
				attribute.Value = &types.NftAttributeValue_NumberAttributeValue{NumberAttributeValue: schemaUpdated.GetNumberAttributeValue()}

			case *types.DefaultMintValue_BooleanAttributeValue:
				attribute.Value = &types.NftAttributeValue_BooleanAttributeValue{BooleanAttributeValue: schemaUpdated.GetBooleanAttributeValue()}

			case *types.DefaultMintValue_FloatAttributeValue:
				attribute.Value = &types.NftAttributeValue_FloatAttributeValue{FloatAttributeValue: schemaUpdated.GetFloatAttributeValue()}

			case *types.DefaultMintValue_StringAttributeValue:
				attribute.Value = &types.NftAttributeValue_StringAttributeValue{StringAttributeValue: schemaUpdated.GetStringAttributeValue()}
			}
		}
	}

	return dataOnToken
}