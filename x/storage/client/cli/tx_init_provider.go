package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdInitProvider() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [ip] [totalspace] [keybase-identity]",
		Short: "init provider",
		Long:  "Initialize a provider with given parameters.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argIP := args[0]
			argTotalSpace := args[1]
			argKeybase := args[2]
			space, err := strconv.ParseInt(argTotalSpace, 10, 64)
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInitProvider(
				clientCtx.GetFromAddress().String(),
				argIP,
				space,
				argKeybase,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}

func CmdShutdownProvider() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shutdown",
		Short: "shutdown provider",
		Long:  "Shutdown a provider.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgShutdownProvider(
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}
