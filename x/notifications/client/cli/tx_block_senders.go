package cli

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v4/x/notifications/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBlockSenders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block-senders [sender-ids]",
		Short: "Broadcast message BlockSenders",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSenderIDs := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			senderAddresses := strings.Split(argSenderIDs, ",")

			var senderIDs []string

			for _, v := range senderAddresses {
				if len(v) < 1 {
					continue
				}
				senderIDs = append(senderIDs, v)
			}

			jsonSenders, err := json.Marshal(senderIDs)
			if err != nil {
				return err
			}

			msg := types.NewMsgBlockSenders(
				clientCtx.GetFromAddress().String(),
				string(jsonSenders),
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
