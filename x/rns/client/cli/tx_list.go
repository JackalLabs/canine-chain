package cli

import (
	"strconv"

	types2 "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v5/x/rns/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [name] [price]",
		Short: "List your name for sale.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argPrice := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coin, err := types2.ParseCoinNormalized(argPrice)
			if err != nil {
				return err
			}
			msg := types.NewMsgList(
				clientCtx.GetFromAddress().String(),
				argName,
				coin,
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
