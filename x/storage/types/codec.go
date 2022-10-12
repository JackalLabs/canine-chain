package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgPostContract{}, "storage/PostContract", nil)
	cdc.RegisterConcrete(&MsgCreateContracts{}, "storage/CreateContracts", nil)
	cdc.RegisterConcrete(&MsgUpdateContracts{}, "storage/UpdateContracts", nil)
	cdc.RegisterConcrete(&MsgDeleteContracts{}, "storage/DeleteContracts", nil)
	cdc.RegisterConcrete(&MsgCreateProofs{}, "storage/CreateProofs", nil)
	cdc.RegisterConcrete(&MsgUpdateProofs{}, "storage/UpdateProofs", nil)
	cdc.RegisterConcrete(&MsgDeleteProofs{}, "storage/DeleteProofs", nil)
	cdc.RegisterConcrete(&MsgItem{}, "storage/Item", nil)
	cdc.RegisterConcrete(&MsgPostproof{}, "storage/Postproof", nil)
	cdc.RegisterConcrete(&MsgCreateActiveDeals{}, "storage/CreateActiveDeals", nil)
	cdc.RegisterConcrete(&MsgUpdateActiveDeals{}, "storage/UpdateActiveDeals", nil)
	cdc.RegisterConcrete(&MsgDeleteActiveDeals{}, "storage/DeleteActiveDeals", nil)
	cdc.RegisterConcrete(&MsgSignContract{}, "storage/SignContract", nil)
	cdc.RegisterConcrete(&MsgCreateProviders{}, "storage/CreateProviders", nil)
	cdc.RegisterConcrete(&MsgUpdateProviders{}, "storage/UpdateProviders", nil)
	cdc.RegisterConcrete(&MsgDeleteProviders{}, "storage/DeleteProviders", nil)
	cdc.RegisterConcrete(&MsgSetProviderIp{}, "storage/SetProviderIp", nil)
	cdc.RegisterConcrete(&MsgSetProviderTotalspace{}, "storage/SetProviderTotalspace", nil)
	cdc.RegisterConcrete(&MsgInitProvider{}, "storage/InitProvider", nil)
	cdc.RegisterConcrete(&MsgCancelContract{}, "storage/CancelContract", nil)
	cdc.RegisterConcrete(&MsgBuyStorage{}, "storage/BuyStorage", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPostContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateContracts{},
		&MsgUpdateContracts{},
		&MsgDeleteContracts{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProofs{},
		&MsgUpdateProofs{},
		&MsgDeleteProofs{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgItem{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPostproof{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateActiveDeals{},
		&MsgUpdateActiveDeals{},
		&MsgDeleteActiveDeals{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSignContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProviders{},
		&MsgUpdateProviders{},
		&MsgDeleteProviders{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetProviderIp{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetProviderTotalspace{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitProvider{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyStorage{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
