package cli

import (
	"context"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRemoveEditors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-editors [editor-ids] [file path] [file owner]",
		Short: "remove an address from the files editing permissions",
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

			for _, v := range editorAddresses {
				if len(v) < 1 {
					continue
				}

				params := &types.QueryFileRequest{
					Address:      merklePath,
					OwnerAddress: ownerChainAddress,
				}

				file, err := fileQueryClient.File(context.Background(), params)
				if err != nil {
					return types.ErrFileNotFound
				}

				newEditorID := keeper.MakeEditorAddress(file.File.TrackingNumber, v) // This used to just be argAddress
				editorIds = append(editorIds, newEditorID)

			}

			msg := types.NewMsgRemoveEditors(
				clientCtx.GetFromAddress().String(),
				strings.Join(editorIds, ","),
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
