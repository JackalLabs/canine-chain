package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBuyStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-storage [for-address] [duration] [bytes] [payment-denom]",
		Short: "Buy or upgrade storage plan",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argForAddress := args[0]
			argDuration := args[1]
			argBytes := args[2]
			argPaymentDenom := args[3]

			bytes, err := strconv.ParseInt(argBytes, 10, 64)
			if err != nil {
				return err
			}

			duration, err := strconv.ParseInt(argDuration, 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyStorage(
				clientCtx.GetFromAddress().String(),
				argForAddress,
				duration,
				bytes,
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
