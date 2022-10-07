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

func TestNftCollectionQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNftCollection(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNftCollectionRequest
		response *types.QueryGetNftCollectionResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetNftCollectionRequest{
				NftSchemaCode: msgs[0].NftSchemaCode,
			},
			response: &types.QueryGetNftCollectionResponse{NftCollection: msgs[0].NftDatas},
		},
		{
			desc: "Second",
			request: &types.QueryGetNftCollectionRequest{
				NftSchemaCode: msgs[1].NftSchemaCode,
			},
			response: &types.QueryGetNftCollectionResponse{NftCollection: msgs[1].NftDatas},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetNftCollectionRequest{
				NftSchemaCode: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NftCollection(wctx, tc.request)
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

func TestNftCollectionQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNftCollection(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryGetNftCollectionRequest {
		return &types.QueryGetNftCollectionRequest{
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
			resp, err := keeper.NftCollection(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NftCollection), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NftCollection),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.NftCollection(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NftCollection), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NftCollection),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.NftCollection(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.NftCollection),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.NftCollection(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
