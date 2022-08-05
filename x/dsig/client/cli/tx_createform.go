package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/dsig/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateform() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "createform [fid] [signees]",
		Short: "Broadcast message createform",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFid := args[0]
			argSignees := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateform(
				clientCtx.GetFromAddress().String(),
				argFid,
				argSignees,
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
