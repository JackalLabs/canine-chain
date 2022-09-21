package cli

import (
	"context"
	"crypto/sha256"
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
	"github.com/jackal-dao/canine/x/filetree/keeper"
	"github.com/jackal-dao/canine/x/filetree/types"
	filetypes "github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddViewers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-viewers [viewer-ids] [file path] [file owner]",
		Short: "add an address to the files viewing permisisons",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argViewerIds := args[0]
			argAddress := args[1]
			argOwner := args[2] //may be named to accountAddress

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fileQueryClient := types.NewQueryClient(clientCtx)

			merklePath := types.MerklePath(argAddress)
			//In next commit, this should all be packed into a fully fledged file query that does permission checking
			h := sha256.New()
			h.Write([]byte(fmt.Sprintf("o%s%s", merklePath, argOwner)))
			hash := h.Sum(nil)
			ownerString := fmt.Sprintf("%x", hash)

			viewerAddresses := strings.Split(argViewerIds, ",")

			var viewerIds []string
			var viewerKeys []string

			for _, v := range viewerAddresses {
				if len(v) < 1 {
					continue
				}
				key, err := sdk.AccAddressFromBech32(v)
				if err != nil {
					fmt.Printf("address: %s\n", v)
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
					OwnerAddress: ownerString,
				}

				file, err := fileQueryClient.Files(context.Background(), params)
				if err != nil {
					return types.ErrFileNotFound
				}

				viewers := file.Files.ViewingAccess
				var m map[string]string

				json.Unmarshal([]byte(viewers), &m) //Unmarshal file's viewing access array into address of this newly declared map 'm'

				ownerViewingAddress := keeper.MakeViewerAddress(argAddress, argOwner)

				hexMessage, err := hex.DecodeString(m[ownerViewingAddress])
				if err != nil {
					return err
				}

				//May need to use "clientCtx.from?"
				ownerPrivateKey, err := MakePrivateKey(clientCtx)
				if err != nil {
					return err
				}

				decrypt, err := eciesgo.Decrypt(ownerPrivateKey, hexMessage)
				if err != nil {
					fmt.Printf("%v\n", hexMessage)
					return err
				}

				//encrypt using viewer's public key
				encrypted, err := clientCtx.Keyring.Encrypt(pkey.Bytes(false), []byte(decrypt))
				if err != nil {
					return err
				}

				newViewerID := keeper.MakeViewerAddress(argAddress, v)
				viewerIds = append(viewerIds, newViewerID)
				viewerKeys = append(viewerKeys, fmt.Sprintf("%x", encrypted))

			}

			msg := types.NewMsgAddViewers(
				clientCtx.GetFromAddress().String(), //msg caller who wants to add viewers
				strings.Join(viewerIds, ","),
				strings.Join(viewerKeys, ","),
				merklePath,
				ownerString,
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
