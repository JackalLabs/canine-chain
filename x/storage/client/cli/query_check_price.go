package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCheckPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check-price [duration] [bytes]",
		Short: "get price of storage for duration",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDuration := args[0]
			reqBytes := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			i, err := strconv.ParseInt(reqBytes, 10, 64)
			if err != nil {
				return err
			}
			dur, err := strconv.ParseInt(reqDuration, 10, 64)
			if err != nil {
				return err
			}
			params := &types.QueryPriceCheck{
				Duration: dur,
				Bytes:    i,
			}

			res, err := queryClient.PriceCheck(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
