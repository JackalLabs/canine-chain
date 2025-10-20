package cli

import (
	"context"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v4/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/v4/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRemoveViewers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-viewers [viewer-ids] [file path] [file owner]",
		Short: "remove an address from the files viewing permissions",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argViewerIDs := args[0]
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

			viewerAddresses := strings.Split(argViewerIDs, ",")
			var viewerIDs []string

			for _, v := range viewerAddresses {
				if len(v) < 1 {
					continue
				}

				params := &types.QueryFile{
					Address:      merklePath,
					OwnerAddress: ownerChainAddress,
				}

				file, err := fileQueryClient.File(context.Background(), params)
				if err != nil {
					return types.ErrFileNotFound
				}

				newViewerID := keeper.MakeViewerAddress(file.File.TrackingNumber, v) // This used to just be argAddress
				viewerIDs = append(viewerIDs, newViewerID)

			}

			// viewerIDs supposed to be JSON marshalled as well?
			msg := types.NewMsgRemoveViewers(
				clientCtx.GetFromAddress().String(),
				strings.Join(viewerIDs, ","),
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
