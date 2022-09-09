package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "sixnft/testutil/keeper"
	"sixnft/testutil/nullify"
	"sixnft/x/evmsupport/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAddressBindingQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EvmsupportKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAddressBinding(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetAddressBindingRequest
		response *types.QueryGetAddressBindingResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetAddressBindingRequest{
				EthAddress:    msgs[0].EthAddress,
				NativeAddress: msgs[0].NativeAddress,
			},
			response: &types.QueryGetAddressBindingResponse{AddressBinding: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetAddressBindingRequest{
				EthAddress:    msgs[1].EthAddress,
				NativeAddress: msgs[1].NativeAddress,
			},
			response: &types.QueryGetAddressBindingResponse{AddressBinding: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetAddressBindingRequest{
				EthAddress:    strconv.Itoa(100000),
				NativeAddress: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.AddressBinding(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestAddressBindingQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EvmsupportKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAddressBinding(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllAddressBindingRequest {
		return &types.QueryAllAddressBindingRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AddressBindingAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.AddressBinding), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.AddressBinding),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AddressBindingAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.AddressBinding), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.AddressBinding),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AddressBindingAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.AddressBinding),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AddressBindingAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
