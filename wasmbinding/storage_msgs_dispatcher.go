package wasmbinding

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/wasmbinding/bindings"
)

// STUB
// Another dispatcher function that can be used to organise dispatching
// the storage module's messages in a different file
func (m *CustomMessenger) DispatchStorageMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		// only handle the happy path where this is really posting files
		// leave everything else for the wrapped version
		var contractMsg bindings.JackalMsg
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, errorsmod.Wrap(err, "Jackal msg")
		}
		// TO DO
		// Replace with storage module's messages
		if contractMsg.MakeRoot != nil {
			return m.makeRoot(ctx, contractAddr, contractMsg.MakeRoot) // need this
		}

	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}
