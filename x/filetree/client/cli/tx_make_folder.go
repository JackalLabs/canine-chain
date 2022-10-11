package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMakeFolder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make-folder [account] [root-hash-path] [contents] [editors] [viewers] [tracking-number]",
		Short: "Broadcast message makeFolder",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAccount := args[0]
			argRootHashPath := args[1]
			argContents := args[2]
			argEditors := args[3]
			argViewers := args[4]
			argTrackingNumber, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMakeFolder(
				clientCtx.GetFromAddress().String(),
				argAccount,
				argRootHashPath,
				argContents,
				argEditors,
				argViewers,
				argTrackingNumber,
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
