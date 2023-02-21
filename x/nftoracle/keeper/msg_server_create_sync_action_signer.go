package keeper

import (
	"context"
	"strconv"
	// "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

func (k msgServer) CreateSyncActionSigner(goCtx context.Context, msg *types.MsgCreateSyncActionSigner) (*types.MsgCreateSyncActionSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	oracleConfig, found := k.GetOracleConfig(ctx)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOracleConfigNotFound, "")
	}

	// Verify msg.RequiredConfirmations is less than or equal to oracleConfig.MinimumConfirmation
	if int32(msg.RequiredConfirm) < oracleConfig.MinimumConfirmation {
		return nil, sdkerrors.Wrap(types.ErrRequiredConfirmTooLess, strconv.Itoa(int(oracleConfig.MinimumConfirmation)))
	}

	// createdAt := ctx.BlockTime()
	// endTime := createdAt.Add(k.ActionSignerActiveDuration(ctx))

	return &types.MsgCreateSyncActionSignerResponse{}, nil
}
