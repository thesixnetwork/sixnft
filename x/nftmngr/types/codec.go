package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateNFTSchema{}, "nftmngr/CreateNFTSchema", nil)
	cdc.RegisterConcrete(&MsgCreateMetadata{}, "nftmngr/CreateMetadata", nil)
	cdc.RegisterConcrete(&MsgPerformActionByAdmin{}, "nftmngr/PerformActionByAdmin", nil)
	cdc.RegisterConcrete(&MsgSetNFTAttribute{}, "nftmngr/SetNFTAttribute", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateNFTSchema{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMetadata{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPerformActionByAdmin{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetNFTAttribute{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
