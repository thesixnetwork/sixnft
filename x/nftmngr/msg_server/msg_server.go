package msg_server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/thesixnetwork/sixnft/x/nftmngr/keeper"
	"github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

type msg_server struct {
	keeper.Keeper
	cdc            codec.BinaryCodec
	nftadminKeeper types.NftadminKeeper
	bankKeeper     types.BankKeeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper keeper.Keeper) types.MsgServer {
	return &msg_server{Keeper: keeper}
}

var _ types.MsgServer = msg_server{}
