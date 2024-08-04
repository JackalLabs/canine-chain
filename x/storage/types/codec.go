package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgPostFile{}, "storage/PostFile", nil)
	cdc.RegisterConcrete(&MsgPostProof{}, "storage/PostProof", nil)
	cdc.RegisterConcrete(&MsgSetProviderIP{}, "storage/SetProviderIp", nil)
	cdc.RegisterConcrete(&MsgSetProviderTotalSpace{}, "storage/SetProviderTotalSpace", nil)
	cdc.RegisterConcrete(&MsgInitProvider{}, "storage/InitProvider", nil)
	cdc.RegisterConcrete(&MsgDeleteFile{}, "storage/DeleteFile", nil)
	cdc.RegisterConcrete(&MsgBuyStorage{}, "storage/BuyStorage", nil)
	cdc.RegisterConcrete(&MsgSetProviderKeybase{}, "storage/SetProviderKeybase", nil)
	cdc.RegisterConcrete(&MsgAddClaimer{}, "storage/AddClaimer", nil)
	cdc.RegisterConcrete(&MsgRemoveClaimer{}, "storage/RemoveClaimer", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPostFile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPostProof{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetProviderIP{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetProviderTotalSpace{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitProvider{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteFile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyStorage{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetProviderKeybase{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddClaimer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveClaimer{},
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
