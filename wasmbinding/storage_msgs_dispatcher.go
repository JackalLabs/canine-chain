package wasmbinding

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/wasmbinding/bindings"
	storagekeeper "github.com/jackalLabs/canine-chain/x/storage/keeper"
	storagetypes "github.com/jackalLabs/canine-chain/x/storage/types"
)

// STUB
// Remember to add in the sender
// Another dispatcher function that can be used to organise dispatching
// the storage module's messages in a different file

func (m *CustomMessenger) DispatchStorageMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg, sender string) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		// only handle the happy path where this is really posting files
		// leave everything else for the wrapped version
		var contractMsg bindings.JackalMsg
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, errorsmod.Wrap(err, "Jackal msg")
		}

		if contractMsg.BuyStorage != nil {
			return m.buyStorage(ctx, contractAddr, contractMsg.BuyStorage, sender)
		}

	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg, sender)
}

func (m *CustomMessenger) buyStorage(ctx sdk.Context, contractAddr sdk.AccAddress, buyStorage *bindings.BuyStorage, sender string) ([]sdk.Event, [][]byte, error) {
	err := PerformBuyStorage(m.storage, ctx, contractAddr, buyStorage, sender)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "perform buy storage")
	}
	return nil, nil, nil
}

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
