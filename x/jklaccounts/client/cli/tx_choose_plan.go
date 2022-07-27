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
		Use:   "choose-plan [tb-count] [payment-denom]",
		Short: "Choose a payment plan for Jackal Storage",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTbCount := args[0]
			argdenom := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgChoosePlan(
				clientCtx.GetFromAddress().String(),
				argTbCount,
				argdenom,
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
