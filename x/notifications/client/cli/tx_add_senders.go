package cli

import (
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

				// queryClient := notifications.NewQueryClient(clientCtx)
				// //Get user's notiCounter
				// params := &types.QueryGetNotiCounterRequest{
				// 	Address: clientCtx.GetFromAddress().String(),
				// }

				// notiCounter, err := queryClient.NotiCounter(context.Background(), params)
				// if err != nil {
				// 	return types.ErrNotiCounterNotFound
				// }

				// //purpose of this?....
				// senders := notiCounter.NotiCounter.PermittedSenders
				// var m map[string]string

				// json.Unmarshal([]byte(senders), &m)

				senderIds = append(senderIds, v)
			}

			msg := types.NewMsgAddSenders(
				clientCtx.GetFromAddress().String(),
				strings.Join(senderIds, ","),
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
