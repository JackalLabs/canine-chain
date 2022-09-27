package cli

import (
	"strconv"

	"github.com/jackal-dao/canine/x/lp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = strconv.Itoa(0)

func CmdCreateLPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-l-pool \"{amount0}{denom0},...,{amountN){denomN}\" [invariant model id] \"swap fee multiplier\" [pool token lock duration (int64)] \"withdraw penalty multiplier\"",
		Short: "Broadcast message createLPool",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			// Parse args

			argCoins, err := sdk.ParseDecCoins(args[0])
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
			}

			invModelId, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
			}

			sfm, err := sdk.NewDecFromStr(args[2])
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
			}

			lockDuration, err := strconv.ParseInt(args[3], 10, 64)
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
			}

			wpm, err := sdk.NewDecFromStr(args[4])
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
			}

			msg := types.NewMsgCreateLPool(
				clientCtx.GetFromAddress().String(),
				argCoins,
				uint32(invModelId),
				sfm,
				lockDuration,
				wpm,
			)
			if err := msg.ValidateBasic(); err != nil {
				return sdkerrors.Wrapf(err, "Pool create message validation failure")
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
