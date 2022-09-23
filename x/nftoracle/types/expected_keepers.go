package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	admintypes "github.com/thesixnetwork/sixnft/x/admin/types"
	nftmngrtypes "github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

type NftmngrKeep interface {
	GetNFTSchema(
		ctx sdk.Context,
		code string,

	) (val nftmngrtypes.NFTSchema, found bool)

	GetNftData(
		ctx sdk.Context,
		nftSchemaCode string,
		tokenId string,

	) (val nftmngrtypes.NftData, found bool)

	// ValidateNFTData(data *nftmngrtypes.NftData, schema *nftmngrtypes.NFTSchema) (bool, error)
	SetNftData(ctx sdk.Context, nftData nftmngrtypes.NftData)
	GetActionByRefId(
		ctx sdk.Context,
		refId string,

	) (val nftmngrtypes.ActionByRefId, found bool)
	SetActionByRefId(ctx sdk.Context, actionByRefId nftmngrtypes.ActionByRefId)
}

type AdminKeeper interface {
	GetAuthorization(ctx sdk.Context) (val admintypes.Authorization, found bool)
	HasPermission(ctx sdk.Context, name string, addr sdk.AccAddress) bool
}
