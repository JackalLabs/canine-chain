// Package rpc contains RPC handler methods and utilities to start
// Ethermint's Web3-compatibly JSON-RPC server.
package liteemrpc

import (
	"fmt"

	servertypes "github.com/cosmos/cosmos-sdk/server"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/ethereum/go-ethereum/rpc"

	rpcethtypes "github.com/jackalLabs/canine-chain/x/ethercoat/liteemrpc/types"

	"github.com/jackalLabs/canine-chain/x/ethercoat/liteemrpc/eth"
	"github.com/jackalLabs/canine-chain/x/ethercoat/liteemrpc/eth/backend"

	rpcclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
)

// RPC namespaces and API version
const (
	CosmosNamespace = "cosmos"
	EthNamespace    = "eth"

	apiVersion = "1.0"
)

// APICreator creates the JSON-RPC API implementations.
type APICreator = func(
	ctx *servertypes.Context,
	clientCtx client.Context,
	tendermintWebsocketClient *rpcclient.WSClient,
	allowUnprotectedTxs bool,
) []rpc.API

// apiCreators defines the JSON-RPC API namespaces.
var apiCreators map[string]APICreator

// populating apiCreators
func init() {
	apiCreators = map[string]APICreator{
		EthNamespace: func(ctx *servertypes.Context, clientCtx client.Context, tmWSClient *rpcclient.WSClient, allowUnprotectedTxs bool) []rpc.API {
			nonceLock := new(rpcethtypes.AddrLocker)
			evmBackend := backend.NewBackend(ctx, ctx.Logger, clientCtx, allowUnprotectedTxs)

			return []rpc.API{
				{
					Namespace: EthNamespace,
					Version:   apiVersion,
					Service:   eth.NewPublicAPI(ctx.Logger, clientCtx, evmBackend, nonceLock),
					Public:    true,
				},
			}
		},
	}
}

// GetRPCAPIs returns the list of all APIs
func GetRPCAPIs(ctx *servertypes.Context, clientCtx client.Context, tmWSClient *rpcclient.WSClient, allowUnprotectedTxs bool, selectedAPIs []string) []rpc.API {
	var apis []rpc.API
	for _, ns := range selectedAPIs {
		if creator, ok := apiCreators[ns]; ok {
			apis = append(apis, creator(ctx, clientCtx, tmWSClient, allowUnprotectedTxs)...)
		} else {
			ctx.Logger.Error("invalid namespace value", "namespace", ns)
		}
	}

	return apis
}

// RegisterAPINamespace registers a new API namespace with the API creator.
// This function fails if the namespace is already registered.
func RegisterAPINamespace(ns string, creator APICreator) error {
	if _, ok := apiCreators[ns]; ok {
		return fmt.Errorf("duplicated api namespace %s", ns)
	}
	apiCreators[ns] = creator
	return nil
}
