package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBuyStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-storage [for-address] [duration] [bytes] [payment-denom]",
		Short: "Broadcast message buy-storage",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argForAddress := args[0]
			argDuration := args[1]
			argBytes := args[2]
			argPaymentDenom := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyStorage(
				clientCtx.GetFromAddress().String(),
				argForAddress,
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

func CmdBuyStorageToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-storage-token [amount] [payment-denom]",
		Short: "Broadcast message buy-storage",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]
			argPaymentDenom := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := strconv.ParseInt(argAmount, 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyStorageToken(
				clientCtx.GetFromAddress().String(),
				amount,
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
