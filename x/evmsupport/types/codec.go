package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBindAddress{}, "evmsupport/BindAddress", nil)
	cdc.RegisterConcrete(&MsgRemoveBinding{}, "evmsupport/RemoveBinding", nil)
	cdc.RegisterConcrete(&MsgCreateActionSigner{}, "evmsupport/CreateActionSigner", nil)
	cdc.RegisterConcrete(&MsgUpdateActionSigner{}, "evmsupport/UpdateActionSigner", nil)
	cdc.RegisterConcrete(&MsgDeleteActionSigner{}, "evmsupport/DeleteActionSigner", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBindAddress{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveBinding{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateActionSigner{},
		&MsgUpdateActionSigner{},
		&MsgDeleteActionSigner{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
