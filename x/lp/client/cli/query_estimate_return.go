package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackal-dao/canine/x/lp/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEstimateSwapOut() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-swap-out [pool-name] [input-coins]",
		Short: "Estimate coin output from a swap. Fees are not considered.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPoolName := args[0]
			reqDepositCoins := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryEstimateSwapOutRequest{
				PoolName:  reqPoolName,
				InputCoin: reqDepositCoins,
			}

			res, err := queryClient.EstimateSwapOut(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
