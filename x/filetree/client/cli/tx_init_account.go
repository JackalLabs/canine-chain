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
	uuid "github.com/google/uuid"
	"github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

// This is a test tx that needs to be carefully removed soon
func CmdInitAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init-account [root-hashpath] [account] [editors]",
		Short: "Broadcast message InitAccount",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRootHashpath := args[0]
			argAccount := args[1]
			argEditors := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			newKey, err := MakePrivateKey(clientCtx)

			if err != nil {
				return err
			}

			pubKey := newKey.PublicKey.Bytes(false)
			//In the keeper, the merklePath function will trim the trailing slash for us but let's just do it anyways to be safe.
			trimMerklePath := strings.TrimSuffix(argRootHashpath, "/")
			merklePath := types.MerklePath(trimMerklePath)

			editors := make(map[string]string)
			editorAddresses := strings.Split(argEditors, ",")
			editorAddresses = append(editorAddresses, clientCtx.GetFromAddress().String())

			trackingNumber := uuid.NewString()

			for _, v := range editorAddresses {
				if len(v) < 1 {
					continue
				}

				//This root folder is not supposed to hold the file's AES key, so there is nothing to encrypt. The purpose
				//Of the list of editors is to allow a user to invite others to write to their root folder.

				h := sha256.New()
				h.Write([]byte(fmt.Sprintf("e%s%s", trackingNumber, v)))
				hash := h.Sum(nil)

				addressString := fmt.Sprintf("%x", hash)

				editors[addressString] = fmt.Sprintf("%x", "NoKeyHere") //No need to store a key here
			}

			jsonEditors, err := json.Marshal(editors)
			if err != nil {
				return err
			}

			h := sha256.New()
			h.Write([]byte(fmt.Sprintf("%s", argAccount)))
			hash := h.Sum(nil)

			accountHash := fmt.Sprintf("%x", hash)
			//FE will init their own root folders, but we are creating home/ in the CLI for visualizing the work flow
			msgInitRoot := types.NewMsgInitAccount(
				clientCtx.GetFromAddress().String(),
				accountHash,
				merklePath,
				string(jsonEditors),
				fmt.Sprintf("%x", pubKey),
				trackingNumber,
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
