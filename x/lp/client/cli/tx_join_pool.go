package cli

import (
	"strconv"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackal-dao/canine/x/lp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdJoinPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "join-pool \"Pool Name\" \"{amount0}{denom0},...,{amountN}{denomN} ...\" [lock duration (sec)]",
		Short: "join a liquidity pool by depositing pool coins.\n ",
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

			msg := types.NewMsgJoinPool(
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
