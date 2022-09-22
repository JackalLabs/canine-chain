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
	"github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdInitAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init-account [root-hashpath] [editors]",
		Short: "Broadcast message InitAccount",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRootHashpath := args[0]
			argEditors := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			newKey, err := MakePrivateKey(clientCtx)

			if err != nil {
				return err
			}

			pubKey := newKey.PublicKey.Bytes(false)

			merklePath := types.MerklePath(argRootHashpath)
			editors := make(map[string]string)
			editorAddresses := strings.Split(argEditors, ",")
			editorAddresses = append(editorAddresses, clientCtx.GetFromAddress().String())

			for _, v := range editorAddresses {
				if len(v) < 1 {
					continue
				}

				//This root folder is not supposed to hold the file's AES key, so there is nothing to encrypt. The purpose
				//Of the list of editors is to allow a user to invite others to write to their root folder. We are still using the JSON map to ensure that
				//The "is editor" function works the same here

				h := sha256.New()
				h.Write([]byte(fmt.Sprintf("e%s%s", merklePath, v))) //this used to be pathString
				hash := h.Sum(nil)

				addressString := fmt.Sprintf("%x", hash)

				editors[addressString] = fmt.Sprintf("%x", "NoKeyHere") //No need to store a key here
			}

			jsonEditors, err := json.Marshal(editors)
			if err != nil {
				return err
			}

			msgInitRoot := types.NewMsgInitAccount(
				clientCtx.GetFromAddress().String(),
				merklePath,
				string(jsonEditors),
				fmt.Sprintf("%x", pubKey),
			)

			if err := msgInitRoot.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgInitRoot)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
