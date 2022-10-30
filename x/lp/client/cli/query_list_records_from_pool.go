package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackal-dao/canine/x/lp/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdListRecordsFromPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-records-from-pool [pool-name]",
		Short: "List all LProviderRecords of the pool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPoolName := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListRecordsFromPoolRequest{
				PoolName: reqPoolName,
			}

			res, err := queryClient.ListRecordsFromPool(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
