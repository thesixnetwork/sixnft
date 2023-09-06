package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/sixnft/testutil/keeper"
	"github.com/thesixnetwork/sixnft/testutil/nullify"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAttributeOfSchemaQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAttributeOfSchema(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetAttributeOfSchemaRequest
		response *types.QueryGetAttributeOfSchemaResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetAttributeOfSchemaRequest{
				NftSchemaCode: msgs[0].NftSchemaCode,
			},
			response: &types.QueryGetAttributeOfSchemaResponse{},
		},
		{
			desc: "Second",
			request: &types.QueryGetAttributeOfSchemaRequest{
				NftSchemaCode: msgs[1].NftSchemaCode,
			},
			response: &types.QueryGetAttributeOfSchemaResponse{},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetAttributeOfSchemaRequest{
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
			response, err := keeper.AttributeOfSchema(wctx, tc.request)
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
