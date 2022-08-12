package cli

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddViewers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-viewers [viewer-ids] [key] [file path] [file owner]",
		Short: "add an address to the files viewing permisisons",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argViewerIds := args[0]
			argViewerKeys := args[1]
			argAddress := args[2]
			argOwner := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			h := sha256.New()
			h.Write([]byte(argAddress))
			hash := h.Sum(nil)

			pathString := fmt.Sprintf("%x", hash)

			h = sha256.New()
			h.Write([]byte(fmt.Sprintf("%s%s", argOwner, pathString)))
			hash = h.Sum(nil)
			pathString = fmt.Sprintf("%x", hash)

			viewerAddresses := strings.Split(argViewerIds, ",")

			var viewerIds []string
			var viewerKeys []string

			for _, v := range viewerAddresses {
				if len(v) < 1 {
					continue
				}
				key, err := sdk.AccAddressFromBech32(v)
				if err != nil {
					return err
				}

				queryClient := authtypes.NewQueryClient(clientCtx)
				res, err := queryClient.Account(cmd.Context(), &authtypes.QueryAccountRequest{Address: key.String()})
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

				encrypted, err := clientCtx.Keyring.Encrypt(pkey.Key, []byte(argViewerKeys))
				if err != nil {
					return err
				}

				h = sha256.New()
				h.Write([]byte(fmt.Sprintf("v%s%s", argAddress, v)))
				hash = h.Sum(nil)

				addressString := fmt.Sprintf("%x", hash)

				viewerIds = append(viewerIds, addressString)
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
