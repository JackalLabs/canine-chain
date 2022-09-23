package cli

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"

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

// Need to input the full hex formatted merkleHash address and the ownerAddress for this to work
func CmdShowFiles() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-files [address] [ownerAddress]",
		Short: "shows a files",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]
			argOwnerAddress := args[1]

			params := &types.QueryGetFilesRequest{
				Address:      argAddress,
				OwnerAddress: argOwnerAddress,
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
// Should really be called "CmdShowFileFromPathAndOwner", but I opted to not make the name so long given that this is just a testing function
func CmdShowFileFromPath() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-file-from-path [path] [owner]",
		Short: "shows a file given human readable path and owner's bech32 address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]
			argOwnerAddress := args[1]
			trimMerklePath := strings.TrimSuffix(argAddress, "/")
			merklePath := types.MerklePath(trimMerklePath)
			fmt.Println("The merklePath is", merklePath)
			//hash the owner address alone
			h := sha256.New()
			h.Write([]byte(fmt.Sprintf("%s", argOwnerAddress)))
			hash := h.Sum(nil)
			accountHash := fmt.Sprintf("%x", hash)

			fmt.Println("accountHash is", accountHash)

			//make the full OwnerAddress
			H := sha256.New()
			H.Write([]byte(fmt.Sprintf("o%s%s", merklePath, accountHash)))
			Hash := H.Sum(nil)
			ownerAddress := fmt.Sprintf("%x", Hash)

			params := &types.QueryGetFilesRequest{
				Address:      merklePath,
				OwnerAddress: ownerAddress,
			}
			fmt.Println("The owner bech32 is", argOwnerAddress)
			fmt.Println("owner Address is", ownerAddress)
			fmt.Println("path Address is", merklePath)
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
