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

func CmdClaimSave() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-save [saveindex] [key]",
		Short: "Broadcast message claim_save",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSaveindex := args[0]
			argKey := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgClaimSave(
				clientCtx.GetFromAddress().String(),
				argSaveindex,
				argKey,
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
