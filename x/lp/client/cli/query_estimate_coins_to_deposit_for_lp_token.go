package cli

import (
	"strconv"

	"github.com/jackal-dao/canine/x/lp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEstimateContribution() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-contribution [pool-name] [desired-amount]",
		Short: "Estimate coins to contribute to get desired amount of LPToken",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPoolName := args[0]
			reqDesiredAmount := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryEstimateContributionRequest{
				PoolName:      reqPoolName,
				DesiredAmount: reqDesiredAmount,
			}

			res, err := queryClient.EstimateContribution(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
