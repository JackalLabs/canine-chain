package cli

import (
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v3/x/notifications/types"
	"github.com/spf13/cobra"
)

func CmdCreateNotification() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-notifications [to] [content]",
		Short: "Create a new notifications",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateNotification(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
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

func CmdDeleteNotification() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-notification [from] [time]",
		Short: "Delete a notification",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			ts, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			t := time.UnixMicro(ts)

			msg := types.NewMsgDeleteNotification(
				clientCtx.GetFromAddress().String(),
				args[0],
				t,
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
