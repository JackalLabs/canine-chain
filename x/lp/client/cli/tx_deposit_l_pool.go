package cli

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackal-dao/canine/x/lp/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDepositLPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-l-pool \"Pool Name\" \"{amount0}{denom0},...,{amountN}{denomN} ...\" [lock duration (sec)]",
		Short: "Broadcast message depositLPool",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
			}

			lockDuration, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
			}

			deposit, err := sdk.ParseDecCoins(args[1])
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
			}

			msg := types.NewMsgDepositLPool(
				clientCtx.GetFromAddress().String(),
				args[0],
				deposit,
				lockDuration,
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
