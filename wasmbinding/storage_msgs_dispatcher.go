package wasmbinding

import (
	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/wasmbinding/bindings"
	storagekeeper "github.com/jackalLabs/canine-chain/v3/x/storage/keeper"
	storagetypes "github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (m *CustomMessenger) buyStorage(ctx sdk.Context, contractAddr sdk.AccAddress, buyStorage *bindings.BuyStorage, sender string) ([]sdk.Event, [][]byte, error) {
	err := PerformBuyStorage(m.storage, ctx, contractAddr, buyStorage, sender)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "perform buy storage")
	}
	return nil, nil, nil
}

// Do we care that this function is public? everything is gated by our fork of wasmd anyways
func PerformBuyStorage(s *storagekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, buyStorage *bindings.BuyStorage, sender string) error {
	if buyStorage == nil {
		return wasmvmtypes.InvalidRequest{Err: "buyStorage is null"}
	}

	msgServer := storagekeeper.NewMsgServerImpl(*s)

	msgBuyStorage := storagetypes.NewMsgBuyStorage(
		sender,
		buyStorage.ForAddress,
		buyStorage.Duration,
		buyStorage.Bytes,
		buyStorage.PaymentDenom,
	)

	if err := msgBuyStorage.ValidateBasic(); err != nil {
		return errorsmod.Wrap(err, "failed validating msgBuyStorage")
	}

	// Buy Storage
	_, err := msgServer.BuyStorage(
		sdk.WrapSDKContext(ctx),
		msgBuyStorage,
	)
	if err != nil {
		return errorsmod.Wrap(err, "failed to buy storage:")
	}
	return nil
}
