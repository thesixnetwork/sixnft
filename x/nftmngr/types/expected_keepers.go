package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	evmtypes "sixnft/x/evmsupport/types"
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

type EvmsupportKeeper interface {
	SetAddressBinding(ctx sdk.Context, addressBinding evmtypes.AddressBinding)
	GetAddressBinding(
		ctx sdk.Context,
		ethAddress string,
	) (val evmtypes.AddressBinding, found bool)
}
