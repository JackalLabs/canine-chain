package liteemserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	sdkserver "github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/types"
	emconfig "github.com/jackalLabs/canine-chain/x/ethercoat/liteemserver/config"
	"github.com/spf13/cobra"
	tcmd "github.com/tendermint/tendermint/cmd/tendermint/commands"
	"github.com/tendermint/tendermint/node"
)

const (
	flagTraceStore = "trace-store"
)

// adding JSON-RPC server commands
func AddCommands(rootCmd *cobra.Command, defaultNodeHome string, appCreator types.AppCreator, appExport types.AppExporter, addStartFlags types.ModuleInitFlags) {
	startCmd := StartCmd(appCreator, defaultNodeHome)

	rootCmd.AddCommand(startCmd)
}

// StartCmd runs the service passed in, either stand-alone or in-process with
// Tendermint.
func StartCmd(appCreator types.AppCreator, defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ethcoat",
		Short: "Run the JSON-RPC Server",
		Long:  `Running the JSON-RPC Server`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			serverCtx := GetServerContextFromCmd(cmd)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			// starting the jsonrpc server
			err = startConfig(serverCtx, clientCtx, appCreator)
			errCode, ok := err.(ErrorCode)
			if !ok {
				return err
			}

			serverCtx.Logger.Debug(fmt.Sprintf("received quit signal: %d", errCode.Code))
			return err
		},
	}
	// TODO: Add cmd flags
	tcmd.AddNodeFlags(cmd)
	return cmd
}

func startConfig(ctx *sdkserver.Context, clientCtx client.Context, appCreator types.AppCreator) error {
	cfg := ctx.Config
	home := cfg.RootDir

	emcfg := emconfig.DefaultConfig()

	genDocProvider := node.DefaultGenesisDocProviderFunc(cfg)

	var (
		httpSrv     *http.Server
		httpSrvDone chan struct{}
	)

	if emcfg.JSONRPC.Enable {
		genDoc, err := genDocProvider()
		if err != nil {
			return err
		}

		// swapping the clientctx
		// clientCtx := clientCtx.WithChainID(genDoc.ChainID)
		clientCtx := clientCtx.WithHomeDir(home).WithChainID(genDoc.ChainID)

		tmEndpoint := "/websocket"
		tmRPCAddr := cfg.RPC.ListenAddress
		httpSrv, httpSrvDone, err = StartJSONRPC(ctx, clientCtx, tmRPCAddr, tmEndpoint, emcfg)
		if err != nil {
			return err
		}
	}

	defer func() {
		if httpSrv != nil {
			shutdownCtx, cancelFn := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancelFn()

			if err := httpSrv.Shutdown(shutdownCtx); err != nil {
				ctx.Logger.Error("HTTP server shutdown produced a warning", "error", err.Error())
			} else {
				ctx.Logger.Info("HTTP server shut down, waiting 5 sec")
				select {
				case <-time.Tick(5 * time.Second):
				case <-httpSrvDone:
				}
			}
		}

		ctx.Logger.Info("exiting...")
	}()

	// wait for signal capture and gracefully return
	return WaitForQuitSignals()
}
