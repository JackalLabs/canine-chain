package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackal-dao/canine/x/storage/types"
	"github.com/spf13/cobra"
)

func CmdListProofs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-proofs",
		Short: "list all proofs",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllProofsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ProofsAll(context.Background(), params)
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

func CmdShowProofs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-proofs [cid]",
		Short: "shows a proofs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argCid := args[0]

			params := &types.QueryGetProofsRequest{
				Cid: argCid,
			}

			res, err := queryClient.Proofs(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
