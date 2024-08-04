package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDelRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del-record [name]",
		Short: "Broadcast message DelRecord",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDelRecord(
				clientCtx.GetFromAddress().String(),
				argName,
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
