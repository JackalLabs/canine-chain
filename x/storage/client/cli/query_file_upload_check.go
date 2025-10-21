package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/jackalLabs/canine-chain/v5/x/storage/types"
	"github.com/spf13/cobra"
)

func CmdFileUploadCheck() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "file-upload-check [address] [bytes]",
		Short: "check user can upload file based on size",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			argAddr := args[0]
			argByte := args[1]

			reqByte, err := strconv.ParseInt(argByte, 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryFileUploadCheck{
				Address: argAddr,
				Bytes:   reqByte,
			}

			res, err := queryClient.FileUploadCheck(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
