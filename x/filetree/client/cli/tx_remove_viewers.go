package cli

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
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
			argOwner := args[2] //may be named to accountAddress

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fileQueryClient := types.NewQueryClient(clientCtx)
			trimPath := strings.TrimSuffix(argHashpath, "/")
			merklePath := types.MerklePath(trimPath)

			h := sha256.New()
			h.Write([]byte(argOwner))
			hash := h.Sum(nil)
			accountHash := fmt.Sprintf("%x", hash)
			ownerString := MakeOwnerAddress(merklePath, accountHash)

			viewerAddresses := strings.Split(argViewerIds, ",")

			var viewerIds []string
			var viewersToNotify []string

			for _, v := range viewerAddresses {
				if len(v) < 1 {
					continue
				}

				params := &types.QueryGetFilesRequest{
					Address:      merklePath,
					OwnerAddress: ownerString,
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
				ownerString,
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
