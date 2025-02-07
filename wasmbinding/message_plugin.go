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
	storagetypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"
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

// NOTE: I think the CosmWasm bindings contract can call this multiple times in a single contract.execute()
// This would be great because we wouldn't need to change the chain code too much
func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	// If the factory contract calls one of its 'child' bindings contracts, the 'sender' field will automatically be filled in with the factory contract's address

	if msg.Custom != nil {
		// only handle the happy path where this is really posting files
		// leave everything else for the wrapped version

		var contractMsg bindings.JackalMsg

		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, sdkerrors.Wrap(err, "Failed to unmarshal CosmosMsg enum variant 'Custom' into jackal msg")
		}

		if contractMsg.PostFile != nil {
			return m.postFile(ctx, contractAddr, contractMsg.PostFile)
		}
		if contractMsg.DeleteFile != nil {
			return m.deleteFile(ctx, contractAddr, contractMsg.DeleteFile)
		}
		if contractMsg.BuyStorage != nil {
			return m.buyStorage(ctx, contractAddr, contractMsg.BuyStorage)
		}
		if contractMsg.RequestReportForm != nil {
			return m.requestReportForm(ctx, contractAddr, contractMsg.RequestReportForm)
		}
		// Filetree msgs start here
		if contractMsg.PostFileTree != nil {
			return m.postFileTree(ctx, contractAddr, contractMsg.PostFileTree)
		}
		if contractMsg.AddViewers != nil {
			return m.addViewers(ctx, contractAddr, contractMsg.AddViewers)
		}
		if contractMsg.PostKey != nil {
			return m.postKey(ctx, contractAddr, contractMsg.PostKey)
		}
		if contractMsg.DeleteFileTree != nil {
			return m.deleteFileTree(ctx, contractAddr, contractMsg.DeleteFileTree)
		}
		if contractMsg.RemoveViewers != nil {
			return m.removeViewers(ctx, contractAddr, contractMsg.RemoveViewers)
		}
		if contractMsg.ProvisionFileTree != nil {
			return m.provisionFileTree(ctx, contractAddr, contractMsg.ProvisionFileTree)
		}
		if contractMsg.AddEditors != nil {
			return m.addEditors(ctx, contractAddr, contractMsg.AddEditors)
		}
	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

func (m *CustomMessenger) postFile(ctx sdk.Context, contractAddr sdk.AccAddress, postFile *storagetypes.MsgPostFile) ([]sdk.Event, [][]byte, error) {
	err := PerformPostFile(m.storage, ctx, contractAddr, postFile)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform post file")
	}
	return nil, nil, nil
}

func (m *CustomMessenger) deleteFile(ctx sdk.Context, contractAddr sdk.AccAddress, deleteFile *storagetypes.MsgDeleteFile) ([]sdk.Event, [][]byte, error) {
	err := PerformDeleteFile(m.storage, ctx, contractAddr, deleteFile)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform delete file")
	}
	return nil, nil, nil
}

func (m *CustomMessenger) buyStorage(ctx sdk.Context, contractAddr sdk.AccAddress, buyStorage *storagetypes.MsgBuyStorage) ([]sdk.Event, [][]byte, error) {
	err := PerformBuyStorage(m.storage, ctx, contractAddr, buyStorage)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform buy storage")
	}
	return nil, nil, nil
}

func (m *CustomMessenger) requestReportForm(ctx sdk.Context, contractAddr sdk.AccAddress, requestReportForm *storagetypes.MsgRequestReportForm) ([]sdk.Event, [][]byte, error) {
	err := PerformRequestReportForm(m.storage, ctx, contractAddr, requestReportForm)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform request report form")
	}
	return nil, nil, nil
}

// Filetree starts here
func (m *CustomMessenger) postFileTree(ctx sdk.Context, contractAddr sdk.AccAddress, postFileTree *filetreetypes.MsgPostFile) ([]sdk.Event, [][]byte, error) {
	err := PerformPostFileTree(m.filetree, ctx, contractAddr, postFileTree)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform post file tree")
	}
	return nil, nil, nil
}

func (m *CustomMessenger) addViewers(ctx sdk.Context, contractAddr sdk.AccAddress, addViewers *filetreetypes.MsgAddViewers) ([]sdk.Event, [][]byte, error) {
	err := PerformAddViewers(m.filetree, ctx, contractAddr, addViewers)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform add viewers")
	}
	return nil, nil, nil
}

func (m *CustomMessenger) postKey(ctx sdk.Context, contractAddr sdk.AccAddress, postKey *filetreetypes.MsgPostKey) ([]sdk.Event, [][]byte, error) {
	err := PerformPostKey(m.filetree, ctx, contractAddr, postKey)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform post key")
	}
	return nil, nil, nil
}

func (m *CustomMessenger) deleteFileTree(ctx sdk.Context, contractAddr sdk.AccAddress, deleteFileTree *filetreetypes.MsgDeleteFile) ([]sdk.Event, [][]byte, error) {
	err := PerformDeleteFileTree(m.filetree, ctx, contractAddr, deleteFileTree)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform delete file tree")
	}
	return nil, nil, nil
}

func (m *CustomMessenger) removeViewers(ctx sdk.Context, contractAddr sdk.AccAddress, removeViewers *filetreetypes.MsgRemoveViewers) ([]sdk.Event, [][]byte, error) {
	err := PerformRemoveViewers(m.filetree, ctx, contractAddr, removeViewers)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform remove viewers")
	}
	return nil, nil, nil
}

func (m *CustomMessenger) provisionFileTree(ctx sdk.Context, contractAddr sdk.AccAddress, provisionFileTree *filetreetypes.MsgProvisionFileTree) ([]sdk.Event, [][]byte, error) {
	err := PerformProvisionFileTree(m.filetree, ctx, contractAddr, provisionFileTree)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform provision filetree")
	}
	return nil, nil, nil
}

func (m *CustomMessenger) addEditors(ctx sdk.Context, contractAddr sdk.AccAddress, addEditors *filetreetypes.MsgAddEditors) ([]sdk.Event, [][]byte, error) {
	err := PerformAddEditors(m.filetree, ctx, contractAddr, addEditors)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform add editors")
	}
	return nil, nil, nil
}
