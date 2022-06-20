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
	cdc.RegisterConcrete(&MsgCreateMiners{}, "jklmining/CreateMiners", nil)
	cdc.RegisterConcrete(&MsgUpdateMiners{}, "jklmining/UpdateMiners", nil)
	cdc.RegisterConcrete(&MsgDeleteMiners{}, "jklmining/DeleteMiners", nil)
	cdc.RegisterConcrete(&MsgClaimSave{}, "jklmining/ClaimSave", nil)
	cdc.RegisterConcrete(&MsgCreateMinerClaims{}, "jklmining/CreateMinerClaims", nil)
	cdc.RegisterConcrete(&MsgUpdateMinerClaims{}, "jklmining/UpdateMinerClaims", nil)
	cdc.RegisterConcrete(&MsgDeleteMinerClaims{}, "jklmining/DeleteMinerClaims", nil)
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
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMiners{},
		&MsgUpdateMiners{},
		&MsgDeleteMiners{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaimSave{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMinerClaims{},
		&MsgUpdateMinerClaims{},
		&MsgDeleteMinerClaims{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
