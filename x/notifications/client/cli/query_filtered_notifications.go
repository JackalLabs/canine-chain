package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackal-dao/canine/x/notifications/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

// Get all notifications that belong to a user
func CmdFilteredNotifications() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "filtered-notifications [address]",
		Short: "Query filteredNotifications",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			address := args[0]

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryFilteredNotificationsRequest{
				Address: address,
			}

			res, err := queryClient.FilteredNotifications(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
