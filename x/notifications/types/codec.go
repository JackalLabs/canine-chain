package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateNotifications{}, "notifications/CreateNotifications", nil)
	cdc.RegisterConcrete(&MsgUpdateNotifications{}, "notifications/UpdateNotifications", nil)
	cdc.RegisterConcrete(&MsgDeleteNotifications{}, "notifications/DeleteNotifications", nil)
	cdc.RegisterConcrete(&MsgSetCounter{}, "notifications/SetCounter", nil)
	cdc.RegisterConcrete(&MsgAddSenders{}, "notifications/AddSenders", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateNotifications{},
		&MsgUpdateNotifications{},
		&MsgDeleteNotifications{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetCounter{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddSenders{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
