package cli

import (
	"strconv"

	types2 "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/v4/x/rns/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bid [name] [bid]",
		Short: "Bid on someone elses name.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argBid := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coin, err := types2.ParseCoinNormalized(argBid)
			if err != nil {
				return err
			}
			msg := types.NewMsgBid(
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
