package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgPostContract{}, "storage/PostContract", nil)
	cdc.RegisterConcrete(&MsgPostproof{}, "storage/Postproof", nil)
	cdc.RegisterConcrete(&MsgSignContract{}, "storage/SignContract", nil)
	cdc.RegisterConcrete(&MsgSetProviderIP{}, "storage/SetProviderIp", nil)
	cdc.RegisterConcrete(&MsgSetProviderTotalspace{}, "storage/SetProviderTotalspace", nil)
	cdc.RegisterConcrete(&MsgInitProvider{}, "storage/InitProvider", nil)
	cdc.RegisterConcrete(&MsgCancelContract{}, "storage/CancelContract", nil)
	cdc.RegisterConcrete(&MsgBuyStorage{}, "storage/BuyStorage", nil)
	cdc.RegisterConcrete(&MsgClaimStray{}, "storage/ClaimStray", nil)
	cdc.RegisterConcrete(&MsgUpgradeStorage{}, "storage/UpgradeStorage", nil)
	cdc.RegisterConcrete(&MsgSetProviderKeybase{}, "storage/SetProviderKeybase", nil)
	cdc.RegisterConcrete(&MsgAddClaimer{}, "storage/AddClaimer", nil)
	cdc.RegisterConcrete(&MsgRemoveClaimer{}, "storage/RemoveClaimer", nil)
	cdc.RegisterConcrete(&MsgAttest{}, "storage/Attest", nil)
	cdc.RegisterConcrete(&MsgReport{}, "storage/Report", nil)
	cdc.RegisterConcrete(&MsgRequestAttestationForm{}, "storage/RequestAttestationForm", nil)
	cdc.RegisterConcrete(&MsgRequestReportForm{}, "storage/RequestReportForm", nil)
	cdc.RegisterConcrete(&MsgShutdownProvider{}, "storage/ShutdownProvider", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPostContract{},
		&MsgPostproof{},
		&MsgSignContract{},
		&MsgSetProviderIP{},
		&MsgSetProviderTotalspace{},
		&MsgInitProvider{},
		&MsgCancelContract{},
		&MsgBuyStorage{},
		&MsgClaimStray{},
		&MsgUpgradeStorage{},
		&MsgSetProviderKeybase{},
		&MsgAddClaimer{},
		&MsgRemoveClaimer{},
		&MsgAttest{},
		&MsgReport{},
		&MsgRequestAttestationForm{},
		&MsgRequestReportForm{},
		&MsgShutdownProvider{},
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
