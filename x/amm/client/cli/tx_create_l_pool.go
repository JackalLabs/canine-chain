package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ = strconv.Itoa(0)

func CmdCreatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool \"{amount0}{denom0},...,{amountN){denomN}\" [amm_id] \"swap fee multiplier\" [lock duration (blocks)] \"lock penalty multiplier\"",
		Short: "Broadcast message createLPool",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			// Parse args

			argCoins, err := sdk.ParseCoinsNormalized(args[0])
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

			msg := types.NewMsgCreatePool(
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
