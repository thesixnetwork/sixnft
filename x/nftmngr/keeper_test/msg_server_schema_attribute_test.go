package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/sixnft/testutil/keeper"
	"github.com/thesixnetwork/sixnft/x/nftmngr/keeper"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestSchemaAttributeMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateSchemaAttribute{Creator: creator,
			NftSchemaCode: strconv.Itoa(i),
			Name:          strconv.Itoa(i),
		}
		_, err := srv.CreateSchemaAttribute(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetSchemaAttribute(ctx,
			expected.NftSchemaCode,
			expected.Name,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestSchemaAttributeMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateSchemaAttribute
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateSchemaAttribute{Creator: creator,
				NftSchemaCode: strconv.Itoa(0),
				Name:          strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateSchemaAttribute{Creator: "B",
				NftSchemaCode: strconv.Itoa(0),
				Name:          strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateSchemaAttribute{Creator: creator,
				NftSchemaCode: strconv.Itoa(100000),
				Name:          strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NftmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateSchemaAttribute{Creator: creator,
				NftSchemaCode: strconv.Itoa(0),
				Name:          strconv.Itoa(0),
			}
			_, err := srv.CreateSchemaAttribute(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateSchemaAttribute(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetSchemaAttribute(ctx,
					expected.NftSchemaCode,
					expected.Name,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestSchemaAttributeMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteSchemaAttribute
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteSchemaAttribute{Creator: creator,
				NftSchemaCode: strconv.Itoa(0),
				Name:          strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteSchemaAttribute{Creator: "B",
				NftSchemaCode: strconv.Itoa(0),
				Name:          strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteSchemaAttribute{Creator: creator,
				NftSchemaCode: strconv.Itoa(100000),
				Name:          strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NftmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateSchemaAttribute(wctx, &types.MsgCreateSchemaAttribute{Creator: creator,
				NftSchemaCode: strconv.Itoa(0),
				Name:          strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteSchemaAttribute(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetSchemaAttribute(ctx,
					tc.request.NftSchemaCode,
					tc.request.Name,
				)
				require.False(t, found)
			}
		})
	}
}
