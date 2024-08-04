package cli

import (
	"context"
	"encoding/hex"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	"github.com/spf13/cobra"
)

func CmdListActiveDeals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "files",
		Short: "list all files",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllFiles{
				Pagination: pageReq,
			}

			res, err := queryClient.AllFiles(context.Background(), params)
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

func CmdShowActiveDeals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "file [merkle] [owner] [start]",
		Short: "shows a file",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argMerkle := args[0]
			argOwner := args[1]
			argStart := args[2]

			start, err := strconv.ParseInt(argStart, 10, 64)
			if err != nil {
				panic(err)
			}
			merkle, err := hex.DecodeString(argMerkle)
			if err != nil {
				panic(err)
			}

			params := &types.QueryFile{
				Merkle: merkle,
				Start:  start,
				Owner:  argOwner,
			}

			res, err := queryClient.File(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proof [prover] [owner] [merkle] [start]",
		Short: "shows a proof",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			argProver := args[0]
			argMerkle := args[2]
			argOwner := args[1]
			argStart := args[3]

			start, err := strconv.ParseInt(argStart, 10, 64)
			if err != nil {
				panic(err)
			}
			merkle, err := hex.DecodeString(argMerkle)
			if err != nil {
				panic(err)
			}

			params := &types.QueryProof{
				ProviderAddress: argProver,
				Merkle:          merkle,
				Start:           start,
				Owner:           argOwner,
			}

			res, err := queryClient.Proof(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
