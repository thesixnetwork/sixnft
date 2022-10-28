package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	evmtypes "github.com/thesixnetwork/sixnft/x/evmsupport/types"
	nftadmintypes "github.com/thesixnetwork/sixnft/x/nftadmin/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule string, recipientModule string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}

type EvmsupportKeeper interface {
	SetAddressBinding(ctx sdk.Context, addressBinding evmtypes.AddressBinding)
	GetAddressBinding(
		ctx sdk.Context,
		ethAddress string,
	) (val evmtypes.AddressBinding, found bool)
}

type NftadminKeeper interface {
	GetAuthorization(ctx sdk.Context) (val nftadmintypes.Authorization, found bool)
	HasPermission(ctx sdk.Context, name string, addr sdk.AccAddress) bool
}

type StakingKeeper interface {
	ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingtypes.ValidatorI // get a particular validator by consensus address
}

type DistributionKeeper interface {
	AllocateTokensToValidator(ctx sdk.Context, val stakingtypes.ValidatorI, tokens sdk.DecCoins)
}
