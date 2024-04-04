package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	"github.com/spf13/cobra"
)

func CmdListProviders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-providers",
		Short: "list all providers",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllProviders{
				Pagination: pageReq,
			}

			res, err := queryClient.AllProviders(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListActiveProviders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-active-providers",
		Short: "list all active providers",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryActiveProviders{}
			res, err := queryClient.ActiveProviders(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowProviders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-providers [address]",
		Short: "shows a providers",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]

			params := &types.QueryProvider{
				Address: argAddress,
			}

			res, err := queryClient.Provider(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
