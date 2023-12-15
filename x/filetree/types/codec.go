package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgPostFile{}, "filetree/PostFile", nil)
	cdc.RegisterConcrete(&MsgAddViewers{}, "filetree/AddViewers", nil)
	cdc.RegisterConcrete(&MsgPostKey{}, "filetree/PostKey", nil)
	cdc.RegisterConcrete(&MsgDeleteFile{}, "filetree/DeleteFile", nil)
	cdc.RegisterConcrete(&MsgRemoveViewers{}, "filetree/RemoveViewers", nil)
	cdc.RegisterConcrete(&MsgProvisionFileTree{}, "filetree/ProvisionFileTree", nil)
	cdc.RegisterConcrete(&MsgAddEditors{}, "filetree/AddEditors", nil)
	cdc.RegisterConcrete(&MsgRemoveEditors{}, "filetree/RemoveEditors", nil)
	cdc.RegisterConcrete(&MsgResetEditors{}, "filetree/ResetEditors", nil)
	cdc.RegisterConcrete(&MsgResetViewers{}, "filetree/ResetViewers", nil)
	cdc.RegisterConcrete(&MsgChangeOwner{}, "filetree/ChangeOwner", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPostFile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddViewers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPostKey{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteFile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveViewers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProvisionFileTree{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddEditors{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveEditors{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgResetEditors{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgResetViewers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgChangeOwner{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	Amino.Seal()
}
