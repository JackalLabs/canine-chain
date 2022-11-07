package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/x/storage/types"
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
