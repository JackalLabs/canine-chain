package cli

import (
	"crypto/sha256"
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

func CmdPostFile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post-file [path] [contents] [keys] [viewers] [editors]",
		Short: "post a new file to your file explorer",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argHashpath := args[0]
			argContents := args[1]
			argKeys := args[2]
			argViewers := args[3]
			argEditors := args[4]

			viewerAddresses := strings.Split(argViewers, ",")
			editorAddresses := strings.Split(argEditors, ",")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pathString := keeper.MakeAddress(argHashpath, clientCtx.FromAddress.String())

			viewers := make(map[string]string)
			editors := make(map[string]string)

			viewerAddresses = append(viewerAddresses, clientCtx.GetFromAddress().String())
			editorAddresses = append(editorAddresses, clientCtx.GetFromAddress().String())

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

				if acc.PubKey == nil {
					return fmt.Errorf("pub key not found")
				}

				err = pkey.Unmarshal(acc.PubKey.Value)
				if err != nil {
					return err
				}

				encrypted, err := clientCtx.Keyring.Encrypt(pkey.Key, []byte(argKeys))
				if err != nil {
					return err
				}

				h := sha256.New()
				h.Write([]byte(fmt.Sprintf("v%s%s", argHashpath, v)))
				hash := h.Sum(nil)

				addressString := fmt.Sprintf("%x", hash)

				viewers[addressString] = fmt.Sprintf("%x", encrypted)
			}

			for _, v := range editorAddresses {
				if len(v) < 1 {
					continue
				}
				fmt.Println(v)

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

				encrypted, err := clientCtx.Keyring.Encrypt(pkey.Key, []byte(argKeys))
				if err != nil {
					return err
				}

				h := sha256.New()
				h.Write([]byte(fmt.Sprintf("e%s%s", pathString, v)))
				hash := h.Sum(nil)

				addressString := fmt.Sprintf("%x", hash)

				editors[addressString] = fmt.Sprintf("%x", encrypted)
			}

			jsonViewers, err := json.Marshal(viewers)
			if err != nil {
				return err
			}

			jsonEditors, err := json.Marshal(editors)
			if err != nil {
				return err
			}

			msg := types.NewMsgPostFile(
				clientCtx.GetFromAddress().String(),
				pathString,
				argContents,
				string(jsonViewers),
				string(jsonEditors),
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
