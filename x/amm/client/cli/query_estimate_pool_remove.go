package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEstimatePoolRemove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-pool-remove [pool-name] [amount]",
		Short: "Estimate the amount of pool shares that will be removed from the pool when removing liquidity",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPoolName := args[0]
			reqAmount := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryEstimatePoolRemoveRequest{
				PoolName: reqPoolName,
				Amount:   reqAmount,
			}

			res, err := queryClient.EstimatePoolRemove(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
