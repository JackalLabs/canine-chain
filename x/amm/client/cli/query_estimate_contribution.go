package cli

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEstimatePoolJoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-pool-join [pool-name] [pool-coins]",
		Short: "Estimate amount of share returned by adding liquidity",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			poolId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			liq, err := sdk.ParseCoinsNormalized(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryEstimatePoolJoinRequest{
				PoolId: poolId,
				Liquidity: liq,
			}

			res, err := queryClient.EstimatePoolJoin(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
