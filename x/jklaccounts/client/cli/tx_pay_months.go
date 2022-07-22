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

func CmdPayMonths() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pay-months [address] [months] [payment-denom]",
		Short: "Broadcast message pay-months",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argMonths := args[1]
			argPaymentDenom := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPayMonths(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argMonths,
				argPaymentDenom,
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
