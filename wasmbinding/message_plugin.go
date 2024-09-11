package wasmbinding

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v4/wasmbinding/bindings"
	filetreekeeper "github.com/jackalLabs/canine-chain/v4/x/filetree/keeper"

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
	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

// postFile posts a File to the storage module
func (m *CustomMessenger) postFile(ctx sdk.Context, contractAddr sdk.AccAddress, postFile *bindings.PostFile) ([]sdk.Event, [][]byte, error) {
	err := PerformPostFile(m.storage, ctx, contractAddr, postFile)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform post file")
	}
	return nil, nil, nil
}

func PerformPostFile(s *storagekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, PostFile *bindings.PostFile) error {

	if PostFile == nil {
		return wasmvmtypes.InvalidRequest{Err: "post file null error"}
	}

	sdkMsg := storagetypes.NewMsgPostFile(
		contractAddr.String(),
		PostFile.Merkle,
		PostFile.FileSize,
		PostFile.ProofInterval,
		PostFile.ProofType,
		PostFile.MaxProofs,
		PostFile.Note,
	)

	sdkMsg.Expires = PostFile.Expires

	if err := sdkMsg.ValidateBasic(); err != nil {
		return err
	}

	msgServer := storagekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.PostFile(sdk.WrapSDKContext(ctx), sdkMsg)
	if err != nil {
		return sdkerrors.Wrap(err, "post file error from message")
	}

	return nil
}
