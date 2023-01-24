package liteemserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	emconfig "github.com/jackalLabs/canine-chain/x/ethercoat/liteemserver/config"

	servertypes "github.com/cosmos/cosmos-sdk/server"

	"github.com/cosmos/cosmos-sdk/client"
	ethlog "github.com/ethereum/go-ethereum/log"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	rpc "github.com/jackalLabs/canine-chain/x/ethercoat/liteemrpc"
)

// StartJSONRPC starts the JSON-RPC server
func StartJSONRPC(ctx *servertypes.Context, clientCtx client.Context, tmRPCAddr, tmEndpoint string, config *emconfig.Config) (*http.Server, chan struct{}, error) {
	tmWsClient := ConnectTmWS(tmRPCAddr, tmEndpoint, ctx.Logger)

	logger := ctx.Logger.With("module", "geth")
	ethlog.Root().SetHandler(ethlog.FuncHandler(func(r *ethlog.Record) error {
		switch r.Lvl {
		case ethlog.LvlTrace, ethlog.LvlDebug:
			logger.Debug(r.Msg, r.Ctx...)
		case ethlog.LvlInfo, ethlog.LvlWarn:
			logger.Info(r.Msg, r.Ctx...)
		case ethlog.LvlError, ethlog.LvlCrit:
			logger.Error(r.Msg, r.Ctx...)
		}
		return nil
	}))

	rpcServer := ethrpc.NewServer()

	// allowUnprotectedTxs := config.JSONRPC.AllowUnprotectedTxs
	allowUnprotectedTxs := true
	rpcAPIArr := config.JSONRPC.API

	apis := rpc.GetRPCAPIs(ctx, clientCtx, tmWsClient, allowUnprotectedTxs, rpcAPIArr)
	for _, api := range apis {
		if err := rpcServer.RegisterName(api.Namespace, api.Service); err != nil {
			ctx.Logger.Error(
				"failed to register service in JSON RPC namespace",
				"namespace", api.Namespace,
				"service", api.Service,
			)
			return nil, nil, err
		}
	}
	r := mux.NewRouter()
	r.HandleFunc("/", rpcServer.ServeHTTP).Methods("POST")

	handlerWithCors := cors.Default()

	// creating an http server with a rpc handler
	httpSrv := &http.Server{
		Addr:         config.JSONRPC.Address,
		Handler:      handlerWithCors.Handler(r),
		ReadTimeout:  config.JSONRPC.HTTPTimeout,
		WriteTimeout: config.JSONRPC.HTTPTimeout,
		IdleTimeout:  config.JSONRPC.HTTPIdleTimeout,
	}
	httpSrvDone := make(chan struct{}, 1)

	errCh := make(chan error)
	go func() {
		ctx.Logger.Error("Starting JSON-RPC server", "address", config.JSONRPC.Address)
		if err := httpSrv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				close(httpSrvDone)
				return
			}

			ctx.Logger.Error("failed to start JSON-RPC server", "error", err.Error())
			errCh <- err
		}
	}()
	return httpSrv, httpSrvDone, nil
}
