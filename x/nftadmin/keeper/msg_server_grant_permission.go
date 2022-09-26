package keeper

import (
	"context"

	"github.com/thesixnetwork/sixnft/x/nftadmin/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Grant a permission
// CLI: sixnftd tx nftadmin grant-permission minter 6nft1q3566qhn4hnjf8l0zt90daew2ade2yc6l5usaq --from alice
func (k msgServer) GrantPermission(goCtx context.Context, msg *types.MsgGrantPermission) (*types.MsgGrantPermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	auth, found := k.Keeper.GetAuthorization(ctx)
	if !found {
		return nil, types.ErrAuthorizationNotFound
	}

	if msg.Creator != auth.RootAdmin {
		return nil, types.ErrUnauthorized
	}

	if auth.Permissions == nil {
		mapAddress := make(map[string]*types.AddressList)
		mapAddress[msg.Name] = &types.AddressList{
			Addresses: []string{msg.Grantee},
		}
		auth.Permissions = &types.Permissions{
			MapName: mapAddress,
		}
	} else {
		if auth.Permissions.MapName[msg.Name] == nil {
			auth.Permissions.MapName[msg.Name] = &types.AddressList{
				Addresses: []string{msg.Grantee},
			}
		} else {
			mapAll := make(map[string]string)
			for _, addr := range auth.Permissions.MapName[msg.Name].Addresses {
				mapAll[addr] = addr
			}
			if _, found := mapAll[msg.Grantee]; !found {
				auth.Permissions.MapName[msg.Name].Addresses = append(auth.Permissions.MapName[msg.Name].Addresses, msg.Grantee)
			} else {
				return nil, types.ErrGranteeExists
			}
		}
	}

	k.Keeper.SetAuthorization(ctx, auth)

	return &types.MsgGrantPermissionResponse{
		Grantee: msg.Grantee,
	}, nil
}
