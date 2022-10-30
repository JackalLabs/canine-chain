package cli

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/x/lp/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap \"Pool Name\" \"[coin-input]\" \"[min-coin-output]\"",
		Short: "Broadcast message swap",
		Long:  "Broadcast message swap.\nCoin input field format: {amount}{denom}",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			coinIn, err := sdk.ParseDecCoin(args[1])
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
			}
			minCoinOut, err := sdk.ParseDecCoin(args[2])
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
			}

			msg := types.NewMsgSwap(
				clientCtx.GetFromAddress().String(),
				args[0],
				coinIn,
				minCoinOut,
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
