package cli

import (
	"context"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRemoveViewers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-viewers [viewer-ids] [file path] [file owner]",
		Short: "remove an address from the files viewing permissions",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argViewerIds := args[0]
			argHashpath := args[1]
			argOwner := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fileQueryClient := types.NewQueryClient(clientCtx)
			trimPath := strings.TrimSuffix(argHashpath, "/")
			merklePath := types.MerklePath(trimPath)
			ownerChainAddress := MakeOwnerAddress(merklePath, argOwner)

			viewerAddresses := strings.Split(argViewerIds, ",")
			var viewerIds []string

			for _, v := range viewerAddresses {
				if len(v) < 1 {
					continue
				}

				params := &types.QueryFileRequest{
					Address:      merklePath,
					OwnerAddress: ownerChainAddress,
				}

				file, err := fileQueryClient.Files(context.Background(), params)
				if err != nil {
					return types.ErrFileNotFound
				}

				newViewerID := keeper.MakeViewerAddress(file.Files.TrackingNumber, v) // This used to just be argAddress
				viewerIds = append(viewerIds, newViewerID)

			}

			// viewerIds supposed to be JSON marshalled aswell?
			msg := types.NewMsgRemoveViewers(
				clientCtx.GetFromAddress().String(),
				strings.Join(viewerIds, ","),
				merklePath,
				ownerChainAddress,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
