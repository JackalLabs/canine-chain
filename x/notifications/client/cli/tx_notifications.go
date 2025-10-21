package cli

import (
	"encoding/hex"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v5/x/notifications/types"
	"github.com/spf13/cobra"
)

func CmdCreateNotification() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-notifications [to] [content] [private-contents]",
		Short: "Create a new notifications",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			byteData, err := hex.DecodeString(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateNotification(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
				byteData,
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

			msg := types.NewMsgDeleteNotification(
				clientCtx.GetFromAddress().String(),
				args[0],
				ts,
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
