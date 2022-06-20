package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackal-dao/canine/x/jklmining/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetMinerIndex() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-miner-index [index]",
		Short: "Query get_miner_index",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqIndex := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetMinerIndexRequest{

				Index: reqIndex,
			}

			res, err := queryClient.GetMinerIndex(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
