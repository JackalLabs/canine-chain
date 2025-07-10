package cli

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	eciesgo "github.com/ecies/go/v2"
	testUtil "github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/v4/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddEditors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-editors [editor-ids] [file path] [file owner]",
		Short: "add an address to the files editing permissions",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEditorIDs := args[0]
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

			editorAddresses := strings.Split(argEditorIDs, ",")

			var editorIDs []string
			var editorKeys []string
			logger, logFile := testUtil.CreateLogger()
			for _, v := range editorAddresses {
				if len(v) < 1 {
					continue
				}
				key, err := sdk.AccAddressFromBech32(v) // I think this isn't needed
				if err != nil {
					return err
				}

				queryClient := types.NewQueryClient(clientCtx)
				res, err := queryClient.PubKey(cmd.Context(), &types.QueryPubKey{Address: key.String()})
				if err != nil {
					return types.ErrPubKeyNotFound
				}

				pkey, err := eciesgo.NewPublicKeyFromHex(res.PubKey.Key)
				if err != nil {
					return err
				}
				// Perhaps below file query should be replaced with fully fledged 'query file' function that checks permissions first
				params := &types.QueryFile{
					Address:      merklePath,
					OwnerAddress: ownerChainAddress,
				}

				file, err := fileQueryClient.File(context.Background(), params)
				if err != nil {
					return types.ErrFileNotFound
				}

				editors := file.File.EditAccess
				var m map[string]string

				err = json.Unmarshal([]byte(editors), &m)
				if err != nil {
					return types.ErrCantUnmarshall
				}

				ownerEditorAddress := keeper.MakeEditorAddress(file.File.TrackingNumber, argOwner)
				logger.Println("m[ownerEditorAddress] =", m[ownerEditorAddress])

				hexMessage, err := hex.DecodeString(m[ownerEditorAddress])
				if err != nil {
					return err
				}
				logger.Println("hex message is", string(hexMessage))

				// May need to use "clientCtx.from?"
				ownerPrivateKey, err := MakePrivateKey(clientCtx)
				if err != nil {
					return err
				}

				decrypt, err := eciesgo.Decrypt(ownerPrivateKey, hexMessage)
				if err != nil {
					fmt.Printf("%v\n", hexMessage)
					logger.Println("error is", err)
					return err
				}
				logFile.Close()
				// encrypt using editor's public key
				encrypted, err := eciesgo.Encrypt(pkey, decrypt)
				if err != nil {
					return err
				}

				newEditorID := keeper.MakeEditorAddress(file.File.TrackingNumber, v)
				editorIDs = append(editorIDs, newEditorID)
				editorKeys = append(editorKeys, fmt.Sprintf("%x", encrypted))

			}

			msg := types.NewMsgAddEditors(
				clientCtx.GetFromAddress().String(),
				strings.Join(editorIDs, ","),
				strings.Join(editorKeys, ","),
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
