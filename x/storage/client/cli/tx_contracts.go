package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/storage/types"
	"github.com/spf13/cobra"
)

func CmdCreateContracts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-contracts [cid] [priceamt] [pricedenom] [chunks] [merkle] [signee] [duration] [filesize] [fid]",
		Short: "Create a new contracts",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexCid := args[0]

			// Get value arguments
			argChunks := args[1]
			argMerkle := args[2]
			argSignee := args[3]
			argDuration := args[4]
			argFilesize := args[5]
			argFid := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateContracts(
				clientCtx.GetFromAddress().String(),
				indexCid,
				argChunks,
				argMerkle,
				argSignee,
				argDuration,
				argFilesize,
				argFid,
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

func CmdUpdateContracts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-contracts [cid] [chunks] [merkle] [signee] [duration] [filesize] [fid]",
		Short: "Update a contracts",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexCid := args[0]

			// Get value arguments
			argChunks := args[1]
			argMerkle := args[2]
			argSignee := args[3]
			argDuration := args[4]
			argFilesize := args[5]
			argFid := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateContracts(
				clientCtx.GetFromAddress().String(),
				indexCid,
				argChunks,
				argMerkle,
				argSignee,
				argDuration,
				argFilesize,
				argFid,
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

func CmdDeleteContracts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-contracts [cid]",
		Short: "Delete a contracts",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexCid := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteContracts(
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
