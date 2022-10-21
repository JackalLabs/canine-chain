package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdChangeOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "change-owner [file path] [fileOwner] [newOwner]",
		Short: "Broadcast message changeOwner",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argHashpath := args[0]
			argFileOwner := args[1]
			argNewOwner := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			trimPath := strings.TrimSuffix(argHashpath, "/")
			merklePath := types.MerklePath(trimPath)
			ownerChainAddress := MakeOwnerAddress(merklePath, argFileOwner)
			newOwnerChainAddress := MakeOwnerAddress(merklePath, argNewOwner)

			msg := types.NewMsgChangeOwner(
				clientCtx.GetFromAddress().String(),
				merklePath,
				ownerChainAddress,
				newOwnerChainAddress,
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
