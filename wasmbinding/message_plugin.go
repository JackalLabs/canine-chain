package wasmbinding

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	wasmkeeper "github.com/JackalLabs/jackal-wasmd/x/wasm/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackalLabs/canine-chain/v3/wasmbinding/bindings"
	filetreekeeper "github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
	filetreetypes "github.com/jackalLabs/canine-chain/v3/x/filetree/types"
	storagekeeper "github.com/jackalLabs/canine-chain/v3/x/storage/keeper"
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
	handler  *SDKMessageHandler
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

// DispatchMsg executes on the contractMsg.
func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg, sender string) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		// only handle the happy path where this is really posting files
		// leave everything else for the wrapped version
		var contractMsg bindings.JackalMsg
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, errorsmod.Wrap(err, "Jackal msg")
		}
		if contractMsg.PostKey != nil {
			return m.postKey(ctx, contractAddr, contractMsg.PostKey, sender)
		}
		if contractMsg.MakeRoot != nil {
			// We forked wasmd to get the bech32 address of the user who executed the contract :D
			// Perhaps we can take this up a notch and make wasmd consume the contract executor as sdk.AccAddress
			// and build a signature verifcation function? To investigate

			return m.makeRoot(ctx, contractAddr, contractMsg.MakeRoot, sender)
		}
		if contractMsg.PostFiles != nil {
			return m.postFiles(ctx, contractAddr, contractMsg.PostFiles, sender)
		}
		if contractMsg.DeleteFile != nil {
			return m.deleteFile(ctx, contractAddr, contractMsg.DeleteFile, sender)
		}
		if contractMsg.BuyStorage != nil {
			return m.buyStorage(ctx, contractAddr, contractMsg.BuyStorage, sender)
		}
		if contractMsg.PostAndSign != nil {
			return m.postAndSign(ctx, contractAddr, contractMsg.PostAndSign, sender)
		}
	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg, sender)
}

// postKey posts a user's public key on chain for the encryption scheme
func (m *CustomMessenger) postKey(ctx sdk.Context, contractAddr sdk.AccAddress, postKey *bindings.PostKey, sender string) ([]sdk.Event, [][]byte, error) {
	err := PerformPostKey(m.filetree, ctx, contractAddr, postKey, sender)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "perform post key")
	}
	return nil, nil, nil
}

// PerformMakeRoot is used with makeRoot to post a root Files struct to chain, as described above; validates the msgMakeRoot.
func PerformPostKey(f *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, postKey *bindings.PostKey, sender string) error {
	if postKey == nil {
		return wasmvmtypes.InvalidRequest{Err: "post key null error"}
	}

	sdkMsg := filetreetypes.NewMsgPostkey(
		sender,
		postKey.Key,
	)
	if err := sdkMsg.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*f)
	_, err := msgServer.Postkey(sdk.WrapSDKContext(ctx), sdkMsg)
	if err != nil {
		return errorsmod.Wrap(err, "post key error from message")
	}

	return nil
}

// makeRoot posts a Files struct on chain that serves as the root directory for the user's filetree
// it is the merklePath of 's'
func (m *CustomMessenger) makeRoot(ctx sdk.Context, contractAddr sdk.AccAddress, makeRoot *bindings.MakeRoot, sender string) ([]sdk.Event, [][]byte, error) {
	err := PerformMakeRoot(m.filetree, ctx, contractAddr, makeRoot, sender)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "perform make root")
	}
	return nil, nil, nil
}

// PerformMakeRoot is used with makeRoot to post a root Files struct to chain, as described above; validates the msgMakeRoot.
func PerformMakeRoot(f *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, makeRoot *bindings.MakeRoot, sender string) error {
	if makeRoot == nil {
		return wasmvmtypes.InvalidRequest{Err: "make root null make root"}
	}

	// Silly me, we can just pass in the executor of the contract to make a root

	sdkMsg := filetreetypes.NewMsgMakeRootV2(
		sender,
		makeRoot.Editors,
		makeRoot.Viewers,
		makeRoot.TrackingNumber,
	)
	if err := sdkMsg.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*f)
	_, err := msgServer.MakeRootV2(sdk.WrapSDKContext(ctx), sdkMsg)
	if err != nil {
		return errorsmod.Wrap(err, "making root from message")
	}

	return nil
}

// postFiles posts a Files struct to chain
func (m *CustomMessenger) postFiles(ctx sdk.Context, contractAddr sdk.AccAddress, postFiles *bindings.PostFiles, sender string) ([]sdk.Event, [][]byte, error) {
	err := PerformPostFiles(m.filetree, ctx, contractAddr, postFiles, sender)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "perform post files")
	}
	return nil, nil, nil
}

// PerformPostFiles is used with postFiles to post a Files struct to chian; validates the msgPostFiles.
func PerformPostFiles(f *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, postFiles *bindings.PostFiles, sender string) error {
	if postFiles == nil {
		return wasmvmtypes.InvalidRequest{Err: "post files null post files"}
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*f)

	msgPostFiles := filetreetypes.NewMsgPostFile(
		sender,
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
		return errorsmod.Wrap(err, "failed to post file:")
	}
	return nil
}

// deleteFile deletes a Files struct on chain
func (m *CustomMessenger) deleteFile(ctx sdk.Context, contractAddr sdk.AccAddress, deleteFile *bindings.DeleteFile, sender string) ([]sdk.Event, [][]byte, error) {
	err := PerformDeleteFile(m.filetree, ctx, contractAddr, deleteFile, sender)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "perform post files")
	}
	return nil, nil, nil
}

// PerformDeleteFile is used with deleteFile to delete a Files struct on chain; validates the msgDeleteFile.
func PerformDeleteFile(f *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, deleteFile *bindings.DeleteFile, sender string) error {
	if deleteFile == nil {
		return wasmvmtypes.InvalidRequest{Err: "delete file is null"}
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*f)

	msgDeleteFile := filetreetypes.NewMsgDeleteFile(
		sender,
		deleteFile.HashPath,
		deleteFile.Account,
	)

	if err := msgDeleteFile.ValidateBasic(); err != nil {
		return errorsmod.Wrap(err, "failed validating msgDeleteFile")
	}

	// Delete File
	_, err := msgServer.DeleteFile(
		sdk.WrapSDKContext(ctx),
		msgDeleteFile,
	)
	if err != nil {
		return errorsmod.Wrap(err, "failed to delete file:")
	}
	return nil
}

// Leave here because we might need this

// parseAddress parses address from bech32 string and verifies its format.
// func parseAddress(addr string) (sdk.AccAddress, error) {
// 	parsed, err := sdk.AccAddressFromBech32(addr)
// 	if err != nil {
// 		return nil, errorsmod.Wrap(err, "address from bech32")
// 	}
// 	err = sdk.VerifyAddressFormat(parsed)
// 	if err != nil {
// 		return nil, errorsmod.Wrap(err, "verify address format")
// 	}
// 	return parsed, nil
// }
