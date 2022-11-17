package cli

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDeleteFile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-file [path] [account]",
		Short: "Delete a file from an account's file explorer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argHashpath := args[0]
			argAccount := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			//Cut out the / at the end for compatibility with types/merkle-paths.go
			trimPath := strings.TrimSuffix(argHashpath, "/")
			merklePath := types.MerklePath(trimPath)

			h := sha256.New()
			h.Write([]byte(fmt.Sprintf("%s", argAccount)))
			hash := h.Sum(nil)

			accountHash := fmt.Sprintf("%x", hash)

			msg := types.NewMsgDeleteFile(
				clientCtx.GetFromAddress().String(),
				merklePath,
				accountHash,
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
