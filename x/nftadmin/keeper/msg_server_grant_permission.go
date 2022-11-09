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

	// validate grantee format as 6x0000000000000000 or not
	_, err := sdk.AccAddressFromBech32(msg.Grantee)
	if err != nil {
		return nil, types.ErrInvalidGrantee
	}

	if auth.Permissions == nil {
		auth.Permissions = &types.Permissions{
			Permissions: []*types.Permission{
				{
					Name: msg.Name,
					Addresses: &types.AddressList{
						Addresses: []string{msg.Grantee},
					},
				},
			},
		}
	} else {

		// check if the permission already exists
		// if it does, append the address to the list
		// if it doesn't, create a new permission
		permissionExists := false
		for _, v := range auth.Permissions.Permissions {
			if v.Name == msg.Name {
				// check if msg.Grantee already exists in the list
				mapAll := make(map[string]string)
				for _, addr := range v.Addresses.Addresses {
					mapAll[addr] = addr
				}
				 _, found := mapAll[msg.Grantee]; 
				if !found {
					v.Addresses.Addresses = append(v.Addresses.Addresses, msg.Grantee)
				}else{
					return nil, types.ErrGranteeAlreadyExists
				}
				
				permissionExists = true
				break
			}
		}

		if !permissionExists {
			auth.Permissions.Permissions = append(auth.Permissions.Permissions, &types.Permission{
				Name: msg.Name,
				Addresses: &types.AddressList{
					Addresses: []string{msg.Grantee},
				},
			})
		}
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeGrantPermission,
			sdk.NewAttribute(types.AttributeKeyPermissionType, msg.Name),
			sdk.NewAttribute(types.AttributeKeyPermissionAddress, msg.Grantee),
			sdk.NewAttribute(types.AttributeKeyGrantPermissionStatus, "success"),
		),
	})

	k.Keeper.SetAuthorization(ctx, auth)

	return &types.MsgGrantPermissionResponse{
		Grantee: msg.Grantee,
	}, nil
}
