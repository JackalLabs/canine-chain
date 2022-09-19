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
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/jackal-dao/canine/x/filetree/keeper"
	"github.com/jackal-dao/canine/x/filetree/types"
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
			argOwner := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			authQueryClient := authtypes.NewQueryClient(clientCtx)
			fileQueryClient := types.NewQueryClient(clientCtx)

			pathString := keeper.MakeChainAddress(argAddress, argOwner)

			fmt.Println(pathString)

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

				res, err := authQueryClient.Account(cmd.Context(), &authtypes.QueryAccountRequest{Address: key.String()})
				if err != nil {
					return err
				}

				var acc authtypes.BaseAccount

				err = acc.Unmarshal(res.Account.Value)
				if err != nil {
					return err
				}
				var pkey secp256k1.PubKey

				err = pkey.Unmarshal(acc.PubKey.Value)
				if err != nil {
					return err
				}

				params := &types.QueryGetFilesRequest{
					Address: pathString,
				}

				file, err := fileQueryClient.Files(context.Background(), params)
				if err != nil {
					return err
				}

				viewers := file.Files.ViewingAccess
				var m map[string]string

				json.Unmarshal([]byte(viewers), &m)

				aString := keeper.MakeViewerAddress(argAddress, clientCtx.GetFromAddress().String())

				hexMessage, err := hex.DecodeString(m[aString])
				if err != nil {
					return err
				}

				from := clientCtx.From

				decrypt, _, err := clientCtx.Keyring.Decrypt(from, hexMessage)
				if err != nil {
					fmt.Println("cannot decrypt keys")
					return err
				}

				encrypted, err := clientCtx.Keyring.Encrypt(pkey.Key, []byte(decrypt))
				if err != nil {
					return err
				}
				newViewerID := keeper.MakeViewerAddress(argAddress, v)
				viewerIds = append(viewerIds, newViewerID)
				viewerKeys = append(viewerKeys, fmt.Sprintf("%x", encrypted))
			}

			msg := types.NewMsgAddViewers(
				clientCtx.GetFromAddress().String(),
				strings.Join(viewerIds, ","),
				strings.Join(viewerKeys, ","),
				pathString,
				argOwner,
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
