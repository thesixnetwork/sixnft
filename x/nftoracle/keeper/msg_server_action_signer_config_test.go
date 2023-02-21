package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/sixnft/testutil/keeper"
	"github.com/thesixnetwork/sixnft/x/nftoracle/keeper"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestActionSignerConfigMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateActionSignerConfig{Creator: creator,
			Chain: strconv.Itoa(i),
		}
		_, err := srv.CreateActionSignerConfig(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetActionSignerConfig(ctx,
			expected.Chain,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestActionSignerConfigMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateActionSignerConfig
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateActionSignerConfig{Creator: creator,
				Chain: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateActionSignerConfig{Creator: "B",
				Chain: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateActionSignerConfig{Creator: creator,
				Chain: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NftoracleKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateActionSignerConfig{Creator: creator,
				Chain: strconv.Itoa(0),
			}
			_, err := srv.CreateActionSignerConfig(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateActionSignerConfig(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetActionSignerConfig(ctx,
					expected.Chain,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestActionSignerConfigMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteActionSignerConfig
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteActionSignerConfig{Creator: creator,
				Chain: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteActionSignerConfig{Creator: "B",
				Chain: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteActionSignerConfig{Creator: creator,
				Chain: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NftoracleKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateActionSignerConfig(wctx, &types.MsgCreateActionSignerConfig{Creator: creator,
				Chain: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteActionSignerConfig(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetActionSignerConfig(ctx,
					tc.request.Chain,
				)
				require.False(t, found)
			}
		})
	}
}
