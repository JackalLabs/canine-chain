package wasmbinding

import (
	"github.com/JackalLabs/jackal-wasmd/x/wasm"
	wasmkeeper "github.com/JackalLabs/jackal-wasmd/x/wasm/keeper"

	filetreekeeper "github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
	storagekeeper "github.com/jackalLabs/canine-chain/v3/x/storage/keeper"
)

func RegisterCustomPlugins(
	// we can add in more keepers here if needed
	// bank *bankkeeper.BaseKeeper,
	filetree *filetreekeeper.Keeper,
	storage *storagekeeper.Keeper,
) []wasmkeeper.Option {
	wasmQueryPlugin := NewQueryPlugin(filetree)

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(filetree, storage),
	)

	return []wasm.Option{
		queryPluginOpt,
		messengerDecoratorOpt,
	}
}
