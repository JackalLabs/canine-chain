package cli

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	uuid "github.com/google/uuid"
	"github.com/jackalLabs/canine-chain/v4/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMakeRootV2() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provision",
		Short: "Provision your file tree",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			editors := make(map[string]string)

			trackingNumber := uuid.NewString()

			// This root folder is the master root and has no file key, so there is nothing to encrypt.
			// We include the creator of this root as an editor so that they can add children--folders or files

			h := sha256.New()
			fmt.Fprintf(h, "e%s%s", trackingNumber, clientCtx.GetFromAddress().String())
			hash := h.Sum(nil)

			addressString := fmt.Sprintf("%x", hash)

			editors[addressString] = fmt.Sprintf("%x", "Placeholder key") // Determine if we need a place holder key

			jsonEditors, err := json.Marshal(editors)
			if err != nil {
				return err
			}

			msg := types.NewMsgProvisionFileTree(
				clientCtx.GetFromAddress().String(),
				string(jsonEditors),
				"Viewers",
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
