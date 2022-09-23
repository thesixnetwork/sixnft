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

func (k Keeper) OrganizationAll(c context.Context, req *types.QueryAllOrganizationRequest) (*types.QueryAllOrganizationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var organizations []types.Organization
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	organizationStore := prefix.NewStore(store, types.KeyPrefix(types.OrganizationKeyPrefix))

	pageRes, err := query.Paginate(organizationStore, req.Pagination, func(key []byte, value []byte) error {
		var organization types.Organization
		if err := k.cdc.Unmarshal(value, &organization); err != nil {
			return err
		}

		organizations = append(organizations, organization)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOrganizationResponse{Organization: organizations, Pagination: pageRes}, nil
}

func (k Keeper) Organization(c context.Context, req *types.QueryGetOrganizationRequest) (*types.QueryGetOrganizationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetOrganization(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOrganizationResponse{Organization: val}, nil
}
