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
	cdc.RegisterConcrete(&MsgAddAttribute{}, "nftmngr/AddAttribute", nil)
	cdc.RegisterConcrete(&MsgAddAction{}, "nftmngr/AddAction", nil)
	cdc.RegisterConcrete(&MsgSetNFTAttribute{}, "nftmngr/SetNFTAttribute", nil)
	cdc.RegisterConcrete(&MsgSetBaseUri{}, "nftmngr/SetBaseUri", nil)
	cdc.RegisterConcrete(&MsgToggleAction{}, "nftmngr/ToggleAction", nil)
	cdc.RegisterConcrete(&MsgChangeSchemaOwner{}, "nftmngr/ChangeSchemaOwner", nil)
	cdc.RegisterConcrete(&MsgAddSystemActioner{}, "nftmngr/AddSystemActioner", nil)
	cdc.RegisterConcrete(&MsgRemoveSystemActioner{}, "nftmngr/RemoveSystemActioner", nil)
	cdc.RegisterConcrete(&MsgResyncAttributes{}, "nftmngr/ResyncAttributes", nil)
	cdc.RegisterConcrete(&MsgShowAttributes{}, "nftmngr/ShowAttributes", nil)
	cdc.RegisterConcrete(&MsgSetFeeConfig{}, "nftmngr/SetFeeConfig", nil)
	cdc.RegisterConcrete(&MsgSetMintauth{}, "nftmngr/SetMintauth", nil)
	cdc.RegisterConcrete(&MsgChageOrgOwner{}, "nftmngr/ChageOrgOwner", nil)
	cdc.RegisterConcrete(&MsgShowNftAttributes{}, "nftmngr/ShowNftAttributes", nil)
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
		&MsgAddAttribute{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddAction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetNFTAttribute{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgToggleAction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetBaseUri{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgChangeSchemaOwner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddSystemActioner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveSystemActioner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgResyncAttributes{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgShowAttributes{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetFeeConfig{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetMintauth{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgChageOrgOwner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgShowNftAttributes{},
	)
	// this line is used by starport scaffolding # 3
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
