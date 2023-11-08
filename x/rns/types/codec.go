package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegister{}, "rns/Register", nil)
	cdc.RegisterConcrete(&MsgBid{}, "rns/Bid", nil)
	cdc.RegisterConcrete(&MsgAcceptBid{}, "rns/AcceptBid", nil)
	cdc.RegisterConcrete(&MsgCancelBid{}, "rns/CancelBid", nil)
	cdc.RegisterConcrete(&MsgList{}, "rns/List", nil)
	cdc.RegisterConcrete(&MsgBuy{}, "rns/Buy", nil)
	cdc.RegisterConcrete(&MsgDeList{}, "rns/DeList", nil)
	cdc.RegisterConcrete(&MsgTransfer{}, "rns/Transfer", nil)
	cdc.RegisterConcrete(&MsgAddRecord{}, "rns/AddRecord", nil)
	cdc.RegisterConcrete(&MsgDelRecord{}, "rns/DelRecord", nil)
	cdc.RegisterConcrete(&MsgInit{}, "rns/Init", nil)
	cdc.RegisterConcrete(&MsgUpdate{}, "rns/Update", nil)
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
		&MsgDeList{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransfer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddRecord{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDelRecord{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdate{},
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
