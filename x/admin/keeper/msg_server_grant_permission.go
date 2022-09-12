package keeper

import (
	"context"

	"sixnft/x/admin/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
			auth.Permissions.MapName[msg.Name].Addresses = append(auth.Permissions.MapName[msg.Name].Addresses, msg.Grantee)
		}
	}

	k.Keeper.SetAuthorization(ctx, auth)

	return &types.MsgGrantPermissionResponse{
		Grantee: msg.Grantee,
	}, nil
}
