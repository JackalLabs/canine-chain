package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/rns/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-record [name] [record] [value] [data]",
		Short: "Broadcast message addRecord",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argRecord := args[1]
			argValue := args[2]
			argData := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddRecord(
				clientCtx.GetFromAddress().String(),
				argName,
				argRecord,
				argValue,
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
