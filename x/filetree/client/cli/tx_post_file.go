package cli

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	uuid "github.com/google/uuid"
	filetypes "github.com/jackalLabs/canine-chain/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPostFile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post-file [path] [account] [contents] [keys] [viewers] [editors]",
		Short: "post a new file to an account's file explorer",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argHashpath := args[0]
			argAccount := args[1]
			argContents := args[2]
			argKeys := args[3]
			argViewers := args[4]
			argEditors := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fromAddress, err := getCallerAddress(clientCtx, cmd)
			if err != nil {
				return err
			}

			if argKeys == "" {
				return filetypes.ErrMissingAESKey
			}

			viewerAddresses := strings.Split(argViewers, ",")
			editorAddresses := strings.Split(argEditors, ",")

			parentHash, childHash := merkleHelper(argHashpath)

			trackingNumber := uuid.NewString()

			viewers := make(map[string]string)
			editors := make(map[string]string)

			viewerAddresses = append(viewerAddresses, *fromAddress)
			editorAddresses = append(editorAddresses, *fromAddress)

			for _, v := range viewerAddresses {
				if len(v) < 1 {
					continue
				}

				encrypted, err := encryptFileAESKey(cmd, v, argKeys)
				if err != nil {
					return err
				}

				h := sha256.New()
				h.Write([]byte(fmt.Sprintf("v%s%s", trackingNumber, v)))
				hash := h.Sum(nil)
				addressString := fmt.Sprintf("%x", hash)

				viewers[addressString] = fmt.Sprintf("%x", encrypted)
			}

			for _, v := range editorAddresses {
				if len(v) < 1 {
					continue
				}

				encrypted, err := encryptFileAESKey(cmd, v, argKeys)
				if err != nil {
					return err
				}

				h := sha256.New()
				h.Write([]byte(fmt.Sprintf("e%s%s", trackingNumber, v)))
				hash := h.Sum(nil)
				addressString := fmt.Sprintf("%x", hash)

				editors[addressString] = fmt.Sprintf("%x", encrypted)
			}

			//Marshall everybody
			jsonViewers, jsonEditors, err := JSONMarshalViewersAndEditors(viewers, editors)
			if err != nil {
				return err
			}
			H := sha256.New()
			H.Write([]byte(argAccount))
			hash := H.Sum(nil)
			accountHash := fmt.Sprintf("%x", hash)

			msg := filetypes.NewMsgPostFile(
				clientCtx.GetFromAddress().String(),
				accountHash,
				parentHash,
				childHash,
				argContents,
				string(jsonViewers),
				string(jsonEditors),
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
