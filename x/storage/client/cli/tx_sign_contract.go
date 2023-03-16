package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

const FlagPayUpfront = "pay-upfront"

func CmdSignContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-contract [cid]",
		Short: "Broadcast message sign-contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCid := args[0]

			pay, err := cmd.Flags().GetBool(FlagPayUpfront)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSignContract(
				clientCtx.GetFromAddress().String(),
				argCid,
				pay,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Bool(FlagPayUpfront, false, "Pay for the contract in advance.")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
