package cli

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/notifications/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddSenders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-senders [sender-ids]",
		Short: "Broadcast message addSenders",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSenderIds := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			senderAddresses := strings.Split(argSenderIds, ",")

			var senderIds []string

			for _, v := range senderAddresses {
				if len(v) < 1 {
					continue
				}
				senderIds = append(senderIds, v)
			}

			jsonSenders, err := json.Marshal(senderIds)
			if err != nil {
				return err
			}

			fmt.Println(string(jsonSenders))

			msg := types.NewMsgAddSenders(
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
