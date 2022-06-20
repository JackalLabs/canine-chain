package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/jklmining/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAllowSave() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allow-save [passkey] [size]",
		Short: "Allow the next file to be saved on the senders address by the mining network.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPasskey := args[0]
			argSize := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAllowSave(
				clientCtx.GetFromAddress().String(),
				argPasskey,
				argSize,
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
