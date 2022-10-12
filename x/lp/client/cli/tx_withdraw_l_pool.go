package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/lp/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdWithdrawLPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-l-pool \"Pool Name\" [shares]",
		Short: "Broadcast message withdrawLPool",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			burnShare, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawLPool(
				clientCtx.GetFromAddress().String(),
				args[0],
				burnShare,
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
