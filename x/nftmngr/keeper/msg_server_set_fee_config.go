package keeper

import (
	"context"
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k msgServer) SetFeeConfig(goCtx context.Context, msg *types.MsgSetFeeConfig) (*types.MsgSetFeeConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate if creator has permission to set fee config
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	granted := k.nftadminKeeper.HasPermission(ctx, types.KeyPermissionNftFeeAdmin, creator)
	if !granted {
		return nil, sdkerrors.Wrap(types.ErrNoNftFeeAdminPermission, msg.Creator)
	}

	json, err := base64.StdEncoding.DecodeString(msg.NewFeeConfigBase64)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}
	// Check if msg.Subject is CREATE_NFT_SCHEMA
	if msg.FeeSubject == types.FeeSubject_CREATE_NFT_SCHEMA {

		feeConfig := types.NFTFeeConfig{}
		err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(json, &feeConfig)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrParsingSchemaMessage, err.Error())
		}

		err = k.ValidateFeeConfig(&feeConfig)
		if err != nil {
			return nil, err
		}

		// Set fee config
		k.Keeper.SetNFTFeeConfig(ctx, feeConfig)

		// Emit event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeSetFeeConfig,
				sdk.NewAttribute(types.AttributeKeyFeeSubject, msg.FeeSubject.String()),
				sdk.NewAttribute(types.AttributeKeyFeeConfig, string(json)),
			),
		)
	}
	return &types.MsgSetFeeConfigResponse{}, nil
}

func (k Keeper) ValidateFeeConfig(feeConfig *types.NFTFeeConfig) error {
	// Validate Amount
	feeAmount, err := sdk.ParseCoinNormalized(feeConfig.SchemaFee.FeeAmount)
	if err != nil {
		return sdkerrors.Wrap(types.ErrInvalidFeeAmount, err.Error())
	}
	// check if feeAmount.Amount > 0
	if feeAmount.Amount.LTE(sdk.NewInt(0)) {
		return sdkerrors.Wrap(types.ErrInvalidFeeAmount, "Fee amount must be greater than 0")
	}
	// loop over feeConfig.SchemaFee.FeeDistributions
	totalPortion := float32(0)
	for _, feeDistribution := range feeConfig.SchemaFee.FeeDistributions {
		// validate portion
		if feeDistribution.Portion <= 0 || feeDistribution.Portion > 1 {
			return sdkerrors.Wrap(types.ErrInvalidFeePortion, "Fee portion must be between 0 and 1")
		}
		totalPortion += feeDistribution.Portion
	}
	if totalPortion != 1 {
		return sdkerrors.Wrap(types.ErrInvalidFeePortion, "Total fee portion must be equal to 1")
	}
	return nil
}
