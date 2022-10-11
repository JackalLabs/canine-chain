package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgPostFile{}, "filetree/PostFile", nil)
	cdc.RegisterConcrete(&MsgAddViewers{}, "filetree/AddViewers", nil)
	cdc.RegisterConcrete(&MsgPostkey{}, "filetree/Postkey", nil)
	cdc.RegisterConcrete(&MsgInitAccount{}, "filetree/InitAccount", nil)
	cdc.RegisterConcrete(&MsgDeleteFile{}, "filetree/DeleteFile", nil)
	cdc.RegisterConcrete(&MsgInitAll{}, "filetree/InitAll", nil)
	cdc.RegisterConcrete(&MsgRemoveViewers{}, "filetree/RemoveViewers", nil)
	cdc.RegisterConcrete(&MsgMakeFolder{}, "filetree/MakeFolder", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPostFile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddViewers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPostkey{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteFile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitAll{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveViewers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMakeFolder{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
