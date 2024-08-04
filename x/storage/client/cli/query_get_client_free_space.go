package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetClientFreeSpace() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-client-free-space [address]",
		Short: "Queries the amount of storage a user has available",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryClientFreeSpace{
				Address: reqAddress,
			}

			res, err := queryClient.GetClientFreeSpace(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
