package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/jklaccounts/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdChoosePlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "choose-plan [tb-count]",
		Short: "Broadcast message choose-plan",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTbCount := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgChoosePlan(
				clientCtx.GetFromAddress().String(),
				argTbCount,
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
