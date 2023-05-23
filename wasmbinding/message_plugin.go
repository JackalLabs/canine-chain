package wasmbinding

import (
	"encoding/json"
	"errors"

	errorsmod "cosmossdk.io/errors"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackalLabs/canine-chain/wasmbinding/bindings"
	filetreekeeper "github.com/jackalLabs/canine-chain/x/filetree/keeper"
	filetreetypes "github.com/jackalLabs/canine-chain/x/filetree/types"
)

// CustomMessageDecorator returns decorator for custom CosmWasm bindings messages
func CustomMessageDecorator(filetree *filetreekeeper.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped:  old,
			filetree: filetree,
		}
	}
}

type CustomMessenger struct {
	wrapped  wasmkeeper.Messenger
	filetree *filetreekeeper.Keeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

// DispatchMsg executes on the contractMsg.
func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		// only handle the happy path where this is really posting files
		// leave everything else for the wrapped version
		var contractMsg bindings.JackalMsg
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, errorsmod.Wrap(err, "Jackal msg")
		}
		if contractMsg.PostFiles != nil {
			return m.postFiles(ctx, contractAddr, contractMsg.PostFiles) // need this
		}

	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

// postFiles posts a Files struct to chain
func (m *CustomMessenger) postFiles(ctx sdk.Context, contractAddr sdk.AccAddress, postFiles *bindings.PostFiles) ([]sdk.Event, [][]byte, error) {
	err := PerformPostFiles(m.filetree, ctx, contractAddr, postFiles)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "perform post files")
	}
	return nil, nil, nil
}

// PerformPostFiles is used with postFiles to post a Files struct to chian; validates the msgPostFiles.
func PerformPostFiles(f *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, postFiles *bindings.PostFiles) error {
	if postFiles == nil {
		return wasmvmtypes.InvalidRequest{Err: "post files null post files"}
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*f)

	msgPostFiles := filetreetypes.NewMsgPostFile(
		postFiles.Creator,
		postFiles.Account,
		postFiles.HashParent,
		postFiles.HashChild,
		postFiles.Contents,
		postFiles.Viewers,
		postFiles.Editors,
		postFiles.TrackingNumber,
	)

	if err := msgPostFiles.ValidateBasic(); err != nil {
		return errorsmod.Wrap(err, "failed validating MsgPostFiles")
	}

	// Post files
	_, err := msgServer.PostFile(
		sdk.WrapSDKContext(ctx),
		msgPostFiles,
	)
	if err != nil {
		return errorsmod.Wrap(err, "creating denom")
	}
	return nil
}

// PROBABLY DON'T NEED THIS BELOW

// Change to GetFiles
// RetrieveFiles is a function, not method, so the message_plugin can use it
// address is the on-chain address/path of the file and owner is the owner address
func RetrieveFiles(ctx sdk.Context, filetree *filetreekeeper.Keeper, owner string, address string) (string, error) {
	// Address validation
	if _, err := parseAddress(owner); err != nil {
		return "", err
	}
	// Careful to make sure it only errors if not found
	// WrapGetFiles uses built in String() method returned by files.pb.go class file
	// Be warned that the struct will be outputted on one line
	Files, err := wrapGetFiles(ctx, filetree, owner, address)
	if err != nil {
		return "", errorsmod.Wrap(err, "validate sub-denom")
	}

	return Files, nil
}

// TO REPLACE
func wrapGetFiles(ctx sdk.Context, filetree *filetreekeeper.Keeper, owner string, address string) (string, error) {
	Files, found := filetree.GetFiles(ctx, address, owner)
	if !found {
		return "", errors.New("files not found")
	}

	return Files.String(), nil
}

// parseAddress parses address from bech32 string and verifies its format.
func parseAddress(addr string) (sdk.AccAddress, error) {
	parsed, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, errorsmod.Wrap(err, "address from bech32")
	}
	err = sdk.VerifyAddressFormat(parsed)
	if err != nil {
		return nil, errorsmod.Wrap(err, "verify address format")
	}
	return parsed, nil
}
