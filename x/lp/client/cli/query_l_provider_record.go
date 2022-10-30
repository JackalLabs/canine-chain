package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/x/lp/types"
	"github.com/spf13/cobra"
)

func CmdShowLProviderRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-l-provider-record [poolName] [provider_addr]",
		Short: "shows a LProviderRecord",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argPoolName := args[0]
			argProviderAddr := args[1]

			params := &types.QueryGetLProviderRecordRequest{
				PoolName: argPoolName,
				Provider: argProviderAddr,
			}

			res, err := queryClient.LProviderRecord(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
