package keeper

import (
	"context"
	"strconv"

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

	createdAt := ctx.BlockTime()
	endTime := createdAt.Add(k.ActionSignerActiveDuration(ctx)) // endtime of request; not expire time of actionSigner and set till now + 5 minutes(300 seconds)

	id := k.AppendSyncActionSigner(ctx, types.SyncActionSigner{
		Chain:           msg.Chain,
		ActorAddress:    msg.ActorAddress,
		OwnerAddress:    msg.OwnerAddress,
		Caller:          msg.Creator,
		RequiredConfirm: msg.RequiredConfirm,
		Status:          types.RequestStatus_PENDING,
		CurrentConfirm:  0,
		CreatedAt:       createdAt,
		ValidUntil:      endTime,
		Confirmers:      make([]string, 0),
		DataHashes:      make([]*types.ContractInfoHash, 0),
	})

	k.Keeper.InsertActiveSyncActionSignerQueue(ctx, id, endTime)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSyncActionReqeustCreated,
			sdk.NewAttribute(types.AttributeKeySyncActionRequestID, strconv.FormatUint(id, 10)),
			sdk.NewAttribute(types.AttributeKeySyncActionRequestActorAddress, msg.ActorAddress),
			sdk.NewAttribute(types.AttributeKeySyncActionRequestOwnerAddress, msg.OwnerAddress),
			sdk.NewAttribute(types.AttributeKeyRequiredConfirm, strconv.FormatUint(msg.RequiredConfirm, 10)),
		),
	})

	return &types.MsgCreateSyncActionSignerResponse{}, nil
}
