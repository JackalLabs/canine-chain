package cli

import (
	"context"
	"encoding/hex"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	"github.com/spf13/cobra"
)

func CmdListAttestForms() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-attest-forms",
		Short: "list attest forms",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAttestations{
				Pagination: pageReq,
			}

			res, err := queryClient.AllAttestations(context.Background(), params)
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

func CmdShowAttestForms() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attest-form [prover] [merkle] [owner] [start]",
		Short: "shows an attest form",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argProver := args[0]
			argMerkle := args[1]
			argOwner := args[2]
			argStart := args[3]

			start, err := strconv.ParseInt(argStart, 10, 64)
			if err != nil {
				panic(err)
			}
			merkle, err := hex.DecodeString(argMerkle)
			if err != nil {
				panic(err)
			}

			params := &types.QueryAttestation{
				Prover: argProver,
				Merkle: merkle,
				Owner:  argOwner,
				Start:  start,
			}

			res, err := queryClient.Attestation(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListReportForms() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-report-forms",
		Short: "list report forms",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllReports{
				Pagination: pageReq,
			}

			res, err := queryClient.AllReports(context.Background(), params)
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

func CmdShowReportForms() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report-form [prover] [merkle] [owner] [start]",
		Short: "shows an report form",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argProver := args[0]
			argMerkle := args[1]
			argOwner := args[2]
			argStart := args[3]

			start, err := strconv.ParseInt(argStart, 10, 64)
			if err != nil {
				panic(err)
			}
			merkle, err := hex.DecodeString(argMerkle)
			if err != nil {
				panic(err)
			}

			params := &types.QueryReport{
				Prover: argProver,
				Merkle: merkle,
				Owner:  argOwner,
				Start:  start,
			}

			res, err := queryClient.Report(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
