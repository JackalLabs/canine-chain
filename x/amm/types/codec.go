package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateLPool{}, "lp/CreateLPool", nil)
	cdc.RegisterConcrete(&MsgJoinPool{}, "lp/JoinPool", nil)
	cdc.RegisterConcrete(&MsgExitPool{}, "lp/ExitPool", nil)
	cdc.RegisterConcrete(&MsgSwap{}, "lp/Swap", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateLPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgJoinPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgExitPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSwap{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
