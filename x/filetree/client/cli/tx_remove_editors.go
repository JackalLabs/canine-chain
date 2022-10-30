package cli

import (
	"context"
	"encoding/json"
	"fmt"
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

func CmdRemoveEditors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-editors [editor-ids] [file path] [file owner]",
		Short: "remove an address from the files editing permisisons",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEditorIds := args[0]
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

			editorAddresses := strings.Split(argEditorIds, ",")
			var editorIds []string
			var editorsToNotify []string

			for _, v := range editorAddresses {
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

				newEditorID := keeper.MakeEditorAddress(file.Files.TrackingNumber, v) // This used to just be argAddress
				editorIds = append(editorIds, newEditorID)
				editorsToNotify = append(editorsToNotify, v)

			}

			jsonEditorsToNotify, err := json.Marshal(editorsToNotify)
			if err != nil {
				return err
			}

			notiForEditors := fmt.Sprintf("%s has removed edit access to %s", clientCtx.GetFromAddress().String(), argHashpath)

			msg := types.NewMsgRemoveEditors(
				clientCtx.GetFromAddress().String(),
				strings.Join(editorIds, ","),
				merklePath,
				ownerChainAddress,
				string(jsonEditorsToNotify),
				notiForEditors,
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
