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
	cdc.RegisterConcrete(&MsgPostkey{}, "filetree/Postkey", nil)
	cdc.RegisterConcrete(&MsgDeleteFile{}, "filetree/DeleteFile", nil)
	cdc.RegisterConcrete(&MsgRemoveViewers{}, "filetree/RemoveViewers", nil)
	cdc.RegisterConcrete(&MsgMakeRoot{}, "filetree/MakeRoot", nil)
	cdc.RegisterConcrete(&MsgMakeRootV2{}, "filetree/MakeRootV2", nil)
	cdc.RegisterConcrete(&MsgAddEditors{}, "filetree/AddEditors", nil)
	cdc.RegisterConcrete(&MsgRemoveEditors{}, "filetree/RemoveEditors", nil)
	cdc.RegisterConcrete(&MsgResetEditors{}, "filetree/ResetEditors", nil)
	cdc.RegisterConcrete(&MsgResetViewers{}, "filetree/ResetViewers", nil)
	cdc.RegisterConcrete(&MsgChangeOwner{}, "filetree/ChangeOwner", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPostFile{},
		&MsgAddViewers{},
		&MsgPostkey{},
		&MsgDeleteFile{},
		&MsgRemoveViewers{},
		&MsgMakeRoot{},
		&MsgMakeRootV2{},
		&MsgAddEditors{},
		&MsgRemoveEditors{},
		&MsgResetEditors{},
		&MsgResetViewers{},
		&MsgChangeOwner{},
	)

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
