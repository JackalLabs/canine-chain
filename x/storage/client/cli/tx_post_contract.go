package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPostContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post-contract [priceamt] [pricedenom] [hashes] [signee] [duration] [filesize] [fid]",
		Short: "Broadcast message post-contract",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argHashes := args[0]
			argSignee := args[1]
			argDuration := args[2]
			argFilesize := args[3]
			argFid := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPostContract(
				clientCtx.GetFromAddress().String(),
				argHashes,
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
