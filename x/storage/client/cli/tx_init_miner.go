package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/storage/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdInitProvider() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init-provider [ip] [totalspace]",
		Short: "Broadcast message init-provider",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argIp := args[0]
			argTotalspace := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInitProvider(
				clientCtx.GetFromAddress().String(),
				argIp,
				argTotalspace,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
