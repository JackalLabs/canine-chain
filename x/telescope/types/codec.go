package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegister{}, "telescope/Register", nil)
	cdc.RegisterConcrete(&MsgBid{}, "telescope/Bid", nil)
	cdc.RegisterConcrete(&MsgAcceptBid{}, "telescope/AcceptBid", nil)
	cdc.RegisterConcrete(&MsgCancelBid{}, "telescope/CancelBid", nil)
	cdc.RegisterConcrete(&MsgList{}, "telescope/List", nil)
	cdc.RegisterConcrete(&MsgBuy{}, "telescope/Buy", nil)
	cdc.RegisterConcrete(&MsgDelist{}, "telescope/Delist", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegister{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBid{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAcceptBid{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelBid{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgList{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuy{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDelist{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
