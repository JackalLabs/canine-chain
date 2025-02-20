package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	filetreekeeper "github.com/jackalLabs/canine-chain/v4/x/filetree/keeper"
	notificationskeeper "github.com/jackalLabs/canine-chain/v4/x/notifications/keeper"
	storagekeeper "github.com/jackalLabs/canine-chain/v4/x/storage/keeper"
)

func RegisterCustomPlugins(
	filetree *filetreekeeper.Keeper,
	storage *storagekeeper.Keeper,
	notifications *notificationskeeper.Keeper,
) []wasmkeeper.Option {
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(filetree, storage, notifications),
	)

	return []wasm.Option{
		messengerDecoratorOpt,
	}
}
