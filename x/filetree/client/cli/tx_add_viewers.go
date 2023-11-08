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
	"github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
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

			for _, v := range viewerAddresses {
				if len(v) < 1 {
					continue
				}
				key, err := sdk.AccAddressFromBech32(v) // I think this isn't needed
				if err != nil {
					return err
				}

				queryClient := types.NewQueryClient(clientCtx)
				res, err := queryClient.PubKey(cmd.Context(), &types.QueryPubKeyRequest{Address: key.String()})
				if err != nil {
					return types.ErrPubKeyNotFound
				}

				pkey, err := eciesgo.NewPublicKeyFromHex(res.PubKey.Key)
				if err != nil {
					return err
				}
				// Perhaps below file query should be replaced with fully fledged 'query file' function that checks permissions first
				params := &types.QueryFileRequest{
					Address:      merklePath,
					OwnerAddress: ownerChainAddress,
				}

				file, err := fileQueryClient.File(context.Background(), params)
				if err != nil {
					return types.ErrFileNotFound
				}

				viewers := file.File.ViewingAccess
				var m map[string]string

				err = json.Unmarshal([]byte(viewers), &m)
				if err != nil {
					return types.ErrCantUnmarshall
				}

				ownerViewingAddress := keeper.MakeViewerAddress(file.File.TrackingNumber, argOwner)

				hexMessage, err := hex.DecodeString(m[ownerViewingAddress])
				if err != nil {
					return err
				}

				// May need to use "clientCtx.from?"
				ownerPrivateKey, err := MakePrivateKey(clientCtx)
				if err != nil {
					return err
				}

				decrypt, err := eciesgo.Decrypt(ownerPrivateKey, hexMessage)
				if err != nil {
					fmt.Printf("%v\n", hexMessage)
					return err
				}

				// encrypt using viewer's public key
				encrypted, err := eciesgo.Encrypt(pkey, decrypt)
				if err != nil {
					return err
				}

				newViewerID := keeper.MakeViewerAddress(file.File.TrackingNumber, v)
				viewerIds = append(viewerIds, newViewerID)
				viewerKeys = append(viewerKeys, fmt.Sprintf("%x", encrypted))

			}

			msg := types.NewMsgAddViewers(
				clientCtx.GetFromAddress().String(),
				strings.Join(viewerIds, ","),
				strings.Join(viewerKeys, ","),
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
