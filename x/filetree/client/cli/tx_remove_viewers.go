package cli

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/filetree/keeper"
	"github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRemoveViewers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-viewers [viewer-ids] [file path] [file owner]",
		Short: "remove an address from the files viewing permisisons",
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
			var viewersToNotify []string

			for _, v := range viewerAddresses {
				if len(v) < 1 {
					continue
				}

				params := &types.QueryGetFilesRequest{
					Address:      merklePath,
					OwnerAddress: ownerChainAddress,
				}

				file, err := fileQueryClient.Files(context.Background(), params)
				if err != nil {
					return types.ErrFileNotFound
				}

				newViewerID := keeper.MakeViewerAddress(file.Files.TrackingNumber, v) //This used to just be argAddress
				viewerIds = append(viewerIds, newViewerID)
				viewersToNotify = append(viewersToNotify, v)

			}

			jsonViewersToNotify, err := json.Marshal(viewersToNotify)
			if err != nil {
				return err
			}

			//viewerIds supposed to be JSON marshalled aswell?
			msg := types.NewMsgRemoveViewers(
				clientCtx.GetFromAddress().String(),
				strings.Join(viewerIds, ","),
				merklePath,
				ownerChainAddress,
				string(jsonViewersToNotify),
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
