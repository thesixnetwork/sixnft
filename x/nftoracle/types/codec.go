package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateMintRequest{}, "nftoracle/CreateMintRequest", nil)
	cdc.RegisterConcrete(&MsgSubmitMintResponse{}, "nftoracle/SubmitMintResponse", nil)
	cdc.RegisterConcrete(&MsgCreateActionRequest{}, "nftoracle/CreateActionRequest", nil)
	cdc.RegisterConcrete(&MsgSubmitActionResponse{}, "nftoracle/SubmitActionResponse", nil)
	cdc.RegisterConcrete(&MsgCreateVerifyCollectionOwnerRequest{}, "nftoracle/CreateVerifyCollectionOwnerRequest", nil)
	cdc.RegisterConcrete(&MsgSubmitVerifyCollectionOwner{}, "nftoracle/SubmitVerifyCollectionOwner", nil)
	cdc.RegisterConcrete(&MsgSetMinimumConfirmation{}, "nftoracle/SetMinimumConfirmation", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMintRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitMintResponse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateActionRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitActionResponse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVerifyCollectionOwnerRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitVerifyCollectionOwner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetMinimumConfirmation{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
