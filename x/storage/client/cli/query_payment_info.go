package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"github.com/spf13/cobra"
)

func CmdListStoragePaymentInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-storage-payment-info",
		Short: "list all storage_payment_info",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllStoragePaymentInfo{
				Pagination: pageReq,
			}

			res, err := queryClient.AllStoragePaymentInfo(context.Background(), params)
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

func CmdShowStoragePaymentInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-storage-payment-info [address]",
		Short: "shows a storage_payment_info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]

			params := &types.QueryStoragePaymentInfo{
				Address: argAddress,
			}

			res, err := queryClient.StoragePaymentInfo(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
