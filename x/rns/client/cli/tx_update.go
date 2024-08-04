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

func CmdUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [name] [data]",
		Short: "Broadcast message update",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argData := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdate(
				clientCtx.GetFromAddress().String(),
				argName,
				argData,
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
