package keeper

import (
	"context"
	"strconv"

	"github.com/thesixnetwork/sixnft/x/nftadmin/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Mint mints new tokens
// CLI: sixnftd tx nftadmin mint 1000000000 stake --from alice
func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	minter, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Check if the the msg sender has permission to mint
	hasMintPermission := k.Keeper.HasPermission(ctx, types.KeyPermissionMinter, minter)
	if !hasMintPermission {
		return nil, types.ErrUnauthorized
	}

	tokens := sdk.Coin{
		Denom:  msg.Token,
		Amount: sdk.NewIntFromUint64(msg.Amount),
	}

	if err := k.bankKeeper.MintCoins(
		ctx, types.ModuleName, sdk.NewCoins(tokens),
	); err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, minter, sdk.NewCoins(tokens),
	)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBurn,
			sdk.NewAttribute(types.EventTypeBurn, msg.Token),
			sdk.NewAttribute(types.AttributeKeyAmount, strconv.FormatUint(msg.Amount, 10)),
			sdk.NewAttribute(types.AttributeKeyMintStatus, "success"),
		),
	})

	return &types.MsgMintResponse{
		Amount: strconv.FormatUint(msg.Amount, 10),
		Token:  msg.Token,
	}, nil
}
