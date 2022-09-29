package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/notifications/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateNotifications() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-notifications [count] [notification] [address]",
		Short: "Create a new notifications",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexCount, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			// Get value arguments
			argNotification := args[1]
			argAddress := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateNotifications(
				clientCtx.GetFromAddress().String(),
				indexCount,
				argNotification,
				argAddress,
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

func CmdUpdateNotifications() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-notifications [count] [notification] [address]",
		Short: "Update a notifications",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexCount, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			// Get value arguments
			argNotification := args[1]
			argAddress := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateNotifications(
				clientCtx.GetFromAddress().String(),
				indexCount,
				argNotification,
				argAddress,
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

func CmdDeleteNotifications() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-notifications [count]",
		Short: "Delete a notifications",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexCount, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteNotifications(
				clientCtx.GetFromAddress().String(),
				indexCount,
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
