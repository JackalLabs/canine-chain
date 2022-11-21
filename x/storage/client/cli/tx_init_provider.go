package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/x/storage/types"
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
			argTotalspace := args[1]
			argKeybase := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInitProvider(
				clientCtx.GetFromAddress().String(),
				argIP,
				argTotalspace,
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
