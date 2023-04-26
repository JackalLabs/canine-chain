package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdListRecordsFromPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-records-from-pool [pool-id]",
		Short: "List all ProviderRecords of the pool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			poolId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil{
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListRecordsFromPoolRequest{

				PoolId: poolId,
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
