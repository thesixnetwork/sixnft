package keeper

import (
	"context"

	"sixnft/x/evmsupport/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AddressBindingAll(c context.Context, req *types.QueryAllAddressBindingRequest) (*types.QueryAllAddressBindingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var addressBindings []types.AddressBinding
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	addressBindingStore := prefix.NewStore(store, types.KeyPrefix(types.AddressBindingKeyPrefix))

	pageRes, err := query.Paginate(addressBindingStore, req.Pagination, func(key []byte, value []byte) error {
		var addressBinding types.AddressBinding
		if err := k.cdc.Unmarshal(value, &addressBinding); err != nil {
			return err
		}

		addressBindings = append(addressBindings, addressBinding)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAddressBindingResponse{AddressBinding: addressBindings, Pagination: pageRes}, nil
}

func (k Keeper) AddressBinding(c context.Context, req *types.QueryGetAddressBindingRequest) (*types.QueryGetAddressBindingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAddressBinding(
		ctx,
		req.EthAddress,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAddressBindingResponse{AddressBinding: val}, nil
}
