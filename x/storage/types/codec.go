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
	cdc.RegisterConcrete(&MsgCreateMiners{}, "storage/CreateMiners", nil)
	cdc.RegisterConcrete(&MsgUpdateMiners{}, "storage/UpdateMiners", nil)
	cdc.RegisterConcrete(&MsgDeleteMiners{}, "storage/DeleteMiners", nil)
	cdc.RegisterConcrete(&MsgSetMinerIp{}, "storage/SetMinerIp", nil)
	cdc.RegisterConcrete(&MsgSetMinerTotalspace{}, "storage/SetMinerTotalspace", nil)
	cdc.RegisterConcrete(&MsgInitMiner{}, "storage/InitMiner", nil)
	cdc.RegisterConcrete(&MsgCancelContract{}, "storage/CancelContract", nil)
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
		&MsgCreateMiners{},
		&MsgUpdateMiners{},
		&MsgDeleteMiners{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetMinerIp{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetMinerTotalspace{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitMiner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelContract{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
