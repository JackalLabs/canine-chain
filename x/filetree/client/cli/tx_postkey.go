package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPostkey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "postkey",
		Short: "Posts a users generated public key for the encryption model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			newKey, err := MakePrivateKey(clientCtx)
			if err != nil {
				return err
			}

			pubKey := newKey.PublicKey.Bytes(false)

			msg := types.NewMsgPostKey(
				clientCtx.GetFromAddress().String(),
				fmt.Sprintf("%x", pubKey),
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
