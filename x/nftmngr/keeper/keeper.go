package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/thesixnetwork/sixnft/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		evmsupportKeeper   types.EvmsupportKeeper
		nftadminKeeper     types.NftadminKeeper
		bankKeeper         types.BankKeeper
		stakingKeeper      types.StakingKeeper
		distributionKeeper types.DistributionKeeper
	}
)

func  NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	evmsupportKeeper types.EvmsupportKeeper,
	nftadminKeeper types.NftadminKeeper,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
	distributionKeeper types.DistributionKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:                cdc,
		storeKey:           storeKey,
		memKey:             memKey,
		paramstore:         ps,
		evmsupportKeeper:   evmsupportKeeper,
		nftadminKeeper:     nftadminKeeper,
		bankKeeper:         bankKeeper,
		stakingKeeper:      stakingKeeper,
		distributionKeeper: distributionKeeper,
	}
}

func (k Keeper) GetCodec() codec.BinaryCodec {
	return k.cdc
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
