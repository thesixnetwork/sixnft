package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/sixnft/testutil/keeper"

	"github.com/thesixnetwork/sixnft/testutil/nullify"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestNftDataQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNftData(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNftDataRequest
		response *types.QueryGetNftDataResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetNftDataRequest{
				NftSchemaCode: msgs[0].NftSchemaCode,
				TokenId:       msgs[0].TokenId,
			},
			response: &types.QueryGetNftDataResponse{NftData: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetNftDataRequest{
				NftSchemaCode: msgs[1].NftSchemaCode,
				TokenId:       msgs[1].TokenId,
			},
			response: &types.QueryGetNftDataResponse{NftData: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetNftDataRequest{
				NftSchemaCode: strconv.Itoa(100000),
				TokenId:       strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NftData(wctx, tc.request)
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

func TestNftDataQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNftData(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllNftDataRequest {
		return &types.QueryAllNftDataRequest{
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
			resp, err := keeper.NftDataAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NftData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NftData),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.NftDataAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NftData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NftData),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.NftDataAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.NftData),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.NftDataAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
