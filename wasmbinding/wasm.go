package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	filetreekeeper "github.com/jackalLabs/canine-chain/v4/x/filetree/keeper"
	storagekeeper "github.com/jackalLabs/canine-chain/v4/x/storage/keeper"
)

// WARNING: we're using wasmd from confio and not our fork atm
// wasmd sends custom CosmWasm messages WITHOUT ensuring that the broadcaster is in fact the 'sender' arg of the msg.

func RegisterCustomPlugins(
	// We can add in more keepers kere if needed
	filetree *filetreekeeper.Keeper,
	storage *storagekeeper.Keeper,
) []wasmkeeper.Option {
	// NOTE: can we add this in later so we have less code to manage?

	// wasmQueryPlugin := NewQueryPlugin(filetree)

	// queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
	// 	Custom: CustomQuerier(wasmQueryPlugin),
	// })

	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(filetree, storage),
	)

	return []wasm.Option{
		messengerDecoratorOpt,
	}

}
