package cli

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	uuid "github.com/google/uuid"
	filetypes "github.com/jackalLabs/canine-chain/v3/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPostFilePublic() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post-file-no-key [path] [account] [contents]",
		Short: "post a new file to an account's file explorer",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argHashpath := args[0]
			argAccount := args[1]
			argContents := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			parentHash, childHash := merkleHelper(argHashpath)

			trackingNumber := uuid.NewString()

			viewers := make(map[string]string)
			editors := make(map[string]string)

			// Marshall everybody
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
