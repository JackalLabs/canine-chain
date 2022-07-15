package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/storage/types"
	"github.com/spf13/cobra"
)

func CmdCreateProofs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-proofs [cid] [item] [hashes]",
		Short: "Create a new proofs",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexCid := args[0]

			// Get value arguments
			argItem := args[1]
			argHashes := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateProofs(
				clientCtx.GetFromAddress().String(),
				indexCid,
				argItem,
				argHashes,
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

func CmdUpdateProofs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-proofs [cid] [item] [hashes]",
		Short: "Update a proofs",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexCid := args[0]

			// Get value arguments
			argItem := args[1]
			argHashes := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateProofs(
				clientCtx.GetFromAddress().String(),
				indexCid,
				argItem,
				argHashes,
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

func CmdDeleteProofs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-proofs [cid]",
		Short: "Delete a proofs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexCid := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteProofs(
				clientCtx.GetFromAddress().String(),
				indexCid,
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
