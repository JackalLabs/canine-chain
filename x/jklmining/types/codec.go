package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAllowSave{}, "jklmining/AllowSave", nil)
	cdc.RegisterConcrete(&MsgCreateSaveRequests{}, "jklmining/CreateSaveRequests", nil)
	cdc.RegisterConcrete(&MsgUpdateSaveRequests{}, "jklmining/UpdateSaveRequests", nil)
	cdc.RegisterConcrete(&MsgDeleteSaveRequests{}, "jklmining/DeleteSaveRequests", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAllowSave{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSaveRequests{},
		&MsgUpdateSaveRequests{},
		&MsgDeleteSaveRequests{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
