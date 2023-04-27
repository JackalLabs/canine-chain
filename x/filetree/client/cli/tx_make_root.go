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

/*
Deprecated: CMDMakeRoot is being replaced by CMDMakeRootV2
*/
func CmdMakeRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make-root [account] [root-hash-path]",
		Short: "Broadcast message make root",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			fmt.Println("make-root is deprecated as of v2.0.0, please consider using make-root-v2.")
			argAccount := args[0]
			argRootHashPath := args[1]

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

			trackingNumber := uuid.NewString()

			// This root folder is the master root and has no file key, so there is nothing to encrypt.
			// We include the creator of this root as an editor so that they can add children--folders or files

			h := sha256.New()
			h.Write([]byte(fmt.Sprintf("e%s%s", trackingNumber, clientCtx.GetFromAddress().String())))
			hash := h.Sum(nil)

			addressString := fmt.Sprintf("%x", hash)

			editors[addressString] = fmt.Sprintf("%x", "Placeholder key") // Determine if we need a place holder key

			jsonEditors, err := json.Marshal(editors)
			if err != nil {
				return err
			}

			h1 := sha256.New()
			h1.Write([]byte(argAccount))
			hash1 := h1.Sum(nil)

			accountHash := fmt.Sprintf("%x", hash1)
			// FE will init their own root folders

			msg := types.NewMsgMakeRoot(
				clientCtx.GetFromAddress().String(),
				accountHash,
				merklePath,
				"Contents",
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

func CmdMakeRootV2() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make-root-v2",
		Short: "Broadcast message make root",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
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
			h.Write([]byte(fmt.Sprintf("e%s%s", trackingNumber, clientCtx.GetFromAddress().String())))
			hash := h.Sum(nil)

			addressString := fmt.Sprintf("%x", hash)

			editors[addressString] = fmt.Sprintf("%x", "Placeholder key") // Determine if we need a place holder key

			jsonEditors, err := json.Marshal(editors)
			if err != nil {
				return err
			}

			msg := types.NewMsgMakeRootV2(
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
