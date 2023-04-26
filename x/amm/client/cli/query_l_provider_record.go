package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"github.com/spf13/cobra"
)

func CmdShowProviderRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-provider-record [pool-id] [provider_addr]",
		Short: "shows a ProviderRecord",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argProviderAddr := args[1]

			poolId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil{
				return err
			}

			params := &types.QueryGetProviderRecordRequest{
				PoolId: poolId,
				Provider: argProviderAddr,
			}

			res, err := queryClient.ProviderRecord(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
