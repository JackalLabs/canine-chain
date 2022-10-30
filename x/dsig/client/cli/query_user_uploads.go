package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/x/dsig/types"
	"github.com/spf13/cobra"
)

func CmdListUserUploads() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-user-uploads",
		Short: "list all UserUploads",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllUserUploadsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.UserUploadsAll(context.Background(), params)
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

func CmdShowUserUploads() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-user-uploads [fid]",
		Short: "shows a UserUploads",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argFid := args[0]

			params := &types.QueryGetUserUploadsRequest{
				Fid: argFid,
			}

			res, err := queryClient.UserUploads(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
