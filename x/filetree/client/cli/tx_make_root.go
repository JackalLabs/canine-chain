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
	"github.com/jackalLabs/canine-chain/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMakeRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make-root [account] [root-hash-path] [contents] [editors] [viewers]",
		Short: "Broadcast message makeFolder",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAccount := args[0]
			argRootHashPath := args[1]
			argContents := args[2]
			argEditors := args[3]
			argViewers := args[4]

			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// In the keeper, the merklePath function will trim the trailing slash for us but let's just do it anyways to be safe.
			trimMerklePath := strings.TrimSuffix(argRootHashPath, "/")
			merklePath := types.MerklePath(trimMerklePath)

			editors := make(map[string]string)
			editorAddresses := strings.Split(argEditors, ",")
			editorAddresses = append(editorAddresses, clientCtx.GetFromAddress().String())

			trackingNumber := uuid.NewString()

			for _, v := range editorAddresses {
				if len(v) < 1 {
					continue
				}

				// This root folder is the master root and has no file key, so there is nothing to encrypt. The purpose
				// Of the list of editors is to allow a user to invite others to write to their root folder.

				h := sha256.New()
				h.Write([]byte(fmt.Sprintf("e%s%s", trackingNumber, v)))
				hash := h.Sum(nil)

				addressString := fmt.Sprintf("%x", hash)

				editors[addressString] = fmt.Sprintf("%x", "NoKeyHere") // No need to store a key here
			}

			jsonEditors, err := json.Marshal(editors)
			if err != nil {
				return err
			}

			h := sha256.New()
			h.Write([]byte(argAccount))
			hash := h.Sum(nil)

			accountHash := fmt.Sprintf("%x", hash)
			// FE will init their own root folders

			msg := types.NewMsgMakeRoot(
				clientCtx.GetFromAddress().String(),
				accountHash,
				merklePath,
				argContents,
				string(jsonEditors),
				argViewers,
				trackingNumber,
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
