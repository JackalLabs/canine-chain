package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	filetreekeeper "github.com/jackalLabs/canine-chain/x/filetree/keeper"
)

func RegisterCustomPlugins(
	bank *bankkeeper.BaseKeeper,
	filetree *filetreekeeper.Keeper,
) []wasmkeeper.Option {
	wasmQueryPlugin := NewQueryPlugin(filetree)

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})
	// messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
	// 	CustomMessageDecorator(bank, tokenFactory),
	// )

	return []wasm.Option{
		queryPluginOpt,
		// messengerDecoratorOpt,
	}
}
