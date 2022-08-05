package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackal-dao/canine/x/dsig/types"
	"github.com/spf13/cobra"
)

func CmdListForm() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-form",
		Short: "list all Form",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllFormRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.FormAll(context.Background(), params)
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

func CmdShowForm() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-form [ffid]",
		Short: "shows a Form",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argFfid := args[0]

			params := &types.QueryGetFormRequest{
				Ffid: argFfid,
			}

			res, err := queryClient.Form(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
