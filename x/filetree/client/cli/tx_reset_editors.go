package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v4/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdResetEditors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset-editors [file path] [fileOwner]",
		Short: "Broadcast message resetEditors",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argHashpath := args[0]
			argFileowner := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			trimPath := strings.TrimSuffix(argHashpath, "/")
			merklePath := types.MerklePath(trimPath)
			ownerChainAddress := MakeOwnerAddress(merklePath, argFileowner)
			// how to notify everyone who is losing edit access?
			msg := types.NewMsgResetEditors(
				clientCtx.GetFromAddress().String(),
				merklePath,
				ownerChainAddress,
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
