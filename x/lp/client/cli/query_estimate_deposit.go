package cli

import (
	"strconv"

	"github.com/jackal-dao/canine/x/lp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEstimateSwapIn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-swap-in [pool-name] [desired-coin]",
		Short: "Estimate coin input to get desired output from a swap. Fees are" +
		" not considered.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPoolName := args[0]
			reqDesiredCoin := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryEstimateSwapInRequest{
				PoolName:     reqPoolName,
				OutputCoins: reqDesiredCoin,
			}

			res, err := queryClient.EstimateSwapIn(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
