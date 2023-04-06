package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"github.com/spf13/cobra"
)

func CmdShowProviderRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-provider-record [poolName] [provider_addr]",
		Short: "shows a LProviderRecord",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argPoolName := args[0]
			argProviderAddr := args[1]

			params := &types.QueryGetProviderRecordRequest{
				PoolName: argPoolName,
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
