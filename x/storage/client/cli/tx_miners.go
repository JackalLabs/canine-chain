package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/storage/types"
	"github.com/spf13/cobra"
)

func CmdCreateMiners() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-miners [address] [ip] [totalspace]",
		Short: "Create a new miners",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexAddress := args[0]

			// Get value arguments
			argIp := args[1]
			argTotalspace := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateMiners(
				clientCtx.GetFromAddress().String(),
				indexAddress,
				argIp,
				argTotalspace,
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

func CmdUpdateMiners() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-miners [address] [ip] [totalspace]",
		Short: "Update a miners",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexAddress := args[0]

			// Get value arguments
			argIp := args[1]
			argTotalspace := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateMiners(
				clientCtx.GetFromAddress().String(),
				indexAddress,
				argIp,
				argTotalspace,
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

func CmdDeleteMiners() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-miners [address]",
		Short: "Delete a miners",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteMiners(
				clientCtx.GetFromAddress().String(),
				indexAddress,
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
