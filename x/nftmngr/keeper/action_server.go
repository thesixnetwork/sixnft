package keeper

import (

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (k Keeper) AddAction(ctx sdk.Context, signer sdk.AccAddress, nftSchemaName string, newAction types.Action) error {
	creator := signer.String()

	// get existing action in schema
	schema, schemaFound := k.GetNFTSchema(ctx, nftSchemaName)
	if !schemaFound {
		return sdkerrors.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if creator != schema.Owner {
		return sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}


	// validate Action data
  err := ValidateAction(&newAction, &schema)
	if err != nil {
		return sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	// append new action
	schema.OnchainData.Actions = append(schema.OnchainData.Actions, &newAction)

	// save index of action
	k.SetActionOfSchema(ctx, types.ActionOfSchema{
		Name:          newAction.Name,
		NftSchemaCode: schema.Code,
		Index:         uint64(len(schema.OnchainData.Actions) - 1),
	})

	// save schema
	k.SetNFTSchema(ctx, schema)

	return nil
}
