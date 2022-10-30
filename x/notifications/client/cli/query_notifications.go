package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListNotifications() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-notifications",
		Short: "list all notifications",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllNotificationsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.NotificationsAll(context.Background(), params)
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

func CmdShowNotifications() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-notifications [count]",
		Short: "shows a notifications",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argCount, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetNotificationsRequest{
				Count: argCount,
			}

			res, err := queryClient.Notifications(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
