package wasmbinding

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v4/wasmbinding/bindings"
	filetreekeeper "github.com/jackalLabs/canine-chain/v4/x/filetree/keeper"
	filetreetypes "github.com/jackalLabs/canine-chain/v4/x/filetree/types"

	storagekeeper "github.com/jackalLabs/canine-chain/v4/x/storage/keeper"
)

// CustomMessageDecorator returns decorator for custom CosmWasm bindings messages

func CustomMessageDecorator(filetree *filetreekeeper.Keeper, storage *storagekeeper.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped:  old,
			filetree: filetree,
			storage:  storage,
		}
	}
}

type CustomMessenger struct {
	wrapped  wasmkeeper.Messenger
	filetree *filetreekeeper.Keeper
	storage  *storagekeeper.Keeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

// DispatchMsg executes on the contractMsg.
// NOTE: did we ever use the 'contractIBCPortID' before?
// If we can give each bindings contract--owned by a user--an IBC port ID, perhaps we can use that for authenticating the sender?

// NOTE: The last arg--'sender' string--was something we could use when running our fork of wasmd because we added it to wasmd's Messenger
// interface

// NOTE: I think the CosmWasm bindings contract can call this multiple times in a single contract.execute()
// This would be great because we wouldn't need to change the chain code too much
func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg /*sender string*/) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		// only handle the happy path where this is really posting files
		// leave everything else for the wrapped version

		// TODO: retrieve the tx signer from the ctx and authenticate it against the 'sender' param of the CosmosMsg above
		var contractMsg bindings.JackalMsg

		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, sdkerrors.Wrap(err, "Failed to unmarshal CosmosMsg enum variant 'Custom' into jackal msg")
		}

		if contractMsg.PostKey != nil {
			return m.postKey(ctx, contractAddr, contractMsg.PostKey, contractMsg.PostKey.Sender)

		}
	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)

}

// postKey posts a user's public key on chain for the encryption scheme
func (m *CustomMessenger) postKey(ctx sdk.Context, contractAddr sdk.AccAddress, postKey *bindings.PostKey, sender string) ([]sdk.Event, [][]byte, error) {
	err := PerformPostKey(m.filetree, ctx, contractAddr, postKey, sender)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform post key")
	}
	return nil, nil, nil
}

func PerformPostKey(f *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, postKey *bindings.PostKey, sender string) error {
	if postKey == nil {
		return wasmvmtypes.InvalidRequest{Err: "post key null error"}
	}

	sdkMsg := filetreetypes.NewMsgPostKey(
		sender,
		postKey.Key,
	)

	if err := sdkMsg.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*f)
	_, err := msgServer.PostKey(sdk.WrapSDKContext(ctx), sdkMsg)
	if err != nil {
		return sdkerrors.Wrap(err, "post key error from message")
	}

	return nil
}
