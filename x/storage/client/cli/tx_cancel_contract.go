package cli

import (
	"encoding/hex"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCancelContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-file [merkle] [start]",
		Short: "Delete a file on chain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMerkle := args[0]
			argStart := args[1]

			start, err := strconv.ParseInt(argStart, 10, 64)
			if err != nil {
				panic(err)
			}
			merkle, err := hex.DecodeString(argMerkle)
			if err != nil {
				panic(err)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteFile(
				clientCtx.GetFromAddress().String(),
				merkle,
				start,
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
