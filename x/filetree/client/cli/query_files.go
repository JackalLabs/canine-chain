package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cobra"
)

func CmdListFiles() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-files",
		Short: "list all files",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllFilesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.FilesAll(context.Background(), params)
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

// Need to input the full hex formatted merkleHash address for this to work
func CmdShowFiles() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-files [address]",
		Short: "shows a files",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]

			params := &types.QueryGetFilesRequest{
				Address: argAddress,
			}

			res, err := queryClient.Files(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// Query using human readable path to show that merklePath() function works as intended
func CmdShowFileFromPath() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-file-from-path [path]",
		Short: "shows a file given human readable path",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]
			merklePath := types.MerklePath(argAddress)

			params := &types.QueryGetFilesRequest{
				Address: merklePath,
			}

			res, err := queryClient.Files(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
