package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/v5/x/jklmint/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetInflation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-inflation",
		Short: "get the inflation rate of the network",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryInflation{}

			res, err := queryClient.Inflation(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
