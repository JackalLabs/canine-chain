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
	testUtil "github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
	filetypes "github.com/jackalLabs/canine-chain/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddEditors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-editors [editor-ids] [file path] [file owner]",
		Short: "add an address to the files editing permisisons",
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
			var editorKeys []string
			logger, logFile := testUtil.CreateLogger()
			for _, v := range editorAddresses {
				if len(v) < 1 {
					continue
				}
				key, err := sdk.AccAddressFromBech32(v) //I think this isn't needed
				if err != nil {
					return err
				}

				queryClient := filetypes.NewQueryClient(clientCtx)
				res, err := queryClient.Pubkey(cmd.Context(), &filetypes.QueryGetPubkeyRequest{Address: key.String()})
				if err != nil {
					return types.ErrPubKeyNotFound
				}

				pkey, err := eciesgo.NewPublicKeyFromHex(res.Pubkey.Key)
				if err != nil {
					return err
				}
				//Perhaps below file query should be replaced with fully fledged 'query file' function that checks permissions first
				params := &types.QueryGetFilesRequest{
					Address:      merklePath,
					OwnerAddress: ownerChainAddress,
				}

				file, err := fileQueryClient.Files(context.Background(), params)
				if err != nil {
					return types.ErrFileNotFound
				}

				editors := file.Files.EditAccess
				var m map[string]string

				json.Unmarshal([]byte(editors), &m)

				ownerEditorAddress := keeper.MakeEditorAddress(file.Files.TrackingNumber, argOwner)
				logger.Println("m[ownerEditorAddress] =", m[ownerEditorAddress])

				hexMessage, err := hex.DecodeString(m[ownerEditorAddress])
				if err != nil {
					return err
				}
				logger.Println("hex message is", hexMessage)

				//May need to use "clientCtx.from?"
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
				//encrypt using editor's public key
				encrypted, err := eciesgo.Encrypt(pkey, []byte(decrypt))
				if err != nil {
					return err
				}

				newEditorID := keeper.MakeEditorAddress(file.Files.TrackingNumber, v)
				editorIds = append(editorIds, newEditorID)
				editorKeys = append(editorKeys, fmt.Sprintf("%x", encrypted))

			}

			msg := types.NewMsgAddEditors(
				clientCtx.GetFromAddress().String(),
				strings.Join(editorIds, ","),
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
