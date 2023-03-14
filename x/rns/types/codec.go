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
	cdc.RegisterConcrete(&MsgDelist{}, "rns/Delist", nil)
	cdc.RegisterConcrete(&MsgTransfer{}, "rns/Transfer", nil)
	cdc.RegisterConcrete(&MsgAddRecord{}, "rns/AddRecord", nil)
	cdc.RegisterConcrete(&MsgDelRecord{}, "rns/DelRecord", nil)
	cdc.RegisterConcrete(&MsgInit{}, "rns/Init", nil)
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

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
