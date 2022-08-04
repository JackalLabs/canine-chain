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

func CmdBuyStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-storage [for-address] [start-block] [duration] [bytes] [payment-denom]",
		Short: "Broadcast message buy-storage",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argForAddress := args[0]
			argStartBlock := args[1]
			argDuration := args[2]
			argBytes := args[3]
			argPaymentDenom := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyStorage(
				clientCtx.GetFromAddress().String(),
				argForAddress,
				argStartBlock,
				argDuration,
				argBytes,
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
