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
			argHashpath := args[1]
			argOwner := args[2] //may be named to accountAddress

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fileQueryClient := types.NewQueryClient(clientCtx)
			trimPath := strings.TrimSuffix(argHashpath, "/")

			merklePath := types.MerklePath(trimPath)

			//Can't use helper functions in access.go so just build ownerString
			//Not working. Need to build the hash of the owner first.

			h := sha256.New()
			h.Write([]byte(fmt.Sprintf("%s", argOwner)))
			hash := h.Sum(nil)

			accountHash := fmt.Sprintf("%x", hash)

			H := sha256.New()
			H.Write([]byte(fmt.Sprintf("o%s%s", merklePath, accountHash)))
			Hash := H.Sum(nil)
			ownerString := fmt.Sprintf("%x", Hash)

			viewerAddresses := strings.Split(argViewerIds, ",")

			var viewerIds []string
			var viewerKeys []string
			var viewersToNotify []string

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

				json.Unmarshal([]byte(viewers), &m)

				ownerViewingAddress := keeper.MakeViewerAddress(file.Files.TrackingNumber, argOwner)

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
				encrypted, err := eciesgo.Encrypt(pkey, []byte(decrypt))
				if err != nil {
					return err
				}

				newViewerID := keeper.MakeViewerAddress(file.Files.TrackingNumber, v) //This used to just be argAddress
				viewerIds = append(viewerIds, newViewerID)
				viewerKeys = append(viewerKeys, fmt.Sprintf("%x", encrypted))
				viewersToNotify = append(viewersToNotify, v)

			}

			jsonViewersToNotify, err := json.Marshal(viewersToNotify)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddViewers(
				clientCtx.GetFromAddress().String(),
				strings.Join(viewerIds, ","),
				strings.Join(viewerKeys, ","),
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
