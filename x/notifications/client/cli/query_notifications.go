package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/v3/x/notifications/types"
	"github.com/spf13/cobra"
)

func CmdListNotifications() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-notifications",
		Short: "list all notifications",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllNotifications{
				Pagination: pageReq,
			}

			res, err := queryClient.AllNotifications(context.Background(), params)
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

func CmdListNotificationsByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-notifications-by-address [address]",
		Short: "list all notifications for an address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllNotificationsByAddress{
				Pagination: pageReq,
				To:         args[0],
			}

			res, err := queryClient.AllNotificationsByAddress(context.Background(), params)
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

func CmdShowNotification() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-notification [to] [from] [unix-timestamp]",
		Short: "shows a notifications",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			ts, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryNotification{
				To:   args[0],
				From: args[1],
				Time: ts,
			}

			res, err := queryClient.Notification(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
