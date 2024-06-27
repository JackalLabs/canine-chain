package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSetProviderKeybase() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-key [keybase-identity]",
		Short: "Broadcast message set-provider-ip",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argKey := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetProviderKeybase(
				clientCtx.GetFromAddress().String(),
				argKey,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}
