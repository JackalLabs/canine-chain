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
	"github.com/jackal-dao/canine/x/filetree/keeper"
	"github.com/jackal-dao/canine/x/filetree/types"
	filetypes "github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddViewers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-viewers [viewer-ids] [file path] [file owner]",
		Short: "add an address to the files viewing permissions",
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
			var viewerKeys []string
			var viewersToNotify []string

			for _, v := range viewerAddresses {
				if len(v) < 1 {
					continue
				}

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

				params := &types.QueryGetFilesRequest{
					Address:      merklePath,
					OwnerAddress: ownerChainAddress,
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


				ownerPrivateKey, err := MakePrivateKey(clientCtx)
				if err != nil {
					return err
				}

				decrypt, err := eciesgo.Decrypt(ownerPrivateKey, hexMessage)
				if err != nil {
					fmt.Printf("%v\n", hexMessage)
					return err
				}


				encrypted, err := eciesgo.Encrypt(pkey, []byte(decrypt))
				if err != nil {
					return err
				}

				newViewerID := keeper.MakeViewerAddress(file.Files.TrackingNumber, v)
				viewerIds = append(viewerIds, newViewerID)
				viewerKeys = append(viewerKeys, fmt.Sprintf("%x", encrypted))
				viewersToNotify = append(viewersToNotify, v)

			}

			jsonViewersToNotify, err := json.Marshal(viewersToNotify)
			if err != nil {
				return err
			}

			notiForViewers := fmt.Sprintf("%s has given you read access to %s", clientCtx.GetFromAddress().String(), argHashpath)

			msg := types.NewMsgAddViewers(
				clientCtx.GetFromAddress().String(),
				strings.Join(viewerIds, ","),
				strings.Join(viewerKeys, ","),
				merklePath,
				ownerChainAddress,
				string(jsonViewersToNotify),
				notiForViewers,
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
