package keeper

import (
	"context"
	"strconv"

	"sixnft/x/admin/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Burn burns tokens
// CLI: sixnftd tx admin burn 1000000000 stake --from alice
func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	burner, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Check if the the msg sender has permission to mint
	hasBurnPermission := k.Keeper.HasPermission(ctx, types.KeyPermissionBurner, burner)
	if !hasBurnPermission {
		return nil, types.ErrUnauthorized
	}

	tokens := sdk.Coin{
		Denom:  msg.Token,
		Amount: sdk.NewIntFromUint64(msg.Amount),
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, burner, types.ModuleName, sdk.NewCoins(tokens),
	)

	if err != nil {
		return nil, err
	}

	if err := k.bankKeeper.BurnCoins(
		ctx, types.ModuleName, sdk.NewCoins(tokens),
	); err != nil {
		return nil, err
	}

	return &types.MsgBurnResponse{
		Amount: strconv.FormatUint(msg.Amount, 10),
		Token:  msg.Token,
	}, nil
}
